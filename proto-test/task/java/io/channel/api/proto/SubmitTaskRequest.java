// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: task/task.proto

package io.channel.api.proto;

/**
 * Protobuf type {@code task.SubmitTaskRequest}
 */
public final class SubmitTaskRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:task.SubmitTaskRequest)
    SubmitTaskRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use SubmitTaskRequest.newBuilder() to construct.
  private SubmitTaskRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private SubmitTaskRequest() {
    queueName_ = 0;
    data_ = "";
    delay_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new SubmitTaskRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.channel.api.proto.TaskOuterClass.internal_static_task_SubmitTaskRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.channel.api.proto.TaskOuterClass.internal_static_task_SubmitTaskRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.channel.api.proto.SubmitTaskRequest.class, io.channel.api.proto.SubmitTaskRequest.Builder.class);
  }

  public static final int QUEUE_NAME_FIELD_NUMBER = 1;
  private int queueName_;
  /**
   * <code>.task.QueueName queue_name = 1;</code>
   * @return The enum numeric value on the wire for queueName.
   */
  @java.lang.Override public int getQueueNameValue() {
    return queueName_;
  }
  /**
   * <code>.task.QueueName queue_name = 1;</code>
   * @return The queueName.
   */
  @java.lang.Override public io.channel.api.proto.QueueName getQueueName() {
    @SuppressWarnings("deprecation")
    io.channel.api.proto.QueueName result = io.channel.api.proto.QueueName.valueOf(queueName_);
    return result == null ? io.channel.api.proto.QueueName.UNRECOGNIZED : result;
  }

  public static final int DATA_FIELD_NUMBER = 2;
  private volatile java.lang.Object data_;
  /**
   * <code>string data = 2;</code>
   * @return The data.
   */
  @java.lang.Override
  public java.lang.String getData() {
    java.lang.Object ref = data_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      data_ = s;
      return s;
    }
  }
  /**
   * <code>string data = 2;</code>
   * @return The bytes for data.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getDataBytes() {
    java.lang.Object ref = data_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      data_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int DELAY_FIELD_NUMBER = 3;
  private volatile java.lang.Object delay_;
  /**
   * <code>string delay = 3;</code>
   * @return The delay.
   */
  @java.lang.Override
  public java.lang.String getDelay() {
    java.lang.Object ref = delay_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      delay_ = s;
      return s;
    }
  }
  /**
   * <code>string delay = 3;</code>
   * @return The bytes for delay.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getDelayBytes() {
    java.lang.Object ref = delay_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      delay_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (queueName_ != io.channel.api.proto.QueueName.AUTOMATION_RULE.getNumber()) {
      output.writeEnum(1, queueName_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(data_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, data_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(delay_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, delay_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (queueName_ != io.channel.api.proto.QueueName.AUTOMATION_RULE.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, queueName_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(data_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, data_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(delay_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, delay_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.channel.api.proto.SubmitTaskRequest)) {
      return super.equals(obj);
    }
    io.channel.api.proto.SubmitTaskRequest other = (io.channel.api.proto.SubmitTaskRequest) obj;

    if (queueName_ != other.queueName_) return false;
    if (!getData()
        .equals(other.getData())) return false;
    if (!getDelay()
        .equals(other.getDelay())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + QUEUE_NAME_FIELD_NUMBER;
    hash = (53 * hash) + queueName_;
    hash = (37 * hash) + DATA_FIELD_NUMBER;
    hash = (53 * hash) + getData().hashCode();
    hash = (37 * hash) + DELAY_FIELD_NUMBER;
    hash = (53 * hash) + getDelay().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.SubmitTaskRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(io.channel.api.proto.SubmitTaskRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code task.SubmitTaskRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:task.SubmitTaskRequest)
      io.channel.api.proto.SubmitTaskRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.channel.api.proto.TaskOuterClass.internal_static_task_SubmitTaskRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.channel.api.proto.TaskOuterClass.internal_static_task_SubmitTaskRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.channel.api.proto.SubmitTaskRequest.class, io.channel.api.proto.SubmitTaskRequest.Builder.class);
    }

    // Construct using io.channel.api.proto.SubmitTaskRequest.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      queueName_ = 0;

      data_ = "";

      delay_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.channel.api.proto.TaskOuterClass.internal_static_task_SubmitTaskRequest_descriptor;
    }

    @java.lang.Override
    public io.channel.api.proto.SubmitTaskRequest getDefaultInstanceForType() {
      return io.channel.api.proto.SubmitTaskRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.channel.api.proto.SubmitTaskRequest build() {
      io.channel.api.proto.SubmitTaskRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.channel.api.proto.SubmitTaskRequest buildPartial() {
      io.channel.api.proto.SubmitTaskRequest result = new io.channel.api.proto.SubmitTaskRequest(this);
      result.queueName_ = queueName_;
      result.data_ = data_;
      result.delay_ = delay_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof io.channel.api.proto.SubmitTaskRequest) {
        return mergeFrom((io.channel.api.proto.SubmitTaskRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.channel.api.proto.SubmitTaskRequest other) {
      if (other == io.channel.api.proto.SubmitTaskRequest.getDefaultInstance()) return this;
      if (other.queueName_ != 0) {
        setQueueNameValue(other.getQueueNameValue());
      }
      if (!other.getData().isEmpty()) {
        data_ = other.data_;
        onChanged();
      }
      if (!other.getDelay().isEmpty()) {
        delay_ = other.delay_;
        onChanged();
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 8: {
              queueName_ = input.readEnum();

              break;
            } // case 8
            case 18: {
              data_ = input.readStringRequireUtf8();

              break;
            } // case 18
            case 26: {
              delay_ = input.readStringRequireUtf8();

              break;
            } // case 26
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private int queueName_ = 0;
    /**
     * <code>.task.QueueName queue_name = 1;</code>
     * @return The enum numeric value on the wire for queueName.
     */
    @java.lang.Override public int getQueueNameValue() {
      return queueName_;
    }
    /**
     * <code>.task.QueueName queue_name = 1;</code>
     * @param value The enum numeric value on the wire for queueName to set.
     * @return This builder for chaining.
     */
    public Builder setQueueNameValue(int value) {
      
      queueName_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.task.QueueName queue_name = 1;</code>
     * @return The queueName.
     */
    @java.lang.Override
    public io.channel.api.proto.QueueName getQueueName() {
      @SuppressWarnings("deprecation")
      io.channel.api.proto.QueueName result = io.channel.api.proto.QueueName.valueOf(queueName_);
      return result == null ? io.channel.api.proto.QueueName.UNRECOGNIZED : result;
    }
    /**
     * <code>.task.QueueName queue_name = 1;</code>
     * @param value The queueName to set.
     * @return This builder for chaining.
     */
    public Builder setQueueName(io.channel.api.proto.QueueName value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      queueName_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.task.QueueName queue_name = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearQueueName() {
      
      queueName_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object data_ = "";
    /**
     * <code>string data = 2;</code>
     * @return The data.
     */
    public java.lang.String getData() {
      java.lang.Object ref = data_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        data_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string data = 2;</code>
     * @return The bytes for data.
     */
    public com.google.protobuf.ByteString
        getDataBytes() {
      java.lang.Object ref = data_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        data_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string data = 2;</code>
     * @param value The data to set.
     * @return This builder for chaining.
     */
    public Builder setData(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      data_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string data = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearData() {
      
      data_ = getDefaultInstance().getData();
      onChanged();
      return this;
    }
    /**
     * <code>string data = 2;</code>
     * @param value The bytes for data to set.
     * @return This builder for chaining.
     */
    public Builder setDataBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      data_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object delay_ = "";
    /**
     * <code>string delay = 3;</code>
     * @return The delay.
     */
    public java.lang.String getDelay() {
      java.lang.Object ref = delay_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        delay_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string delay = 3;</code>
     * @return The bytes for delay.
     */
    public com.google.protobuf.ByteString
        getDelayBytes() {
      java.lang.Object ref = delay_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        delay_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string delay = 3;</code>
     * @param value The delay to set.
     * @return This builder for chaining.
     */
    public Builder setDelay(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      delay_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string delay = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearDelay() {
      
      delay_ = getDefaultInstance().getDelay();
      onChanged();
      return this;
    }
    /**
     * <code>string delay = 3;</code>
     * @param value The bytes for delay to set.
     * @return This builder for chaining.
     */
    public Builder setDelayBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      delay_ = value;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:task.SubmitTaskRequest)
  }

  // @@protoc_insertion_point(class_scope:task.SubmitTaskRequest)
  private static final io.channel.api.proto.SubmitTaskRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.channel.api.proto.SubmitTaskRequest();
  }

  public static io.channel.api.proto.SubmitTaskRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<SubmitTaskRequest>
      PARSER = new com.google.protobuf.AbstractParser<SubmitTaskRequest>() {
    @java.lang.Override
    public SubmitTaskRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<SubmitTaskRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<SubmitTaskRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.channel.api.proto.SubmitTaskRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

