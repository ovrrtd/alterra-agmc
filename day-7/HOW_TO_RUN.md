#HOW TO RUN

At First Please go to app/ directory.

## Run the App
You can run this app by simply use : `make run`, and it will automatically build and run the binary
Or you can run with docker with `make docker-build` and then `make docker-api`.

## Migration
You can run migration using `make migrate` or if you prefer docker you can use `make docker-migrate-up`.

# Pull image
Optionally, you can pull this application image from docker hub, you can pull with `docker pull ovrrtd/agmc:latest`. 
With this image you can run the as already explained above how to run this app.