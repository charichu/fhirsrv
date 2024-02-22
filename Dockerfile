# use official Golang image
FROM golang:1.22

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o fhirsrv .

#EXPOSE the port
EXPOSE 8080

# Run the executable
CMD ["./fhirsrv"]