build:
	hugo

watch:
	open http://127.0.0.1:1313
	hugo server --buildDrafts --watch

init:
	brew update
	brew install hugo
	git clone --recursive https://github.com/spf13/hugoThemes themes
