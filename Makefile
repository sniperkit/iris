.PHONY: update-fork register-upstream pull-upstream

update-fork: pull-upstream

register-upstream:
	@git remote add upstream git://github.com/kataras/iris
	@git fetch upstream

pull-upstream:
	@git pull upstream master
	@git add .
	@git commit -am "Updating fork with original repo to keep up with their changes"
	@git push