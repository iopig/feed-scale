// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: fs_gw.proto

package cn.iopig.feed.scale.grpc;

public interface PigHouseInfoOrBuilder extends
    // @@protoc_insertion_point(interface_extends:fsapi.PigHouseInfo)
    com.google.protobuf.MessageLiteOrBuilder {

  /**
   * <pre>
   *猪舍ID
   * </pre>
   *
   * <code>optional string HouseId = 1;</code>
   */
  java.lang.String getHouseId();
  /**
   * <pre>
   *猪舍ID
   * </pre>
   *
   * <code>optional string HouseId = 1;</code>
   */
  com.google.protobuf.ByteString
      getHouseIdBytes();

  /**
   * <pre>
   *猪圈信息
   * </pre>
   *
   * <code>repeated .fsapi.PigstyInfo pigstyInfo = 2;</code>
   */
  java.util.List<cn.iopig.feed.scale.grpc.PigstyInfo> 
      getPigstyInfoList();
  /**
   * <pre>
   *猪圈信息
   * </pre>
   *
   * <code>repeated .fsapi.PigstyInfo pigstyInfo = 2;</code>
   */
  cn.iopig.feed.scale.grpc.PigstyInfo getPigstyInfo(int index);
  /**
   * <pre>
   *猪圈信息
   * </pre>
   *
   * <code>repeated .fsapi.PigstyInfo pigstyInfo = 2;</code>
   */
  int getPigstyInfoCount();
}
