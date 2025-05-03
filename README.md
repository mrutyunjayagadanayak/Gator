Gator
Gator is a backend service developed in Go, utilizing SQLC for type-safe database interactions. It is designed to provide a robust and efficient foundation for applications requiring structured database access.

Features
Go Backend: Leverages Go's concurrency and performance features.

SQLC Integration: Ensures type-safe and efficient database queries.

Modular Structure: Organized codebase for scalability and maintainability.

Getting Started
Prerequisites
Go 1.16 or higher

SQLC installed

A compatible SQL database (e.g., PostgreSQL)

Installation
Clone the repository:
git clone https://github.com/mrutyunjayagadanayak/Gator.git

Install dependencies:
go mod tidy

Generate Go code from SQL:
sqlc generate

CLI Commands
Gator operates via a command-line interface. You invoke commands as:
go run main.go <command> [args...]
