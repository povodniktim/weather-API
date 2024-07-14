# weather-API
The Weather API is a Go application that fetches weather data for specified cities from a remote API (https://meteo.arso.gov.si/), processes the data, and provides it through a RESTful API. The application includes a caching mechanism to reduce redundant API calls and improve performance. It supports Docker for easy deployment.

## Features
- Fetch weather data for individual cities
- Fetch weather data for multiple cities
- Caching mechanism to minimize redundant API calls
- Error handling and logging
- Docker support for containerized deployment

## Requirements
- Go go1.22.5 or later
- Docker (optional, for Docker support)

## Installation

1. **Clone the repository:**

```sh
git clone https://github.com/yourusername/weather-api.git
cd weather-api
```

2. **Install dependencies:**
```sh
go mod tidy
```
## Usage
### Build and Run

**To build and run the application locally:**
```sh
go build -o weather-api .
./weather-api
```

**Or you can directly run the application using:**
```sh
go run .
```

**Or using CompileDaemon (Recommended):**
```sh
CompileDaemon -command="./weather-api"
```

## Docker

**To build and run the Docker container:**

1. **Build the Docker image:**
```sh
docker build -t weather-api .
```

1. **Run the Docker container:**
```sh
docker run -p 8080:8080 weather-api
```

## API Endpoints

- GET /weather  
  Fetches weather data for all configured cities.  
  Example: `GET /weather`

- GET /weather/:city  
  Fetches weather data for the specified city.  
  Example: `GET /weather/Reka`

## Configuration

The cities and their corresponding URLs are configured in the data/cities.go file. You can add or remove cities by modifying the city URL-s map.
