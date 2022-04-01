FROM golang AS build
ENV LOCATION /go/src/mission-generator
ENV APP_NAME "grpc-mission-generator-service"

WORKDIR ${LOCATION}/

ADD ./ ${LOCATION}/

RUN go mod download -x

RUN CGO_ENABLED=0 go build -o /bin/${APP_NAME} ./cmd/main.go


FROM scratch
COPY --from=build /bin/${APP_NAME} /bin/${APP_NAME}
ENTRYPOINT ["/bin/grpc-mission-generator-service"]
EXPOSE ${APP_PORT}