FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY . /app

# build the go app with the name main
RUN go build -o main .

ENV JWT_SECRET_KEY=$JWT_SECRET_KEY

ENV MONGO_URI=$MONGO_URI

ENV DB_NAME=$DB_NAME

EXPOSE 8080

CMD ["./main"]