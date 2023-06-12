# TikTok Tech Immersion Backend Assignment Demo

![Tests](https://github.com/kxrt/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

## Description

This project extends the [TikTok Tech Immersion Backend Assignment](https://github.com/TikTokTechImmersion/assignment_demo_2023) with the following changes:
- Added a database service in `docker-compose.yml` to save the chat messages in a PostgreSQL database
- Connected the HTTP server to the RPC server to send and receive chat messages
- Connected the RPC server to the database service to save and retrieve chat messages

## Usage

### Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Go](https://golang.org/doc/install)

### Running the application

1. Clone the repository
2. Run `docker-compose up -d` in the root directory

### API Routes

- `GET /ping` - Check if the server is running
- `POST /api/send` - Send a chat message
  - Query parameters:
    - `sender` - REQUIRED: The sender of the message
    - `receiver` - REQUIRED: The receiver of the message
    - `text` - REQUIRED: The text of the message
- `GET /api/pull` - Get all chat messages between two users
  - Query parameters:
    - `chat` - REQUIRED: The chat ID of the chat in the form of `user1:user2`
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