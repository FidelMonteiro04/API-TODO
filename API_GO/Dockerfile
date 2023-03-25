FROM golang:latest

WORKDIR /app

COPY . /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq
RUN go get -u github.com/golang-migrate/migrate/v4/cmd/migrate

EXPOSE 9000


ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=user_todo
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=api_todo

RUN go build -o main .

CMD ["./main", "migrate"]

CMD ["./main", "serve"]
