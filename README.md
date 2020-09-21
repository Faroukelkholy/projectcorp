# project_Corp


## Table of contents
* [Documentation](#Documentation)
* [Technologies](#Technologies)
* [Deployment](#Deployment)
* [Example](#Example)
* [Testing](#Testing)
* [Future_Work](#Future_Work)


## Documentation

> documentation resides at ./doc folder. It includes rest api && database schema specification.  

## Technologies

Project is created with:

* Go
* Gin
* Postgres
* Docker

## Deployment

> Note: The app will start with a seeder. 

```
$ docker-compose up -d
```

## Example

>Application will be up and running on port 3000 and endpoints available are:

```
$ curl http://localhost:3000/projects
$ curl -d '{"name":"newProject","department":"sales","owner": "b82522f0-8644-4c65-a552-9c6b8a9e4b6f"}' -H "Content-Type: application/json" -X POST http://localhost:3000/projects
$ curl -d '{"name":"oldProject","department":"sales","owner": "b82522f0-8644-4c65-a552-9c6b8a9e4b6f"}' -H "Content-Type: application/json" -X PUT http://localhost:3000/projects/9eb8e365-3d7e-4994-960d-e51275343f23
$ curl -d '{"id":"a0d5e87a-af04-473d-b1f5-3105bbf986c8","department":"sales","role":"employee"}' -H "Content-Type: application/json" -X POST http://localhost:3000/projects/9eb8e365-3d7e-4994-960d-e51275343f23/participants
```
>More details on endpoint specification in the doc folder.
	
## Testing

> 1.first get a bash shell in the container 
> 2.excute the test cases

```
$ docker exec -it <container> /bin/bash
$ go test ./adapters/rest/controllers/project -v
$ go test ./domain/useCases/project -v
```

## Future_Work

* Implement more coverage with Unit Test
* Implement intergration tests
* Implement functional tests