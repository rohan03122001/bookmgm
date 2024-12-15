
# BookStore API

A simple REST API built with Go to manage books. Perfect for learning Go web development and understanding basic CRUD operations.

## Quick Start 🚀

### Prerequisites
- Go 1.19+
- MySQL
- Git

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/bookstore.git
   cd bookstore
   ```

2. Update the database connection in `pkg/config/app.go`:
   ```go
   "your-username:your-password@tcp(localhost:3306)/your-database"
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:9010`

## API Endpoints

### Get all books
```http
GET /book/
```

### Get a specific book
```http
GET /book/{bookId}
```

### Add a new book
```http
POST /book/
```

#### Request Body:
```json
{
    "name": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "publication": "Addison-Wesley"
}
```

### Update a book
```http
PUT /book/{bookId}
```

#### Request Body:
```json
{
    "name": "Updated Name",
    "author": "Updated Author",
    "publication": "Updated Publication"
}
```

### Delete a book
```http
DELETE /book/{bookId}
```

## Project Structure
```
bookstore/
├── main.go                 # Entry point
├── pkg/
    ├── config/            # Database configuration
    ├── controllers/       # Request handlers
    ├── models/            # Data models
    ├── routes/            # API routes
    └── utils/             # Utility functions
```

## Technologies Used
- Gorilla Mux for routing
- GORM as ORM
- MySQL database

## Contributing
Feel free to fork and submit pull requests!

## License
MIT License
