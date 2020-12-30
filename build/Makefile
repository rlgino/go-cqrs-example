deps-install: CMD=install
deps-update: CMD=update
# make deps-require package=zaraza
deps-require: CMD=require $(package)

deps-install deps-update deps-require:
	@echo "deps" $(CMD)