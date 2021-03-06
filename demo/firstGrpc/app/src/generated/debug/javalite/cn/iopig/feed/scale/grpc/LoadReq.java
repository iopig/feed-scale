// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: fs_gw.proto

package cn.iopig.feed.scale.grpc;

/**
 * <pre>
 *LoadReq
 * </pre>
 *
 * Protobuf type {@code fsapi.LoadReq}
 */
public  final class LoadReq extends
    com.google.protobuf.GeneratedMessageLite<
        LoadReq, LoadReq.Builder> implements
    // @@protoc_insertion_point(message_implements:fsapi.LoadReq)
    LoadReqOrBuilder {
  private LoadReq() {
    currentWeight_ = "";
  }
  public static final int REQHEADER_FIELD_NUMBER = 1;
  private cn.iopig.feed.scale.grpc.ReqHeader reqHeader_;
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  public boolean hasReqHeader() {
    return reqHeader_ != null;
  }
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  public cn.iopig.feed.scale.grpc.ReqHeader getReqHeader() {
    return reqHeader_ == null ? cn.iopig.feed.scale.grpc.ReqHeader.getDefaultInstance() : reqHeader_;
  }
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  private void setReqHeader(cn.iopig.feed.scale.grpc.ReqHeader value) {
    if (value == null) {
      throw new NullPointerException();
    }
    reqHeader_ = value;
    
    }
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  private void setReqHeader(
      cn.iopig.feed.scale.grpc.ReqHeader.Builder builderForValue) {
    reqHeader_ = builderForValue.build();
    
  }
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  private void mergeReqHeader(cn.iopig.feed.scale.grpc.ReqHeader value) {
    if (reqHeader_ != null &&
        reqHeader_ != cn.iopig.feed.scale.grpc.ReqHeader.getDefaultInstance()) {
      reqHeader_ =
        cn.iopig.feed.scale.grpc.ReqHeader.newBuilder(reqHeader_).mergeFrom(value).buildPartial();
    } else {
      reqHeader_ = value;
    }
    
  }
  /**
   * <pre>
   *Version ,Devid ,Timestamp
   * </pre>
   *
   * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
   */
  private void clearReqHeader() {  reqHeader_ = null;
    
  }

  public static final int CURRENTWEIGHT_FIELD_NUMBER = 2;
  private java.lang.String currentWeight_;
  /**
   * <pre>
   *当前重量
   * </pre>
   *
   * <code>optional string CurrentWeight = 2;</code>
   */
  public java.lang.String getCurrentWeight() {
    return currentWeight_;
  }
  /**
   * <pre>
   *当前重量
   * </pre>
   *
   * <code>optional string CurrentWeight = 2;</code>
   */
  public com.google.protobuf.ByteString
      getCurrentWeightBytes() {
    return com.google.protobuf.ByteString.copyFromUtf8(currentWeight_);
  }
  /**
   * <pre>
   *当前重量
   * </pre>
   *
   * <code>optional string CurrentWeight = 2;</code>
   */
  private void setCurrentWeight(
      java.lang.String value) {
    if (value == null) {
    throw new NullPointerException();
  }
  
    currentWeight_ = value;
  }
  /**
   * <pre>
   *当前重量
   * </pre>
   *
   * <code>optional string CurrentWeight = 2;</code>
   */
  private void clearCurrentWeight() {
    
    currentWeight_ = getDefaultInstance().getCurrentWeight();
  }
  /**
   * <pre>
   *当前重量
   * </pre>
   *
   * <code>optional string CurrentWeight = 2;</code>
   */
  private void setCurrentWeightBytes(
      com.google.protobuf.ByteString value) {
    if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
    
    currentWeight_ = value.toStringUtf8();
  }

  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (reqHeader_ != null) {
      output.writeMessage(1, getReqHeader());
    }
    if (!currentWeight_.isEmpty()) {
      output.writeString(2, getCurrentWeight());
    }
  }

  public int getSerializedSize() {
    int size = memoizedSerializedSize;
    if (size != -1) return size;

    size = 0;
    if (reqHeader_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getReqHeader());
    }
    if (!currentWeight_.isEmpty()) {
      size += com.google.protobuf.CodedOutputStream
        .computeStringSize(2, getCurrentWeight());
    }
    memoizedSerializedSize = size;
    return size;
  }

  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static cn.iopig.feed.scale.grpc.LoadReq parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }

  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(cn.iopig.feed.scale.grpc.LoadReq prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }

  /**
   * <pre>
   *LoadReq
   * </pre>
   *
   * Protobuf type {@code fsapi.LoadReq}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageLite.Builder<
        cn.iopig.feed.scale.grpc.LoadReq, Builder> implements
      // @@protoc_insertion_point(builder_implements:fsapi.LoadReq)
      cn.iopig.feed.scale.grpc.LoadReqOrBuilder {
    // Construct using cn.iopig.feed.scale.grpc.LoadReq.newBuilder()
    private Builder() {
      super(DEFAULT_INSTANCE);
    }


    /**
     * <pre>
     *Version ,Devid ,Timestamp
     * </pre>
     *
     * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
     */
    public boolean hasReqHeader() {
      return instance.hasReqHeader();
    }
    /**
     * <pre>
     *Version ,Devid ,Timestamp
     * </pre>
     *
     * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
     */
    public cn.iopig.feed.scale.grpc.ReqHeader getReqHeader() {
      return instance.getReqHeader();
    }
    /**
     * <pre>
     *Version ,Devid ,Timestamp
     * </pre>
     *
     * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
     */
    public Builder setReqHeader(cn.iopig.feed.scale.grpc.ReqHeader value) {
      copyOnWrite();
      instance.setReqHeader(value);
      return this;
      }
    /**
     * <pre>
     *Version ,Devid ,Timestamp
     * </pre>
     *
     * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
     */
    public Builder setReqHeader(
        cn.iopig.feed.scale.grpc.ReqHeader.Builder builderForValue) {
      copyOnWrite();
      instance.setReqHeader(builderForValue);
      return this;
    }
    /**
     * <pre>
     *Version ,Devid ,Timestamp
     * </pre>
     *
     * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
     */
    public Builder mergeReqHeader(cn.iopig.feed.scale.grpc.ReqHeader value) {
      copyOnWrite();
      instance.mergeReqHeader(value);
      return this;
    }
    /**
     * <pre>
     *Version ,Devid ,Timestamp
     * </pre>
     *
     * <code>optional .fsapi.ReqHeader ReqHeader = 1;</code>
     */
    public Builder clearReqHeader() {  copyOnWrite();
      instance.clearReqHeader();
      return this;
    }

    /**
     * <pre>
     *当前重量
     * </pre>
     *
     * <code>optional string CurrentWeight = 2;</code>
     */
    public java.lang.String getCurrentWeight() {
      return instance.getCurrentWeight();
    }
    /**
     * <pre>
     *当前重量
     * </pre>
     *
     * <code>optional string CurrentWeight = 2;</code>
     */
    public com.google.protobuf.ByteString
        getCurrentWeightBytes() {
      return instance.getCurrentWeightBytes();
    }
    /**
     * <pre>
     *当前重量
     * </pre>
     *
     * <code>optional string CurrentWeight = 2;</code>
     */
    public Builder setCurrentWeight(
        java.lang.String value) {
      copyOnWrite();
      instance.setCurrentWeight(value);
      return this;
    }
    /**
     * <pre>
     *当前重量
     * </pre>
     *
     * <code>optional string CurrentWeight = 2;</code>
     */
    public Builder clearCurrentWeight() {
      copyOnWrite();
      instance.clearCurrentWeight();
      return this;
    }
    /**
     * <pre>
     *当前重量
     * </pre>
     *
     * <code>optional string CurrentWeight = 2;</code>
     */
    public Builder setCurrentWeightBytes(
        com.google.protobuf.ByteString value) {
      copyOnWrite();
      instance.setCurrentWeightBytes(value);
      return this;
    }

    // @@protoc_insertion_point(builder_scope:fsapi.LoadReq)
  }
  protected final Object dynamicMethod(
      com.google.protobuf.GeneratedMessageLite.MethodToInvoke method,
      Object arg0, Object arg1) {
    switch (method) {
      case NEW_MUTABLE_INSTANCE: {
        return new cn.iopig.feed.scale.grpc.LoadReq();
      }
      case IS_INITIALIZED: {
        return DEFAULT_INSTANCE;
      }
      case MAKE_IMMUTABLE: {
        return null;
      }
      case NEW_BUILDER: {
        return new Builder();
      }
      case VISIT: {
        Visitor visitor = (Visitor) arg0;
        cn.iopig.feed.scale.grpc.LoadReq other = (cn.iopig.feed.scale.grpc.LoadReq) arg1;
        reqHeader_ = visitor.visitMessage(reqHeader_, other.reqHeader_);
        currentWeight_ = visitor.visitString(!currentWeight_.isEmpty(), currentWeight_,
            !other.currentWeight_.isEmpty(), other.currentWeight_);
        if (visitor == com.google.protobuf.GeneratedMessageLite.MergeFromVisitor
            .INSTANCE) {
        }
        return this;
      }
      case MERGE_FROM_STREAM: {
        com.google.protobuf.CodedInputStream input =
            (com.google.protobuf.CodedInputStream) arg0;
        com.google.protobuf.ExtensionRegistryLite extensionRegistry =
            (com.google.protobuf.ExtensionRegistryLite) arg1;
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              default: {
                if (!input.skipField(tag)) {
                  done = true;
                }
                break;
              }
              case 10: {
                cn.iopig.feed.scale.grpc.ReqHeader.Builder subBuilder = null;
                if (reqHeader_ != null) {
                  subBuilder = reqHeader_.toBuilder();
                }
                reqHeader_ = input.readMessage(cn.iopig.feed.scale.grpc.ReqHeader.parser(), extensionRegistry);
                if (subBuilder != null) {
                  subBuilder.mergeFrom(reqHeader_);
                  reqHeader_ = subBuilder.buildPartial();
                }

                break;
              }
              case 18: {
                String s = input.readStringRequireUtf8();

                currentWeight_ = s;
                break;
              }
            }
          }
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw new RuntimeException(e.setUnfinishedMessage(this));
        } catch (java.io.IOException e) {
          throw new RuntimeException(
              new com.google.protobuf.InvalidProtocolBufferException(
                  e.getMessage()).setUnfinishedMessage(this));
        } finally {
        }
      }
      case GET_DEFAULT_INSTANCE: {
        return DEFAULT_INSTANCE;
      }
      case GET_PARSER: {
        if (PARSER == null) {    synchronized (cn.iopig.feed.scale.grpc.LoadReq.class) {
            if (PARSER == null) {
              PARSER = new DefaultInstanceBasedParser(DEFAULT_INSTANCE);
            }
          }
        }
        return PARSER;
      }
    }
    throw new UnsupportedOperationException();
  }


  // @@protoc_insertion_point(class_scope:fsapi.LoadReq)
  private static final cn.iopig.feed.scale.grpc.LoadReq DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new LoadReq();
    DEFAULT_INSTANCE.makeImmutable();
  }

  public static cn.iopig.feed.scale.grpc.LoadReq getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static volatile com.google.protobuf.Parser<LoadReq> PARSER;

  public static com.google.protobuf.Parser<LoadReq> parser() {
    return DEFAULT_INSTANCE.getParserForType();
  }
}

