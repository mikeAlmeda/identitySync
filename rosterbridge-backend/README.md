# identitySync
Pulls roster data from an external API, stores and updates records

## Local Development

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop) (for MongoDB)
- Go 1.25.3+ (or version in go.mod)

### Setup

1. **Start MongoDB**
   ```sh
   docker compose up -d mongo
   docker compose ps
   ```

2. **Configure environment**
   ```sh
   cp .env.example .env
   # Edit .env if needed (defaults work for local Docker)
   ```

3. **Run the server**
   ```sh
   go run ./cmd/server
   ```

4. **Test the API**
   ```sh
   curl -i http://localhost:8080/students
   ```

### Stopping services
```sh
docker compose down
```
