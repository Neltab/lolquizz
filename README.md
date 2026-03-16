
# lolquizz

A real-time multiplayer quiz game built with Go and React inspired by [KCulture](https://kculture.kgames.fr/).

This project serves as a way to learn Go, Domain Driven Design, websockets and reinforce my knowledge about web development in general. 
I want to use this project as a way to discover new technologies even though they may not be the correct choice in a real world application. For example I plan to implement services like Kafka, RabbitMQ or Redis, or at least explore how they could be used, even though I don't think I'll reach the scale needed to justify using those technologies. 

## Getting Started


### Prerequisites

- Go 1.23+
- BunJs
- (Optional) Docker

### Development
Start the backend
```bash
go run main.go
```
Start the frontend
```bash
cd  web
```
```bash
bun install
```
```bash
bun dev
```

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- WebSocket: ws://localhost:8080/ws?token=...

### Production
To be determined

### Docker
To be determined
