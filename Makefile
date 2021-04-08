
.PHONY: run-infra
run-infra:
	docker-compose up --build

.PHONY: curl-hello
curl-hello:
	curl localhost:3001/hello