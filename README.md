# Final Project

## Project Overview

This project is a fully functional system for managing customers, authors, books, orders, and sales. It is built using Go and provides a REST API to perform CRUD operations and generate sales reports. The system implements robust error handling and logical constraints to maintain data integrity.

### Key Features

1. **Customer Management**:
   - Create, update, retrieve, delete, and search customers.
   - A customer cannot be deleted if they are linked to an order.

2. **Author Management**:
   - Create, update, retrieve, delete, and search authors.
   - When an author is deleted, their books are deleted unless the books are part of an order.

3. **Book Management**:
   - Create, update, retrieve, delete, and search books.
   - Books with stock `0` cannot be used in new orders, but they remain available for historical orders.

4. **Order Management**:
   - Create, update, retrieve, delete, and search orders.
   - Requires existing customers and books to create an order.
   - Deleting an order increases the stock of the associated books.

5. **Sales Reports**:
   - Generate and retrieve sales reports for a specific date range or generate a report instantly.
   - Reports summarize sales within the last 24 hours, even if no orders exist.

---

## Steps to Run the Project

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/kawtarlabzae/Introduction-to-Go.git
   ```

2. **Navigate to the Final Project Folder**:
   ```bash
   cd Final-project
   ```

3. **Run the Main File**:
   ```bash
   go run main.go
   ```

---

## Folder Structure

The `Final-project` folder is organized as follows:

```
Final-project/
├── Controllers/         # Logic for handling requests and routing
├── Documentation/       # Project documentation files
├── InmemoryStores/      # In-memory data storage (e.g., for testing purposes)
├── Interfaces/          # Interface definitions for abstractions
├── StructureData/       # Data structures (e.g., structs for Customers, Books, etc.)
├── swaggerfiles/        # Swagger API definitions
├── utils/               # Utility functions or helpers
├── authors.json         # Sample data for authors
├── books.json           # Sample data for books
├── customers.json       # Sample data for customers
├── go.mod               # Go module file for dependencies
├── go.sum               # Go dependency checksum file
├── main.go              # Main entry point for the application
├── orders.json          # Sample data for orders
└── sales_reports.json   # Sample sales report data
```

This structure ensures a clean and modular organization for the project.

---

## How to Use the System

### 1. **Customer Management**
   - Create a customer before proceeding with orders.
   - Example JSON for creating a customer:
     ```json
     {
       "name": "John Doe",
       "email": "john.doe@example.com",
       "address": {
         "street": "123 Elm St",
         "city": "Springfield",
         "state": "IL",
         "postal_code": "62701"
       }
     }
     ```

### 2. **Author Management**
   - Create an author or let the system automatically create one when adding a book.
   - Authors can be deleted, but their books will remain if part of an order.

### 3. **Book Management**
   - Add books for an author. If the author does not exist, the system will create the author.
   - Example JSON for creating a book:
     ```json
     {
       "title": "The Great Adventure",
       "price": 19.99,
       "stock": 100,
       "author": {
         "first_name": "Jane",
         "last_name": "Doe",
         "bio": "Famous author of adventurous tales."
       }
     }
     ```

### 4. **Order Management**
   - Create an order with existing customers and books. The stock of books will decrease by the ordered quantity.
   - If a book’s stock is `0`, it will not be included in the order.
   - Example JSON for creating an order:
     ```json
     {
       "customer": {
         "id": 1
       },
       "items": [
         {
           "book": {
             "id": 1
           },
           "quantity": 2
         }
       ]
     }
     ```
   - Deleting an order restores the stock of the associated books.

### 5. **Sales Reports**
   - View sales reports for all orders or a specific date range.
   - Example query to retrieve sales for a date range:
     ```http
     GET /reports/sales?start_date=2025-01-01&end_date=2025-01-31
     ```
   - Generate a sales report instantly:
     ```http
     POST /reports/sales/generate
     ```

---

## Search Criteria

### General Search Notes
1. Path parameters like `id` are used for specific resource retrieval.
2. For complex filters, use the POST method with a JSON body containing a comprehensive search criteria structure.

Example search criteria:
```json
{
  "ids": [1, 2],
  "names": ["John Doe"],
  "emails": ["john.doe@example.com"],
  "min_created_at": "2025-01-01T00:00:00Z",
  "max_created_at": "2025-01-12T23:59:59Z"
}
```

---

## Constraints and Business Logic

1. **Orders**:
   - Must reference existing customers and books.
   - Cannot include books with stock `0`.
   - Deleting an order adjusts the stock of the associated books.

2. **Authors**:
   - Deleting an author deletes their books unless the books are part of an order.

3. **Books**:
   - Books with stock `0` cannot be used in new orders.

4. **Customers**:
   - Cannot be deleted if they exist in an order.

---

