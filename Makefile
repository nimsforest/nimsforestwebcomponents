css:
	npx tailwindcss -i static/input.css -o static/style.css --minify

css-watch:
	npx tailwindcss -i static/input.css -o static/style.css --watch

.PHONY: css css-watch
