# Warehouse Management System
This is a simple implementation of a warehouse management system

## Assumptions
I am making an assumptions such that a single product in the database represents an item stored in a warehouse, in other words, an inventory. This means a warehouse can have multiple inventories as long as the total sum of all quantities of the inventories stored in the warehouse does not exceed the maximum capacity of the warehouse.

## Prerequisites
1. Ensure that you have installed Go >= 1.24 and Postgres
2. Populate a .env file, you can duplicate the .env.example and fill it with your postgres credential and jwt configuration
3. Install all the dependencies with go mod download
4. Ensure that you are in the root directory of the project
5. Run the service with go run main.go
6. You can use the included postman collection to try out the APIs