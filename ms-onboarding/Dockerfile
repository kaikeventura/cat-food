FROM golang
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-gs-ping

EXPOSE 5000

CMD [ "/docker-gs-ping" ]