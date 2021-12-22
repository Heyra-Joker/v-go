default: help

include script/make-rules/tools.mk

BIN_NAME=v-go

# Usage, 这样定义之后以Tab开头的都是变量了 :-)
define USAGE_OPTIONS

Options:
	GOOS Target Operating System, more information: https://go.dev/doc/install/source#environment
	GOARCH Target Platform
	UPX Ultimate Packer for eXecutables, more information: https://upx.github.io/
endef

# 导出环境变量,让任意 mk 文件均可使用
export USAGE_OPTIONS

## checkDir: Make ./bin dir if not exists
.PHONY: checkDir
checkDir:
	@mkdir -p ./bin

## verifyEnv: Verification your environment
.PHONY: verifyEnv
verifyEnv: checkDir tools.verify.ffmpeg

## build: Only build binary file from main.go
.PHONY: build
build: verifyEnv
	@env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/${BIN_NAME} -ldflags '-w' main.go
ifneq ($(origin UPX), undefined)
	@upx ./bin/${BIN_NAME}
endif

## clean: Remove all
.PHONY: clean
clean:
	@-rm -rf ./bin

## help: Show help info
.PHONY: help
help: Makefile
	@echo  "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"