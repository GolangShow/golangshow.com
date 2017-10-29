build:
	rm -fr public
	hugo --theme=golangshow
	go run tools/strip-whitespace/main.go

watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch --theme=golangshow

init:
	brew update
	brew install hugo

upload:
	git diff --exit-code
	git push
	rsync -avz -e ssh public/ root@golangshow.com:golangshow-data/www
