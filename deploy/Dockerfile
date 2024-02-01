# This Dockerfile is only used by CI/CD Pipeline system
FROM debian:bookworm-slim AS runtime

RUN apt-get update \
  && apt-get install -y ca-certificates \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

ARG TARGETOS
ARG TARGETARCH

ARG TARGET_APP
ARG TARGET_PORT

WORKDIR /${TARGET_APP}

# k8s 환경에서 환경변수 설정은 .env 파일이 아닌 configmap & secret 에서 가져옴.
COPY ./target/bin/${TARGET_APP}.$TARGETOS.$TARGETARCH ./${TARGET_APP}

# Add configuration
COPY ./config/*.yaml ./config/

EXPOSE ${TARGET_PORT}

ENV TARGET_APP ${TARGET_APP}
CMD ./${TARGET_APP}
