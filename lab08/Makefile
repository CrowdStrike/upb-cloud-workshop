dockerize:
	docker build -t upb-postgress-db ./
run-postgress:
	docker run -p 5432:5432 --name pg-upb upb-postgress-db
dockerize-clean:
	docker stop pg-upb
	docker rm pg-upb
	docker rmi upb-postgress-db

TARGETS=dockerize run-postgress


all:$(TARGETS) 
clean: dockerize-clean

