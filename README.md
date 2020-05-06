# Bank Service - POC

This project was intended as a POC to learn Go.

I tried to replicate most of my development workflow in other languages/frameworks. 

But what's this workflow?

For me, some project has some requirements to have a better productivity:

1. **Easy to Run**, preferably, one comand to run. So we could ramp up someone new fast
2. **Tests running from scratch**, the tests should not depend on nothing outside of the provided here
3. **BDD workflow**, specially with integration tests
4. **Coverage control on CI**, that should be implemented using a Continuous Integration Tool. It's not the proposal of this POC
5. **Code Quality Check**, such as [![Maintainability](https://api.codeclimate.com/v1/badges/0f7c53f39dba077e3a50/maintainability)](https://codeclimate.com/github/nelsonmhjr/bank_service/maintainability)

Why BDD and not TDD? Since it's an Rest API, the Requests and JSON Responses are the contract. So It's importante to have that in check always

## Instalation and Running

There is a makefile to automate some tasks, including building and running this app. 
Either using docker or running locally.
Why both? Well, tests run faster locally, but I prefer to spin the webserver using docker.
But there is all flavors to you...
  
### Running the app using Docker

Docker image for the app run using [CompileDaemon](https://github.com/githubnemo/CompileDaemon), it comes with HotReload Feature. It keep monitoring .go files for changes and rebuild the whole app

```bash
make # build docker image, if you wanna force 
```

```bash
make up # run docker-compose up in background (start the servers)
```

Access in http://localhost:8080/ping

```bash
make down # run docker-compose down (stop the servers)
```

### Running Tests

```bash
make test # run tests using docker images
make test_doc # run tests using docker images and print all tests documentation (Ginkgo Descriptions, Contexts and Its texts)
make clean # clean docker images and volumes BEWARE
```


###

