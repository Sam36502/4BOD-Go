
BIN = 4bod
BIN_DIR = bin/lin
SRC_DIR = src

build: build-lin build-win
	@echo '---> Building Linux binary...'
	GOOS=linux go build -o $(BIN_DIR)/$(BIN) $(SRC_DIR)/*.go

run: build
	@echo '---> Running...'
	@echo ''
	@$(BIN_DIR)/$(BIN)