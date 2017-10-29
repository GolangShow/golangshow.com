.PHONY: html
html:
	rm -fr public
	hugo --theme=golangshow

.PHONY: whitespace
whitespace:
	go run tools/strip-whitespace/main.go

.PHONY: build
build: html whitespace
	$(info -------------------------------------------------------------------------------------)
	$(info Yeah! The code is ready for upload, please don't forget to check, commit and push it.)
	$(info After that, you can run `make upload` to publish the site.)

.PHONY: watch
watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch --theme=golangshow

.PHONY: init
init:
	brew update
	brew install hugo

.PHONY: upload
upload:
	rsync -avz -e ssh public/ root@golangshow.com:golangshow-data/www
