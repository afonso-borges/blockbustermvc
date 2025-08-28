# Go Blockbuster MVC

A movie rental management system built with Go, Gin framework, and PostgreSQL database.

## Architecture

This project follows the MVC (Model-View-Controller) pattern with clean architecture principles:

- **Models**: Domain entities, DTOs, and interfaces
- **Repositories**: Data access layer with PostgreSQL integration
- **Services**: Business logic layer
- **Controllers**: HTTP request handlers using Gin framework

## Features

- **Movie Management**: Create, read, update, delete movies
- **User Management**: User registration and management
- **Loan System**: Movie borrowing and returning functionality
- **Database Migrations**: Automated database schema management
- **Environment Configuration**: Flexible configuration via environment variables

## Prerequisites

- Go 1.24.1 or higher
- PostgreSQL 15+ (latest version recommended)
- Git

## Database Setup

1. Install PostgreSQL and create a database:
```sql
CREATE DATABASE blockbuster;
```

2. Copy the environment configuration:
```bash
cp .env.example .env
```

3. Update `.env` with your database credentials:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blockbuster
DB_SSL_MODE=disable
SERVER_PORT=8080
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-blockbuster-mvc
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run cmd/main.go
```

The application will:
- Connect to PostgreSQL database
- Run migrations automatically
- Start the server on port 8080 (or configured port)

## API Endpoints

### Movies
- `POST /movies` - Create a new movie
- `GET /movies` - Get all movies
- `GET /movies/:id` - Get movie by ID
- `PUT /movies/:id` - Update movie
- `DELETE /movies/:id` - Delete movie

### Users
- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user

### Loans
- `POST /loans` - Create a new loan
- `GET /loans` - Get all loans
- `GET /loans/:id` - Get loan by ID
- `POST /loans/:id/return` - Return a movie
- `GET /loans/users/:userId/loans` - Get user's loans

## Database Schema

The application uses PostgreSQL with the following tables:
- `users` - User information
- `movies` - Movie catalog with quantity tracking
- `loans` - Movie borrowing records

## Technologies Used

- **Go 1.24.1** - Programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **pgx/v5** - PostgreSQL driver
- **golang-migrate** - Database migrations
- **UUID** - Unique identifiers

---

**Status**: ✅ Production Ready
