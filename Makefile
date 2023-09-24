.PHONY: all build run clean

all: build

build:
	docker build -t my_app_image .

run:
	docker-compose up -d

clean:
	docker-compose down -v --remove-orphans
