run:
	@npx tailwindcss -m tailwind.css -o ./internal/public/styles.css
	@templ generate
	@go run .
