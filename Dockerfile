FROM golang:alpine AS builder
RUN apk update && apk add --no-cache gcc musl-dev
RUN mkdir /work
WORKDIR /work
COPY . .
RUN cd src && go get -d -v
RUN cd src && go build -o /work/bibleserver

FROM scratch
COPY --from=builder /work/bibleserver/ /bin/bibleserver
COPY --from=builder /work/bibles /bibles

EXPOSE 8000

ENTRYPOINT ["/bin/bibleserver"]