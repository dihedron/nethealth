.PHONY: binary
binary: 
#	@goreleaser build --single-target --snapshot --rm-dist
	@goreleaser build --snapshot --rm-dist
	@cp dist/nethealth_linux_amd64_v3/nethealth .
	@sudo setcap cap_net_raw=+ep ./nethealth

.PHONY: clean
clean:
	@rm -rf graphstack dist/

.PHONY: run
run: binary
	@NETHEALTH_LOG_LEVEL=warn ./nethealth start