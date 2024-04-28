run:
	@npx tailwindcss -m tailwind.css -o api/public/styles.css
	@templ generate
	@go run .
