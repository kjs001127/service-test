# This Dockerfile is only used by CI/CD Pipeline system
FROM public.ecr.aws/i8a4b9p4/circleci-base/debian:bookworm-slim-curl AS runtime

RUN apt-get update \
  && apt-get install -y ca-certificates \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

ARG TARGETOS
ARG TARGETARCH

ARG TARGET_APP
ARG TARGET_PUBLIC_PORT
ARG TARGET_ADMIN_PORT

WORKDIR /$TARGET_APP

# k8s 환경에서 환경변수 설정은 .env 파일이 아닌 configmap & secret 에서 가져옴.
COPY ./target/bin/$TARGET_APP-general.$TARGETOS.$TARGETARCH ./general
COPY ./target/bin/$TARGET_APP-admin.$TARGETOS.$TARGETARCH ./admin
COPY ./deploy/run.sh ./run.sh

# Add configuration
COPY ./config/*.yaml ./config/

# front, genenral, desk 모두 하나의 포트만 사용 할 것이므로
EXPOSE $TARGET_PUBLIC_PORT
EXPOSE $TARGET_ADMIN_PORT

CMD ["./run.sh"]
