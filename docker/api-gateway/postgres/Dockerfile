FROM postgres:latest

ARG UID=1000
ARG APP_NAME=app

ENV UID=${UID:-1000} \
    APP_NAME=${APP_NAME:-1000}

RUN useradd -u ${UID} ${APP_NAME} &>/dev/null

USER ${UID}

EXPOSE 5432