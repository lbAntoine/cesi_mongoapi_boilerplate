# Go Gin MongoDB Boilerplate

A production-ready boilerplate for building RESTful APIs with Go, using Gin framework and MongoDB. This boilerplate implements generic repositories and provides a complete example with user management.

## Features

- Clean and scalable project structure
- Generic repository pattern for MongoDB
- Configuration management using YAML
- RESTful API implementation with Gin
- Graceful shutdown handling
- Example implementation with User model
- Type-safe generic implementations
- Built-in logging and recovery middleware

## Prerequisites

- Go 1.21 or higher
- MongoDB 4.0 or higher
- Make (optional, for using Makefile commands)

## Project Structure

```
.
├── README.md
├── cmd
│   └── server
│       └── main.go           # Application entry point
├── config.yaml              # Configuration file
├── go.mod                   # Go modules file
└── internal
    ├── api
    │   ├── app.go           # Application setup and HTTP handlers
    │   └── config.go        # Configuration management
    ├── db
    │   └── mongo.go         # MongoDB connection management
    ├── models
    │   ├── mod_user.go      # User model
    │   └── model.go         # Base model interface and implementation
    └── repositories
        ├── rep_user.go      # User-specific repository
        └── repository.go     # Generic repository interface and implementation
```

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-gin-mongo-boilerplate.git
   cd go-gin-mongo-boilerplate
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Configure the application:

   - Copy `config.yaml.example` to `config.yaml`
   - Modify the configuration values as needed

4. Run the application:

   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### User Management

- `POST /api/users` - Create a new user

  ```json
  {
    "email": "user@example.com",
    "name": "John Doe",
    "password": "securepassword"
  }
  ```

- `GET /api/users/:id` - Get user by ID
- `GET /api/users/email/:email` - Get user by email
- `GET /api/users` - List all users
- `PUT /api/users/:id` - Update user

  ```json
  {
    "name": "Updated Name"
  }
  ```

- `DELETE /api/users/:id` - Delete user

## Using the Generic Repository

The boilerplate provides a generic repository pattern that can be used with any model that implements the `Model` interface. Here's how to create a new repository for your model:

1. Create your model in `internal/models`:

```go
type YourModel struct {
    BaseModel `bson:",inline"`
    // Add your fields here
}
```

2. Create a specific repository in `internal/repositories` (optional):

```go
type YourModelRepository interface {
    Repository[*YourModel]
    // Add custom methods here
}

type MongoYourModelRepository struct {
    *MongoRepository[*YourModel]
}

func NewYourModelRepository(collection *mongo.Collection) YourModelRepository {
    return &MongoYourModelRepository{
        MongoRepository: NewMongoRepository[*YourModel](collection).(*MongoRepository[*YourModel]),
    }
}
```

3. Initialize the repository in `app.go`:

```go
yourModelRepo := repositories.NewYourModelRepository(mongodb.GetCollection("your_models"))
```

## Configuration

The application uses YAML for configuration. Here's an example configuration:

```yaml
server:
  port: 8080
  mode: debug # debug or release

mongodb:
  uri: "mongodb://localhost:27017"
  database: "boilerplate_db"
  timeout: 10 # seconds
```

## Best Practices

1. Always implement the `Model` interface for new models
2. Use the generic repository for basic CRUD operations
3. Create specific repositories only when custom functionality is needed
4. Use context for timeouts and cancellation
5. Implement proper error handling
6. Use proper logging
7. Follow RESTful API conventions

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
