ASSETS_DIR ?= $(shell pwd)/assets

test:
	mkdir -p cover
	go test -coverprofile=cover/coverage.out ./...
	go tool cover -html=cover/coverage.out

check:
	@! gometalinter \
		--vendor \
		./...

build:
	docker build -t phoneval .

grun:
	go run main.go -a assets/area_codes.txt -n assets/phone_numbers.txt
run:
 	 run -v $(ASSETS_DIR):/assets phoneval -a assets/area_codes.txt -n assets/phone_numbers.txt
