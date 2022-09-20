# Base Go commands.
GO_CMD     := go
GO_FMT     := $(GO_CMD) fmt
GO_GET     := $(GO_CMD) get
GO_INSTALL := $(GO_CMD) install
GO_BUILD   := $(GO_CMD) build
GO_RUN     := $(GO_CMD) run


BINARY_NAME := main

.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

.PHONY: clean
clean:
	@$(GO_CLEAN) .

.PHONY: build
build: clean fmt
	$(GO_BUILD) -o $(BINARY_NAME) -v .

# Build and migrate database
.PHONY: migrate
migrate: build
	./$(BINARY_NAME) -migrate=migrate


# Build and run the binary.
.PHONY: run
run: build
	./$(BINARY_NAME) run