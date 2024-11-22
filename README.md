# System Login REST API

The System Login REST API provides a secure login system using JWT tokens. The API supports operations for user registration, login, and managing user data, including security features like token blacklisting for user logout.

## Key Features

- **User Registration**: New users can register via the `/register` endpoint.
- **User Login**: Users can log in with their email and password to receive a JWT token for authentication.
- **Authentication**: The API uses JWT tokens to authenticate users on restricted endpoints.
- **User Management**: Users can view, update, and delete their data through API endpoints protected by JWT tokens.
- **Logout with Token Blacklist**: Users can log out and blacklist their JWT token, ensuring it cannot be used again.

## API Endpoints

### Public Endpoints (No Authentication Required)

- `POST /register`: Register a new user.
- `POST /login`: Login a user and obtain a JWT token.
- `GET /public`: A public endpoint that can be accessed without authentication.

### Private Endpoints (Requires Authentication)

- `GET /users`: Retrieve a list of registered users (requires JWT token).
- `GET /{id}`: Get details of a user by ID (requires JWT token).
- `PUT /{id}`: Update user data by ID (requires JWT token).
- `DELETE /{id}`: Delete a user by ID (requires JWT token).
- `POST /logout`: Logout and blacklist the JWT token (requires JWT token).

## Installation

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/username/system-login-rest-api.git
cd system-login-rest-api
```

### 2. Instal dependensi
Run the following command to install the required dependencies:
```bash
go mod tidy
```

### 3. Run the Application
After the dependencies are installed, run the application with the following command:
```bash
go run main.go
```
The API will now be running and ready for use.

## API Usage
After the API is running, you can interact with it using tools like Postman or CURL.

### 1. User Registration
To register a new user, send a POST request to /register with the following JSON data:

```bash
{
  "username": "username",
  "email": "user@example.com",
  "password": "password123"
}
```
### 2. User Login
To log in, send a POST request to /login with the following JSON data:

```bash
{
  "email": "user@example.com",
  "password": "password123"
}
```
Response akan memberikan token JWT yang dapat digunakan untuk autentikasi:
```bash
{
  "token": "your-jwt-token"
}
```
Project Directory Structure
```bash
├── controllers/       # Contains handlers for API endpoints
├── middleware/        # Contains middleware for authentication and logging
├── models/            # Contains data models (e.g., User)
├── config/            # Contains database configurations and application setup
├── utils/             # Contains utility functions like JSON responses
├── main.go            # The main file to run the server
├── routes/            # Contains the routing logic for the API
└── README.md          # Project documentation

```
### Teknologi 
- Go: The main programming language for the backend.
- JWT (JSON Web Tokens): Used for authentication and authorization.
- Gorilla Mux: Router for handling HTTP requests.
- GORM: ORM for interacting with the database.
- PostgreSQL / MySQL: The database used for storing user data (optional, can be replaced   with another database).
### Author
- Name : Chandra F.A
- Github : https://github.com/cpzz18
- email : robbincfa@gmail.com

This version of the README is more polished, with improved structure, clarity, and formatting in English. It includes all necessary instructions for installation, usage, and contribution.
