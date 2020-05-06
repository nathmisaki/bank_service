# Bank Service - POC

This project was intended as a POC to learn Go.

I tried to replicate most of my development workflow in other languages/frameworks. 

But what's this workflow?

For me, some project has some requirements to have a better productivity:

1. **Easy to Run**, preferably, one comand to run. So we could ramp up someone new fast
2. **Tests running from scratch**, the tests should not depend on nothing outside of the provided here
3. **BDD workflow**, specially with integration tests
4. **Tests control on CI**, such as [![CircleCI](https://circleci.com/gh/nelsonmhjr/bank_service.svg?style=shield)](<LINK>) :point_left: click to see details
5. **Code Quality Check**, such as [![Maintainability](https://api.codeclimate.com/v1/badges/0f7c53f39dba077e3a50/maintainability)](https://codeclimate.com/github/nelsonmhjr/bank_service/maintainability) :point_left: click to see details

Why BDD and not TDD? Since it's an Rest API, the Requests and JSON Responses are the contract. So It's importante to have integration tests.
But Unit Test are also a must have on the project.

## What I used to built this project, and some _whys?_

* [**PostgreSQL**](https://www.postgresql.org/) _(version 12.2)_
* [**Go**](https://golang.org/) _(version 1.14.2)_
* [**Gin Framework**](https://gin-gonic.com/) _(version 1.6.3)_
* [**Ginkgo BDD**](https://onsi.github.io/ginkgo/) _(version 1.12.0)_

**Go**: The proposal was to use one of the 3 languages: **Java**, **Groovy** or **Go**.
And I don't have much experience in any of them (except Java, but not for backend services).
So my choice was based on the language culture. Go has a simplicity, and a philosophy to have a
focus on productivity, avoiding endless discussions about how to format your code or where you put your brackets.
That catch my interest, since it's a culture much shared with Ruby. And that culture
made programmers much more productive and projects much more valuables, in my career experience. 

**Gin**: I now that a http service can be done using just std lib in Go. And there is plenty of examples. But, using a framework, can save sometime doing stuff that are already been solved. AKA _don't reinvent the wheel_. Searching for it, I found that Gin is pretty popular. So I choose it. 

**Ginkgo**: Why not just use `testing` lib? Well, that's a personal preference on the syntax of BDD. I came from Ruby, and [RSpec](https://rspec.info/) is a great library to test and document your code behaviour, even in Unit Testing. So I searched for a lib that simulated the same syntax of RSpec.


## Instalation and Running

There is a [makefile](https://github.com/nelsonmhjr/bank_service/blob/master/makefile) to automate some tasks, including building and running this app. 

Either using docker or running locally.

Why both? Well, tests run faster locally, but I prefer to spin the webserver using docker.
But there is all flavors to you...
  
### Running the app using Docker

Docker runs the app using [CompileDaemon](https://github.com/githubnemo/CompileDaemon), it comes with HotReload Feature.
It keeps monitoring `.go` files for changes and `go build` the whole app inside the container.

The Docker Compose share the root of this directory with the container as a shared volume.
So it make faster to change something locally and check the change in the service, without having to `docker build` the image.

With the two combined, we have a fast development flow using docker and go!!

*You will need to rebuild if you install new packages or upgrade them. Basically, any changes to what's inside [Dockerfile](https://github.com/nelsonmhjr/bank_service/blob/master/Dockerfile)

#### Handy Comands

```bash
make # build docker image, if you wanna force 
```

```bash
make up # run docker-compose up in background (start the servers)
```

Access in http://localhost:8080/ping

```bash
make down # run docker-compose down (stop the servers)
make logs # run docker-compose logs -f --tail=100 api (show the logs on screen and keep showing)
```

### Running Tests

```bash
make test # run tests using docker images
make test_doc # run tests using docker images and print all tests 
make clean # clean docker images and volumes. BEWARE can erase the local Database
```

### Running locally

If you prefere to run locally, or run the tests locally, I recommend to install [godotenv](https://github.com/joho/godotenv)
Also, it will need Go installed locally too. The version used in this project is listed above.
Please refer to [Go-lang Website](https://golang.org/) to download the distribution to your OS.

```bash
go get github.com/joho/godotenv/cmd/godotenv
```

**godotenv** automates loading environment variables

Then create a `.env` file in the root of the repository

```bash
tee -a .env <<-EOS
DB_HOST=localhost
DB_USER=bank_service
DB_PASSWORD=yQGKSuS8tC
DB_PORT=5432
DB_NAME_PREFIX=bank_service
DB_SSLMODE=disable
EOS
```
Remember to replace for values for the one you used when installed your Postgres.

These values above reflect the same values on docker-compose. So that you can
only spin the db up `docker-compose up db` 
and execute the go using local machine.

With the `.env` file inplace, you can run the usual commands using `godotenv` before them.

Example:

```bash
godotenv go run main.go
godotenv go test -v ./...
godotenv ginkgo -r -cover # this needs to install ginkgo locally too
```

To make it easier there is some...

#### Handy Commands

```bash
make test_local # Run all pkg tests 
make test_local_debug # Run all pkg tests with debug logs (show querys and requests)
make cover_report # show in your browser the coverage report (need to run some make test before)
```

## Author

### Nelson Haraguchi <nelsonmhjr@gmail.com> [@nelsonmhjr](https://github.com/nelsonmhjr)


If you found this useful in some way, please send me a message or an email!
