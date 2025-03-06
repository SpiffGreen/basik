# Basik - Multi-Tenant SaaS Application

A scalable SaaS platform where companies can manage users and tasks built with Go and ReactJS.

## Features

- Multi-tenant architecture with PostgreSQL
- OAuth 2.0 + JWT authentication
- Role-based access control (RBAC)
- Prometheus/Grafana monitoring

## Prerequisites

- Go 1.19+
- PostgreSQL 14+
- Redis
- Node.js 16+

## Project Structure

```
/
├── frontend/    # React frontend
├── ./           # Go/Gin backend
└── k8s/         # Kubernetes configs
```

## Quick Start

1. Start backend:
```sh
go run main.go
```

2. Run migrations:
```sh
go run migrate/migrate.go
```

3. Start frontend:
```sh
cd frontend
npm install
npm start
```

## API Endpoints

- `POST /tenants` - Register company
- `POST /register` - Register user
- `POST /login` - Authenticate user
- `GET /tasks` - List tasks
- More endpoints documented in API docs

See Postman docs [here](./basik_postman.json)

## Docker

```sh
docker-compose up
```

## License

MIT