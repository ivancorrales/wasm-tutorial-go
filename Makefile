PROJECT:=github.com/ivancorrales/wasm-pos-tagging-service
DOCKER_IMG:=docker.pkg.$(PROJECT)

build:
	docker build -t $(DOCKER_IMG) .
run:
	docker run -p 3000:3000 -it $(DOCKER_IMG)