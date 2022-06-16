package io.channel.api.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.48.0-SNAPSHOT)",
    comments = "Source: meet/meet.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class MeetServiceGrpc {

  private MeetServiceGrpc() {}

  public static final String SERVICE_NAME = "meet.MeetService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.OutboundMeetRequest,
      io.channel.api.proto.BareResponse> getCreateOutboundMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateOutboundMeet",
      requestType = io.channel.api.proto.OutboundMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.OutboundMeetRequest,
      io.channel.api.proto.BareResponse> getCreateOutboundMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.OutboundMeetRequest, io.channel.api.proto.BareResponse> getCreateOutboundMeetMethod;
    if ((getCreateOutboundMeetMethod = MeetServiceGrpc.getCreateOutboundMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getCreateOutboundMeetMethod = MeetServiceGrpc.getCreateOutboundMeetMethod) == null) {
          MeetServiceGrpc.getCreateOutboundMeetMethod = getCreateOutboundMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.OutboundMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateOutboundMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.OutboundMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("CreateOutboundMeet"))
              .build();
        }
      }
    }
    return getCreateOutboundMeetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.PrivateMeetRequest,
      io.channel.api.proto.BareResponse> getCreatePrivateMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreatePrivateMeet",
      requestType = io.channel.api.proto.PrivateMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.PrivateMeetRequest,
      io.channel.api.proto.BareResponse> getCreatePrivateMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.PrivateMeetRequest, io.channel.api.proto.BareResponse> getCreatePrivateMeetMethod;
    if ((getCreatePrivateMeetMethod = MeetServiceGrpc.getCreatePrivateMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getCreatePrivateMeetMethod = MeetServiceGrpc.getCreatePrivateMeetMethod) == null) {
          MeetServiceGrpc.getCreatePrivateMeetMethod = getCreatePrivateMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.PrivateMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreatePrivateMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.PrivateMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("CreatePrivateMeet"))
              .build();
        }
      }
    }
    return getCreatePrivateMeetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.HangUpMeetByManagerRequest,
      io.channel.api.proto.BareResponse> getHangUpMeetByManagerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "HangUpMeetByManager",
      requestType = io.channel.api.proto.HangUpMeetByManagerRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.HangUpMeetByManagerRequest,
      io.channel.api.proto.BareResponse> getHangUpMeetByManagerMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.HangUpMeetByManagerRequest, io.channel.api.proto.BareResponse> getHangUpMeetByManagerMethod;
    if ((getHangUpMeetByManagerMethod = MeetServiceGrpc.getHangUpMeetByManagerMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getHangUpMeetByManagerMethod = MeetServiceGrpc.getHangUpMeetByManagerMethod) == null) {
          MeetServiceGrpc.getHangUpMeetByManagerMethod = getHangUpMeetByManagerMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.HangUpMeetByManagerRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "HangUpMeetByManager"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.HangUpMeetByManagerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("HangUpMeetByManager"))
              .build();
        }
      }
    }
    return getHangUpMeetByManagerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetByManagerRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByManagerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "JoinMeetByManager",
      requestType = io.channel.api.proto.JoinMeetByManagerRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetByManagerRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByManagerMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetByManagerRequest, io.channel.api.proto.BareResponse> getJoinMeetByManagerMethod;
    if ((getJoinMeetByManagerMethod = MeetServiceGrpc.getJoinMeetByManagerMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getJoinMeetByManagerMethod = MeetServiceGrpc.getJoinMeetByManagerMethod) == null) {
          MeetServiceGrpc.getJoinMeetByManagerMethod = getJoinMeetByManagerMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.JoinMeetByManagerRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "JoinMeetByManager"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.JoinMeetByManagerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("JoinMeetByManager"))
              .build();
        }
      }
    }
    return getJoinMeetByManagerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.TerminateMeetRequest,
      io.channel.api.proto.BareResponse> getTerminateMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "TerminateMeet",
      requestType = io.channel.api.proto.TerminateMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.TerminateMeetRequest,
      io.channel.api.proto.BareResponse> getTerminateMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.TerminateMeetRequest, io.channel.api.proto.BareResponse> getTerminateMeetMethod;
    if ((getTerminateMeetMethod = MeetServiceGrpc.getTerminateMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getTerminateMeetMethod = MeetServiceGrpc.getTerminateMeetMethod) == null) {
          MeetServiceGrpc.getTerminateMeetMethod = getTerminateMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.TerminateMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "TerminateMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.TerminateMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("TerminateMeet"))
              .build();
        }
      }
    }
    return getTerminateMeetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.InboundMeetRequest,
      io.channel.api.proto.InboundMeetResponse> getCreateInboundMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateInboundMeet",
      requestType = io.channel.api.proto.InboundMeetRequest.class,
      responseType = io.channel.api.proto.InboundMeetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.InboundMeetRequest,
      io.channel.api.proto.InboundMeetResponse> getCreateInboundMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.InboundMeetRequest, io.channel.api.proto.InboundMeetResponse> getCreateInboundMeetMethod;
    if ((getCreateInboundMeetMethod = MeetServiceGrpc.getCreateInboundMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getCreateInboundMeetMethod = MeetServiceGrpc.getCreateInboundMeetMethod) == null) {
          MeetServiceGrpc.getCreateInboundMeetMethod = getCreateInboundMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.InboundMeetRequest, io.channel.api.proto.InboundMeetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateInboundMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.InboundMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.InboundMeetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("CreateInboundMeet"))
              .build();
        }
      }
    }
    return getCreateInboundMeetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetByUserRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "JoinMeetByUser",
      requestType = io.channel.api.proto.JoinMeetByUserRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetByUserRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByUserMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetByUserRequest, io.channel.api.proto.BareResponse> getJoinMeetByUserMethod;
    if ((getJoinMeetByUserMethod = MeetServiceGrpc.getJoinMeetByUserMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getJoinMeetByUserMethod = MeetServiceGrpc.getJoinMeetByUserMethod) == null) {
          MeetServiceGrpc.getJoinMeetByUserMethod = getJoinMeetByUserMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.JoinMeetByUserRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "JoinMeetByUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.JoinMeetByUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("JoinMeetByUser"))
              .build();
        }
      }
    }
    return getJoinMeetByUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.HangUpMeetByUserRequest,
      io.channel.api.proto.BareResponse> getHangUpMeetByUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "HangUpMeetByUser",
      requestType = io.channel.api.proto.HangUpMeetByUserRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.HangUpMeetByUserRequest,
      io.channel.api.proto.BareResponse> getHangUpMeetByUserMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.HangUpMeetByUserRequest, io.channel.api.proto.BareResponse> getHangUpMeetByUserMethod;
    if ((getHangUpMeetByUserMethod = MeetServiceGrpc.getHangUpMeetByUserMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getHangUpMeetByUserMethod = MeetServiceGrpc.getHangUpMeetByUserMethod) == null) {
          MeetServiceGrpc.getHangUpMeetByUserMethod = getHangUpMeetByUserMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.HangUpMeetByUserRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "HangUpMeetByUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.HangUpMeetByUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("HangUpMeetByUser"))
              .build();
        }
      }
    }
    return getHangUpMeetByUserMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MeetServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MeetServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MeetServiceStub>() {
        @java.lang.Override
        public MeetServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MeetServiceStub(channel, callOptions);
        }
      };
    return MeetServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MeetServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MeetServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MeetServiceBlockingStub>() {
        @java.lang.Override
        public MeetServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MeetServiceBlockingStub(channel, callOptions);
        }
      };
    return MeetServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MeetServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MeetServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MeetServiceFutureStub>() {
        @java.lang.Override
        public MeetServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MeetServiceFutureStub(channel, callOptions);
        }
      };
    return MeetServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class MeetServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public void createOutboundMeet(io.channel.api.proto.OutboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateOutboundMeetMethod(), responseObserver);
    }

    /**
     */
    public void createPrivateMeet(io.channel.api.proto.PrivateMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreatePrivateMeetMethod(), responseObserver);
    }

    /**
     */
    public void hangUpMeetByManager(io.channel.api.proto.HangUpMeetByManagerRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getHangUpMeetByManagerMethod(), responseObserver);
    }

    /**
     */
    public void joinMeetByManager(io.channel.api.proto.JoinMeetByManagerRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getJoinMeetByManagerMethod(), responseObserver);
    }

    /**
     */
    public void terminateMeet(io.channel.api.proto.TerminateMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getTerminateMeetMethod(), responseObserver);
    }

    /**
     * <pre>
     * sfu -&gt; dw
     * </pre>
     */
    public void createInboundMeet(io.channel.api.proto.InboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.InboundMeetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateInboundMeetMethod(), responseObserver);
    }

    /**
     */
    public void joinMeetByUser(io.channel.api.proto.JoinMeetByUserRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getJoinMeetByUserMethod(), responseObserver);
    }

    /**
     */
    public void hangUpMeetByUser(io.channel.api.proto.HangUpMeetByUserRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getHangUpMeetByUserMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCreateOutboundMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.OutboundMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_CREATE_OUTBOUND_MEET)))
          .addMethod(
            getCreatePrivateMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.PrivateMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_CREATE_PRIVATE_MEET)))
          .addMethod(
            getHangUpMeetByManagerMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.HangUpMeetByManagerRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_HANG_UP_MEET_BY_MANAGER)))
          .addMethod(
            getJoinMeetByManagerMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.JoinMeetByManagerRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_JOIN_MEET_BY_MANAGER)))
          .addMethod(
            getTerminateMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.TerminateMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_TERMINATE_MEET)))
          .addMethod(
            getCreateInboundMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.InboundMeetRequest,
                io.channel.api.proto.InboundMeetResponse>(
                  this, METHODID_CREATE_INBOUND_MEET)))
          .addMethod(
            getJoinMeetByUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.JoinMeetByUserRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_JOIN_MEET_BY_USER)))
          .addMethod(
            getHangUpMeetByUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.HangUpMeetByUserRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_HANG_UP_MEET_BY_USER)))
          .build();
    }
  }

  /**
   */
  public static final class MeetServiceStub extends io.grpc.stub.AbstractAsyncStub<MeetServiceStub> {
    private MeetServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MeetServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MeetServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public void createOutboundMeet(io.channel.api.proto.OutboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateOutboundMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createPrivateMeet(io.channel.api.proto.PrivateMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreatePrivateMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void hangUpMeetByManager(io.channel.api.proto.HangUpMeetByManagerRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getHangUpMeetByManagerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void joinMeetByManager(io.channel.api.proto.JoinMeetByManagerRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getJoinMeetByManagerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void terminateMeet(io.channel.api.proto.TerminateMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getTerminateMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * sfu -&gt; dw
     * </pre>
     */
    public void createInboundMeet(io.channel.api.proto.InboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.InboundMeetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateInboundMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void joinMeetByUser(io.channel.api.proto.JoinMeetByUserRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getJoinMeetByUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void hangUpMeetByUser(io.channel.api.proto.HangUpMeetByUserRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getHangUpMeetByUserMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class MeetServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<MeetServiceBlockingStub> {
    private MeetServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MeetServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MeetServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public io.channel.api.proto.BareResponse createOutboundMeet(io.channel.api.proto.OutboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateOutboundMeetMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse createPrivateMeet(io.channel.api.proto.PrivateMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreatePrivateMeetMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse hangUpMeetByManager(io.channel.api.proto.HangUpMeetByManagerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getHangUpMeetByManagerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse joinMeetByManager(io.channel.api.proto.JoinMeetByManagerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getJoinMeetByManagerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse terminateMeet(io.channel.api.proto.TerminateMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getTerminateMeetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * sfu -&gt; dw
     * </pre>
     */
    public io.channel.api.proto.InboundMeetResponse createInboundMeet(io.channel.api.proto.InboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateInboundMeetMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse joinMeetByUser(io.channel.api.proto.JoinMeetByUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getJoinMeetByUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse hangUpMeetByUser(io.channel.api.proto.HangUpMeetByUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getHangUpMeetByUserMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class MeetServiceFutureStub extends io.grpc.stub.AbstractFutureStub<MeetServiceFutureStub> {
    private MeetServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MeetServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MeetServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> createOutboundMeet(
        io.channel.api.proto.OutboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateOutboundMeetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> createPrivateMeet(
        io.channel.api.proto.PrivateMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreatePrivateMeetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> hangUpMeetByManager(
        io.channel.api.proto.HangUpMeetByManagerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getHangUpMeetByManagerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> joinMeetByManager(
        io.channel.api.proto.JoinMeetByManagerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getJoinMeetByManagerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> terminateMeet(
        io.channel.api.proto.TerminateMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getTerminateMeetMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * sfu -&gt; dw
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.InboundMeetResponse> createInboundMeet(
        io.channel.api.proto.InboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateInboundMeetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> joinMeetByUser(
        io.channel.api.proto.JoinMeetByUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getJoinMeetByUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> hangUpMeetByUser(
        io.channel.api.proto.HangUpMeetByUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getHangUpMeetByUserMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_OUTBOUND_MEET = 0;
  private static final int METHODID_CREATE_PRIVATE_MEET = 1;
  private static final int METHODID_HANG_UP_MEET_BY_MANAGER = 2;
  private static final int METHODID_JOIN_MEET_BY_MANAGER = 3;
  private static final int METHODID_TERMINATE_MEET = 4;
  private static final int METHODID_CREATE_INBOUND_MEET = 5;
  private static final int METHODID_JOIN_MEET_BY_USER = 6;
  private static final int METHODID_HANG_UP_MEET_BY_USER = 7;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final MeetServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(MeetServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_OUTBOUND_MEET:
          serviceImpl.createOutboundMeet((io.channel.api.proto.OutboundMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_CREATE_PRIVATE_MEET:
          serviceImpl.createPrivateMeet((io.channel.api.proto.PrivateMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_HANG_UP_MEET_BY_MANAGER:
          serviceImpl.hangUpMeetByManager((io.channel.api.proto.HangUpMeetByManagerRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_JOIN_MEET_BY_MANAGER:
          serviceImpl.joinMeetByManager((io.channel.api.proto.JoinMeetByManagerRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_TERMINATE_MEET:
          serviceImpl.terminateMeet((io.channel.api.proto.TerminateMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_CREATE_INBOUND_MEET:
          serviceImpl.createInboundMeet((io.channel.api.proto.InboundMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.InboundMeetResponse>) responseObserver);
          break;
        case METHODID_JOIN_MEET_BY_USER:
          serviceImpl.joinMeetByUser((io.channel.api.proto.JoinMeetByUserRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_HANG_UP_MEET_BY_USER:
          serviceImpl.hangUpMeetByUser((io.channel.api.proto.HangUpMeetByUserRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class MeetServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MeetServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.channel.api.proto.Meet.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MeetService");
    }
  }

  private static final class MeetServiceFileDescriptorSupplier
      extends MeetServiceBaseDescriptorSupplier {
    MeetServiceFileDescriptorSupplier() {}
  }

  private static final class MeetServiceMethodDescriptorSupplier
      extends MeetServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    MeetServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (MeetServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MeetServiceFileDescriptorSupplier())
              .addMethod(getCreateOutboundMeetMethod())
              .addMethod(getCreatePrivateMeetMethod())
              .addMethod(getHangUpMeetByManagerMethod())
              .addMethod(getJoinMeetByManagerMethod())
              .addMethod(getTerminateMeetMethod())
              .addMethod(getCreateInboundMeetMethod())
              .addMethod(getJoinMeetByUserMethod())
              .addMethod(getHangUpMeetByUserMethod())
              .build();
        }
      }
    }
    return result;
  }
}
