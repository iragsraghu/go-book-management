# Go CRUD API for Book Management

This project is a simple **CRUD (Create, Read, Update, Delete) API** built using **Go (Golang)**.  
It includes endpoints to manage items and provides **Swagger API documentation**.  

## 🚀 Prerequisites
- [Go](https://go.dev/dl/) installed
- [Docker](https://docs.docker.com/get-docker/) installed
- [swag](https://github.com/swaggo/swag) installed for Swagger generation

## Running the Application

Clone the repository:

```bash
git clone git@github.com:iragsraghu/go-book-management.git

cd go-book-management

## Generate Swagger, Build, and Run in one step:
make all

or 

## Run individually:
make swagger       # Generate Swagger docs
make docker-build  # Build Docker image
make docker-run    # Run Docker container
```

## Access the Application

Swagger Documentation → http://localhost:8080/swagger/index.html