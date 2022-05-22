# bookkeeper

This repository contains a web service example implementation build with golang using gofiber showing basic use cases. 

## Architecture
It leverages a clean architecture with the following layers:

- domain: contains the domain entities
- uc: contains the use cases
- store: handles persistence
- router: handles API requests

The architecture requires that the first 2 layers model the pure business logic without any external dependencies. 
The required functionality for storing and access is provided by the 3rd and 4th layer implementing the interfaces (Dependency Inversion Principle).

## Features
- clean architecture
- JWT Authentication
- Web API
- connection to PostgreSQL database
- migration with gomigrate
- testcontainers for integration tests
- Docker support

## Improvements
- add documentation
- add database transaction handling to use cases where appropriate
- add/improve test cases

## Inspired by
- https://github.com/err0r500/go-realworld-clean/
