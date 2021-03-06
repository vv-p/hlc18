FROM golang:latest

WORKDIR /go/src/github.com/vv-p/hlc18/
COPY main.go index_id.go index_sex.go http.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

EXPOSE 80

WORKDIR /root/
COPY --from=0 /go/src/github.com/vv-p/hlc18/app .
COPY run.sh .
CMD ["./run.sh"]
