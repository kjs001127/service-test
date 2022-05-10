package io.channel.api.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0-SNAPSHOT)",
    comments = "Source: meet/meet.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class MeetServiceGrpc {

  private MeetServiceGrpc() {}

  public static final String SERVICE_NAME = "meet.MeetService";

  // Static method descriptors that strictly reflect the proto.
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

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.LeaveMeetRequest,
      io.channel.api.proto.BareResponse> getLeaveMeetByUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LeaveMeetByUser",
      requestType = io.channel.api.proto.LeaveMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.LeaveMeetRequest,
      io.channel.api.proto.BareResponse> getLeaveMeetByUserMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.LeaveMeetRequest, io.channel.api.proto.BareResponse> getLeaveMeetByUserMethod;
    if ((getLeaveMeetByUserMethod = MeetServiceGrpc.getLeaveMeetByUserMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getLeaveMeetByUserMethod = MeetServiceGrpc.getLeaveMeetByUserMethod) == null) {
          MeetServiceGrpc.getLeaveMeetByUserMethod = getLeaveMeetByUserMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.LeaveMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "LeaveMeetByUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.LeaveMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("LeaveMeetByUser"))
              .build();
        }
      }
    }
    return getLeaveMeetByUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.LeaveMeetRequest,
      io.channel.api.proto.BareResponse> getLeaveMeetByManagerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LeaveMeetByManager",
      requestType = io.channel.api.proto.LeaveMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.LeaveMeetRequest,
      io.channel.api.proto.BareResponse> getLeaveMeetByManagerMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.LeaveMeetRequest, io.channel.api.proto.BareResponse> getLeaveMeetByManagerMethod;
    if ((getLeaveMeetByManagerMethod = MeetServiceGrpc.getLeaveMeetByManagerMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getLeaveMeetByManagerMethod = MeetServiceGrpc.getLeaveMeetByManagerMethod) == null) {
          MeetServiceGrpc.getLeaveMeetByManagerMethod = getLeaveMeetByManagerMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.LeaveMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "LeaveMeetByManager"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.LeaveMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("LeaveMeetByManager"))
              .build();
        }
      }
    }
    return getLeaveMeetByManagerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.AddPeersRequest,
      io.channel.api.proto.AddPeersResponse> getAddPeersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AddPeers",
      requestType = io.channel.api.proto.AddPeersRequest.class,
      responseType = io.channel.api.proto.AddPeersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.AddPeersRequest,
      io.channel.api.proto.AddPeersResponse> getAddPeersMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.AddPeersRequest, io.channel.api.proto.AddPeersResponse> getAddPeersMethod;
    if ((getAddPeersMethod = MeetServiceGrpc.getAddPeersMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getAddPeersMethod = MeetServiceGrpc.getAddPeersMethod) == null) {
          MeetServiceGrpc.getAddPeersMethod = getAddPeersMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.AddPeersRequest, io.channel.api.proto.AddPeersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AddPeers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.AddPeersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.AddPeersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("AddPeers"))
              .build();
        }
      }
    }
    return getAddPeersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "JoinMeetByUser",
      requestType = io.channel.api.proto.JoinMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByUserMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetRequest, io.channel.api.proto.BareResponse> getJoinMeetByUserMethod;
    if ((getJoinMeetByUserMethod = MeetServiceGrpc.getJoinMeetByUserMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getJoinMeetByUserMethod = MeetServiceGrpc.getJoinMeetByUserMethod) == null) {
          MeetServiceGrpc.getJoinMeetByUserMethod = getJoinMeetByUserMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.JoinMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "JoinMeetByUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.JoinMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("JoinMeetByUser"))
              .build();
        }
      }
    }
    return getJoinMeetByUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByManagerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "JoinMeetByManager",
      requestType = io.channel.api.proto.JoinMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetRequest,
      io.channel.api.proto.BareResponse> getJoinMeetByManagerMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.JoinMeetRequest, io.channel.api.proto.BareResponse> getJoinMeetByManagerMethod;
    if ((getJoinMeetByManagerMethod = MeetServiceGrpc.getJoinMeetByManagerMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getJoinMeetByManagerMethod = MeetServiceGrpc.getJoinMeetByManagerMethod) == null) {
          MeetServiceGrpc.getJoinMeetByManagerMethod = getJoinMeetByManagerMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.JoinMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "JoinMeetByManager"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.JoinMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("JoinMeetByManager"))
              .build();
        }
      }
    }
    return getJoinMeetByManagerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.CloseMeetRequest,
      io.channel.api.proto.BareResponse> getCloseMeetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CloseMeet",
      requestType = io.channel.api.proto.CloseMeetRequest.class,
      responseType = io.channel.api.proto.BareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.CloseMeetRequest,
      io.channel.api.proto.BareResponse> getCloseMeetMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.CloseMeetRequest, io.channel.api.proto.BareResponse> getCloseMeetMethod;
    if ((getCloseMeetMethod = MeetServiceGrpc.getCloseMeetMethod) == null) {
      synchronized (MeetServiceGrpc.class) {
        if ((getCloseMeetMethod = MeetServiceGrpc.getCloseMeetMethod) == null) {
          MeetServiceGrpc.getCloseMeetMethod = getCloseMeetMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.CloseMeetRequest, io.channel.api.proto.BareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CloseMeet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.CloseMeetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.BareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MeetServiceMethodDescriptorSupplier("CloseMeet"))
              .build();
        }
      }
    }
    return getCloseMeetMethod;
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
     */
    public void createInboundMeet(io.channel.api.proto.InboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.InboundMeetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateInboundMeetMethod(), responseObserver);
    }

    /**
     */
    public void createOutboundMeet(io.channel.api.proto.OutboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateOutboundMeetMethod(), responseObserver);
    }

    /**
     */
    public void leaveMeetByUser(io.channel.api.proto.LeaveMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLeaveMeetByUserMethod(), responseObserver);
    }

    /**
     */
    public void leaveMeetByManager(io.channel.api.proto.LeaveMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLeaveMeetByManagerMethod(), responseObserver);
    }

    /**
     */
    public void addPeers(io.channel.api.proto.AddPeersRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.AddPeersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAddPeersMethod(), responseObserver);
    }

    /**
     */
    public void joinMeetByUser(io.channel.api.proto.JoinMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getJoinMeetByUserMethod(), responseObserver);
    }

    /**
     */
    public void joinMeetByManager(io.channel.api.proto.JoinMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getJoinMeetByManagerMethod(), responseObserver);
    }

    /**
     */
    public void closeMeet(io.channel.api.proto.CloseMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCloseMeetMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCreateInboundMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.InboundMeetRequest,
                io.channel.api.proto.InboundMeetResponse>(
                  this, METHODID_CREATE_INBOUND_MEET)))
          .addMethod(
            getCreateOutboundMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.OutboundMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_CREATE_OUTBOUND_MEET)))
          .addMethod(
            getLeaveMeetByUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.LeaveMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_LEAVE_MEET_BY_USER)))
          .addMethod(
            getLeaveMeetByManagerMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.LeaveMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_LEAVE_MEET_BY_MANAGER)))
          .addMethod(
            getAddPeersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.AddPeersRequest,
                io.channel.api.proto.AddPeersResponse>(
                  this, METHODID_ADD_PEERS)))
          .addMethod(
            getJoinMeetByUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.JoinMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_JOIN_MEET_BY_USER)))
          .addMethod(
            getJoinMeetByManagerMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.JoinMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_JOIN_MEET_BY_MANAGER)))
          .addMethod(
            getCloseMeetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.CloseMeetRequest,
                io.channel.api.proto.BareResponse>(
                  this, METHODID_CLOSE_MEET)))
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
     */
    public void createInboundMeet(io.channel.api.proto.InboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.InboundMeetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateInboundMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createOutboundMeet(io.channel.api.proto.OutboundMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateOutboundMeetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void leaveMeetByUser(io.channel.api.proto.LeaveMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLeaveMeetByUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void leaveMeetByManager(io.channel.api.proto.LeaveMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLeaveMeetByManagerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void addPeers(io.channel.api.proto.AddPeersRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.AddPeersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAddPeersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void joinMeetByUser(io.channel.api.proto.JoinMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getJoinMeetByUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void joinMeetByManager(io.channel.api.proto.JoinMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getJoinMeetByManagerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void closeMeet(io.channel.api.proto.CloseMeetRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCloseMeetMethod(), getCallOptions()), request, responseObserver);
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
     */
    public io.channel.api.proto.InboundMeetResponse createInboundMeet(io.channel.api.proto.InboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateInboundMeetMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse createOutboundMeet(io.channel.api.proto.OutboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateOutboundMeetMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse leaveMeetByUser(io.channel.api.proto.LeaveMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLeaveMeetByUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse leaveMeetByManager(io.channel.api.proto.LeaveMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLeaveMeetByManagerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.AddPeersResponse addPeers(io.channel.api.proto.AddPeersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAddPeersMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse joinMeetByUser(io.channel.api.proto.JoinMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getJoinMeetByUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse joinMeetByManager(io.channel.api.proto.JoinMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getJoinMeetByManagerMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.BareResponse closeMeet(io.channel.api.proto.CloseMeetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCloseMeetMethod(), getCallOptions(), request);
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
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.InboundMeetResponse> createInboundMeet(
        io.channel.api.proto.InboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateInboundMeetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> createOutboundMeet(
        io.channel.api.proto.OutboundMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateOutboundMeetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> leaveMeetByUser(
        io.channel.api.proto.LeaveMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLeaveMeetByUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> leaveMeetByManager(
        io.channel.api.proto.LeaveMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLeaveMeetByManagerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.AddPeersResponse> addPeers(
        io.channel.api.proto.AddPeersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAddPeersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> joinMeetByUser(
        io.channel.api.proto.JoinMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getJoinMeetByUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> joinMeetByManager(
        io.channel.api.proto.JoinMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getJoinMeetByManagerMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.BareResponse> closeMeet(
        io.channel.api.proto.CloseMeetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCloseMeetMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_INBOUND_MEET = 0;
  private static final int METHODID_CREATE_OUTBOUND_MEET = 1;
  private static final int METHODID_LEAVE_MEET_BY_USER = 2;
  private static final int METHODID_LEAVE_MEET_BY_MANAGER = 3;
  private static final int METHODID_ADD_PEERS = 4;
  private static final int METHODID_JOIN_MEET_BY_USER = 5;
  private static final int METHODID_JOIN_MEET_BY_MANAGER = 6;
  private static final int METHODID_CLOSE_MEET = 7;

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
        case METHODID_CREATE_INBOUND_MEET:
          serviceImpl.createInboundMeet((io.channel.api.proto.InboundMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.InboundMeetResponse>) responseObserver);
          break;
        case METHODID_CREATE_OUTBOUND_MEET:
          serviceImpl.createOutboundMeet((io.channel.api.proto.OutboundMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_LEAVE_MEET_BY_USER:
          serviceImpl.leaveMeetByUser((io.channel.api.proto.LeaveMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_LEAVE_MEET_BY_MANAGER:
          serviceImpl.leaveMeetByManager((io.channel.api.proto.LeaveMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_ADD_PEERS:
          serviceImpl.addPeers((io.channel.api.proto.AddPeersRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.AddPeersResponse>) responseObserver);
          break;
        case METHODID_JOIN_MEET_BY_USER:
          serviceImpl.joinMeetByUser((io.channel.api.proto.JoinMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_JOIN_MEET_BY_MANAGER:
          serviceImpl.joinMeetByManager((io.channel.api.proto.JoinMeetRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.BareResponse>) responseObserver);
          break;
        case METHODID_CLOSE_MEET:
          serviceImpl.closeMeet((io.channel.api.proto.CloseMeetRequest) request,
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
              .addMethod(getCreateInboundMeetMethod())
              .addMethod(getCreateOutboundMeetMethod())
              .addMethod(getLeaveMeetByUserMethod())
              .addMethod(getLeaveMeetByManagerMethod())
              .addMethod(getAddPeersMethod())
              .addMethod(getJoinMeetByUserMethod())
              .addMethod(getJoinMeetByManagerMethod())
              .addMethod(getCloseMeetMethod())
              .build();
        }
      }
    }
    return result;
  }
}
