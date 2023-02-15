FROM golang:alpine3.16 as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash openssh musl-dev make pkgconfig

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ARG GIT_SHA
ARG VERSION
ARG GOARCH

RUN make test
RUN GOOS=linux GOARCH=${GOARCH} go build -mod=readonly -ldflags="-w -s -X 'main.version=${VERSION}' -X 'main.commit=${GIT_SHA}' -X 'main.date=$(date)'" -o /go/bin/app ./cmd/srv/

FROM golang:alpine3.16
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN true
COPY --from=builder /go/bin/app /go/bin/app

RUN addgroup -S app && adduser -S -g -D -H app app
USER app

ENV TEMPLATE_DIR="/template/"
COPY --from=builder /go/src/app/template/* ${TEMPLATE_DIR}
ENTRYPOINT ["/go/bin/app"]