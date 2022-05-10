# syntax = docker/dockerfile:1

FROM --platform=${BUILDPLATFORM} golang:1.18-bullseye AS build
WORKDIR .
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/rawleydotxyz .

FROM scratch AS bin
COPY --from=build /bin/rawleydotxyz /

FROM golang:1.18-bullseye AS run
EXPOSE 8080/tcp
CMD ["/rawleydotxyz"]
