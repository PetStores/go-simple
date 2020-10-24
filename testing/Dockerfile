FROM golangci/golangci-lint:v1.31-alpine

RUN mkdir -p /petstore
ADD ./app /petstore
WORKDIR /petstore

RUN golangci-lint run --issues-exit-code=1 --deadline=600s ./...

RUN go test -v ./...