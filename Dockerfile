FROM golang:1.22.2-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o relactions .

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build /app/relactions .

ENTRYPOINT ["/app/relactions"]
CMD [ "jira", "-h" ]