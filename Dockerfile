FROM golang
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .


RUN go build .
VOLUME /database/todos.db
EXPOSE 8080
CMD ["./TO_DO_PROJECT"]