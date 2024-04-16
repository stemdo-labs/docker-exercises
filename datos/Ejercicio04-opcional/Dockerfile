# build and run golang app

FROM golang:latest as builder
WORKDIR /app
COPY ./app .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM registry.access.redhat.com/ubi8-minimal:latest as runner
ARG USER_NAME=dummy
ARG USER_UID=1001
ARG USER_SHELL=/bin/bash
RUN microdnf -y distro-sync && microdnf -y install shadow-utils && \
    adduser -U -u "${USER_UID}" -s "${USER_SHELL}" -M "${USER_NAME}"  && \
    microdnf -y remove shadow-utils && \
    microdnf clean all
WORKDIR /app
COPY --chmod=555 --chown=${USER_NAME}:${USER_NAME} --from=builder /app/main .
EXPOSE 8080
USER ${USER_UID}
CMD ["./main"]