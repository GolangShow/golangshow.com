build:
	hugo --theme=golangshow

watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch --theme=golangshow

init:
	brew update
	brew install hugo
	gsutil config

upload_gc: build
	gsutil -m rsync -d -r public/ gs://golangshow.com

upload: build
	rsync -avz -e ssh public/ root@golangshow.com:golangshow-data/www