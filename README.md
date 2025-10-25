# JWT Auth Backend (Express.js, Go-Fiber and others)

This project demonstrates a complete JWT (JSON Web Token) based authentication backend, uniquely implemented in two popular technologies for comparison: **Node.js with Express.js & TypeScript** and **Go with Fiber**.

Both applications provide the same core API for user login and profile retrieval, allowing for a direct comparison of syntax, project structure, performance, and ecosystem.

## Technologies Used

![Node.js](https://img.shields.io/badge/Node.js-339933?style=for-the-badge&logo=nodedotjs&logoColor=white)
![Express.js](https://img.shields.io/badge/Express.js-000000?style=for-the-badge&logo=express&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

## Project Structure

The repository is organized into two main directories, one for each implementation:

- **`/express-ts`**: Contains the backend application built with Express.js and TypeScript.
- **`/gofiber`**: Contains the backend application built with Go and the Fiber framework.

Both follow a modular structure, separating concerns for domains like `auth` and `user`.

## API Documentation

A standardized OpenAPI 3.0 specification for the API is available in the `/docs` directory.

- **`docs/openapi.yaml`**

You can use tools like the [Swagger Editor](https://editor.swagger.io/) to visualize the API documentation and interact with the endpoints.

## Getting Started

Below are the instructions to run both applications, either directly on your machine or using Docker.

### Prerequisites

- [Node.js](https://nodejs.org/en/) (v18 or later)
- [Go](https://go.dev/doc/install) (v1.21 or later)
- [Docker](https://www.docker.com/get-started)

--- 

### Running without Docker

#### 1. Express.js + TypeScript App

```bash
# Navigate to the express-ts directory
cd express-ts

# Install dependencies
npm install

# Start the server (usually on port 3000 )

# Dev Run 
npm run dev

# Prod Run
npm run build
npm start
```

#### 2. Go + Fiber App

```bash
# Navigate to the gofiber directory
cd gofiber

# Run the application (usually on port 3000 )
go run ./cmd/main.go
```

--- 

### Running with Docker

#### 1. Go + Fiber App

The `gofiber` application is fully containerized.

```bash
# From the project root, navigate to the gofiber directory
cd gofiber

# Build the Docker image
docker build -t gofiber-app -f docker/Dockerfile .

# Run the container
# This will map the container's port 3000 to your local machine's port 3000
docker run -p 3000:3000 gofiber-app
```

#### 2. Express.js + TypeScript App
The `express` application is fully containerized.

```bash
# From the project root, navigate to the gofiber directory
cd express-ts

# Build the Docker image
docker build -t express-app -f docker/Dockerfile .

# Run the container
# This will map the container's port 3000 to your local machine's port 3000
docker run -p 3000:3000 express-app
```