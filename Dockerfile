FROM golang:1.22.1-alpine3.19 AS build-stage
ADD . /app/
WORKDIR /app/
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate 
RUN go build -o bin/main cmd/goshop/main.go

FROM build-stage AS run-stage
WORKDIR /app/
EXPOSE 8080
ENTRYPOINT ["./bin/main"]

