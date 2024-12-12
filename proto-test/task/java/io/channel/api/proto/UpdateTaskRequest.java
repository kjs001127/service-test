// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: task/task.proto

package io.channel.api.proto;

/**
 * Protobuf type {@code task.UpdateTaskRequest}
 */
public final class UpdateTaskRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:task.UpdateTaskRequest)
    UpdateTaskRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use UpdateTaskRequest.newBuilder() to construct.
  private UpdateTaskRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private UpdateTaskRequest() {
    queueName_ = 0;
    id_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new UpdateTaskRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.channel.api.proto.TaskOuterClass.internal_static_task_UpdateTaskRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.channel.api.proto.TaskOuterClass.internal_static_task_UpdateTaskRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.channel.api.proto.UpdateTaskRequest.class, io.channel.api.proto.UpdateTaskRequest.Builder.class);
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

  public static final int ID_FIELD_NUMBER = 2;
  private volatile java.lang.Object id_;
  /**
   * <code>string id = 2;</code>
   * @return The id.
   */
  @java.lang.Override
  public java.lang.String getId() {
    java.lang.Object ref = id_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      id_ = s;
      return s;
    }
  }
  /**
   * <code>string id = 2;</code>
   * @return The bytes for id.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getIdBytes() {
    java.lang.Object ref = id_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      id_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int DATA_FIELD_NUMBER = 3;
  private com.google.protobuf.Duration data_;
  /**
   * <code>.google.protobuf.Duration data = 3;</code>
   * @return Whether the data field is set.
   */
  @java.lang.Override
  public boolean hasData() {
    return data_ != null;
  }
  /**
   * <code>.google.protobuf.Duration data = 3;</code>
   * @return The data.
   */
  @java.lang.Override
  public com.google.protobuf.Duration getData() {
    return data_ == null ? com.google.protobuf.Duration.getDefaultInstance() : data_;
  }
  /**
   * <code>.google.protobuf.Duration data = 3;</code>
   */
  @java.lang.Override
  public com.google.protobuf.DurationOrBuilder getDataOrBuilder() {
    return getData();
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
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(id_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, id_);
    }
    if (data_ != null) {
      output.writeMessage(3, getData());
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
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(id_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, id_);
    }
    if (data_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(3, getData());
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
    if (!(obj instanceof io.channel.api.proto.UpdateTaskRequest)) {
      return super.equals(obj);
    }
    io.channel.api.proto.UpdateTaskRequest other = (io.channel.api.proto.UpdateTaskRequest) obj;

    if (queueName_ != other.queueName_) return false;
    if (!getId()
        .equals(other.getId())) return false;
    if (hasData() != other.hasData()) return false;
    if (hasData()) {
      if (!getData()
          .equals(other.getData())) return false;
    }
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
    hash = (37 * hash) + ID_FIELD_NUMBER;
    hash = (53 * hash) + getId().hashCode();
    if (hasData()) {
      hash = (37 * hash) + DATA_FIELD_NUMBER;
      hash = (53 * hash) + getData().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.UpdateTaskRequest parseFrom(
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
  public static Builder newBuilder(io.channel.api.proto.UpdateTaskRequest prototype) {
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
   * Protobuf type {@code task.UpdateTaskRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:task.UpdateTaskRequest)
      io.channel.api.proto.UpdateTaskRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.channel.api.proto.TaskOuterClass.internal_static_task_UpdateTaskRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.channel.api.proto.TaskOuterClass.internal_static_task_UpdateTaskRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.channel.api.proto.UpdateTaskRequest.class, io.channel.api.proto.UpdateTaskRequest.Builder.class);
    }

    // Construct using io.channel.api.proto.UpdateTaskRequest.newBuilder()
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

      id_ = "";

      if (dataBuilder_ == null) {
        data_ = null;
      } else {
        data_ = null;
        dataBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.channel.api.proto.TaskOuterClass.internal_static_task_UpdateTaskRequest_descriptor;
    }

    @java.lang.Override
    public io.channel.api.proto.UpdateTaskRequest getDefaultInstanceForType() {
      return io.channel.api.proto.UpdateTaskRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.channel.api.proto.UpdateTaskRequest build() {
      io.channel.api.proto.UpdateTaskRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.channel.api.proto.UpdateTaskRequest buildPartial() {
      io.channel.api.proto.UpdateTaskRequest result = new io.channel.api.proto.UpdateTaskRequest(this);
      result.queueName_ = queueName_;
      result.id_ = id_;
      if (dataBuilder_ == null) {
        result.data_ = data_;
      } else {
        result.data_ = dataBuilder_.build();
      }
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
      if (other instanceof io.channel.api.proto.UpdateTaskRequest) {
        return mergeFrom((io.channel.api.proto.UpdateTaskRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.channel.api.proto.UpdateTaskRequest other) {
      if (other == io.channel.api.proto.UpdateTaskRequest.getDefaultInstance()) return this;
      if (other.queueName_ != 0) {
        setQueueNameValue(other.getQueueNameValue());
      }
      if (!other.getId().isEmpty()) {
        id_ = other.id_;
        onChanged();
      }
      if (other.hasData()) {
        mergeData(other.getData());
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
              id_ = input.readStringRequireUtf8();

              break;
            } // case 18
            case 26: {
              input.readMessage(
                  getDataFieldBuilder().getBuilder(),
                  extensionRegistry);

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

    private java.lang.Object id_ = "";
    /**
     * <code>string id = 2;</code>
     * @return The id.
     */
    public java.lang.String getId() {
      java.lang.Object ref = id_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        id_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string id = 2;</code>
     * @return The bytes for id.
     */
    public com.google.protobuf.ByteString
        getIdBytes() {
      java.lang.Object ref = id_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        id_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string id = 2;</code>
     * @param value The id to set.
     * @return This builder for chaining.
     */
    public Builder setId(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      id_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string id = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearId() {
      
      id_ = getDefaultInstance().getId();
      onChanged();
      return this;
    }
    /**
     * <code>string id = 2;</code>
     * @param value The bytes for id to set.
     * @return This builder for chaining.
     */
    public Builder setIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      id_ = value;
      onChanged();
      return this;
    }

    private com.google.protobuf.Duration data_;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.google.protobuf.Duration, com.google.protobuf.Duration.Builder, com.google.protobuf.DurationOrBuilder> dataBuilder_;
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     * @return Whether the data field is set.
     */
    public boolean hasData() {
      return dataBuilder_ != null || data_ != null;
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     * @return The data.
     */
    public com.google.protobuf.Duration getData() {
      if (dataBuilder_ == null) {
        return data_ == null ? com.google.protobuf.Duration.getDefaultInstance() : data_;
      } else {
        return dataBuilder_.getMessage();
      }
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    public Builder setData(com.google.protobuf.Duration value) {
      if (dataBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        data_ = value;
        onChanged();
      } else {
        dataBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    public Builder setData(
        com.google.protobuf.Duration.Builder builderForValue) {
      if (dataBuilder_ == null) {
        data_ = builderForValue.build();
        onChanged();
      } else {
        dataBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    public Builder mergeData(com.google.protobuf.Duration value) {
      if (dataBuilder_ == null) {
        if (data_ != null) {
          data_ =
            com.google.protobuf.Duration.newBuilder(data_).mergeFrom(value).buildPartial();
        } else {
          data_ = value;
        }
        onChanged();
      } else {
        dataBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    public Builder clearData() {
      if (dataBuilder_ == null) {
        data_ = null;
        onChanged();
      } else {
        data_ = null;
        dataBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    public com.google.protobuf.Duration.Builder getDataBuilder() {
      
      onChanged();
      return getDataFieldBuilder().getBuilder();
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    public com.google.protobuf.DurationOrBuilder getDataOrBuilder() {
      if (dataBuilder_ != null) {
        return dataBuilder_.getMessageOrBuilder();
      } else {
        return data_ == null ?
            com.google.protobuf.Duration.getDefaultInstance() : data_;
      }
    }
    /**
     * <code>.google.protobuf.Duration data = 3;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.google.protobuf.Duration, com.google.protobuf.Duration.Builder, com.google.protobuf.DurationOrBuilder> 
        getDataFieldBuilder() {
      if (dataBuilder_ == null) {
        dataBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.google.protobuf.Duration, com.google.protobuf.Duration.Builder, com.google.protobuf.DurationOrBuilder>(
                getData(),
                getParentForChildren(),
                isClean());
        data_ = null;
      }
      return dataBuilder_;
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


    // @@protoc_insertion_point(builder_scope:task.UpdateTaskRequest)
  }

  // @@protoc_insertion_point(class_scope:task.UpdateTaskRequest)
  private static final io.channel.api.proto.UpdateTaskRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.channel.api.proto.UpdateTaskRequest();
  }

  public static io.channel.api.proto.UpdateTaskRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<UpdateTaskRequest>
      PARSER = new com.google.protobuf.AbstractParser<UpdateTaskRequest>() {
    @java.lang.Override
    public UpdateTaskRequest parsePartialFrom(
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

  public static com.google.protobuf.Parser<UpdateTaskRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<UpdateTaskRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.channel.api.proto.UpdateTaskRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

