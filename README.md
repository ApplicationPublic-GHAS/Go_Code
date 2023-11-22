# Sample Go App 

[![CircleCI Build Status](https://circleci.com/gh/CircleCI-Public/sample-go-cfd.svg?style=shield)](https://app.circleci.com/pipelines/github/CircleCI-Public/sample-go-cfd) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://app.circleci.com/pipelines/github/CircleCI-Public/dotnet-sample-cfd)

## Description

The sample Go app here is designed to demonstrate what a typical Go CI workflow may look on CircleCI.

You can see the CI pipelines for this application running [live on CircleCI](https://app.circleci.com/pipelines/github/CircleCI-Public/sample-go-cfd/?branch=main).

In this sample config, we have a single workflow `build-and-test` which will install and cache our required Go packages, and then run tests by running `go test -v ./...`.

## Getting Started

If you would like to copy the [config.yml](https://github.com/CircleCI-Public/sample-go-cfd/blob/main/.circleci/config.yml) and adapt it to your project, be sure to read the comments in the config file to ensure it works for your project. For more details, see the [CircleCI configuration reference](https://circleci.com/docs/2.0/configuration-reference/).


## About This App

This sample application is a Go REST API, and utilizes the entity framework to administer a SQL Server database running on the localhost. This application utilitzes the [OpenAPI/Swagger specification](https://swagger.io/specification/).


### Continuous Food Delivery

When you start up the service, you can open [this page](http://localhost:8080/CFD/1.0.0/ui/) in your browser to view the available API endpoints.


### Front-End

CFD(Continuous Food Delivery) is a sample application that relies on a separate UI framework. If you would like to run this project locally with a complete UI, you can use a valid CFD front-end, such as one of the following sample projects:

| Language |  GitHub | Description |
|---|---|---|
|  Javascript (Vue.js) | [Link](https://github.com/CircleCI-Public/sample-javascript-cfd)  | A Javascript Front-End for CFD |

## Run and Test Locally

If you would like to try this application out locally, you can find runtime instructions below.

### Requirements

Postgres

### Run Local Server

To run the server on a Docker container, please execute the following from the root directory:

```cmd
docker compose up
go run main.go
```


### Tests

To launch the unit tests, use the following:

```
go test -v ./...
```


## Additional Resources

* [CircleCI Docs](https://circleci.com/docs/) - The official CircleCI Documentation website.
* [CircleCI Configuration Reference](https://circleci.com/docs/2.0/configuration-reference/#section=configuration) - From CircleCI Docs, the configuration reference page is one of the most useful pages we have.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).
