

# Actions with environment
dev:
	docker-compose --env-file env/envfile -f env/dev/docker-compose.yaml up

prod:
	docker-compose --env-file /env/envfile -f env/prod/docker-compose.yaml up -d

stop:
	docker-compose -f env/dev/docker-compose.yaml stop
	docker-compose -f env/prod/docker-compose.yaml stop

# Actions with db
run_db:
	docker exec -it wave-db psql -U postgres -d wave

dump_db:
	docker exec wave-db pg_dump -U postgres -d wave > db_dump/wave.sql

restore_db:
	docker exec wave-db bash -c 'cd tmp && psql -U postgres -c "drop database wave with (FORCE)" && psql -U postgres -c "create database wave" && psql -U postgres -c "grant all privileges on database wave to postgres" && psql -U postgres -d wave -1 -f wave.sql'

