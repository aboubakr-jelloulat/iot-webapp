FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY backend/ ./backend/
COPY frontend/ ./frontend/
RUN CGO_ENABLED=0 go build -trimpath -o /webapp ./backend

FROM scratch
COPY --from=build /webapp /webapp
USER 10001:10001
EXPOSE 8080
ENTRYPOINT ["/webapp"]
