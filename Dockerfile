FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o module-nine .

EXPOSE 8080

ENTRYPOINT [ "./module-nine" ]
