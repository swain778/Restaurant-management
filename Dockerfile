FROM golang:1.25.0-alpine3.22

# Set environment variables (optional, if not used in app logic)
ENV POSTGRES_USER=postgres \
    POSTGRES_PASSWORD=swain@123

# Set working directory
WORKDIR /restaurant

# Copy everything into the container
COPY . .

# Build the Go app
RUN go build -o restaurant-app ./main.go

# Run the binary
ENTRYPOINT ["./restaurant-app"]
