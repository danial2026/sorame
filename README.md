# Sorame

A minimalist URL shortener service built with Go and Redis.

🌐 **[Live Demo](https://sorame.danials.space)**

## Features

- ⚡️ Fast URL shortening
- 🔗 Clean, shortened URLs
- 📱 Responsive, minimal design
- 📋 Automatic clipboard copy
- ⏳ 30-day link expiration
- 🔒 URL validation
- 🌍 Hosted at [sorame.danials.space](https://sorame.danials.space)

## Tech Stack

- **Backend**: Go - Fast, reliable server implementation
- **Database**: Redis - In-memory data store for quick access
- **Frontend**: Vanilla JavaScript, HTML, CSS - Lightweight and performant

## API Endpoints

- `POST /link` - Create a shortened URL
  ```json
  {
    "data": "https://your-long-url.com"
  }
  ```
  Returns a shortened URL ID

- `GET /link/{shareID}` - Redirect to original URL
  - Fast redirect to the original destination
  - Validates URL before redirect

## Prerequisites

- Go 1.19 or higher
- Docker
- Docker Compose  

## Environment Variables

The project uses environment variables for configuration. 
To set them up, Run the provided script to generate secure credentials:
```bash
bash generate_credentials.sh
```

This will:
- Create `.env` file from template if not exists
- Generate secure random passwords
- Set up required Docker network
- Configure Redis credentials

## Installation & Development

1. Clone the repository
   ```bash
   git clone https://github.com/danial2026/sorame.git
   cd sorame
   ```

2. Install dependencies
   ```bash
   export GOPROXY=direct && export GOSUMDB=off  && go mod tidy
   ```

3. Start Redis
   ```bash
   docker-compose -f docker-compose.yml up -d redis-service redis-commander
   ```

4. Run the application
   ```bash
   go run ./main.go
   ```

The service will be available at `http://localhost:3000`

## Deployment

```bash
# Build for linux production
bash build_linux.sh

# Run with Docker
docker-compose -f docker-compose.yml up -d
```

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## License

This project is open source and available under the [GNU GENERAL PUBLIC LICENSE](LICENSE).