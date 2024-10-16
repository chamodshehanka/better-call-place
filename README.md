# Better Call Place

Better Call Place is a lightweight service leveraging the Google Places API for easy integration of location-based data into your applications. Retrieve details about nearby businesses, points of interest, and geolocation seamlessly.

## Features

- Fetch place suggestions based on user input.
- Utilize Google Places API for accurate and up-to-date information.
- Modular architecture with controllers, services, and routes.
- Middleware for logging requests using `zerolog`.

## Project Structure

- `controllers/`: Handles HTTP requests and responses.
- `services/`: Contains the business logic and API calls.
- `routes/`: Maps endpoints to controller functions.
- `middlewares/`: Contains middleware functions for logging.
- `configs/`: Manages configuration and environment variables.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Google Places API Key

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/chamodshehanka/better-call-place.git
    cd better-call-place
    ```

2. Create a `.env` file based on `.env.example` and set your environment variables:
    ```dotenv
    PORT=8080
    GOOGLE_PLACE_API_KEY=your_api_key_here
    ```

3. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Application

1. Start the server:
    ```sh
    go run main.go
    ```

2. The server will start on the port specified in the `.env` file. By default, it will run on port 8080.

### API Endpoints

- **GET /places**: Fetch place suggestions based on the query parameter.

  Example request:
    ```sh
    curl -X GET "http://localhost:8080/places?query=Colombo"
    ```

### Middleware

The project includes a logging middleware using `zerolog` to log request details.

### Configuration

Configuration is managed using environment variables. The `configs` package loads these variables and ensures all required variables are set.

### Example Configuration

```dotenv
PORT=8080
GOOGLE_PLACE_API_KEY=your_api_key_here
```