# Create a stage for building the application.
ARG GO_VERSION=1.22.1
FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /app  

# Copy Go module files and dependencies 
COPY go.mod . 
COPY go.sum . 
RUN go mod download 

# Copy the rest of the source code 
COPY . /app 

# install templ 
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate 

# Build the Go app
RUN CGO_ENABLED=0 go build -o run-app 

# Deploy 
FROM alpine:latest
WORKDIR / 
COPY --from=build /app / 
EXPOSE 4500

# What the container should run when it is started.
CMD [ "/run-app" ]
