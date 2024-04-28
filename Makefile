run:
	@templ generate
	@go run .

.PHONY: tailwind-watch
css:
	./tailwindcss -i static/css/input.css -o public/styles.css
