// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: fs_gw.proto

package cn.iopig.feed.scale.grpc;

public interface PigstyInfoResOrBuilder extends
    // @@protoc_insertion_point(interface_extends:fsapi.PigstyInfoRes)
    com.google.protobuf.MessageLiteOrBuilder {

  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  boolean hasReqHeader();
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  cn.iopig.feed.scale.grpc.ReqHeader getReqHeader();

  /**
   * <pre>
   *猪场ID
   * </pre>
   *
   * <code>optional string PigFarmId = 2;</code>
   */
  java.lang.String getPigFarmId();
  /**
   * <pre>
   *猪场ID
   * </pre>
   *
   * <code>optional string PigFarmId = 2;</code>
   */
  com.google.protobuf.ByteString
      getPigFarmIdBytes();

  /**
   * <pre>
   *猪场主人
   * </pre>
   *
   * <code>optional string farmer = 3;</code>
   */
  java.lang.String getFarmer();
  /**
   * <pre>
   *猪场主人
   * </pre>
   *
   * <code>optional string farmer = 3;</code>
   */
  com.google.protobuf.ByteString
      getFarmerBytes();

  /**
   * <code>repeated .fsapi.PigHouseInfo pigHouseInfo = 4;</code>
   */
  java.util.List<cn.iopig.feed.scale.grpc.PigHouseInfo> 
      getPigHouseInfoList();
  /**
   * <code>repeated .fsapi.PigHouseInfo pigHouseInfo = 4;</code>
   */
  cn.iopig.feed.scale.grpc.PigHouseInfo getPigHouseInfo(int index);
  /**
   * <code>repeated .fsapi.PigHouseInfo pigHouseInfo = 4;</code>
   */
  int getPigHouseInfoCount();
}