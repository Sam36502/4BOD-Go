
BIN = 4bod
BIN_DIR = bin
LIN_DIR = $(BIN_DIR)/lin
WIN_DIR = $(BIN_DIR)/win
SRC_DIR = src

build: build-lin build-win

build-lin:
	@echo '---> Building Linux binary...'
	GOOS=linux go build -o $(LIN_DIR)/$(BIN) $(SRC_DIR)/*.go

build-win:
	@echo '---> Building Windows binary...'
	GOOS=windows go build -o $(WIN_DIR)/$(BIN) $(SRC_DIR)/*.go

run-lin: build-lin
	@echo '---> Running...'
	@echo ''
	@$(LIN_DIR)/$(BIN)

run-win: build-win
	@echo '---> Running...'
	@echo ''
	@$(WIN_DIR)/$(BIN)
