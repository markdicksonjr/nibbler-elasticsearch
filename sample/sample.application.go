package main

import (
	"github.com/markdicksonjr/nibbler"
	nes "github.com/markdicksonjr/nibbler-elasticsearch"
	"log"
)

func main() {
	logger := nibbler.DefaultLogger{}
	config, err := nibbler.LoadConfiguration()
	nibbler.LogFatalNonNil(logger, err)

	// initialize the application, provide config, logger, extensions
	appContext := nibbler.Application{}
	if err = appContext.Init(config, logger, []nibbler.Extension{
		&nes.Extension{},
	}); err != nil {
		log.Fatal(err.Error())
	}

	if err = appContext.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
