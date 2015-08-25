build:
	hugo --theme=golangshow

build-old:
	hugo --theme=golangshow-old --destination=public-old

watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch --theme=golangshow

init:
	brew update
	brew install hugo
	gsutil config

upload: build
	gsutil -m -h "Cache-Control:public, max-age=600" cp -a public-read -r public/* gs://golangshow.com
