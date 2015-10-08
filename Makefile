build:
	hugo --theme=golangshow

watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch --theme=golangshow

init:
	brew update
	brew install hugo

upload: build
	rsync -avz -e ssh public/ root@golangshow.com:golangshow-data/www
