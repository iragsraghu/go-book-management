# Go CRUD API for Book Management

This project is a simple **CRUD (Create, Read, Update, Delete) API** built using **Go (Golang)**.  
It includes endpoints to manage items and provides **Swagger API documentation**.  

## Technologies Used
- **Go (Golang)** – Backend API  
- **Gin Framework** – HTTP web framework  
- **Swagger** – API documentation  
- **Docker** – Containerization  

## Running the Application

Clone the repository:

```bash
git clone git@github.com:iragsraghu/go-book-management.git
cd go-book-management
docker build -t go-book-management .
docker run -p 8080:8080 --env-file .env go-book-management
```

## Access the Application

Swagger Documentation → http://localhost:8080/swagger/index.html