ARG GO_VERSION=1.24.1
FROM golang:${GO_VERSION}-bookworm AS builder

RUN apt-get update && apt-get install -y nodejs npm

WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN npm ci
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN npx @tailwindcss/cli -i ./templates/static/app.css -o ./templates/static/final.css
RUN go build -v -o /run-app .

FROM debian:bookworm

COPY --from=builder /run-app .
# COPY --from=builder /usr/src/app/templates ./templates
CMD ["./run-app"]
