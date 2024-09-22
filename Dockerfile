FROM golang:1.23-alpine AS base

RUN apk add --no-cache make curl

WORKDIR /app

COPY . .
RUN go mod download

FROM base AS dev
RUN make build
EXPOSE ${PORT}
CMD [ "sh", "-c", "echo 'y' | make watch" ]

FROM base AS build
RUN make build

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
COPY --from=build /app/ui /app/ui
EXPOSE ${PORT}
CMD ["./main"]
