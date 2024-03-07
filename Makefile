eal.PHONY: protoc

setup:


protoc:
	@echo "--- Preparing proto output directories ---"
	@echo "--- Compiling all proto files ---"
	@cd ./proto && protoc -I. --go_out=../generated/proto --go-grpc_out=require_unimplemented_servers=false:../generated/proto *.proto