package io.channel.api.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: task/task.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class TaskQueueServiceGrpc {

  private TaskQueueServiceGrpc() {}

  public static final String SERVICE_NAME = "task.TaskQueueService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.ShowTaskRequest,
      io.channel.api.proto.TaskResponse> getShowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "show",
      requestType = io.channel.api.proto.ShowTaskRequest.class,
      responseType = io.channel.api.proto.TaskResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.ShowTaskRequest,
      io.channel.api.proto.TaskResponse> getShowMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.ShowTaskRequest, io.channel.api.proto.TaskResponse> getShowMethod;
    if ((getShowMethod = TaskQueueServiceGrpc.getShowMethod) == null) {
      synchronized (TaskQueueServiceGrpc.class) {
        if ((getShowMethod = TaskQueueServiceGrpc.getShowMethod) == null) {
          TaskQueueServiceGrpc.getShowMethod = getShowMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.ShowTaskRequest, io.channel.api.proto.TaskResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "show"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.ShowTaskRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.TaskResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TaskQueueServiceMethodDescriptorSupplier("show"))
              .build();
        }
      }
    }
    return getShowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.SubmitTaskRequest,
      io.channel.api.proto.TaskResponse> getSubmitMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "submit",
      requestType = io.channel.api.proto.SubmitTaskRequest.class,
      responseType = io.channel.api.proto.TaskResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.SubmitTaskRequest,
      io.channel.api.proto.TaskResponse> getSubmitMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.SubmitTaskRequest, io.channel.api.proto.TaskResponse> getSubmitMethod;
    if ((getSubmitMethod = TaskQueueServiceGrpc.getSubmitMethod) == null) {
      synchronized (TaskQueueServiceGrpc.class) {
        if ((getSubmitMethod = TaskQueueServiceGrpc.getSubmitMethod) == null) {
          TaskQueueServiceGrpc.getSubmitMethod = getSubmitMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.SubmitTaskRequest, io.channel.api.proto.TaskResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "submit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.SubmitTaskRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.TaskResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TaskQueueServiceMethodDescriptorSupplier("submit"))
              .build();
        }
      }
    }
    return getSubmitMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.UpdateTaskRequest,
      io.channel.api.proto.TaskResponse> getUpdateTaskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "updateTask",
      requestType = io.channel.api.proto.UpdateTaskRequest.class,
      responseType = io.channel.api.proto.TaskResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.UpdateTaskRequest,
      io.channel.api.proto.TaskResponse> getUpdateTaskMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.UpdateTaskRequest, io.channel.api.proto.TaskResponse> getUpdateTaskMethod;
    if ((getUpdateTaskMethod = TaskQueueServiceGrpc.getUpdateTaskMethod) == null) {
      synchronized (TaskQueueServiceGrpc.class) {
        if ((getUpdateTaskMethod = TaskQueueServiceGrpc.getUpdateTaskMethod) == null) {
          TaskQueueServiceGrpc.getUpdateTaskMethod = getUpdateTaskMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.UpdateTaskRequest, io.channel.api.proto.TaskResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "updateTask"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.UpdateTaskRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.TaskResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TaskQueueServiceMethodDescriptorSupplier("updateTask"))
              .build();
        }
      }
    }
    return getUpdateTaskMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.channel.api.proto.DeleteTaskRequest,
      io.channel.api.proto.TaskResponse> getDeleteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "delete",
      requestType = io.channel.api.proto.DeleteTaskRequest.class,
      responseType = io.channel.api.proto.TaskResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.channel.api.proto.DeleteTaskRequest,
      io.channel.api.proto.TaskResponse> getDeleteMethod() {
    io.grpc.MethodDescriptor<io.channel.api.proto.DeleteTaskRequest, io.channel.api.proto.TaskResponse> getDeleteMethod;
    if ((getDeleteMethod = TaskQueueServiceGrpc.getDeleteMethod) == null) {
      synchronized (TaskQueueServiceGrpc.class) {
        if ((getDeleteMethod = TaskQueueServiceGrpc.getDeleteMethod) == null) {
          TaskQueueServiceGrpc.getDeleteMethod = getDeleteMethod =
              io.grpc.MethodDescriptor.<io.channel.api.proto.DeleteTaskRequest, io.channel.api.proto.TaskResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "delete"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.DeleteTaskRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.channel.api.proto.TaskResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TaskQueueServiceMethodDescriptorSupplier("delete"))
              .build();
        }
      }
    }
    return getDeleteMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static TaskQueueServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TaskQueueServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TaskQueueServiceStub>() {
        @java.lang.Override
        public TaskQueueServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TaskQueueServiceStub(channel, callOptions);
        }
      };
    return TaskQueueServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static TaskQueueServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TaskQueueServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TaskQueueServiceBlockingStub>() {
        @java.lang.Override
        public TaskQueueServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TaskQueueServiceBlockingStub(channel, callOptions);
        }
      };
    return TaskQueueServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static TaskQueueServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TaskQueueServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TaskQueueServiceFutureStub>() {
        @java.lang.Override
        public TaskQueueServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TaskQueueServiceFutureStub(channel, callOptions);
        }
      };
    return TaskQueueServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class TaskQueueServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * dw -&gt; task queue
     * </pre>
     */
    public void show(io.channel.api.proto.ShowTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getShowMethod(), responseObserver);
    }

    /**
     */
    public void submit(io.channel.api.proto.SubmitTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitMethod(), responseObserver);
    }

    /**
     */
    public void updateTask(io.channel.api.proto.UpdateTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateTaskMethod(), responseObserver);
    }

    /**
     */
    public void delete(io.channel.api.proto.DeleteTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getShowMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.ShowTaskRequest,
                io.channel.api.proto.TaskResponse>(
                  this, METHODID_SHOW)))
          .addMethod(
            getSubmitMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.SubmitTaskRequest,
                io.channel.api.proto.TaskResponse>(
                  this, METHODID_SUBMIT)))
          .addMethod(
            getUpdateTaskMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.UpdateTaskRequest,
                io.channel.api.proto.TaskResponse>(
                  this, METHODID_UPDATE_TASK)))
          .addMethod(
            getDeleteMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                io.channel.api.proto.DeleteTaskRequest,
                io.channel.api.proto.TaskResponse>(
                  this, METHODID_DELETE)))
          .build();
    }
  }

  /**
   */
  public static final class TaskQueueServiceStub extends io.grpc.stub.AbstractAsyncStub<TaskQueueServiceStub> {
    private TaskQueueServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TaskQueueServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TaskQueueServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * dw -&gt; task queue
     * </pre>
     */
    public void show(io.channel.api.proto.ShowTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getShowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void submit(io.channel.api.proto.SubmitTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateTask(io.channel.api.proto.UpdateTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateTaskMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void delete(io.channel.api.proto.DeleteTaskRequest request,
        io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class TaskQueueServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<TaskQueueServiceBlockingStub> {
    private TaskQueueServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TaskQueueServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TaskQueueServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * dw -&gt; task queue
     * </pre>
     */
    public io.channel.api.proto.TaskResponse show(io.channel.api.proto.ShowTaskRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getShowMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.TaskResponse submit(io.channel.api.proto.SubmitTaskRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.TaskResponse updateTask(io.channel.api.proto.UpdateTaskRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateTaskMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.channel.api.proto.TaskResponse delete(io.channel.api.proto.DeleteTaskRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class TaskQueueServiceFutureStub extends io.grpc.stub.AbstractFutureStub<TaskQueueServiceFutureStub> {
    private TaskQueueServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TaskQueueServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TaskQueueServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * dw -&gt; task queue
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.TaskResponse> show(
        io.channel.api.proto.ShowTaskRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getShowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.TaskResponse> submit(
        io.channel.api.proto.SubmitTaskRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.TaskResponse> updateTask(
        io.channel.api.proto.UpdateTaskRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateTaskMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.channel.api.proto.TaskResponse> delete(
        io.channel.api.proto.DeleteTaskRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SHOW = 0;
  private static final int METHODID_SUBMIT = 1;
  private static final int METHODID_UPDATE_TASK = 2;
  private static final int METHODID_DELETE = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final TaskQueueServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(TaskQueueServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SHOW:
          serviceImpl.show((io.channel.api.proto.ShowTaskRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse>) responseObserver);
          break;
        case METHODID_SUBMIT:
          serviceImpl.submit((io.channel.api.proto.SubmitTaskRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse>) responseObserver);
          break;
        case METHODID_UPDATE_TASK:
          serviceImpl.updateTask((io.channel.api.proto.UpdateTaskRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse>) responseObserver);
          break;
        case METHODID_DELETE:
          serviceImpl.delete((io.channel.api.proto.DeleteTaskRequest) request,
              (io.grpc.stub.StreamObserver<io.channel.api.proto.TaskResponse>) responseObserver);
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

  private static abstract class TaskQueueServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    TaskQueueServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.channel.api.proto.TaskOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("TaskQueueService");
    }
  }

  private static final class TaskQueueServiceFileDescriptorSupplier
      extends TaskQueueServiceBaseDescriptorSupplier {
    TaskQueueServiceFileDescriptorSupplier() {}
  }

  private static final class TaskQueueServiceMethodDescriptorSupplier
      extends TaskQueueServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    TaskQueueServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (TaskQueueServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new TaskQueueServiceFileDescriptorSupplier())
              .addMethod(getShowMethod())
              .addMethod(getSubmitMethod())
              .addMethod(getUpdateTaskMethod())
              .addMethod(getDeleteMethod())
              .build();
        }
      }
    }
    return result;
  }
}
