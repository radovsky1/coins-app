FROM golang:1.19-buster AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

# build go app
RUN go mod download; CGO_ENABLED=0 go build -o /coins-app ./cmd/api/main.go


FROM alpine:latest

# copy go app, config and wait-for-postgres.sh
COPY --from=build /coins-app /coins-app
COPY ./configs/ /configs/
COPY ./wait-for-postgres.sh ./

# install psql and make wait-for-postgres.sh executable
RUN apk --no-cache add postgresql-client && chmod +x wait-for-postgres.sh

CMD ["/coins-app"]