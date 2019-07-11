# Nibbler Elasticsearch

A simple extension for Nibbler that connects to elasticsearch and provide an easy interface
for operations.

## Environment Variables

If "Url" is not set on the extension instance, the url will be derived from "elastic.url" 
(ELASTIC_URL), then "database.url" (DATABASE_URL).

Credentials will be pulled (optionally) from "elastic.user" (ELASTIC_USER) and elastic.password" (ELASTIC_PASSWORD)
