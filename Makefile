build:
	rm -f dist/server
	rm  -rf dist/static
	mkdir -p ./dist/static/css
	mkdir -p ./dist/static/js
	cp -r assets/js ./dist/static
	go build -o ./dist/server cmd/server/server.go
	tailwind -i assets/css/input.css -o dist/static/css/main.css --minify

generate:
	find . -type f -name "*.templ" | entr templ generate

run:build
	cd dist && ./server

watch:
	find . -type f -name "*.go" | entr -r make run 