# TikTok Tech Immersion Backend Assignment Demo

![Tests](https://github.com/kxrt/tiktok-assignment/actions/workflows/test.yml/badge.svg)

## Description

This project extends the [TikTok Tech Immersion Backend Assignment](https://github.com/TikTokTechImmersion/assignment_demo_2023) with the following changes:
- Added a database service in `docker-compose.yml` to save the chat messages in a PostgreSQL database
- Connected the HTTP server to the RPC server to send and receive chat messages
- Connected the RPC server to the database service to save and retrieve chat messages

## Usage

### Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Go](https://golang.org/doc/install)

### Quickstart Instructions

1. Clone the repository
2. Navigate to the root directory and run:

```bash
docker-compose up -d
```

3. The API server will be running on `localhost:8080`, and can be accessed using the [sample requests](#sample-requests) below.
4. To stop the application, run:

```bash
docker-compose down
```

### API Routes

- `GET /ping` - Check if the server is running
- `POST /api/send` - Send a chat message
  - Query parameters:
    - `sender` - **REQUIRED**: The sender of the message
    - `receiver` - **REQUIRED**: The receiver of the message
    - `text` - **REQUIRED**: The text of the message
- `GET /api/pull` - Get all chat messages between two users
  - Query parameters:
    - `chat` - **REQUIRED**: The chat ID of the chat in the form of `user1:user2`
    - `limit` - The maximum number of messages to return
    - `reverse` - Whether to return the messages in reverse order
    - `cursor` - The starting time of the messages to return

### Sample requests

- Check if the server is running
```bash
curl localhost:8080/ping
```

- Send a chat message - "hi", from user "a" to user "b"
```bash
curl -X POST 'localhost:8080/api/send?sender=a&receiver=b&text=hi'
```

- Get all chat messages between user "a" and user "b"
```bash
curl 'localhost:8080/api/pull?chat=a%3Ab'
```

- Using additional query parameters
```bash
curl 'localhost:8080/api/pull?chat=a%3Ab&limit=5&reverse=true&cursor=0000000'
```

## Configuration

### PostgreSQL

The PostgreSQL database service is configured using the following environment variables in `docker-compose.yml`:
- `POSTGRES_USER` - The username of the database user
- `POSTGRES_PASSWORD` - The password of the database user
- `POSTGRES_DB` - The name of the database

The RPC server is configured using the following environment variables in `docker-compose.yml`:
- `DB_HOST` - The hostname of the database service
- `DB_PORT` - The port of the database service
- `DB_USER` - The username of the database user
- `DB_PASSWORD` - The password of the database user
- `DB_NAME` - The name of the database

These should be changed if the default values are not desired. The API server does not require any configuration.

To access the database through Docker Desktop, navigate to the running container's terminal and run:

```bash
psql -U {DB_USER} -d {DB_NAME}
```

Replace `{DB_USER}` and `{DB_NAME}` with the values in `docker-compose.yml`. Once logged in, you can run SQL queries, such as:

```sql
SELECT * FROM messages;
```

### API Server

To change the port of the API server, change the `ports` value in `docker-compose.yml`.

### RPC Server

To change the port of the RPC server, change the `ports` value in `docker-compose.yml`.

## Additional Notes

- Messages are restricted to a length of 500 characters. Change this in `db/init.sql` if required.

