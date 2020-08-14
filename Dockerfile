FROM golang:1.14.7-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apk add git --no-cache
RUN go get github.com/gorilla/mux
RUN go get github.com/mdlayher/wol
RUN go get github.com/sparrc/go-ping

COPY "./webapi" /appbuild
WORKDIR /appbuild
RUN go build -o output/bin/app

FROM scratch
COPY --from=builder /appbuild/output/bin/app /app
COPY "./webapi/static" /static
CMD ["/app"]
