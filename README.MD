# Agency-Type Backend

## Overview

This repository contains the backend code for the Online Song Library application. It provides a RESTful API for managing a collection of songs, including user authentication, managing results, and retrieving recordings. The backend is built on Go using the Gin framework for processing HTTP requests and PostgreSQL for database operations.

## Technologies Used

- **Go**: The primary programming language.
- **Gin**: A high performance, extensible, and minimalist web framework for Go.
- **PostgreSQL**: A powerful, open source object-relational database system.

## Prerequisites

- Go (version 1.16 or higher)
- PostgreSQL (version 10 or higher)
- Git

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/lmd1e/song_library
    ```

2. **Set up the environment variables:**

    Create a `.env` file in the root directory and add the following .env.example:

3. **Install dependencies:**

    ```sh
    go mod download
    ```

4. **Run the application:**

    ```sh
    go run .\app\main.go    
    ```

    The server will start running on `http://localhost:8080`.

## Api Endpoints
    ```sh
    http://localhost:8080/swagger/index.html
    ```