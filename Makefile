start-db:
	echo 'db is starting'
	docker compose -f docker-compose.yaml up -d
	echo 'db is started'

stop-db:
	docker compose down
	echo 'db is stopped'