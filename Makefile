APP_NAME=shopping-cart-service

PROTO_DIR=proto
GEN_DIR=gen

.PHONY: proto run

proto:
	protoc -I $(PROTO_DIR) \
	  --go_out=$(GEN_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
	  $(PROTO_DIR)/*.proto
gen:
	sh gen.sh
	
run:
	go run ./cmd/server