## products-api

- Minimal microservice implemented in Go 1.22
- Http service will start on port 8080 and provide twp endpoints
- End point '/' will provide a JSON object of 50 random products
- End point '/:id' will return the product with the corrorsponding id or 'not found'


## Docker
To build and run the microservice, use the following commands:
```
docker build -t products-api .
docker run -d -p 8080:8080 products-api
```


