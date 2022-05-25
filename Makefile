# Actions with environment
dev:
	docker-compose --env-file env/envfile -f env/dev/docker-compose.yaml up --remove-orphans

stop:
	docker-compose -f env/dev/docker-compose.yaml stop

