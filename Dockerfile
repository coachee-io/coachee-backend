FROM pmdcosta/golang:1.13 AS builder
WORKDIR /code

# Add go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Add code and compile it
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app ./cmd/coachee

# Final image
FROM alpine:3.11
COPY --from=builder /app ./
COPY web ./web
ENTRYPOINT ["./app"]