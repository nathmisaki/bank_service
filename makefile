default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t api .

up: default
	@echo "=============starting api locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	docker-compose run api go test -v -cover ./...

test_local:
	godotenv go test -v -cover ./...

test_local_debug:
	DEBUG_TEST=true godotenv go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f