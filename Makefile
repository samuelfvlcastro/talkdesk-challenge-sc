ASSETS_DIR ?= $(shell pwd)/assets

build:
	docker build -t phoneval .

run:
 	 run -v $(ASSETS_DIR):/assets phoneval -a assets/area_codes.txt -n assets/phone_numbers.txt

