pkg_rev := $(shell git log -1 --format=%h_%ct)

all: run

run: 
	go run *.go

freebsd: GOOS=freebsd
freebsd: release_production

linux: GOOS=linux
linux: release_production

release_production:
	@echo "👌 Compiling $(GOOS) binary..."
	@go build -o pp *.go

	@echo "📦 Packaging assets..."
	@tar cf pp_$(GOOS)_$(pkg_rev).tar.gz pp assets/

	@echo "🥳 All done!"
