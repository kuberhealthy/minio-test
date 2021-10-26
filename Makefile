default: build push

build: 
	docker build -t kuberhealthy/minio-test:v1.0.0 .

push: 
	docker push kuberhealthy/minio-test:v1.0.0 
