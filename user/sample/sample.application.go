package main

import (
	"github.com/markdicksonjr/nibbler"
	NibElastic "github.com/markdicksonjr/nibbler-elasticsearch"
	NibUserElastic "github.com/markdicksonjr/nibbler-elasticsearch/user"
	"github.com/markdicksonjr/nibbler/user"
	"log"
)

type UserAndDbExtensions struct {
	UserExtension            *user.Extension
	UserPersistenceExtension user.PersistenceExtension
	DbExtension              nibbler.Extension
}

func allocateEsExtensions() UserAndDbExtensions {
	elasticExtension := NibElastic.Extension{}

	elasticUserExtension := NibUserElastic.Extension{
		ElasticExtension: &elasticExtension,
	}

	return UserAndDbExtensions{
		DbExtension:              &elasticExtension,
		UserPersistenceExtension: &elasticUserExtension,
		UserExtension: &user.Extension{
			PersistenceExtension: &elasticUserExtension,
		},
	}
}

func main() {

	// allocate logger and configuration
	config, err := nibbler.LoadConfiguration(nil)

	sqlExtensions := allocateEsExtensions()

	// initialize the application, provide config, logger, extensions
	appContext := nibbler.Application{}
	if err = appContext.Init(config, nibbler.DefaultLogger{}, []nibbler.Extension{
		sqlExtensions.DbExtension,
		sqlExtensions.UserPersistenceExtension,
		sqlExtensions.UserExtension,
	}); err != nil {
		log.Fatal(err.Error())
	}

	// create a test user
	emailVal := "test@example.com"
	_, errCreate := sqlExtensions.UserExtension.Create(&nibbler.User{
		Email: &emailVal,
	})

	if errCreate != nil {
		log.Fatal(errCreate.Error())
	}

	uV, errFind := sqlExtensions.UserExtension.GetUserByEmail(emailVal)

	if errFind != nil {
		log.Fatal(errFind.Error())
	}

	log.Println(uV)

	firstName := "testfirst"
	lastName := "testlast"
	uV.FirstName = &firstName
	uV.LastName = &lastName

	if err = sqlExtensions.UserExtension.Update(uV); err != nil {
		log.Fatal(err.Error())
	}

	// start the app
	if err = appContext.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
