# Go CRUD Operations API Documentation

## Overview
This project is a product management API built using Go and SQLite. It provides endpoints to create, retrieve, update, and delete products.

## Login to obtain a JWT token
```bash
curl -X POST http://localhost:8000/login -H "Content-Type: application/json" -d '{"username": "username", "password": "password"}'
```

## GET products (requires token)
```bash
curl -X GET http://localhost:8000/products -H "Authorization: Bearer JWT_TOKEN"
```

## Create a product (requires token)
```bash
curl -X POST http://localhost:8000/products \
-H "Authorization: Bearer JWT_TOKEN" \
-H "Content-Type: application/json" \
-d '{"name": "GO Book", "price": 19.99, "description": "Go Programming Language"}'
```

## GET a single product by ID (requires token)
```bash
curl -X GET http://localhost:8000/products/1 -H "Authorization: Bearer JWT_TOKEN"
```

## Update a product by ID (requires token)
```bash
curl -X PUT http://localhost:8000/products/1 \
-H "Authorization: Bearer JWT_TOKEN" \
-H "Content-Type: application/json" \
-d '{"name": "GO Book 2", "price": 29.99}'
```

## Delete a product by ID (requires token)
```bash
curl -X DELETE http://localhost:8000/products/1 -H "Authorization: Bearer JWT_TOKEN"
```

## Docker
```bash
docker build -t store-app .
docker run -p 8000:8000 store-app
```
