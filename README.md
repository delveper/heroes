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
│   └── user.go
├── services/
│   ├── keeper/
│   │   ├── sql/
│   │   │   └── migrations.sql
│   │   ├── keeper.go
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
