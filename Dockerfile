FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
ENV APP_PORT=8080
EXPOSE 8080
COPY --from=build /out/api /api
COPY migrations /migrations
USER nonroot:nonroot
ENTRYPOINT ["/api"]
