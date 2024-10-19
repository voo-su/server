.PHONY: install
install:
	go install github.com/google/wire/cmd/wire@latest \
	&& go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
	&& go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	&& go install github.com/srikrsna/protoc-gen-gotag@latest \
	&& go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: proto
proto:
	protoc --proto_path=./api/proto \
		--go_out=paths=source_relative:./api/pb/ \
		--validate_out=paths=source_relative,lang=go:./api/pb/ ./api/proto/v1/*.proto
	protoc --proto_path=./api/proto \
		--gotag_out=outdir="./api/pb/":./ ./api/proto/v1/*.proto

.PHONY: http
http:
	go run ./cmd/voo-su http --c ./configs/voo-su.yaml

.PHONY: ws
ws:
	go run ./cmd/voo-su ws --c ./configs/voo-su.yaml

.PHONY: cli-cron
cli-cron:
	go run ./cmd/voo-su cli-cron --c ./configs/voo-su.yaml

.PHONY: cli-queue
cli-queue:
	go run ./cmd/voo-su cli-queue --c ./configs/voo-su.yaml

.PHONY: build
build:
	go build -o ./build/voo-su ./cmd/voo-su

