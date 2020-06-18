FROM golang:1.13 as builder

WORKDIR /go/src/github.com/stobias123/wavecmd
COPY . .
RUN go build .

FROM scratch
COPY --from=builder /go/src/github.com/stobias123/wavecmd/wavecmd .
ENTRYPOINT ["wavecmd"]

