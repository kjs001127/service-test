// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

/**
 * Protobuf type {@code meet.InitializeInboundMeetResponse}
 */
public final class InitializeInboundMeetResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:meet.InitializeInboundMeetResponse)
    InitializeInboundMeetResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use InitializeInboundMeetResponse.newBuilder() to construct.
  private InitializeInboundMeetResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private InitializeInboundMeetResponse() {
    guideVoiceUrl_ = "";
    meetId_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new InitializeInboundMeetResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private InitializeInboundMeetResponse(
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

            isMeetAvailable_ = input.readBool();
            break;
          }
          case 18: {
            java.lang.String s = input.readStringRequireUtf8();

            guideVoiceUrl_ = s;
            break;
          }
          case 26: {
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
    return io.channel.api.proto.Meet.internal_static_meet_InitializeInboundMeetResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.channel.api.proto.Meet.internal_static_meet_InitializeInboundMeetResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.channel.api.proto.InitializeInboundMeetResponse.class, io.channel.api.proto.InitializeInboundMeetResponse.Builder.class);
  }

  public static final int IS_MEET_AVAILABLE_FIELD_NUMBER = 1;
  private boolean isMeetAvailable_;
  /**
   * <code>bool is_meet_available = 1;</code>
   * @return The isMeetAvailable.
   */
  @java.lang.Override
  public boolean getIsMeetAvailable() {
    return isMeetAvailable_;
  }

  public static final int GUIDE_VOICE_URL_FIELD_NUMBER = 2;
  private volatile java.lang.Object guideVoiceUrl_;
  /**
   * <code>string guide_voice_url = 2;</code>
   * @return The guideVoiceUrl.
   */
  @java.lang.Override
  public java.lang.String getGuideVoiceUrl() {
    java.lang.Object ref = guideVoiceUrl_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      guideVoiceUrl_ = s;
      return s;
    }
  }
  /**
   * <code>string guide_voice_url = 2;</code>
   * @return The bytes for guideVoiceUrl.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getGuideVoiceUrlBytes() {
    java.lang.Object ref = guideVoiceUrl_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      guideVoiceUrl_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int MEET_ID_FIELD_NUMBER = 3;
  private volatile java.lang.Object meetId_;
  /**
   * <code>string meet_id = 3;</code>
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
   * <code>string meet_id = 3;</code>
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
    if (isMeetAvailable_ != false) {
      output.writeBool(1, isMeetAvailable_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(guideVoiceUrl_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, guideVoiceUrl_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(meetId_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, meetId_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (isMeetAvailable_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(1, isMeetAvailable_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(guideVoiceUrl_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, guideVoiceUrl_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(meetId_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, meetId_);
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
    if (!(obj instanceof io.channel.api.proto.InitializeInboundMeetResponse)) {
      return super.equals(obj);
    }
    io.channel.api.proto.InitializeInboundMeetResponse other = (io.channel.api.proto.InitializeInboundMeetResponse) obj;

    if (getIsMeetAvailable()
        != other.getIsMeetAvailable()) return false;
    if (!getGuideVoiceUrl()
        .equals(other.getGuideVoiceUrl())) return false;
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
    hash = (37 * hash) + IS_MEET_AVAILABLE_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getIsMeetAvailable());
    hash = (37 * hash) + GUIDE_VOICE_URL_FIELD_NUMBER;
    hash = (53 * hash) + getGuideVoiceUrl().hashCode();
    hash = (37 * hash) + MEET_ID_FIELD_NUMBER;
    hash = (53 * hash) + getMeetId().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.InitializeInboundMeetResponse parseFrom(
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
  public static Builder newBuilder(io.channel.api.proto.InitializeInboundMeetResponse prototype) {
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
   * Protobuf type {@code meet.InitializeInboundMeetResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:meet.InitializeInboundMeetResponse)
      io.channel.api.proto.InitializeInboundMeetResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.channel.api.proto.Meet.internal_static_meet_InitializeInboundMeetResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.channel.api.proto.Meet.internal_static_meet_InitializeInboundMeetResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.channel.api.proto.InitializeInboundMeetResponse.class, io.channel.api.proto.InitializeInboundMeetResponse.Builder.class);
    }

    // Construct using io.channel.api.proto.InitializeInboundMeetResponse.newBuilder()
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
      isMeetAvailable_ = false;

      guideVoiceUrl_ = "";

      meetId_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.channel.api.proto.Meet.internal_static_meet_InitializeInboundMeetResponse_descriptor;
    }

    @java.lang.Override
    public io.channel.api.proto.InitializeInboundMeetResponse getDefaultInstanceForType() {
      return io.channel.api.proto.InitializeInboundMeetResponse.getDefaultInstance();
    }

    @java.lang.Override
    public io.channel.api.proto.InitializeInboundMeetResponse build() {
      io.channel.api.proto.InitializeInboundMeetResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.channel.api.proto.InitializeInboundMeetResponse buildPartial() {
      io.channel.api.proto.InitializeInboundMeetResponse result = new io.channel.api.proto.InitializeInboundMeetResponse(this);
      result.isMeetAvailable_ = isMeetAvailable_;
      result.guideVoiceUrl_ = guideVoiceUrl_;
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
      if (other instanceof io.channel.api.proto.InitializeInboundMeetResponse) {
        return mergeFrom((io.channel.api.proto.InitializeInboundMeetResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.channel.api.proto.InitializeInboundMeetResponse other) {
      if (other == io.channel.api.proto.InitializeInboundMeetResponse.getDefaultInstance()) return this;
      if (other.getIsMeetAvailable() != false) {
        setIsMeetAvailable(other.getIsMeetAvailable());
      }
      if (!other.getGuideVoiceUrl().isEmpty()) {
        guideVoiceUrl_ = other.guideVoiceUrl_;
        onChanged();
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
      io.channel.api.proto.InitializeInboundMeetResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.channel.api.proto.InitializeInboundMeetResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private boolean isMeetAvailable_ ;
    /**
     * <code>bool is_meet_available = 1;</code>
     * @return The isMeetAvailable.
     */
    @java.lang.Override
    public boolean getIsMeetAvailable() {
      return isMeetAvailable_;
    }
    /**
     * <code>bool is_meet_available = 1;</code>
     * @param value The isMeetAvailable to set.
     * @return This builder for chaining.
     */
    public Builder setIsMeetAvailable(boolean value) {
      
      isMeetAvailable_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bool is_meet_available = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearIsMeetAvailable() {
      
      isMeetAvailable_ = false;
      onChanged();
      return this;
    }

    private java.lang.Object guideVoiceUrl_ = "";
    /**
     * <code>string guide_voice_url = 2;</code>
     * @return The guideVoiceUrl.
     */
    public java.lang.String getGuideVoiceUrl() {
      java.lang.Object ref = guideVoiceUrl_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        guideVoiceUrl_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string guide_voice_url = 2;</code>
     * @return The bytes for guideVoiceUrl.
     */
    public com.google.protobuf.ByteString
        getGuideVoiceUrlBytes() {
      java.lang.Object ref = guideVoiceUrl_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        guideVoiceUrl_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string guide_voice_url = 2;</code>
     * @param value The guideVoiceUrl to set.
     * @return This builder for chaining.
     */
    public Builder setGuideVoiceUrl(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      guideVoiceUrl_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string guide_voice_url = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearGuideVoiceUrl() {
      
      guideVoiceUrl_ = getDefaultInstance().getGuideVoiceUrl();
      onChanged();
      return this;
    }
    /**
     * <code>string guide_voice_url = 2;</code>
     * @param value The bytes for guideVoiceUrl to set.
     * @return This builder for chaining.
     */
    public Builder setGuideVoiceUrlBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      guideVoiceUrl_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object meetId_ = "";
    /**
     * <code>string meet_id = 3;</code>
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
     * <code>string meet_id = 3;</code>
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
     * <code>string meet_id = 3;</code>
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
     * <code>string meet_id = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearMeetId() {
      
      meetId_ = getDefaultInstance().getMeetId();
      onChanged();
      return this;
    }
    /**
     * <code>string meet_id = 3;</code>
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


    // @@protoc_insertion_point(builder_scope:meet.InitializeInboundMeetResponse)
  }

  // @@protoc_insertion_point(class_scope:meet.InitializeInboundMeetResponse)
  private static final io.channel.api.proto.InitializeInboundMeetResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.channel.api.proto.InitializeInboundMeetResponse();
  }

  public static io.channel.api.proto.InitializeInboundMeetResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<InitializeInboundMeetResponse>
      PARSER = new com.google.protobuf.AbstractParser<InitializeInboundMeetResponse>() {
    @java.lang.Override
    public InitializeInboundMeetResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new InitializeInboundMeetResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<InitializeInboundMeetResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<InitializeInboundMeetResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.channel.api.proto.InitializeInboundMeetResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

