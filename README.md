# CRM-Backend-Project

Description 
The Customer Management API is a simple RESTful service built using Go (Golang) that allows for CRUD (Create, Read, Update, Delete) operations on a list of customers. The API supports the following operations for customers:

GET /customers: Retrieve a list of all customers.
GET /customers/{id}: Retrieve a single customer by their unique identifier.
POST /customers: Create a new customer.
PUT /customers/{id}: Update an existing customer's information.
DELETE /customers/{id}: Delete a customer by their unique identifier.

The API responses are in JSON format, and the application serves static HTML on the home route (/), displaying an overview of the API.

Installation
Install Go and Postman on your local machine 

Launch
Type "go run main.go" into your terminal and click allow after the Windows Security popup appears then go to Postman and enter your desired request

Usage 
GET /customers: Retrieve a list of all customers.
GET /customers/{id}: Retrieve a single customer by their unique identifier.
POST /customers: Create a new customer.
PUT /customers/{id}: Update an existing customer's information.
DELETE /customers/{id}: Delete a customer by their unique identifier.
