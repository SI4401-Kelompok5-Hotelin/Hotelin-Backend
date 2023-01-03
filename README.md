# Hotelin-BE

Hotelin-BE is a backend service for Hotelin, a web application for managing and booking hotels.

## Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Installation

- Clone the repository

```bash
$ git clone https://github.com/SI4401-Kelompok5-Hotelin/Hotelin-BE.git
```

- Install dependencies

```bash
$ go mod download
```

- Create a `.env` file in the root directory of the project and copy the contents of `.env.example` into it

- Run the application

```bash
$ docker-compose up
$ go run main.go
```

Running the application will create a database named `hotelin` in your local PostgreSQL instance.

## Contributing

TL;DR: Please read the [Contributing Guide](CONTRIBUTING.md) before contributing.

## Regards

- Abdi Fatih
