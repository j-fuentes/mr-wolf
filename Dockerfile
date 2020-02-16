# STAGE 1
FROM golang:1.13.4 as builder

WORKDIR /go/j-fuentes/mr-wolf

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .
RUN make install

# STAGE 2
FROM gcr.io/distroless/base:nonroot
COPY --from=builder /go/bin/mr-wolf /bin/mr-wolf
ENTRYPOINT ["mr-wolf"]
