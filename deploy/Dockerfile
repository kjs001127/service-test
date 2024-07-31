# This Dockerfile is only used by CI/CD Pipeline system
FROM public.ecr.aws/i8a4b9p4/circleci-base/debian:bookworm-slim-curl AS runtime

RUN apt-get update \
  && apt-get install -y ca-certificates \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

ARG TARGETOS
ARG TARGETARCH

WORKDIR /ch-app-store

# k8s 환경에서 환경변수 설정은 .env 파일이 아닌 configmap & secret 에서 가져옴.
COPY ./target/bin/ch-app-store-general.$TARGETOS.$TARGETARCH ./general
COPY ./target/bin/ch-app-store-admin.$TARGETOS.$TARGETARCH ./admin
COPY ./deploy/run.sh ./run.sh

# Add configuration
COPY ./config/*.yaml ./config/
# Add i18n
COPY ./i18n/ ./i18n/

EXPOSE 3020
EXPOSE 3021

CMD ["./run.sh"]
