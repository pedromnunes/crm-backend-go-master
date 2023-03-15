
The objective of this project is to develop a REST API(back-end) for a CRM application.
The application provides a set of endpoints to carry out CRUD operations on customer data.

Installation:

After configuring the Go runtime environment:

1. Access the project folder.
2. Run the command: 
    go run main.go


The service will be available on localhost on port 7000

API provides the following endpoints that can be used with cURL or Postman:

    Get a single customer through a /customers/{id}
    Get all the customers through a the /customers
    Creating a customer through a /customers
    Updating a customer through a /customers/{id}
    Deleting a customer through a /customers/{id}