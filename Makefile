build:
	rm -fr public
	hugo --theme=golangshow

watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch --theme=golangshow

init:
	brew update
	brew install hugo

upload: build
	git push
	rsync -avz -e ssh public/ root@golangshow.com:golangshow-data/www

.PHONY: whitespace
whitespace:
	rm -rf $(GOPATH)/src/github.com/AlekSi/trim-whitespace
	mkdir $(GOPATH)/src/github.com/AlekSi/trim-whitespace
	git clone git@gist.github.com:e98f74b1fa3edd0a496b5a0cb514ff0a.git $(GOPATH)/src/github.com/AlekSi/trim-whitespace
	cd $(GOPATH)/src/github.com/AlekSi/trim-whitespace && go get
	trim-whitespace
