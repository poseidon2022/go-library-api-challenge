# node-library-api

## Implement Simple CRUD API for Library Management

Your task is to implement a simple CRUD API using an in-memory database underneath.

### Files to Work On

For this task, focus on the following files:

1. **`book.service.js`**

   - Handles the business logic related to books.

2. **`book.controller.js`**
   - Manages the API endpoints and handles requests related to books.
3. **`book.router.js`**
   - uncomment the code to use the router.

## Details:

1.  The task must be solved using only **express** and provided packages in **package.json**. Any libraries and packages **are prohibited**.
2.  API path `/books`:

    - **GET** `/books` - returns a list of all books. there should be a possibility to filter books by `author` , `title` , and `genre ` query parameters and sort by `year ` query parameters. If no query parameters are provided, all books should be returned.
    - **GET** `/books/:id` - returns a single book by id.
    - **POST** `/books` - creates a new book.
    - **PUT** `/books/:id` - updates a book by id.

3.  The book object should have the following properties:

    - `id` - unique identifier
    - `title` - title of the book
    - `author` - author of the book
    - `genre` - genre of the book
    - `year` - year of publication

4.  Requests to non-existing endpoints (e.g., `/some-non/existing/resource`) should be handled.
5.  Internal server errors should be handled and processed correctly.

### Start Development

```bash
go run main.go
```

### Run Tests

```bash
go test
```