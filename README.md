## ch-proto

채널톡 백엔드 내부에서 활용되는 gRPC proto 관리를 위한 repository
submodule은 채널톡 전체 도메인에서 사용되는 모델을 기준으로 나눔

- `/meet`: 컨퍼런스콜 모델인 `Meet` 과 연관된 rpc
- `/sip`: 전화연동을 위한 프로토콜인 `sip` 를 처리하는 rpc

### Getting started

1. proto 빌드가 가능하도록 로컬에 세팅하기
   - [노션 가이드](https://www.notion.so/channelio/m1-ch-proto-5e09013c24df46219d28701c35ae22aa) 를 통해 `protobuf`와 `grpc-java` 컴파일러 설치

2. `grpc-java` 의 빌드된 executable 을 `GRPC_JAVA_PATH` 의 환경변수로 등록
  ``` shell
# example
export GRPC_JAVA_PATH = /Users/max/desktop/grpc-java/compiler/build/exe/java_plugin
  ```

3. `make all`

4. proto 작성 후 `make grpc`