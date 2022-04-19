package io.channel.api.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.46.0-SNAPSHOT)",
    comments = "Source: meet/meet.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class MeetServiceGrpc {

  private MeetServiceGrpc() {}

  public static final String SERVICE_NAME = "meet.MeetService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.InboundCallRequest,
      io.channel.api.proto.MeetId> getCreateInboundCallMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateInboundCallMeet",
      requestType = io.channel.api.proto.InboundCallRequest.class,
      responseType = io.channel.api.proto.MeetId.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.InboundCallRequest,
      io.channel.api.proto.MeetId> getCreateInboundCallMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.InboundCallRequest, io.channel.api.proto.MeetId> getCreateInboundCallMeetMethod;
    if ((getCreateInboundCallMeetMethod = MeetServiceGrpc.getCreateInboundCallMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getCreateInboundCallMeetMethod = MeetServiceGrpc.getCreateInboundCallMeetMethod) == null) {
          MeetServiceGrpc.getCreateInboundCallMeetMethod = getCreateInboundCallMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.InboundCallRequest, io.channel.api.proto.MeetId>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateInboundCallMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.InboundCallRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.MeetId.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("CreateInboundCallMeet"))
              .build();
        }
      }
    }
    return getCreateInboundCallMeetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.OutboundCallRequest,
      io.channel.api.proto.MeetId> getCreateOutboundCallMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateOutboundCallMeet",
      requestType = io.channel.api.proto.OutboundCallRequest.class,
      responseType = io.channel.api.proto.MeetId.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.OutboundCallRequest,
      io.channel.api.proto.MeetId> getCreateOutboundCallMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.OutboundCallRequest, io.channel.api.proto.MeetId> getCreateOutboundCallMeetMethod;
    if ((getCreateOutboundCallMeetMethod = MeetServiceGrpc.getCreateOutboundCallMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getCreateOutboundCallMeetMethod = MeetServiceGrpc.getCreateOutboundCallMeetMethod) == null) {
          MeetServiceGrpc.getCreateOutboundCallMeetMethod = getCreateOutboundCallMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.OutboundCallRequest, io.channel.api.proto.MeetId>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateOutboundCallMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.OutboundCallRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.MeetId.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("CreateOutboundCallMeet"))
              .build();
        }
      }
    }
    return getCreateOutboundCallMeetMethod;
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
     * sfu -&gt; dw
     * </pre>
     */
    public void createInboundCallMeet(io.channel.api.proto.InboundCallRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.MeetId> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateInboundCallMeetMethod(), responseObserver);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public void createOutboundCallMeet(io.channel.api.proto.OutboundCallRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.MeetId> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateOutboundCallMeetMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCreateInboundCallMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.InboundCallRequest,
                io.channel.api.proto.MeetId>(
                  this, METHODID_CREATE_INBOUND_CALL_MEET)))
          .addMethod(
            getCreateOutboundCallMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.OutboundCallRequest,
                io.channel.api.proto.MeetId>(
                  this, METHODID_CREATE_OUTBOUND_CALL_MEET)))
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
     * sfu -&gt; dw
     * </pre>
     */
    public void createInboundCallMeet(io.channel.api.proto.InboundCallRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.MeetId> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateInboundCallMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public void createOutboundCallMeet(io.channel.api.proto.OutboundCallRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.MeetId> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateOutboundCallMeetMethod(), getCallOptions()), request, responseObserver);
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
     * sfu -&gt; dw
     * </pre>
     */
    public io.channel.api.proto.MeetId createInboundCallMeet(io.channel.api.proto.InboundCallRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateInboundCallMeetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public io.channel.api.proto.MeetId createOutboundCallMeet(io.channel.api.proto.OutboundCallRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateOutboundCallMeetMethod(), getCallOptions(), request);
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
     * sfu -&gt; dw
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.MeetId> createInboundCallMeet(
        io.channel.api.proto.InboundCallRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateInboundCallMeetMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * dw -&gt; sfu
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.MeetId> createOutboundCallMeet(
        io.channel.api.proto.OutboundCallRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateOutboundCallMeetMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_INBOUND_CALL_MEET = 0;
  private static final int METHODID_CREATE_OUTBOUND_CALL_MEET = 1;

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
        case METHODID_CREATE_INBOUND_CALL_MEET:
          serviceImpl.createInboundCallMeet((io.channel.api.proto.InboundCallRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.MeetId>) responseObserver);
          break;
        case METHODID_CREATE_OUTBOUND_CALL_MEET:
          serviceImpl.createOutboundCallMeet((io.channel.api.proto.OutboundCallRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.MeetId>) responseObserver);
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
              .addMethod(getCreateInboundCallMeetMethod())
              .addMethod(getCreateOutboundCallMeetMethod())
              .build();
        }
      }
    }
    return result;
  }
}
