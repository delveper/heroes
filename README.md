# Octagon 

## Arch
```
octagon/
├── api/
│   └── openapi.yml
├── cfg/
│   └── config.go
├── cmd/
│   └── srv/
│       └── main.go 
├── core/
│   ├── error.go 
│   ├── user.go 
│   └── validation.go
├── services/
│   ├─ agent/
│   │   └── user.go
│   ├── mover/
│   │   ├── mover.go
│   │   ├── response.go
│   │   ├── middleware.go
│   │   └── user.go
│   ├── repo/
│   │   ├── sql/
│   │   │   └── migrations.sql
│   │   ├── repo.go
│   │   ├── migrate.go
│   │   └── user.go
│   ├── mover/
│   │   ├── mover.go
│   │   ├── response.go
│   │   ├── middleware.go
│   │   └── user.go
│   └── promoter/
│       └── promoter.go
├── pkg/
├── README.md
└── .env 
```
