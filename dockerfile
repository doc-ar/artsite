FROM golang:1.24.4 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# Install tools
RUN curl -L -o /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && chmod +x /usr/local/bin/tailwindcss
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

# Produce Binary
RUN tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
RUN templ generate
RUN sqlc --file ./internal/db/config/sqlc.yaml generate
RUN CGO_ENABLED=0 go build -o ./app ./cmd

FROM scratch AS production
COPY --from=builder /build/app /
CMD [ "/app" ]
