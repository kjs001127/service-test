// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: meet/meet.proto

package io.channel.api.proto;

public final class Meet {
  private Meet() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_AudioFile_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_AudioFile_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_Person_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_Person_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_Peer_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_Peer_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_BareResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_BareResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_InboundMeetRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_InboundMeetRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_JoinMeetByUserRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_JoinMeetByUserRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_HangUpMeetByManagerRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_HangUpMeetByManagerRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_GetGreetingRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_GetGreetingRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_CreateMeetRecordRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_CreateMeetRecordRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_OutboundMeetRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_OutboundMeetRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_InboundMeetResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_InboundMeetResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_PrivateMeetRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_PrivateMeetRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_JoinMeetByManagerRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_JoinMeetByManagerRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_JoinMeetByUserResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_JoinMeetByUserResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_TerminateMeetRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_TerminateMeetRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_HangUpMeetByUserRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_HangUpMeetByUserRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_meet_GetGreetingResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_meet_GetGreetingResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\017meet/meet.proto\022\004meet\"l\n\tAudioFile\022\016\n\006" +
      "bucket\030\001 \001(\t\022\013\n\003key\030\002 \001(\t\022\014\n\004name\030\003 \001(\t\022" +
      "\024\n\014content_type\030\004 \001(\t\022\020\n\010duration\030\005 \001(\001\022" +
      "\014\n\004size\030\006 \001(\003\"\"\n\006Person\022\014\n\004type\030\001 \001(\t\022\n\n" +
      "\002id\030\002 \001(\t\"7\n\004Peer\022\034\n\006person\030\001 \001(\0132\014.meet" +
      ".Person\022\021\n\tdevice_id\030\002 \001(\t\"9\n\014BareRespon" +
      "se\022)\n\rresponse_code\030\001 \001(\0162\022.meet.Respons" +
      "eCode\"i\n\022InboundMeetRequest\022\014\n\004from\030\001 \001(" +
      "\t\022\n\n\002to\030\002 \001(\t\022\017\n\007carrier\030\003 \001(\t\022\025\n\rsfu_se" +
      "rver_id\030\004 \001(\t\022\021\n\tdevice_id\030\005 \001(\t\"B\n\025Join" +
      "MeetByUserRequest\022\030\n\004peer\030\001 \001(\0132\n.meet.P" +
      "eer\022\017\n\007meet_id\030\002 \001(\t\"G\n\032HangUpMeetByMana" +
      "gerRequest\022\030\n\004peer\030\001 \001(\0132\n.meet.Peer\022\017\n\007" +
      "meet_id\030\002 \001(\t\" \n\022GetGreetingRequest\022\n\n\002t" +
      "o\030\001 \001(\t\"P\n\027CreateMeetRecordRequest\022\017\n\007me" +
      "et_id\030\001 \001(\t\022$\n\013meet_record\030\002 \001(\0132\017.meet." +
      "AudioFile\"\243\001\n\023OutboundMeetRequest\022\017\n\007mee" +
      "t_id\030\001 \001(\t\022\014\n\004from\030\002 \001(\t\022\n\n\002to\030\003 \001(\t\022\017\n\007" +
      "carrier\030\004 \001(\t\022\032\n\004user\030\005 \001(\0132\014.meet.Perso" +
      "n\022\033\n\007manager\030\006 \001(\0132\n.meet.Peer\022\027\n\017guide_" +
      "voice_url\030\007 \001(\t\"\235\001\n\023InboundMeetResponse\022" +
      ")\n\rresponse_code\030\001 \001(\0162\022.meet.ResponseCo" +
      "de\022\017\n\007meet_id\030\002 \001(\t\022\030\n\004peer\030\003 \001(\0132\n.meet" +
      ".Peer\022\034\n\017guide_voice_url\030\004 \001(\tH\000\210\001\001B\022\n\020_" +
      "guide_voice_url\"B\n\022PrivateMeetRequest\022\017\n" +
      "\007meet_id\030\001 \001(\t\022\033\n\007manager\030\002 \001(\0132\n.meet.P" +
      "eer\"E\n\030JoinMeetByManagerRequest\022\030\n\004peer\030" +
      "\001 \001(\0132\n.meet.Peer\022\017\n\007meet_id\030\002 \001(\t\"]\n\026Jo" +
      "inMeetByUserResponse\022)\n\rresponse_code\030\001 " +
      "\001(\0162\022.meet.ResponseCode\022\030\n\004peer\030\002 \001(\0132\n." +
      "meet.Peer\"^\n\024TerminateMeetRequest\022\017\n\007mee" +
      "t_id\030\001 \001(\t\022!\n\004code\030\002 \001(\0162\023.meet.CloseMee" +
      "tCode\022\022\n\nchannel_id\030\003 \001(\t\"D\n\027HangUpMeetB" +
      "yUserRequest\022\030\n\004peer\030\001 \001(\0132\n.meet.Peer\022\017" +
      "\n\007meet_id\030\002 \001(\t\".\n\023GetGreetingResponse\022\027" +
      "\n\017guide_voice_url\030\001 \001(\t*?\n\rCloseMeetCode" +
      "\022\014\n\010COMPLETE\020\000\022\024\n\020NOT_IN_OPERATION\020\001\022\n\n\006" +
      "MISSED\020\002*X\n\014ResponseCode\022\013\n\007SUCCESS\020\000\022\020\n" +
      "\014UNAUTHORIZED\020\001\022\r\n\tFORBIDDEN\020\002\022\r\n\tNOT_FO" +
      "UND\020\003\022\013\n\007UNKNOWN\020\0042\313\005\n\013MeetService\022C\n\022Cr" +
      "eateOutboundMeet\022\031.meet.OutboundMeetRequ" +
      "est\032\022.meet.BareResponse\022A\n\021CreatePrivate" +
      "Meet\022\030.meet.PrivateMeetRequest\032\022.meet.Ba" +
      "reResponse\022K\n\023HangUpMeetByManager\022 .meet" +
      ".HangUpMeetByManagerRequest\032\022.meet.BareR" +
      "esponse\022G\n\021JoinMeetByManager\022\036.meet.Join" +
      "MeetByManagerRequest\032\022.meet.BareResponse" +
      "\022?\n\rTerminateMeet\022\032.meet.TerminateMeetRe" +
      "quest\032\022.meet.BareResponse\022H\n\021CreateInbou" +
      "ndMeet\022\030.meet.InboundMeetRequest\032\031.meet." +
      "InboundMeetResponse\022E\n\020CreateMeetRecord\022" +
      "\035.meet.CreateMeetRecordRequest\032\022.meet.Ba" +
      "reResponse\022B\n\013GetGreeting\022\030.meet.GetGree" +
      "tingRequest\032\031.meet.GetGreetingResponse\022A" +
      "\n\016JoinMeetByUser\022\033.meet.JoinMeetByUserRe" +
      "quest\032\022.meet.BareResponse\022E\n\020HangUpMeetB" +
      "yUser\022\035.meet.HangUpMeetByUserRequest\032\022.m" +
      "eet.BareResponseB=\n\024io.channel.api.proto" +
      "P\001Z#github.com/channel-io/ch-proto/meetb" +
      "\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_meet_AudioFile_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_meet_AudioFile_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_AudioFile_descriptor,
        new java.lang.String[] { "Bucket", "Key", "Name", "ContentType", "Duration", "Size", });
    internal_static_meet_Person_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_meet_Person_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_Person_descriptor,
        new java.lang.String[] { "Type", "Id", });
    internal_static_meet_Peer_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_meet_Peer_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_Peer_descriptor,
        new java.lang.String[] { "Person", "DeviceId", });
    internal_static_meet_BareResponse_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_meet_BareResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_BareResponse_descriptor,
        new java.lang.String[] { "ResponseCode", });
    internal_static_meet_InboundMeetRequest_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_meet_InboundMeetRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_InboundMeetRequest_descriptor,
        new java.lang.String[] { "From", "To", "Carrier", "SfuServerId", "DeviceId", });
    internal_static_meet_JoinMeetByUserRequest_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_meet_JoinMeetByUserRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_JoinMeetByUserRequest_descriptor,
        new java.lang.String[] { "Peer", "MeetId", });
    internal_static_meet_HangUpMeetByManagerRequest_descriptor =
      getDescriptor().getMessageTypes().get(6);
    internal_static_meet_HangUpMeetByManagerRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_HangUpMeetByManagerRequest_descriptor,
        new java.lang.String[] { "Peer", "MeetId", });
    internal_static_meet_GetGreetingRequest_descriptor =
      getDescriptor().getMessageTypes().get(7);
    internal_static_meet_GetGreetingRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_GetGreetingRequest_descriptor,
        new java.lang.String[] { "To", });
    internal_static_meet_CreateMeetRecordRequest_descriptor =
      getDescriptor().getMessageTypes().get(8);
    internal_static_meet_CreateMeetRecordRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_CreateMeetRecordRequest_descriptor,
        new java.lang.String[] { "MeetId", "MeetRecord", });
    internal_static_meet_OutboundMeetRequest_descriptor =
      getDescriptor().getMessageTypes().get(9);
    internal_static_meet_OutboundMeetRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_OutboundMeetRequest_descriptor,
        new java.lang.String[] { "MeetId", "From", "To", "Carrier", "User", "Manager", "GuideVoiceUrl", });
    internal_static_meet_InboundMeetResponse_descriptor =
      getDescriptor().getMessageTypes().get(10);
    internal_static_meet_InboundMeetResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_InboundMeetResponse_descriptor,
        new java.lang.String[] { "ResponseCode", "MeetId", "Peer", "GuideVoiceUrl", "GuideVoiceUrl", });
    internal_static_meet_PrivateMeetRequest_descriptor =
      getDescriptor().getMessageTypes().get(11);
    internal_static_meet_PrivateMeetRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_PrivateMeetRequest_descriptor,
        new java.lang.String[] { "MeetId", "Manager", });
    internal_static_meet_JoinMeetByManagerRequest_descriptor =
      getDescriptor().getMessageTypes().get(12);
    internal_static_meet_JoinMeetByManagerRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_JoinMeetByManagerRequest_descriptor,
        new java.lang.String[] { "Peer", "MeetId", });
    internal_static_meet_JoinMeetByUserResponse_descriptor =
      getDescriptor().getMessageTypes().get(13);
    internal_static_meet_JoinMeetByUserResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_JoinMeetByUserResponse_descriptor,
        new java.lang.String[] { "ResponseCode", "Peer", });
    internal_static_meet_TerminateMeetRequest_descriptor =
      getDescriptor().getMessageTypes().get(14);
    internal_static_meet_TerminateMeetRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_TerminateMeetRequest_descriptor,
        new java.lang.String[] { "MeetId", "Code", "ChannelId", });
    internal_static_meet_HangUpMeetByUserRequest_descriptor =
      getDescriptor().getMessageTypes().get(15);
    internal_static_meet_HangUpMeetByUserRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_HangUpMeetByUserRequest_descriptor,
        new java.lang.String[] { "Peer", "MeetId", });
    internal_static_meet_GetGreetingResponse_descriptor =
      getDescriptor().getMessageTypes().get(16);
    internal_static_meet_GetGreetingResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_meet_GetGreetingResponse_descriptor,
        new java.lang.String[] { "GuideVoiceUrl", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
