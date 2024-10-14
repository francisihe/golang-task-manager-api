# Go Task Manager API
- Francis Ihejirika

This is a basic Task Manager API built with Go, without using any frameworks. It provides endpoints to manage tasks, including creating, reading, updating, and deleting tasks. The API uses JWT for authentication and GORM for database interactions. This project is a simple CRUD application to try out Go.

NOTE: The coebase contains comments with alternative methods of performing certain operations, and the login functionality is just a dummy to see the JWT token creation and checker methods. It's not written to actually work.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/francisihe/golang-task-manager-api.git
    cd golang-task-manager-api
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Create a `.env` file in the root directory and add your environment variables:

    ```env
    DB_HOST=your_db_host
    DB_PORT=your_db_port
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_name
    PORT=your_server_port
    ```

4. Run the application:

    ```sh
    go run main.go
    ```

    OR

     ```sh
    go run .
    ```

## Usage

The API will be available at `http://localhost:8080` or whatever port you specified in your env file. You can use tools like `curl` or Postman to interact with the API.

## API Endpoints

### Authentication

- **Login**

    ```http
    POST /login
    ```

    Request Body:

    ```json
    {
        "username": "admin",
        "password": "password"
    }
    ```

    Response:

    ```json
    {
        "token": "your_jwt_token"
    }
    ```

### Tasks

- **Get All Tasks**

    ```http
    GET /api/tasks
    ```

    Response:

    ```json
    [
        {
            "id": "task_id",
            "title": "Task Title",
            "description": "Task Description",
            "created_at": "timestamp",
            "updated_at": "timestamp"
        }
    ]
    ```

- **Create Task**

    ```http
    POST /api/tasks
    ```

    Request Body:

    ```json
    {
        "title": "Task Title",
        "description": "Task Description"
    }
    ```

    Response:

    ```json
    {
        "id": "task_id",
        "title": "Task Title",
        "description": "Task Description",
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
    ```

- **Update Task**

    ```http
    PUT /api/tasks/{id}
    ```

    Request Body:

    ```json
    {
        "title": "Updated Task Title",
        "description": "Updated Task Description"
    }
    ```

    Response:

    ```json
    {
        "id": "task_id",
        "title": "Updated Task Title",
        "description": "Updated Task Description",
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
    ```

- **Delete Task**

    ```http
    DELETE /api/tasks/{id}
    ```

    Response:

    ```json
    {
        "message": "Task deleted successfully"
    }
    ```

## Environment Variables

- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_NAME`: Database name
- `SECRET_KEY`: Secret key for JWT

## Project Structure

```plaintext
task-manager-api/
│
├── main.go                # Entry point of the application
├── handlers/              # Contains HTTP handler functions
│   └── loginJWTsample.go  # JWT login handler
│   └── task.go            # Task-related handler functions
├── middlewares/           # Contains middleware functions
│   └── authJWTchecker.go  # JWT authentication middleware
│   └── cors.go            # CORS middleware
│   └── ratelimit.go       # Rate limiting middleware
├── models/                # Contains models (e.g., Task, User)
│   └── task.go            # Task model definition
│   └── userCredentials.go # User credentials model definition
├── routes/                # Contains route setup
│   └── router.go          # Sets up all the API routes
├── .env                   # Environment variables
├── .gitignore             # Git ignore file
├── clarityNotes.md        # Clarity notes
├── developerNotes.md      # Developer notes
├── go.mod                 # Go module definition
├── go.sum                 # Go module dependencies
└── README.md              # This file