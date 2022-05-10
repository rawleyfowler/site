FROM --platform=${BUILDPLATFORM} golang:1.18-bullseye AS build
WORKDIR .
ENV CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/rawleydotxyz .

FROM scratch AS bin
COPY --from=build /bin/rawleydotxyz /
EXPOSE 8080/tcp
CMD ["/rawleydotxyz"]
