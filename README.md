# API Documentation

## Overview
This project is a product management API built using Go and SQLite. It provides endpoints to create, retrieve, update, and delete products.

## Base URL

## Endpoints

### 1. Create a Product

- **Endpoint**: `/products`
- **Method**: `POST`
- **Description**: Create a new product.
- **Request Body**:
  - `name` (string, required): The name of the product.
  - `description` (string, required): A description of the product.
  - `price` (float, required): The price of the product.
  - Additional optional fields: `data_created`, `last_updated`, `is_active`, `json_meta`, `type`, `uuid`.

- **Example Request**:
```bash
curl -X POST http://localhost:8000/products \
-H "Content-Type: application/json" \
-d '{
  "name": "Sample Product",
  "description": "This is a sample product.",
  "price": 19.99
}'


curl -X GET http://localhost:8000/products

curl -X GET http://localhost:8000/products/1
