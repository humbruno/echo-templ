run:
	@npx tailwindcss -m tailwind.css -o public/styles.css
	@templ generate
	@go run .
