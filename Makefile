run-db:
	docker compose down && docker compose up -d
run-app:
	go run main.go