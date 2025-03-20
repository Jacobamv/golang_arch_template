SHELL := /bin/bash
PROJECTNAME=$(shell basename "$(PWD)")
PID=/tmp/.$(PROJECTNAME)-api-server.pid

# Colors
NC=\033[0m
RED=\033[91m
GREEN=\033[92m

# Project dependencies
DEPENDENCIES = Golang="go version" \
			Redis="redis-cli -v" \
			Swagger="swagger version" \
			Postgres="psql -V"

# Build
# TODO: Implement loader
# TODO: Add check for installed dependencies
.PHONY: build
build: check-dependencies docs
	@echo "[+] Trying to build the application"
	@cd ./cmd && go build -o $(PWD)/build/app

	@printf "${GREEN}[●]${NC} The project has successfully passed the build! Don't forget to reload the service\n"

# TODO: Explain
check-dependencies: check-go-dependencies
	@echo "[+] Check dependencies versions:"

	@is_installed=1
	@check () {
		@$$(eval $$2 >/dev/null 2>&1)

		@if [ $$? -eq "0" ]; then
			@version=$$(eval $$2 | grep -oP "([0-9]{1,}\.[0-9]{1,}(\.[0-9]{1,})?)|(dev)" | head -1);
			@printf '%s \t: ${GREEN}%s${NC}\n' $$1 $$version | expand -t 15;
		else
			@is_installed=0
			@printf '%s \t: ${RED}%s${NC}\n' $$1 'Not installed' | expand -t 15;
		fi
		return 1
	}

	@IFS='='
	@read -a strarr <<< "$$DEPENDENCIES"
	@for key in $(DEPENDENCIES)
	do
		check $$key >/dev/null 2>&1;
	done

	@if [ $$is_installed -eq 0 ]; then
		for key in $(DEPENDENCIES)
		do
			check $$key;
		done
		exit 1
	fi

	@printf "\t%s\n" "All dependencies are installed" | expand -4;

check-go-dependencies:
	@echo "[+] Checking golang dependencies..."
	@go mod tidy
	@go mod download


# Documentation generation
.PHONY: docs
docs:
	@echo "[+] Generation swagger file"
	@swagger generate spec -m -o $(PWD)/pkg/docs/swagger.yaml

# Run project as dev
.PHONY: run
run:
	@cd ./cmd && go run main.go

.ONESHELL:
	pass
