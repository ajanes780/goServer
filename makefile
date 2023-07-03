
run:
	sass styles:static/css
	echo "CSS compiled"
	go build -o bin/$(APP_NAME) .
	echo "Binary compiled"
