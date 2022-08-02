FROM golang
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .


RUN go build .
VOLUME [ "/database" ]
EXPOSE 8080
CMD ["./to_do_project"]