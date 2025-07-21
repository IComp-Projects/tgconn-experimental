build-prod:
	docker build -f docker/Dockerfile.prod . -t tgconn:latest
run-prod:
	docker run --env-file .env -p 5555:5555 tgconn:latest

# Includes AIR + Volume mounting for live reloading
build-dev:
	docker build -f docker/Dockerfile.dev .  -t tgconn:dev
run-dev:
	docker run --env-file .env -p 5555:5555 -v .:/app tgconn:dev

