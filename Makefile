.PHONY: run

run:
	swag init -g main.go --output docs
	go run .

dev:
	swag init -g main.go --output docs
	air # или "go run ."

watch:
	swag init -g main.go --output docs
	reflex -r '\.go$$' -s -- sh -c 'swag init -g main.go --output docs && go run .'