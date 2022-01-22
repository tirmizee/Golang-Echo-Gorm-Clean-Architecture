start-data:
	echo 'db is starting'
	docker compose -f docker-compose.yaml up -d
	echo 'db is started'

stop-data:
	docker compose down
	echo 'db is stopped'

start-app:
	go run .