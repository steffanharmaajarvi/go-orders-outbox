start:
	cp ./configs/.env.yml.dist ./configs/.env.yml
	make build
build:
	docker-compose up -d --build
up:
	docker-compose up -d