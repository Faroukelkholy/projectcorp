# Project_Corp


## Table of contents
1. [Description](#Description)
2. [Documentation](#Documentation)
3. [Technologies](#Technologies)
4. [Deployment](#Deployment)
5. [Example](#Example)
6. [Testing](#Testing)
7. [Linting](#Linting)
8. [Future_Work](#Future_Work)

## 1. Description

Project_Corp is about project management and its participants.

## 2. Documentation

> documentation resides at ./doc folder.  

## 3. Technologies

Project is created with:

* Go
* Gin
* Postgres
* Docker

## 4. Deployment

> Note: The app will start with a seeder. 

```
$ docker-compose up -d
```

## 5. Example

>Application will be running on port 3000 and endpoints available are:

```
$ curl http://localhost:3000/projects
$ curl -d '{"name":"newProject","department":"sales","owner": "b82522f0-8644-4c65-a552-9c6b8a9e4b6f"}' -H "Content-Type: application/json" -X POST http://localhost:3000/projects
$ curl -d '{"name":"oldProject","department":"sales","owner": "b82522f0-8644-4c65-a552-9c6b8a9e4b6f"}' -H "Content-Type: application/json" -X PUT http://localhost:3000/projects/9eb8e365-3d7e-4994-960d-e51275343f23
$ curl -d '{"id":"a0d5e87a-af04-473d-b1f5-3105bbf986c8","department":"sales","role":"employee"}' -H "Content-Type: application/json" -X POST http://localhost:3000/projects/9eb8e365-3d7e-4994-960d-e51275343f23/participants
```
>More details on endpoint specification in the doc folder.
	
## 6. Testing

> 1\. First get a bash shell in the container

```
$ docker-compose exec corp_go bash
```

> 2\. Execute all test cases with coverage

```
$ go test -cover -v ./...
```

> 3\. Execute test cases for a specific package [optional]

```
$ go test -cover -v ./adapters/rest/controllers/project -v
$ go test -cover -v ./domain/usecase/project -v
```

> 4\. Execute only unit test cases [optional]

```
$ go test -short -cover -v ./...
```

> 5\. Execute only integration test cases [optional]

```
go test -v -run ".Integration" ./...
```

## 7. Linting
Project uses [golangci-lint](https://golangci-lint.run/). It is a go linter aggregator that can enable up to 48 linters.

#### 7.1 Configuartion

golanci-lint configuration is found in .golangci.yml file.

#### 7.2 Installation

```
# binary installation for linux and Windows, binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2
```

Check if the tool is correctly installed

```
golangci-lint --version
```

#### 7.3 Run the tool with the enabled linter

```
golangci-lint run
```

golangci-lint print out warning messages in command-line related to the enabled linters in order to fix them.

#### 7.4 Linters commands to automatically fix warning messages provided

To format all files in the project based on the gofmt linter. [Ref](https://stackoverflow.com/a/13333931/5486622)

```
gofmt -s -w -l .
```

To fix go import packages linting warnings based on goimport linter. [Ref](https://stackoverflow.com/a/59964885/5486622)

```
goimports -local ./ -w .
```
[Guide](https://stackoverflow.com/a/38714480/5486622) How you should group your package based on golang structure.

## 8. Future_Work

* Implement more coverage with Unit Test
* Implement intergration tests
* Implement functional tests