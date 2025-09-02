# Book Management API

A RESTful API backend service built with Go, designed to work with a Vue.js frontend for managing books, authors, and users with authentication capabilities.

## Features

- **User Authentication**: JWT-like token-based authentication system
- **Book Management**: CRUD operations for books with author and genre relationships
- **User Management**: Admin functionality for managing users
- **Author Management**: Support for book authors
- **Genre Support**: Many-to-many relationship between books and genres
- **File Upload**: Cover image upload support for books
- **Middleware**: Authentication and CORS support
- **PostgreSQL Integration**: Full database integration with connection pooling

## Technology Stack

- **Language**: Go
- **Web Framework**: Chi Router v5
- **Database**: PostgreSQL with pgx driver
- **Authentication**: Token-based authentication
- **Image Processing**: Base64 encoding for cover uploads
- **Logging**: Structured logging with built-in Go logger
- **Development**: Docker Compose for local development

## Project Structure

```
.
├── cmd/api/              # Application entry point and HTTP handlers
│   ├── main.go          # Main application entry
│   ├── handlers.go      # HTTP request handlers
│   ├── routes.go        # Route definitions
│   ├── middleware.go    # Authentication middleware
│   └── helpers.go       # Helper functions
├── internal/            # Internal application packages
│   ├── data/           # Data models and database operations
│   │   ├── model.go    # Base models and database connection
│   │   └── books.go    # Book-related data operations
│   └── driver/         # Database connection driver
│       └── driver.go   # PostgreSQL connection setup
├── db-data/            # PostgreSQL configuration
├── docker-compose.yml  # Docker services configuration
└── Makefile           # Build and development commands
```

## Quick Start

1. **Prerequisites**
   - Go 1.21+
   - PostgreSQL 13+
   - Docker and Docker Compose (optional)

2. **Environment Setup**
   ```bash
   # Start PostgreSQL with Docker Compose
   docker-compose up -d
   
   # Set environment variables
   export DSN="host=localhost port=5432 user=postgres password=password dbname=books sslmode=disable timezone=UTC connect_timeout=5"
   export ENV=development
   ```

3. **Build and Run**
   ```bash
   # Build the application
   make build
   
   # Run the application
   make run
   ```

4. **API Access**
   - API Server: `http://localhost:8082`
   - Static Files: `http://localhost:8082/static/`

## API Endpoints

### Authentication
- `POST /users/login` - User login
- `POST /users/logout` - User logout
- `POST /validate-token` - Token validation

### Public Endpoints
- `GET /books` - Get all books
- `GET /books/{slug}` - Get book by slug

### Admin Endpoints (Requires Authentication)
- `POST /admin/users` - Get all users
- `POST /admin/users/save` - Create/Update user
- `POST /admin/users/get/{id}` - Get user by ID
- `POST /admin/users/delete` - Delete user
- `POST /admin/books/save` - Create/Update book
- `POST /admin/books/delete` - Delete book
- `POST /admin/books/{id}` - Get book by ID
- `POST /admin/authors/all` - Get all authors

## Architecture Diagrams

### System Architecture
```mermaid
graph TB
    subgraph "Client Layer"
        VUE[Vue.js Frontend]
        API_CLIENT[API Client]
    end
    
    subgraph "API Layer"
        ROUTER[Chi Router]
        MIDDLEWARE[Auth Middleware]
        CORS[CORS Middleware]
    end
    
    subgraph "Handler Layer"
        AUTH_HANDLER[Auth Handlers]
        BOOK_HANDLER[Book Handlers]
        USER_HANDLER[User Handlers]
    end
    
    subgraph "Service Layer"
        USER_MODEL[User Model]
        BOOK_MODEL[Book Model]
        TOKEN_MODEL[Token Model]
        AUTHOR_MODEL[Author Model]
    end
    
    subgraph "Data Layer"
        DB_DRIVER[Database Driver]
        POSTGRES[(PostgreSQL)]
    end
    
    subgraph "Static Layer"
        STATIC[Static Files]
        COVERS[Book Covers]
    end
    
    VUE --> API_CLIENT
    API_CLIENT --> ROUTER
    ROUTER --> MIDDLEWARE
    ROUTER --> CORS
    MIDDLEWARE --> AUTH_HANDLER
    MIDDLEWARE --> BOOK_HANDLER
    MIDDLEWARE --> USER_HANDLER
    
    AUTH_HANDLER --> USER_MODEL
    AUTH_HANDLER --> TOKEN_MODEL
    BOOK_HANDLER --> BOOK_MODEL
    BOOK_HANDLER --> AUTHOR_MODEL
    USER_HANDLER --> USER_MODEL
    
    USER_MODEL --> DB_DRIVER
    BOOK_MODEL --> DB_DRIVER
    TOKEN_MODEL --> DB_DRIVER
    AUTHOR_MODEL --> DB_DRIVER
    
    DB_DRIVER --> POSTGRES
    BOOK_HANDLER --> STATIC
    STATIC --> COVERS
```

### Entity Relationship Diagram (ERD)
```mermaid
erDiagram
    USERS {
        int id PK
        string email UK
        string password
        string first_name
        string last_name
        int user_active
        timestamp created_at
        timestamp updated_at
    }
    
    TOKENS {
        int id PK
        int user_id FK
        string email
        string token UK
        bytea token_hash
        timestamp created_at
        timestamp updated_at
        timestamp expiry
    }
    
    AUTHORS {
        int id PK
        string author_name
        timestamp created_at
        timestamp updated_at
    }
    
    BOOKS {
        int id PK
        string title
        int author_id FK
        int publication_year
        string slug UK
        text description
        timestamp created_at
        timestamp updated_at
    }
    
    GENRES {
        int id PK
        string genre_name
        timestamp created_at
        timestamp updated_at
    }
    
    BOOKS_GENRES {
        int book_id FK
        int genre_id FK
        timestamp created_at
        timestamp updated_at
    }
    
    USERS ||--o{ TOKENS : "has"
    AUTHORS ||--o{ BOOKS : "writes"
    BOOKS ||--o{ BOOKS_GENRES : "has"
    GENRES ||--o{ BOOKS_GENRES : "categorizes"
```

### Authentication Flow Sequence Diagram
```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant AH as Auth Handler
    participant UM as User Model
    participant TM as Token Model
    participant DB as Database
    
    Note over C,DB: User Login Process
    C->>R: POST /users/login {email, password}
    R->>AH: Login Handler
    AH->>UM: GetByEmail(email)
    UM->>DB: SELECT user by email
    DB-->>UM: User data
    UM-->>AH: User object
    AH->>AH: PasswordMatches(password)
    AH->>TM: GenerateToken(userID, 24h)
    TM-->>AH: Token object
    AH->>TM: Insert(token, user)
    TM->>DB: INSERT token
    DB-->>TM: Success
    TM-->>AH: Success
    AH-->>R: {token, user}
    R-->>C: 200 OK {token, user}
    
    Note over C,DB: Authenticated Request
    C->>R: GET /admin/books (Authorization: Bearer token)
    R->>AH: AuthTokenMiddleware
    AH->>TM: AuthenticateToken(request)
    TM->>DB: SELECT token
    DB-->>TM: Token data
    TM->>TM: Validate expiry
    TM-->>AH: Valid user
    AH->>R: Continue to handler
    R->>AH: Book Handler
    AH-->>C: Book data
```

### Book Management Flow Sequence Diagram
```mermaid
sequenceDiagram
    participant C as Client
    participant R as Router
    participant M as Middleware
    participant BH as Book Handler
    participant BM as Book Model
    participant AM as Author Model
    participant DB as Database
    participant FS as File System
    
    Note over C,FS: Create/Update Book
    C->>R: POST /admin/books/save {book data, cover image}
    R->>M: AuthTokenMiddleware
    M->>BH: Authenticated request
    BH->>BH: Parse request payload
    BH->>BH: Decode base64 cover image
    BH->>FS: Save cover image to /static/covers/
    FS-->>BH: Success
    
    alt New Book (ID = 0)
        BH->>BM: Insert(book)
        BM->>DB: INSERT book
        DB-->>BM: New book ID
        BM->>DB: DELETE old genres
        BM->>DB: INSERT new book_genres
        DB-->>BM: Success
        BM-->>BH: Book ID
    else Update Book (ID > 0)
        BH->>BM: Update(book)
        BM->>DB: UPDATE book
        BM->>DB: DELETE old book_genres
        BM->>DB: INSERT new book_genres
        DB-->>BM: Success
        BM-->>BH: Success
    end
    
    BH-->>R: Success response
    R-->>C: 202 Accepted {success message}
    
    Note over C,DB: Get Book Details
    C->>R: GET /books/{slug}
    R->>BH: OneBook handler
    BH->>BM: GetOneBySlug(slug)
    BM->>DB: SELECT book with author
    DB-->>BM: Book and author data
    BM->>BM: genresForBook(bookID)
    BM->>DB: SELECT genres for book
    DB-->>BM: Genre data
    BM-->>BH: Complete book object
    BH-->>R: Book data
    R-->>C: 200 OK {book with author and genres}
```

### Service Layer Architecture
```mermaid
graph TB
    subgraph "HTTP Layer"
        ROUTES[Route Definitions]
        HANDLERS[HTTP Handlers]
        MIDDLEWARE[Middleware Stack]
    end
    
    subgraph "Business Logic Layer"
        AUTH_SERVICE[Authentication Service]
        BOOK_SERVICE[Book Management Service]
        USER_SERVICE[User Management Service]
        FILE_SERVICE[File Upload Service]
    end
    
    subgraph "Data Access Layer"
        USER_REPO[User Repository]
        TOKEN_REPO[Token Repository]
        BOOK_REPO[Book Repository]
        AUTHOR_REPO[Author Repository]
        GENRE_REPO[Genre Repository]
    end
    
    subgraph "Infrastructure Layer"
        DB_POOL[Connection Pool]
        LOGGER[Logging Service]
        CONFIG[Configuration]
    end
    
    subgraph "External Layer"
        POSTGRES[(PostgreSQL)]
        FILESYSTEM[File System]
    end
    
    ROUTES --> HANDLERS
    HANDLERS --> MIDDLEWARE
    HANDLERS --> AUTH_SERVICE
    HANDLERS --> BOOK_SERVICE
    HANDLERS --> USER_SERVICE
    HANDLERS --> FILE_SERVICE
    
    AUTH_SERVICE --> USER_REPO
    AUTH_SERVICE --> TOKEN_REPO
    BOOK_SERVICE --> BOOK_REPO
    BOOK_SERVICE --> AUTHOR_REPO
    BOOK_SERVICE --> GENRE_REPO
    USER_SERVICE --> USER_REPO
    FILE_SERVICE --> FILESYSTEM
    
    USER_REPO --> DB_POOL
    TOKEN_REPO --> DB_POOL
    BOOK_REPO --> DB_POOL
    AUTHOR_REPO --> DB_POOL
    GENRE_REPO --> DB_POOL
    
    DB_POOL --> POSTGRES
    
    AUTH_SERVICE --> LOGGER
    BOOK_SERVICE --> LOGGER
    USER_SERVICE --> LOGGER
    
    AUTH_SERVICE --> CONFIG
    BOOK_SERVICE --> CONFIG
    USER_SERVICE --> CONFIG
```

## Database Schema

### Users Table
- Primary authentication entity
- Stores user credentials and profile information
- Supports active/inactive status

### Tokens Table
- JWT-like token management
- Links to users for authentication
- Includes expiry management

### Books Table
- Central entity for book information
- Links to authors (many-to-one)
- Includes slug for SEO-friendly URLs
- Supports cover image storage

### Authors Table
- Author information
- One-to-many relationship with books

### Genres Table
- Genre categories
- Many-to-many relationship with books via junction table

### Books_Genres Junction Table
- Implements many-to-many relationship
- Links books to multiple genres

## Development Commands

```bash
# Build application
make build

# Run application  
make run

# Start application
make start

# Stop application
make stop

# Restart application
make restart

# Clean build artifacts
make clean
```

## Configuration

### Environment Variables
- `DSN`: PostgreSQL connection string
- `ENV`: Environment (development/production)

### Default Configuration
- **Port**: 8082
- **Database Timeout**: 3 seconds
- **Token Expiry**: 24 hours
- **Max DB Connections**: 5
- **Connection Lifetime**: 5 minutes

## Security Features

- **Token-based Authentication**: Secure API access
- **Password Hashing**: bcrypt for password security
- **CORS Configuration**: Cross-origin request handling
- **SQL Injection Protection**: Parameterized queries
- **Input Validation**: Request payload validation

## API Response Format

```json
{
    "error": false,
    "message": "Success message",
    "data": {
        // Response data
    }
}
```

## Error Handling

The API uses consistent error response format:
```json
{
    "error": true,
    "message": "Error description"
}
```

Common HTTP status codes:
- `200 OK`: Successful GET requests
- `201 Created`: Successful resource creation
- `202 Accepted`: Successful POST/PUT requests
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Authentication required/failed
- `500 Internal Server Error`: Server-side errors

## Testing

The project includes test files for:
- Handlers (`handlers_test.go`)
- Routes (`routes_test.go`)
- Models (`models_test.go`)
- Helpers (`helpers_test.go`)

Run tests with:
```bash
go test ./...
```
