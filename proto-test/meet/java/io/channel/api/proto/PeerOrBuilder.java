// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

public interface PeerOrBuilder extends
    // @@protoc_insertion_point(interface_extends:meet.Peer)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.meet.Person person = 1;</code>
   * @return Whether the person field is set.
   */
  boolean hasPerson();
  /**
   * <code>.meet.Person person = 1;</code>
   * @return The person.
   */
  io.channel.api.proto.Person getPerson();
  /**
   * <code>.meet.Person person = 1;</code>
   */
  io.channel.api.proto.PersonOrBuilder getPersonOrBuilder();

  /**
   * <code>string device_id = 2;</code>
   * @return The deviceId.
   */
  java.lang.String getDeviceId();
  /**
   * <code>string device_id = 2;</code>
   * @return The bytes for deviceId.
   */
  com.google.protobuf.ByteString
      getDeviceIdBytes();
}
