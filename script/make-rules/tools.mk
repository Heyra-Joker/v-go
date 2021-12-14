
tools.install.%:
	@echo "Install ========> $*"
	@$(MAKE) install.$*

tools.verify.%:
	@if ! which $* &> /dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.ffmpeg
install.ffmpeg:
ifneq ("$(shell whoami)", "root")
	@echo "Current: $(shell whoami), You must be root."
else
	@chmod +x script/tools/install.ffmpeg.sh
	@script/tools/install.ffmpeg.sh --help
endif