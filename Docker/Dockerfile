#FROM golang:1.18.4-alpine
#WORKDIR /go/household/
#COPY src ./src
#COPY go.mod go.sum ./
#RUN apk update && apk --no-cache add git
#RUN go mod tidy && go install github.com/cosmtrek/air@v1.29.0
#WORKDIR /go/household/src
#CMD ["go", "run", "main.go"]
#EXPOSE 80

#FROM golang:1.18.4-alpine
#WORKDIR /go/household/
#COPY ../src ./src
#COPY ../go.mod go.sum ./
#RUN apk update && apk --no-cache add git
#RUN go mod tidy
#WORKDIR /go/household/src
#RUN go build -o app
#CMD ["go","run","main.go"]
#EXPOSE 80

 #builder
#FROM golang:1.16.4-alpine as builder
#WORKDIR /go/household/
#COPY src ./src
#COPY go.mod go.sum ./
#RUN apk update && apk --no-cache add git
#RUN go mod tidy
#WORKDIR /go/household/src
#RUN CGO_ENABLE=0 GOOS=linux go build -o /go/household/binary
#
#
## ##production
#FROM alpine as production
#WORKDIR go/household/production
#RUN apk add --no-cache ca-certificates
#COPY --from=builder /go/household/binary /go/household/production
#ENV PORT=${PORT}
#CMD ["/go/household/production/binary"]

FROM golang:1.18.4-alpine
WORKDIR /go/household/
COPY tools ./tools
RUN apk update && apk --no-cache add git
RUN go install github.com/rubenv/sql-migrate/...@latest
WORKDIR /go/household/tools
CMD ["sql-migrate", "up"]
EXPOSE 80