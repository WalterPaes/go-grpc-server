# go-grpc-crud

## Introduction
This repository contains an implementation of gRPC Server in Go. The service provides functionalities to manage products.

## Installation
To install and run this service, follow these steps:

1. Ensure you have Go installed on your machine. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository to your local machine.
   ```bash
   git clone https://github.com/WalterPaes/go-grpc-crud.git
   ```
3. Navigate to the project directory.
   ```bash
   cd go-grpc-crud
   ```
4. Run the following command to install the dependencies.
   ```bash
   go mod tidy
   ```
5. Finally, run the service.
   ```bash
   go run cmd/server/main.go
   ```

## Features
- **Create Product**: Allows the creation of a new product with name, category, description, and price.
- **Find Product by ID**: Retrieves product details by its ID.

## Usage
This service exposes gRPC endpoints for interacting with the product data. Below are the available endpoints:

### Create Product
To create a new product, call the `Create` method with the necessary product details.

### Find Product by ID
To retrieve product details by ID, call the `FindById` method with the product ID.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.