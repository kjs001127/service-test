// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

/**
 * Protobuf type {@code meet.TerminateMeetRequest}
 */
public final class TerminateMeetRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:meet.TerminateMeetRequest)
    TerminateMeetRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use TerminateMeetRequest.newBuilder() to construct.
  private TerminateMeetRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private TerminateMeetRequest() {
    meetId_ = "";
    channelId_ = "";
    guideVoiceUrl_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new TerminateMeetRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private TerminateMeetRequest(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
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
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            meetId_ = s;
            break;
          }
          case 18: {
            java.lang.String s = input.readStringRequireUtf8();

            channelId_ = s;
            break;
          }
          case 26: {
            java.lang.String s = input.readStringRequireUtf8();
            bitField0_ |= 0x00000001;
            guideVoiceUrl_ = s;
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
    return io.channel.api.proto.Meet.internal_static_meet_TerminateMeetRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.channel.api.proto.Meet.internal_static_meet_TerminateMeetRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.channel.api.proto.TerminateMeetRequest.class, io.channel.api.proto.TerminateMeetRequest.Builder.class);
  }

  private int bitField0_;
  public static final int MEET_ID_FIELD_NUMBER = 1;
  private volatile java.lang.Object meetId_;
  /**
   * <code>string meet_id = 1;</code>
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
   * <code>string meet_id = 1;</code>
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

  public static final int CHANNEL_ID_FIELD_NUMBER = 2;
  private volatile java.lang.Object channelId_;
  /**
   * <code>string channel_id = 2;</code>
   * @return The channelId.
   */
  @java.lang.Override
  public java.lang.String getChannelId() {
    java.lang.Object ref = channelId_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      channelId_ = s;
      return s;
    }
  }
  /**
   * <code>string channel_id = 2;</code>
   * @return The bytes for channelId.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getChannelIdBytes() {
    java.lang.Object ref = channelId_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      channelId_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int GUIDE_VOICE_URL_FIELD_NUMBER = 3;
  private volatile java.lang.Object guideVoiceUrl_;
  /**
   * <code>optional string guide_voice_url = 3;</code>
   * @return Whether the guideVoiceUrl field is set.
   */
  @java.lang.Override
  public boolean hasGuideVoiceUrl() {
    return ((bitField0_ & 0x00000001) != 0);
  }
  /**
   * <code>optional string guide_voice_url = 3;</code>
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
   * <code>optional string guide_voice_url = 3;</code>
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
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(meetId_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, meetId_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(channelId_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, channelId_);
    }
    if (((bitField0_ & 0x00000001) != 0)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, guideVoiceUrl_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(meetId_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, meetId_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(channelId_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, channelId_);
    }
    if (((bitField0_ & 0x00000001) != 0)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, guideVoiceUrl_);
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
    if (!(obj instanceof io.channel.api.proto.TerminateMeetRequest)) {
      return super.equals(obj);
    }
    io.channel.api.proto.TerminateMeetRequest other = (io.channel.api.proto.TerminateMeetRequest) obj;

    if (!getMeetId()
        .equals(other.getMeetId())) return false;
    if (!getChannelId()
        .equals(other.getChannelId())) return false;
    if (hasGuideVoiceUrl() != other.hasGuideVoiceUrl()) return false;
    if (hasGuideVoiceUrl()) {
      if (!getGuideVoiceUrl()
          .equals(other.getGuideVoiceUrl())) return false;
    }
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
    hash = (37 * hash) + MEET_ID_FIELD_NUMBER;
    hash = (53 * hash) + getMeetId().hashCode();
    hash = (37 * hash) + CHANNEL_ID_FIELD_NUMBER;
    hash = (53 * hash) + getChannelId().hashCode();
    if (hasGuideVoiceUrl()) {
      hash = (37 * hash) + GUIDE_VOICE_URL_FIELD_NUMBER;
      hash = (53 * hash) + getGuideVoiceUrl().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.channel.api.proto.TerminateMeetRequest parseFrom(
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
  public static Builder newBuilder(io.channel.api.proto.TerminateMeetRequest prototype) {
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
   * Protobuf type {@code meet.TerminateMeetRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:meet.TerminateMeetRequest)
      io.channel.api.proto.TerminateMeetRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.channel.api.proto.Meet.internal_static_meet_TerminateMeetRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.channel.api.proto.Meet.internal_static_meet_TerminateMeetRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.channel.api.proto.TerminateMeetRequest.class, io.channel.api.proto.TerminateMeetRequest.Builder.class);
    }

    // Construct using io.channel.api.proto.TerminateMeetRequest.newBuilder()
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
      meetId_ = "";

      channelId_ = "";

      guideVoiceUrl_ = "";
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.channel.api.proto.Meet.internal_static_meet_TerminateMeetRequest_descriptor;
    }

    @java.lang.Override
    public io.channel.api.proto.TerminateMeetRequest getDefaultInstanceForType() {
      return io.channel.api.proto.TerminateMeetRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.channel.api.proto.TerminateMeetRequest build() {
      io.channel.api.proto.TerminateMeetRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.channel.api.proto.TerminateMeetRequest buildPartial() {
      io.channel.api.proto.TerminateMeetRequest result = new io.channel.api.proto.TerminateMeetRequest(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      result.meetId_ = meetId_;
      result.channelId_ = channelId_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        to_bitField0_ |= 0x00000001;
      }
      result.guideVoiceUrl_ = guideVoiceUrl_;
      result.bitField0_ = to_bitField0_;
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
      if (other instanceof io.channel.api.proto.TerminateMeetRequest) {
        return mergeFrom((io.channel.api.proto.TerminateMeetRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.channel.api.proto.TerminateMeetRequest other) {
      if (other == io.channel.api.proto.TerminateMeetRequest.getDefaultInstance()) return this;
      if (!other.getMeetId().isEmpty()) {
        meetId_ = other.meetId_;
        onChanged();
      }
      if (!other.getChannelId().isEmpty()) {
        channelId_ = other.channelId_;
        onChanged();
      }
      if (other.hasGuideVoiceUrl()) {
        bitField0_ |= 0x00000001;
        guideVoiceUrl_ = other.guideVoiceUrl_;
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
      io.channel.api.proto.TerminateMeetRequest parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.channel.api.proto.TerminateMeetRequest) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object meetId_ = "";
    /**
     * <code>string meet_id = 1;</code>
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
     * <code>string meet_id = 1;</code>
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
     * <code>string meet_id = 1;</code>
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
     * <code>string meet_id = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearMeetId() {
      
      meetId_ = getDefaultInstance().getMeetId();
      onChanged();
      return this;
    }
    /**
     * <code>string meet_id = 1;</code>
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

    private java.lang.Object channelId_ = "";
    /**
     * <code>string channel_id = 2;</code>
     * @return The channelId.
     */
    public java.lang.String getChannelId() {
      java.lang.Object ref = channelId_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        channelId_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string channel_id = 2;</code>
     * @return The bytes for channelId.
     */
    public com.google.protobuf.ByteString
        getChannelIdBytes() {
      java.lang.Object ref = channelId_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        channelId_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string channel_id = 2;</code>
     * @param value The channelId to set.
     * @return This builder for chaining.
     */
    public Builder setChannelId(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      channelId_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string channel_id = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearChannelId() {
      
      channelId_ = getDefaultInstance().getChannelId();
      onChanged();
      return this;
    }
    /**
     * <code>string channel_id = 2;</code>
     * @param value The bytes for channelId to set.
     * @return This builder for chaining.
     */
    public Builder setChannelIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      channelId_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object guideVoiceUrl_ = "";
    /**
     * <code>optional string guide_voice_url = 3;</code>
     * @return Whether the guideVoiceUrl field is set.
     */
    public boolean hasGuideVoiceUrl() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>optional string guide_voice_url = 3;</code>
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
     * <code>optional string guide_voice_url = 3;</code>
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
     * <code>optional string guide_voice_url = 3;</code>
     * @param value The guideVoiceUrl to set.
     * @return This builder for chaining.
     */
    public Builder setGuideVoiceUrl(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  bitField0_ |= 0x00000001;
      guideVoiceUrl_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>optional string guide_voice_url = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearGuideVoiceUrl() {
      bitField0_ = (bitField0_ & ~0x00000001);
      guideVoiceUrl_ = getDefaultInstance().getGuideVoiceUrl();
      onChanged();
      return this;
    }
    /**
     * <code>optional string guide_voice_url = 3;</code>
     * @param value The bytes for guideVoiceUrl to set.
     * @return This builder for chaining.
     */
    public Builder setGuideVoiceUrlBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      bitField0_ |= 0x00000001;
      guideVoiceUrl_ = value;
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


    // @@protoc_insertion_point(builder_scope:meet.TerminateMeetRequest)
  }

  // @@protoc_insertion_point(class_scope:meet.TerminateMeetRequest)
  private static final io.channel.api.proto.TerminateMeetRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.channel.api.proto.TerminateMeetRequest();
  }

  public static io.channel.api.proto.TerminateMeetRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<TerminateMeetRequest>
      PARSER = new com.google.protobuf.AbstractParser<TerminateMeetRequest>() {
    @java.lang.Override
    public TerminateMeetRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new TerminateMeetRequest(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<TerminateMeetRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<TerminateMeetRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.channel.api.proto.TerminateMeetRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

