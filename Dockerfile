FROM golang:1.23.3-alpine AS build_stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/hw8 ./cmd

FROM alpine AS run_stage

WORKDIR /bin

COPY --from=build_stage /bin/hw8 .

RUN chmod +x ./hw8

#EXPOSE 8000 8081
#
#ENTRYPOINT ["./hw8"]

EXPOSE 8000 8081

CMD [ "hw8" ]
