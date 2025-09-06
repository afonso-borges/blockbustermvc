# Go Blockbuster MVC

A movie rental management system built with Go, implementing clean architecture principles and MVC pattern. This system provides comprehensive functionality for managing movies, users, and loan operations through both REST API and web interface.

## Project Overview

Go Blockbuster MVC is designed as a scalable movie rental platform that demonstrates modern Go development practices. The system handles movie inventory management, user registration, and loan tracking with automated database migrations and environment-based configuration.

## Architecture

The project follows clean architecture principles with clear separation of concerns:

### Layer Structure

- **Models**: Domain entities and data transfer objects
- **Repositories**: Data access layer with PostgreSQL integration
- **Services**: Business logic and domain rules
- **Controllers**: HTTP request handlers and routing
- **Web**: Template-based user interface

### Module Organization

```
internal/
├── database/          # Database configuration and migrations
├── movies/           # Movie management module
├── users/            # User management module
├── loans/            # Loan operations module
└── web/              # Web interface module
```

Each module implements the repository pattern with:

- Repository layer for data persistence
- Service layer for business logic
- Controller layer for HTTP handling

### Technology Stack

- **Go 1.24.1** - Core programming language
- **Gin Framework** - HTTP web framework with middleware support
- **PostgreSQL** - Primary database with ACID compliance
- **pgx/v5** - High-performance PostgreSQL driver
- **golang-migrate** - Database schema versioning
- **UUID** - Unique identifier generation
- **CORS** - Cross-origin resource sharing support
- **godotenv** - Environment variable management

## Features

### Core Functionality

- **Movie Catalog Management**: Complete CRUD operations for movie inventory
- **User Management**: User registration and profile management
- **Loan System**: Movie borrowing and return tracking
- **Inventory Control**: Automatic quantity management and availability checking
- **Database Migrations**: Automated schema management and versioning

### API Capabilities

- RESTful API design with consistent endpoints
- JSON request/response handling
- CORS-enabled for cross-origin requests
- Environment-based configuration
- Connection pooling for database efficiency

## Prerequisites

- **Go**: Version 1.24.1 or higher
- **PostgreSQL**: Version 15+ recommended
- **Git**: For repository management

## Installation

### 1. Clone Repository

```bash
git clone <repository-url>
cd go-blockbuster-mvc
```

### 2. Environment Configuration

Create environment file with database credentials:

```bash
cp .env.example .env
```

Configure the following environment variables:

```env
BLK_DATABASE_HOST=localhost
BLK_DATABASE_PORT=5432
BLK_DATABASE_USER=postgres
BLK_DATABASE_PASSWORD=your_password
BLK_DATABASE_NAME=blockbuster
BLK_DATABASE_SSL_MODE=disable
SERVER_PORT=8080
```

### 3. Database Setup

#### Option A: Local PostgreSQL

```sql
CREATE DATABASE blockbuster;
```

#### Option B: Docker Compose

```bash
docker-compose up -d
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run Application

```bash
go run cmd/api/main.go
```

The application will automatically:

- Connect to PostgreSQL database
- Start HTTP server on configured port

## Database Schema

### Tables

- **users**: User profiles and authentication data
- **movies**: Movie catalog with inventory tracking
- **loans**: Rental records and return status

### Migration Management

Database migrations are located in `internal/database/migrations/` and managed using [tern](https://github.com/jackc/tern). Migrations are **not** executed automatically on startup and must be run manually.

To run migrations:

```bash
go run cmd/terndotenv/main.go
```

This command:

- Loads environment variables from `.env` file using `godotenv`
- Executes `tern migrate` with the proper configuration
- Uses the tern configuration file at `./internal/database/migrations/tern.conf`
- Applies all pending migrations to the database

**Note:** Make sure your database environment variables are properly configured before running migrations.

## API Documentation

### Base URL

```
http://localhost:8080/api
```

### Movies Endpoints

- `POST /movies` - Create new movie
- `GET /movies` - List all movies
- `GET /movies/:id` - Get movie details
- `PUT /movies/:id` - Update movie information
- `DELETE /movies/:id` - Remove movie from catalog

### Users Endpoints

- `POST /users` - Register new user
- `GET /users` - List all users
- `GET /users/:id` - Get user profile
- `PUT /users/:id` - Update user information
- `DELETE /users/:id` - Delete user account

### Loans Endpoints

- `POST /loans` - Create new loan
- `GET /loans` - List all loans
- `GET /loans/:id` - Get loan details
- `POST /loans/:id/return` - Process movie return
- `GET /loans/users/:userId/loans` - Get user's loan history

### Web Interface

- `/` - Dashboard and movie catalog
- `/loans` - Loan management interface

## Development

### Project Structure

```
go-blockbuster-mvc/
├── cmd/
│   ├── api/                 # Application entry point
│   └── terndotenv/         # Migration utility
├── internal/
│   ├── database/           # Database configuration
│   ├── models/             # Domain entities
│   ├── movies/             # Movie module
│   ├── users/              # User module
│   ├── loans/              # Loan module
│   └── web/                # Web interface
├── templates/              # HTML templates
├── docker-compose.yml      # PostgreSQL container
├── go.mod                  # Go module definition
└── README.md              # Project documentation
```

### Code Organization Principles

- Each module is self-contained with its own repository, service, and controller
- Database operations are abstracted through repository interfaces
- Business logic is centralized in service layers
- HTTP concerns are isolated in controller layers
- Configuration is environment-driven with sensible defaults
