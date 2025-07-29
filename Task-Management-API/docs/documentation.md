# Task Management API Documentation

This document provides a comprehensive overview of the Task Management API, including its architecture, endpoints, and instructions for setup and testing.

## 1. Project Architecture

This project is built using Go and follows the principles of **Clean Architecture**. The code is organized into distinct layers, each with a specific responsibility. This separation of concerns makes the application modular, scalable, and easy to maintain.

### Directory Structure

-   **/delivery**: Handles the presentation layer, including the main application entry point.
    -   **/controllers**: Contains the HTTP handlers that parse requests and call the appropriate use cases.
    -   **/routers**: Defines all the API routes and wires them to the controllers.
    -   `main.go`: The main application entry point.
-   **/domain**: Represents the core business entities and rules.
    -   **/task**: Defines the `Task` struct and the `TaskRepository` interface.
    -   **/user**: Defines the `User` struct and the `UserRepository` interface.
-   **/usecases**: Contains the application-specific business logic. It orchestrates the flow of data between the domain and the repositories and defines its own interfaces.
-   **/repositories**: Implements the data access layer. This is where the `UserRepository` and `TaskRepository` interfaces are implemented using MongoDB.
-   **/infrastructure**: Contains external services and frameworks.
    -   `auth_middleware.go`: JWT authentication middleware.
    -   `jwt_service.go`: Logic for generating and validating JWTs.
    -   `password_service.go`: Logic for hashing and verifying passwords.
-   **/docs**: Contains project documentation.

## 2. API Endpoints

The API is divided into two main groups: authentication and tasks.

### Authentication (`/auth`)

These endpoints do not require an authentication token.

#### Register User

-   **Endpoint**: `POST /auth/register`
-   **Description**: Creates a new user account.
-   **Request Body**:
    ```json
    {
      "username": "testuser",
      "password": "password123"
    }
    ```
-   **Responses**:
    -   `201 Created`: User successfully created.
    -   `400 Bad Request`: Invalid request body.
    -   `409 Conflict`: Username already exists.

#### Login User

-   **Endpoint**: `POST /auth/login`
-   **Description**: Authenticates a user and returns a JWT.
-   **Request Body**:
    ```json
    {
      "username": "testuser",
      "password": "password123"
    }
    ```
-   **Responses**:
    -   `200 OK`: Login successful. Returns a JWT.
        ```json
        {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
        }
        ```
    -   `401 Unauthorized`: Invalid credentials.

### Tasks (`/api`)

All endpoints under this group are protected and require a valid JWT to be included in the `Authorization` header.

-   **Header Format**: `Authorization: Bearer <your_jwt_token>`

#### Get All Tasks

-   **Endpoint**: `GET /api/tasks`
-   **Description**: Retrieves a list of all tasks.
-   **Responses**:
    -   `200 OK`: Returns an array of task objects.

#### Create a Task

-   **Endpoint**: `POST /api/tasks`
-   **Description**: Creates a new task.
-   **Request Body**:
    ```json
    {
      "title": "My New Task",
      "description": "Details about the task.",
      "due_date": "2025-12-31T23:59:59Z",
      "status": "Pending"
    }
    ```
-   **Responses**:
    -   `201 Created`: Task successfully created.

#### Get Task by ID

-   **Endpoint**: `GET /api/tasks/:id`
-   **Description**: Retrieves a single task by its unique ID.
-   **Responses**:
    -   `200 OK`: Returns the requested task object.
    -   `404 Not Found`: No task with the specified ID was found.

#### Update a Task

-   **Endpoint**: `PUT /api/tasks/:id`
-   **Description**: Updates an existing task.
-   **Request Body**: (Include any fields you want to update)
    ```json
    {
      "status": "Completed"
    }
    ```
-   **Responses**:
    -   `200 OK`: Task successfully updated.
    -   `404 Not Found`: No task with the specified ID was found.

#### Delete a Task

-   **Endpoint**: `DELETE /api/tasks/:id`
-   **Description**: Deletes a task by its unique ID.
-   **Responses**:
    -   `200 OK`: Task successfully deleted.
    -   `404 Not Found`: No task with the specified ID was found.

## 3. Setup and Installation

To run this project locally, follow these steps:

1.  **Clone the Repository**:
    ```sh
    git clone <repository_url>
    cd Task-Management-API
    ```

2.  **Create a `.env` File**:
    Create a `.env` file in the root of the project and add the following environment variables:
    ```env
    MONGO_URI="your_mongodb_connection_string"
    JWT_SECRET="a_strong_and_secret_key_for_jwt"
    ```

3.  **Install Dependencies**:
    Download the necessary Go modules.
    ```sh
    go mod tidy
    ```

4.  **Run the Application**:
    Start the server from the project root.
    ```sh
    go run delivery/main.go
    ```
    The server will start on `http://localhost:8080`.

## 4. Testing

The project includes a comprehensive suite of unit tests covering the use cases, controllers, and infrastructure.

To run all tests, execute the following command from the project root:

```sh
go test -v ./...
```

### Mock Generation

The project uses `mockery` to generate mocks for interfaces, which is essential for isolating components during testing. If you modify an interface in the `domain` or `usecases` packages, you will need to regenerate the corresponding mock:

```sh
# Example for regenerating the UserRepository mock
mockery --name=UserRepository --dir=domain/user --output=domain/user/mocks

# Example for regenerating the UserUsecaseInterface mock
mockery --name=UserUsecaseInterface --dir=usecases --output=usecases/mocks
```
