// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

public interface InboundMeetResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:meet.InboundMeetResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.meet.ResponseCode response_code = 1;</code>
   * @return The enum numeric value on the wire for responseCode.
   */
  int getResponseCodeValue();
  /**
   * <code>.meet.ResponseCode response_code = 1;</code>
   * @return The responseCode.
   */
  io.channel.api.proto.ResponseCode getResponseCode();

  /**
   * <code>string meet_id = 2;</code>
   * @return The meetId.
   */
  java.lang.String getMeetId();
  /**
   * <code>string meet_id = 2;</code>
   * @return The bytes for meetId.
   */
  com.google.protobuf.ByteString
      getMeetIdBytes();
}
