git:
	git add .
	git commit -m "$(msg)"
	git push origin master

fmt:
	@echo "formatting code"
	go fmt ./...

lint:
	@echo "Linting source code"
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...

graphql:
	go run github.com/99designs/gqlgen generate

run:
	 go run cmd/main.go