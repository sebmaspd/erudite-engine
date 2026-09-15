BIN_DIR := bin

.PHONY: all build add hourly-haze clean

all: build

build: add hourly-haze

add:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/add ./add
	cp add/add.json $(BIN_DIR)/add.json

hourly-haze:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/hourly-haze ./hourly-haze
	cp hourly-haze/haze-hourly.json $(BIN_DIR)/haze-hourly.json

clean:
	rm -rf $(BIN_DIR)
