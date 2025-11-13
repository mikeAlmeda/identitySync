# RosterBridge

A full-stack student roster synchronization platform built with Go, MongoDB, React, and TypeScript with an emphasis on modern backend architecture and API-driven workflows for education technology.

## Overview

RosterBridge simulates the core workflow of platforms like Clever: fetching student roster data from external sources, normalizing records, and exposing them through a REST API and admin dashboard. Built as preparation for the Clever Software Engineer Apprenticeship, this project showcases repository patterns, idempotent sync operations, database migrations with indexes, unit testing with mocks, and type-safe frontend development.

## Tech Stack

**Backend:**
- Go 1.25.3
- Chi router (v5)
- MongoDB with mongo-driver
- godotenv for environment configuration

**Frontend:**
- Reach 19
- TypeScript 5.9
- Vite (dev server & build tool)
- Fetch API for HTTP requests

**Infrastructure:**
- Docker Compose with MongoDB container
- Local development environment with hot reload

## Quick Start

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop) - Must be running
- Go 1.25.3+ ([Download](https://go.dev/dl/))
- Node.js 18+ (LTS) ([Download](https://nodejs.org/))

### Clone & Setup

```sh
git clone https://github.com/mikeAlmeda/identitySync.git
cd identitySync
```

### Run Locally
1. Start MongoDB
```sh
cd rosterbridge-backend
docker compose up -d mongo
```

**2. Start Backend (Terminal 1)**
```sh
cd rosterbridge-backend
cp .env.example .env
go run ./cmd/server
```

**3. Start Frontend (Terminal 2)**
```sh
cd rosterbridge-ui
npm install
npm run dev

# You should see:
# VITE ... ready in ...ms
```

**4. Access the Application**
- Frontend UI: http://localhost:5173
- Backend API: http://localhost:8080/students

```sh
curl -X POST http://localhost:8080/sync
```

Click "Sync Students" in the UI to populate data.

## Project Structure

```
identitySync/
├── rosterbridge-backend/          # Go REST API
│   ├── cmd/
│   │   └── server/                # Application entry point
│   ├── internal/
│   │   ├── db/                    # MongoDB connection & migrations
│   │   ├── handlers/              # HTTP request handlers
│   │   ├── models/                # Student data model
│   │   ├── services/              # Business logic (sync service)
│   │   ├── store/                 # Repository interface & implementations
│   │   └── router/                # Route registration
│   ├── docker-compose.yml         # MongoDB container definition
│   └── .env.example               # Environment template
└── rosterbridge-ui/               # React TypeScript frontend
    └── src/
        ├── api/                   # API client with typed interfaces
        ├── components/            # React components (StudentList)
        ├── App.tsx                # Root component
        └── main.tsx               # Entry point
```

## Screenshots

![Student Directory UI](docs/demo-screenshot.png)

*Student roster dashboard after clicking "Sync Students"—displays 5 mock students with source IDs, names, grades, and timestamps.*

## Why I Built This

I created RosterBridge as a way to get hands-on experience with the same technologies used at Clever: Go, MongoDB, React, and TypeScript. My goal was to build something that feels real: a working system that connects and syncs student roster data like schools actually do. Along the way, I focused on writing clean, reliable backend code and building a simple, intuitive interface that brings everything together.