BIN_DIR := bin

.PHONY: all build add hourly-haze cumulative-discount clean

all: build

build: add hourly-haze cumulative-discount

add:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/add ./add
	cp add/add.json $(BIN_DIR)/add.json

hourly-haze:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/hourly-haze ./hourly-haze
	cp hourly-haze/haze-hourly.json $(BIN_DIR)/haze-hourly.json

cumulative-discount:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/cumulative-discount ./cumulative-discount
	cp cumulative-discount/cumulative-discount.json $(BIN_DIR)/cumulative-discount.json

clean:
	rm -rf $(BIN_DIR)
