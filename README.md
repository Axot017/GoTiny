# GoTiny: A Link Shortener App

GoTiny is a web app that lets you create short and custom URLs from long ones. 
You can set a TTL or a visit limit for your links and track their clicks. 
GoTiny is open source and free to use. Please star or contribute to the project on GitHub if you like it.

You can use the app at [https://dev.goti.one](https://dev.goti.one) (todo)
or check out the [API docs](https://dev.goti.one/api/docs) and [SwaggerUI](https://dev.goti.one/api/swagger-ui)


## Table of Contents
1. [Introduction](#gotiny-a-link-shortener-app)
2. [Get Started](#get-started)
    * [Requirements](#requirements)
    * [Run locally](#run-locally)
    * [Test](#test)
3. [Technologies](#technologies)
    * [Backend](#backend)
    * [Frontend](#frontend)
    * [Infrastructure](#infrastructure)

## Get Started

### Requirements
* Go 1.20+ ([https://go.dev/doc/install](https://go.dev/doc/install))
* AWS CLI ([https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html] (https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html))
* Terraform ([https://developer.hashicorp.com/terraform/downloads](https://developer.hashicorp.com/terraform/downloads))
* Docker ([https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/))
* goswagger ([https://goswagger.io/install.html](https://goswagger.io/install.html))

### Run locally
You can run app directly with go:

  	go run cmd/gotiny/main.go

or using make

    make run
    
### Test
With go:

    go test ./...
    
or with make:

    make test

## Technologies

### Backend

The backend is written in Go, a fast and reliable programming language that is well suited for web development. The backend does not use many external packages, as the goal was to keep the code simple and maintainable. The backend relies mostly on the standard library features and functions. The backend uses the following libraries:
- chi router: A lightweight and idiomatic HTTP router that follows the standard library patterns and interfaces.
- fx: A dependency injection framework that simplifies the wiring and configuration of components and services.
- slog: A structured and leveled logging library that supports multiple output formats and contexts.
- aws sdk: A comprehensive and easy-to-use library that allows interacting with various AWS services.
- testify: A testing toolkit that provides assertions, mocks, suites and helpers for writing clear and concise tests.
- goswagger: A tool that generates API documentation and swagger-ui from annotations in the code.

### Frontend

TODO

### Infrastructure

The project is hosted on AWS, a secure and scalable cloud platform that offers a variety of services and features. The project also uses Terraform, an open-source tool that enables writing infrastructure as code. Terraform allows creating and managing the infrastructure resources in a declarative and consistent way. The project uses the following AWS services:
- App Runner: A fully managed service that makes it easy to deploy and run containerized web applications at scale.
- DynamoDB: A fast and flexible NoSQL database that provides consistent and low-latency performance.
- Route 53: A reliable and cost-effective DNS service that routes traffic to the best endpoint for the application.
- CloudFront: A global content delivery network that improves the speed and security of the web content delivery. (TODO)

## Architecture

The project follows the hexagonal architecture, a design pattern that aims to create loosely coupled and testable components. The project consists of three main parts:
* Core: The core contains the models and the business logic of the application. The core has no dependency on other parts, it only defines interfaces that the data part satisfies.
* Data: The data part handles the connection to the database and third-party services. The data part implements the interfaces defined by the core and provides the data access and manipulation functions.
* API: The API part contains the handlers for different routes. The API part receives the requests from the clients and calls the appropriate functions from the core part. The API part also returns the responses to the clients in a suitable format.
  
The project uses a service locator pattern to inject the dependencies between the parts. The project uses fx, a dependency injection framework, to create and configure the components and services. The project also uses chi router, a lightweight and idiomatic HTTP router, to handle the routing of the requests.
