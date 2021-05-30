build:
	DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose -f development.docker-compose.yml build --progress plain --no-cache
up:
	docker-compose -f development.docker-compose.yml up
down:
	docker-compose -f development.docker-compose.yml down
bash:
	docker-compose -f development.docker-compose.yml run -p 8000:8000 --rm app bash