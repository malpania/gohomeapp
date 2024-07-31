FROM golang:1.22-alpine as dev

WORKDIR /work

FROM golang:1.22-alpine as build

WORKDIR /app
COPY / /app/
RUN ls -al /app/*
RUN go build -o app github.com/malpania/beerproj/cmd/web

FROM alpine as runtime
COPY --from=build /app/templates /templates
COPY --from=build /app/ui /ui
COPY --from=build /app/app /

CMD ./app