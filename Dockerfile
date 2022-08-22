FROM golang:1.18-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
ENV GOPROXY=direct
RUN go mod download

COPY *.go ./

RUN go build -o /gcp-stale-resource-finder

CMD [ "/gcp-stale-resource-finder" ]

## Over-engineer with multi stage and just copy the binary over?
