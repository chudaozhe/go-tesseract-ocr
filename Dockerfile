FROM golang:1.16 AS build
WORKDIR /go/src/demo
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -o app main.go

#FROM build AS development
#RUN apt-get update \
#    && apt-get install -y git
#CMD ["go", "run", "main.go"]

FROM tesseractshadow/tesseract4re
EXPOSE 8000
COPY --from=build /go/src/demo/app /app
CMD ["/app"]