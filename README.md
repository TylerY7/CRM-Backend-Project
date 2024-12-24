# CRM-Backend-Project

Description 
The Customer Management API is a simple RESTful service built using Go (Golang) that allows for CRUD (Create, Read, Update, Delete) operations on a list of customers. The API supports the following operations for customers:

GET /customers: Retrieve a list of all customers.
GET /customers/{id}: Retrieve a single customer by their unique identifier.
POST /customers: Create a new customer.
PUT /customers/{id}: Update an existing customer's information.
DELETE /customers/{id}: Delete a customer by their unique identifier.

The API responses are in JSON format, and the application serves static HTML on the home route (/), displaying an overview of the API. The external packages that were used in this project are github.gorilla/mux to help with routing and github.com/google/uuid for generating unique IDs. The standard Go libraries that I used were encoding/json, net/http, fmt, and io/ioutil.

Installation
Install Go and Postman on your local machine 

Launch
Type "go run main.go" into your terminal and click allow after the Windows Security popup appears then go to Postman and enter your desired request.

Usage 
Endpoint: GET /customers
Description: Retrieve a list of all customers.
Request: No body is required for this request
Response: 
  200 OK Status 
  JSON array of all customers 

Endpoint: GET /customers/{id}
Description: Retrieve a single customer by their unique identifier.
Request: No body required just provide the ID in the URL
Response:
   If the customer is found: 200 OK Status and JSON objecting representing the customer
   If the customer is not found: 404 Not Found Status with a message saying "Customer not found"
   
Endpoint: POST /customers
Description: Create a new customer
Request Body:
  Sample Body:
  {
  "name": "Alice",
  "role": "Sales Representative",
  "email": "alice@example.com",
  "phone": "333-444-5555",
  "contacted": false
}
Response:
  201 Status Created
  JSON object representing the created customer with the auto-generated ID
  
Endpoint: PUT /customers/{id}
Description: Update an existing customer's information.
Response: 
  200 OK Status with a JSON object representing the updated customer

Endpoint: DELETE /customers/{id}
Description: Delete a customer by their unique identifier.
Request: No body is required just provide the Customer's ID in the URL
Response:
  If the customer is found: 200 OK Status
  If the customer is not found: 404 Not Found Status with a message saying "Customer not found"
