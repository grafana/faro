serve:
	docker run -p 80:8080 -e SWAGGER_JSON=/src/openapi.yaml -v ./:/src swaggerapi/swagger-ui
