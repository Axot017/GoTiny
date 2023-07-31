# GoTiny: A Link Shortener App

GoTiny is a web app that lets you create short URLs from long ones. 
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
4. [Architecture](#architecture)

## Get Started

### Requirements
* Go 1.20+ (https://go.dev/doc/install)
* AWS CLI (https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
* Terraform (https://developer.hashicorp.com/terraform/downloads)
* Docker (https://www.docker.com/products/docker-desktop/)
* goswagger (https://goswagger.io/install.html)

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

This section describes the main technologies and tools used to develop and deploy the GoTiny app. It covers both the backend and the frontend aspects of the project.

### Backend

The backend is designed in Go, a fast and reliable programming language ideal for web development. Because the intention was to keep the code basic and maintainable, the backend does not use many external packages. The backend is primarily dependent on standard library features and functions. The following libraries are used by the backend:
- chi router: A lightweight and idiomatic HTTP router that follows the standard library patterns and interfaces. [[Learn more]](https://htmx.org/)
- fx: A dependency injection framework that simplifies the wiring and configuration of components and services. [[Learn more]](https://github.com/uber-go/fx)
- slog: A structured and leveled logging library that supports multiple output formats and contexts. [[Learn more]](https://pkg.go.dev/golang.org/x/exp/slog)
- aws sdk: A comprehensive and easy-to-use library that allows interacting with various AWS services. [[Learn more]](https://github.com/aws/aws-sdk-go-v2)
- testify: A testing toolkit that provides assertions, mocks, suites and helpers for writing clear and concise tests. [[Learn more]](https://github.com/stretchr/testify)
- goswagger: A tool that generates API documentation and swagger-ui from annotations in the code. [[Learn more]](https://goswagger.io)

### Frontend

The frontend is very basic HTML and CSS, using the Tailwind CSS framework. Tailwind CSS is a utility-first framework that provides low-level and customizable classes for styling elements. [[Learn more]](https://github.com/go-chi/chi).

The pages are rendered with Go templates, a powerful and simple templating engine that allows injecting data into HTML files. [[Learn more]](https://pkg.go.dev/html/template).

To add reactivity to the pages, the htmx library is used. htmx is a lightweight and modern library that enables high-performance AJAX without writing any JavaScript code. [[Learn more]](https://htmx.org/).

### Infrastructure

The project is hosted on Amazon Web Services (AWS), a secure and scalable cloud platform with a wide range of services and capabilities. Terraform, an open-source tool for writing infrastructure as code, is also used in the project. Terraform enables the declarative and consistent creation and management of infrastructure resources. The following AWS services are used in the project:
- App Runner: A fully managed service that makes it easy to deploy and run containerized web applications at scale. [[Learn more]](https://aws.amazon.com/apprunner/)
- DynamoDB: A fast and flexible NoSQL database that provides consistent and low-latency performance. [[Learn more]](https://aws.amazon.com/dynamodb/)
- Route 53: A reliable and cost-effective DNS service that routes traffic to the best endpoint for the application. [[Learn more]](https://aws.amazon.com/route53/)
- CloudFront: A global content delivery network that improves the speed and security of the web content delivery. [[Learn more]](https://aws.amazon.com/cloudfront/)

## Architecture

The project follows to hexagonal architecture, a design pattern that aims to create loosely coupled and testable components. The project is divided into three major sections:
* Core: The core contains the models and the business logic of the application. The core has no dependency on other parts, it only defines interfaces that the data part satisfies.
* Data: The data part handles the connection to the database and third-party services. The data part implements the interfaces defined by the core and provides the data access and manipulation functions.
* API: The API part contains the handlers for different routes. The API part receives the requests from the clients and calls the appropriate functions from the core part. The API part also returns the responses to the clients in a suitable format.
  
The project uses a service locator pattern to inject the dependencies between the parts. The project uses fx, a dependency injection framework, to create and configure the components and services. The project also uses chi router, a lightweight and idiomatic HTTP router, to handle the routing of the requests.
