# LoLQuizz - Architecture

## Overview

LoLQuizz is a real-time multiplayer quiz game built with Go (backend) and React/TypeScript (frontend). The architecture follows **Domain-Driven Design (DDD)** with four clean layers.

## Architecture Diagram

```mermaid
---
title: LoLQuizz - Architecture Overview
---
graph TB
    subgraph Client["Client"]
        React["React/TypeScript SPA\n(web/dist)"]
    end

    subgraph Interfaces["Interfaces Layer"]
        subgraph HTTP["HTTP (interfaces/http)"]
            Server["Server\nserver.go"]
            MW["Middleware\nCORS - Logging"]
            Auth["AuthHandler\nPOST /api/auth/login"]
            RoomH["RoomHandler\nPOST /api/rooms\nGET /api/rooms/:code"]
            SPA["SPAHandler\nGET /*"]
        end
        subgraph WS["WebSocket (interfaces/ws)"]
            Hub["Hub\nClient registry\nRoom broadcasting"]
            ClientWS["Client\nReadPump - WritePump"]
            Router["Router\nMessage routing\nEvent subscriptions"]
            Msg["Messages\nIncomingMessage\nOutgoingMessage"]
        end
    end

    subgraph Application["Application Layer"]
        RoomSvc["RoomService\nCreate - Join - Leave\nUpdate Settings"]
        GameSvc["GameService\nStart - SubmitAnswer\nNextQuestion - Timer"]
        SessionSvc["SessionService\nCreate - Validate\nTTL management"]
    end

    subgraph Domain["Domain Layer"]
        subgraph RoomAgg["Room Aggregate"]
            Room["Room\nCode - Status - Players"]
            Player["Player\nName - IsHost - Score"]
            Settings["Settings\nMaxPlayers - QuestionCount"]
            RoomEvents["Events\nPlayerJoined - PlayerLeft\nSettingsUpdated"]
            RoomRepo["Repository Interface"]
        end
        subgraph GameAgg["Game Aggregate"]
            Game["Game\nPhase - Questions - Scores"]
            Question["Question\nText - Image - Duration"]
            Answer["Answer\nPlayerId - Points"]
            GameEvents["Events\nQuestionStarted - AnswerSubmitted\nAnswerJudged - GameFinished"]
        end
        subgraph Shared["Shared"]
            IDs["PlayerId - RoomId - GameId"]
            EventIntf["Event Interface"]
        end
        Session["Session\nPlayerId - ExpiresAt"]
    end

    subgraph Infrastructure["Infrastructure Layer"]
        EventBus["EventBus\nPublish/Subscribe\n(bus/event_bus.go)"]
        MemRoomRepo["In-Memory RoomRepository\n(memory/room_repo.go)"]
        MemQRepo["In-Memory QuestionRepository\n(memory/question_repo.go)"]
    end

    subgraph Infra["Infrastructure"]
        MainGo["main.go\nDependency Wiring"]
        Config["Config\nPORT - ALLOWED_ORIGINS - SESSION_TTL"]
        Docker["Docker\nMulti-stage Build\nNode - Go - Alpine"]
    end

    %% Client connections
    React -->|"REST API"| MW
    React -->|"WebSocket /ws"| Hub

    %% HTTP flow
    MW --> Auth
    MW --> RoomH
    MW --> SPA
    Server --> MW

    %% Handler dependencies
    Auth --> SessionSvc
    RoomH --> RoomSvc
    RoomH --> SessionSvc

    %% WebSocket flow
    Hub <--> ClientWS
    ClientWS --> Router
    Router --> RoomSvc
    Router --> GameSvc

    %% Service dependencies
    RoomSvc --> Room
    RoomSvc --> EventBus
    RoomSvc --> MemRoomRepo
    GameSvc --> Game
    GameSvc --> EventBus
    GameSvc --> MemRoomRepo
    GameSvc --> MemQRepo
    SessionSvc --> Session

    %% Event flow
    EventBus -->|"Domain Events"| Router
    Router -->|"Broadcast to Room"| Hub

    %% Domain relationships
    Room --> Player
    Room --> Settings
    Room --> RoomEvents
    Game --> Question
    Question --> Answer
    Game --> GameEvents
    RoomEvents --> EventIntf
    GameEvents --> EventIntf
    Room --> IDs
    Game --> IDs
    Player --> IDs

    %% Infrastructure implements domain
    MemRoomRepo -.->|"implements"| RoomRepo

    %% Wiring
    MainGo --> Config
    MainGo --> Server

    %% Styling
    classDef domain fill:#e8f5e9,stroke:#2e7d32
    classDef app fill:#e3f2fd,stroke:#1565c0
    classDef infra fill:#fff3e0,stroke:#e65100
    classDef iface fill:#f3e5f5,stroke:#6a1b9a
    classDef client fill:#fce4ec,stroke:#c62828

    class Room,Player,Settings,RoomEvents,RoomRepo,Game,Question,Answer,GameEvents,IDs,EventIntf,Session domain
    class RoomSvc,GameSvc,SessionSvc app
    class EventBus,MemRoomRepo,MemQRepo,MainGo,Config,Docker infra
    class Server,MW,Auth,RoomH,SPA,Hub,ClientWS,Router,Msg iface
    class React client
```

## Layers

| Layer | Directory | Purpose |
|-------|-----------|---------|
| **Domain** | `internal/domain/` | Pure business logic — Room/Game aggregates, events, repository interfaces |
| **Application** | `internal/application/` | Orchestration — RoomService, GameService, SessionService |
| **Infrastructure** | `internal/infrastructure/` | Technical details — in-memory repos, EventBus (pub/sub) |
| **Interfaces** | `internal/interfaces/` | External communication — REST API + WebSocket |

## Data Flows

### REST API

```
Client --> Middleware --> Handler --> Service --> Domain --> Repository
```

### WebSocket

```
Client --> Hub --> Router --> Service --> Domain --> EventBus --> Hub --> Broadcast to room
```

### Domain Events

```
Domain Event --> EventBus.Publish --> Router handler --> Hub.BroadcastToRoom --> All clients in room
```

## Key Design Patterns

- **Domain-Driven Design** — Aggregates (Room, Game), value objects, domain events
- **Repository Pattern** — `room.Repository` interface with in-memory implementation
- **Event-Driven Architecture** — EventBus for decoupled pub/sub communication
- **Hub Pattern** — WebSocket Hub manages clients and room-based broadcasting
- **Dependency Injection** — All dependencies wired explicitly in `main.go`

## External Dependencies

| Dependency | Purpose |
|------------|---------|
| `github.com/google/uuid` | UUID generation for player/room/game IDs |
| `github.com/gorilla/websocket` | WebSocket protocol implementation |

## Infrastructure

- **Docker**: Multi-stage build (Node.js -> Go -> Alpine)
- **Config**: Environment variables (`PORT`, `ALLOWED_ORIGINS`, `SESSION_TTL`)
- **Storage**: Fully in-memory (no database)
- **Concurrency**: `sync.RWMutex` + channels for thread-safe operations
