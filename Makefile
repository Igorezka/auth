include .env

LOCAL_BIN:=$(CURDIR)/bin
APP_NAME:=auth-service
VERSION:=0.0.2

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.24.3

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable" status -v

migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable" up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable" down -v

check-lint-config:
	$(LOCAL_BIN)/golangci-lint config verify --config .golangci.pipeline.yaml

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

test:
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/igorezka/auth/internal/service/...,github.comgithub.com/igorezka/auth/internal/api/... -count 5

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/igorezka/auth/internal/service/...,github.com/igorezka/auth/internal/api/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out -o=coverage.html;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore
	grep -sqFx "/coverage.html" .gitignore || echo "/coverage.html" >> .gitignore

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/app/auth-service cmd/grpc_server/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t $(REGISTRY_HOST)/$(REGISTRY_NAME)/$(APP_NAME):$(VERSION) -f ./server.Dockerfile . --provenance=false

docker-push:
	docker login $(REGISTRY_HOST)
	docker push $(REGISTRY_HOST)/$(REGISTRY_NAME)/$(APP_NAME):$(VERSION)