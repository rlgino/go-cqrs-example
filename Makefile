# Esto busca el step default
.DEFAULT_GOAL := build

# This is a rule
# target: dependency1 dependency2
#(tab) --- shell lines
# El deps hace el download de dependencies
deps: 
	@go mod tidy

# El @ evita mostrar el comando a ejecutar
run: test
	@go run ./src/api/main.go

# Elige el step por delante de la carpeta
# El estandar es siempre arriba del que necesitemos.
.PHONY: build
build:
	@go build ./...
	@echo "Build Done!"

test:
	@go test ./...
	@echo "Tested ok"

build-inside:
	@cd build && make

# Tareas comunes:
# build (La primera y paso previo al deploy)
# install (Instalamos el build en el sistema y ejecuta el comando "build")
# deps (Instalando dependencias y son llamados por build e install)
# start/stop (Levanta y apaga la app, por ejemplo cuando se corre en docker)
# clean (Para limpiar cache y demás cosas)