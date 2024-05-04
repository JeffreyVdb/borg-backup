FROM --platform=$BUILDPLATFORM public.ecr.aws/docker/library/golang:latest as builder

ARG TARGETARCH
ARG TARGETOS
WORKDIR /src
RUN --mount=type=bind,source=.,target=/src,Z \
    env GOARCH=${TARGETARCH} GOOS=${TARGETOS} go build -o /tmp/create-borg-backup ./cmd/create-borg-backup/

FROM public.ecr.aws/ubuntu/ubuntu:latest

RUN set -xe \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get install -qq borgbackup rclone

COPY --from=builder /tmp/create-borg-backup /bin/
ENTRYPOINT [ "/bin/create-borg-backup" ]
