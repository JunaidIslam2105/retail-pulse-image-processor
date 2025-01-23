
FROM golang:1.22.3-alpine AS build

]
WORKDIR /app


COPY go.mod go.sum ./
RUN go mod tidy


COPY . .


WORKDIR /app/cmd


RUN go build -o main .

# Expose the port that the app will run on
EXPOSE 8080

CMD ["./main"]
