BIN_NAME=v-go

.PHONY: checkDir
checkDir:
	@mkdir -p ./bin

.PHONY: clean
clean:
	@ rm -rf ./bin

.PHONY: run
run: checkDir
	@go build -o ./bin/${BIN_NAME} ./main.go


