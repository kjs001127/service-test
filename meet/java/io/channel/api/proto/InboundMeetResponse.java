// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

/**
 * Protobuf type {@code meet.InboundMeetResponse}
 */
public final class InboundMeetResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:meet.InboundMeetResponse)
    InboundMeetResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use InboundMeetResponse.newBuilder() to construct.
  private InboundMeetResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private InboundMeetResponse() {
    responseCode_ = 0;
    meetId_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new InboundMeetResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private InboundMeetResponse(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {
            int rawValue = input.readEnum();

            responseCode_ = rawValue;
            break;
          }
          case 18: {
            java.lang.String s = input.readStringRequireUtf8();

            meetId_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.channel.api.proto.Meet.internal_static_meet_InboundMeetResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.channel.api.proto.Meet.internal_static_meet_InboundMeetResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.channel.api.proto.InboundMeetResponse.class, io.channel.api.proto.InboundMeetResponse.Builder.class);
  }

  public static final int RESPONSE_CODE_FIELD_NUMBER = 1;
  private int responseCode_;
  /**
   * <code>.meet.ResponseCode response_code = 1;</code>
   * @return The enum numeric value on the wire for responseCode.
   */
  @java.lang.Override public int getResponseCodeValue() {
    return responseCode_;
  }
  /**
   * <code>.meet.ResponseCode response_code = 1;</code>
   * @return The responseCode.
   */
  @java.lang.Override public io.channel.api.proto.ResponseCode getResponseCode() {
    @SuppressWarnings("deprecation")
    io.channel.api.proto.ResponseCode result = io.channel.api.proto.ResponseCode.valueOf(responseCode_);
    return result == null ? io.channel.api.proto.ResponseCode.UNRECOGNIZED : result;
  }

  public static final int MEET_ID_FIELD_NUMBER = 2;
  private volatile java.lang.Object meetId_;
  /**
   * <code>string meet_id = 2;</code>
   * @return The meetId.
   */
  @java.lang.Override
  public java.lang.String getMeetId() {
    java.lang.Object ref = meetId_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      meetId_ = s;
      return s;
    }
  }
  /**
   * <code>string meet_id = 2;</code>
   * @return The bytes for meetId.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getMeetIdBytes() {
    java.lang.Object ref = meetId_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      meetId_ = b;
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
    if (responseCode_ != io.channel.api.proto.ResponseCode.SUCCESS.getNumber()) {
      output.writeEnum(1, responseCode_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(meetId_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, meetId_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (responseCode_ != io.channel.api.proto.ResponseCode.SUCCESS.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, responseCode_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(meetId_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, meetId_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.channel.api.proto.InboundMeetResponse)) {
      return super.equals(obj);
    }
    io.channel.api.proto.InboundMeetResponse other = (io.channel.api.proto.InboundMeetResponse) obj;

    if (responseCode_ != other.responseCode_) return false;
    if (!getMeetId()
        .equals(other.getMeetId())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + RESPONSE_CODE_FIELD_NUMBER;
    hash = (53 * hash) + responseCode_;
    hash = (37 * hash) + MEET_ID_FIELD_NUMBER;
    hash = (53 * hash) + getMeetId().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.InboundMeetResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.InboundMeetResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.InboundMeetResponse parseFrom(
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
  public static Builder newBuilder(io.channel.api.proto.InboundMeetResponse prototype) {
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
   * Protobuf type {@code meet.InboundMeetResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:meet.InboundMeetResponse)
      io.channel.api.proto.InboundMeetResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.channel.api.proto.Meet.internal_static_meet_InboundMeetResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.channel.api.proto.Meet.internal_static_meet_InboundMeetResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.channel.api.proto.InboundMeetResponse.class, io.channel.api.proto.InboundMeetResponse.Builder.class);
    }

    // Construct using io.channel.api.proto.InboundMeetResponse.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      responseCode_ = 0;

      meetId_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.channel.api.proto.Meet.internal_static_meet_InboundMeetResponse_descriptor;
    }

    @java.lang.Override
    public io.channel.api.proto.InboundMeetResponse getDefaultInstanceForType() {
      return io.channel.api.proto.InboundMeetResponse.getDefaultInstance();
    }

    @java.lang.Override
    public io.channel.api.proto.InboundMeetResponse build() {
      io.channel.api.proto.InboundMeetResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.channel.api.proto.InboundMeetResponse buildPartial() {
      io.channel.api.proto.InboundMeetResponse result = new io.channel.api.proto.InboundMeetResponse(this);
      result.responseCode_ = responseCode_;
      result.meetId_ = meetId_;
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
      if (other instanceof io.channel.api.proto.InboundMeetResponse) {
        return mergeFrom((io.channel.api.proto.InboundMeetResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.channel.api.proto.InboundMeetResponse other) {
      if (other == io.channel.api.proto.InboundMeetResponse.getDefaultInstance()) return this;
      if (other.responseCode_ != 0) {
        setResponseCodeValue(other.getResponseCodeValue());
      }
      if (!other.getMeetId().isEmpty()) {
        meetId_ = other.meetId_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
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
      io.channel.api.proto.InboundMeetResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.channel.api.proto.InboundMeetResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int responseCode_ = 0;
    /**
     * <code>.meet.ResponseCode response_code = 1;</code>
     * @return The enum numeric value on the wire for responseCode.
     */
    @java.lang.Override public int getResponseCodeValue() {
      return responseCode_;
    }
    /**
     * <code>.meet.ResponseCode response_code = 1;</code>
     * @param value The enum numeric value on the wire for responseCode to set.
     * @return This builder for chaining.
     */
    public Builder setResponseCodeValue(int value) {
      
      responseCode_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.meet.ResponseCode response_code = 1;</code>
     * @return The responseCode.
     */
    @java.lang.Override
    public io.channel.api.proto.ResponseCode getResponseCode() {
      @SuppressWarnings("deprecation")
      io.channel.api.proto.ResponseCode result = io.channel.api.proto.ResponseCode.valueOf(responseCode_);
      return result == null ? io.channel.api.proto.ResponseCode.UNRECOGNIZED : result;
    }
    /**
     * <code>.meet.ResponseCode response_code = 1;</code>
     * @param value The responseCode to set.
     * @return This builder for chaining.
     */
    public Builder setResponseCode(io.channel.api.proto.ResponseCode value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      responseCode_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.meet.ResponseCode response_code = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearResponseCode() {
      
      responseCode_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object meetId_ = "";
    /**
     * <code>string meet_id = 2;</code>
     * @return The meetId.
     */
    public java.lang.String getMeetId() {
      java.lang.Object ref = meetId_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        meetId_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string meet_id = 2;</code>
     * @return The bytes for meetId.
     */
    public com.google.protobuf.ByteString
        getMeetIdBytes() {
      java.lang.Object ref = meetId_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        meetId_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string meet_id = 2;</code>
     * @param value The meetId to set.
     * @return This builder for chaining.
     */
    public Builder setMeetId(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      meetId_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string meet_id = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearMeetId() {
      
      meetId_ = getDefaultInstance().getMeetId();
      onChanged();
      return this;
    }
    /**
     * <code>string meet_id = 2;</code>
     * @param value The bytes for meetId to set.
     * @return This builder for chaining.
     */
    public Builder setMeetIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      meetId_ = value;
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


    // @@protoc_insertion_point(builder_scope:meet.InboundMeetResponse)
  }

  // @@protoc_insertion_point(class_scope:meet.InboundMeetResponse)
  private static final io.channel.api.proto.InboundMeetResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.channel.api.proto.InboundMeetResponse();
  }

  public static io.channel.api.proto.InboundMeetResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<InboundMeetResponse>
      PARSER = new com.google.protobuf.AbstractParser<InboundMeetResponse>() {
    @java.lang.Override
    public InboundMeetResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new InboundMeetResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<InboundMeetResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<InboundMeetResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.channel.api.proto.InboundMeetResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

