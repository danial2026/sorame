FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built Go binary into the image
COPY ./main .

# Copy the static files
COPY ./static ./static

# Expose port to the outside world
EXPOSE $SORAME_SERVICE_PORT

# Command to run the executable
CMD ["./main"]
