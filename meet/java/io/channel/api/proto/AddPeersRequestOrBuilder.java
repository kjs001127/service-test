// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

public interface AddPeersRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:meet.AddPeersRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .meet.Person persons = 1;</code>
   */
  java.util.List<io.channel.api.proto.Person> 
      getPersonsList();
  /**
   * <code>repeated .meet.Person persons = 1;</code>
   */
  io.channel.api.proto.Person getPersons(int index);
  /**
   * <code>repeated .meet.Person persons = 1;</code>
   */
  int getPersonsCount();
  /**
   * <code>repeated .meet.Person persons = 1;</code>
   */
  java.util.List<? extends io.channel.api.proto.PersonOrBuilder> 
      getPersonsOrBuilderList();
  /**
   * <code>repeated .meet.Person persons = 1;</code>
   */
  io.channel.api.proto.PersonOrBuilder getPersonsOrBuilder(
      int index);

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
