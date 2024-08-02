default:
    just --list

run:
   go run ./cmd/server/main.go

run-web:
    cd ./web && npm run dev
