package cn.iopig.feed.scale.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 * <pre>
 * The greeting service definition.
 * interface between feed scale and pad  
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.20.0)",
    comments = "Source: fs_gw.proto")
public final class FsPadGrpc {

  private FsPadGrpc() {}

  public static final String SERVICE_NAME = "fsapi.FsPad";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.DevInfoReq,
      cn.iopig.feed.scale.grpc.PigstyInfoRes> getPadLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PadLogin",
      requestType = cn.iopig.feed.scale.grpc.DevInfoReq.class,
      responseType = cn.iopig.feed.scale.grpc.PigstyInfoRes.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.DevInfoReq,
      cn.iopig.feed.scale.grpc.PigstyInfoRes> getPadLoginMethod() {
    io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.DevInfoReq, cn.iopig.feed.scale.grpc.PigstyInfoRes> getPadLoginMethod;
    if ((getPadLoginMethod = FsPadGrpc.getPadLoginMethod) == null) {
      synchronized (FsPadGrpc.class) {
        if ((getPadLoginMethod = FsPadGrpc.getPadLoginMethod) == null) {
          FsPadGrpc.getPadLoginMethod = getPadLoginMethod = 
              io.grpc.MethodDescriptor.<cn.iopig.feed.scale.grpc.DevInfoReq, cn.iopig.feed.scale.grpc.PigstyInfoRes>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "fsapi.FsPad", "PadLogin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.DevInfoReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.PigstyInfoRes.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getPadLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.LoadReq,
      cn.iopig.feed.scale.grpc.ResHeader> getLoadCmdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LoadCmd",
      requestType = cn.iopig.feed.scale.grpc.LoadReq.class,
      responseType = cn.iopig.feed.scale.grpc.ResHeader.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.LoadReq,
      cn.iopig.feed.scale.grpc.ResHeader> getLoadCmdMethod() {
    io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.LoadReq, cn.iopig.feed.scale.grpc.ResHeader> getLoadCmdMethod;
    if ((getLoadCmdMethod = FsPadGrpc.getLoadCmdMethod) == null) {
      synchronized (FsPadGrpc.class) {
        if ((getLoadCmdMethod = FsPadGrpc.getLoadCmdMethod) == null) {
          FsPadGrpc.getLoadCmdMethod = getLoadCmdMethod = 
              io.grpc.MethodDescriptor.<cn.iopig.feed.scale.grpc.LoadReq, cn.iopig.feed.scale.grpc.ResHeader>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "fsapi.FsPad", "LoadCmd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.LoadReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.ResHeader.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getLoadCmdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.ChoosePigstyReq,
      cn.iopig.feed.scale.grpc.CurrentFedRes> getChoosePigstyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ChoosePigsty",
      requestType = cn.iopig.feed.scale.grpc.ChoosePigstyReq.class,
      responseType = cn.iopig.feed.scale.grpc.CurrentFedRes.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.ChoosePigstyReq,
      cn.iopig.feed.scale.grpc.CurrentFedRes> getChoosePigstyMethod() {
    io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.ChoosePigstyReq, cn.iopig.feed.scale.grpc.CurrentFedRes> getChoosePigstyMethod;
    if ((getChoosePigstyMethod = FsPadGrpc.getChoosePigstyMethod) == null) {
      synchronized (FsPadGrpc.class) {
        if ((getChoosePigstyMethod = FsPadGrpc.getChoosePigstyMethod) == null) {
          FsPadGrpc.getChoosePigstyMethod = getChoosePigstyMethod = 
              io.grpc.MethodDescriptor.<cn.iopig.feed.scale.grpc.ChoosePigstyReq, cn.iopig.feed.scale.grpc.CurrentFedRes>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "fsapi.FsPad", "ChoosePigsty"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.ChoosePigstyReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.CurrentFedRes.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getChoosePigstyMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.ChoosePigstyReq,
      cn.iopig.feed.scale.grpc.CurrentFedRes> getUploadRawInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UploadRawInfo",
      requestType = cn.iopig.feed.scale.grpc.ChoosePigstyReq.class,
      responseType = cn.iopig.feed.scale.grpc.CurrentFedRes.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.ChoosePigstyReq,
      cn.iopig.feed.scale.grpc.CurrentFedRes> getUploadRawInfoMethod() {
    io.grpc.MethodDescriptor<cn.iopig.feed.scale.grpc.ChoosePigstyReq, cn.iopig.feed.scale.grpc.CurrentFedRes> getUploadRawInfoMethod;
    if ((getUploadRawInfoMethod = FsPadGrpc.getUploadRawInfoMethod) == null) {
      synchronized (FsPadGrpc.class) {
        if ((getUploadRawInfoMethod = FsPadGrpc.getUploadRawInfoMethod) == null) {
          FsPadGrpc.getUploadRawInfoMethod = getUploadRawInfoMethod = 
              io.grpc.MethodDescriptor.<cn.iopig.feed.scale.grpc.ChoosePigstyReq, cn.iopig.feed.scale.grpc.CurrentFedRes>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "fsapi.FsPad", "UploadRawInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.ChoosePigstyReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  cn.iopig.feed.scale.grpc.CurrentFedRes.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getUploadRawInfoMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FsPadStub newStub(io.grpc.Channel channel) {
    return new FsPadStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FsPadBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new FsPadBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static FsPadFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new FsPadFutureStub(channel);
  }

  /**
   * <pre>
   * The greeting service definition.
   * interface between feed scale and pad  
   * </pre>
   */
  public static abstract class FsPadImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     *PAD 登录
     * </pre>
     */
    public void padLogin(cn.iopig.feed.scale.grpc.DevInfoReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.PigstyInfoRes> responseObserver) {
      asyncUnimplementedUnaryCall(getPadLoginMethod(), responseObserver);
    }

    /**
     * <pre>
     *上料命令
     * </pre>
     */
    public void loadCmd(cn.iopig.feed.scale.grpc.LoadReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.ResHeader> responseObserver) {
      asyncUnimplementedUnaryCall(getLoadCmdMethod(), responseObserver);
    }

    /**
     * <pre>
     *选择猪圈
     * </pre>
     */
    public void choosePigsty(cn.iopig.feed.scale.grpc.ChoosePigstyReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.CurrentFedRes> responseObserver) {
      asyncUnimplementedUnaryCall(getChoosePigstyMethod(), responseObserver);
    }

    /**
     * <pre>
     *上传当前重量
     * </pre>
     */
    public void uploadRawInfo(cn.iopig.feed.scale.grpc.ChoosePigstyReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.CurrentFedRes> responseObserver) {
      asyncUnimplementedUnaryCall(getUploadRawInfoMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getPadLoginMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                cn.iopig.feed.scale.grpc.DevInfoReq,
                cn.iopig.feed.scale.grpc.PigstyInfoRes>(
                  this, METHODID_PAD_LOGIN)))
          .addMethod(
            getLoadCmdMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                cn.iopig.feed.scale.grpc.LoadReq,
                cn.iopig.feed.scale.grpc.ResHeader>(
                  this, METHODID_LOAD_CMD)))
          .addMethod(
            getChoosePigstyMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                cn.iopig.feed.scale.grpc.ChoosePigstyReq,
                cn.iopig.feed.scale.grpc.CurrentFedRes>(
                  this, METHODID_CHOOSE_PIGSTY)))
          .addMethod(
            getUploadRawInfoMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                cn.iopig.feed.scale.grpc.ChoosePigstyReq,
                cn.iopig.feed.scale.grpc.CurrentFedRes>(
                  this, METHODID_UPLOAD_RAW_INFO)))
          .build();
    }
  }

  /**
   * <pre>
   * The greeting service definition.
   * interface between feed scale and pad  
   * </pre>
   */
  public static final class FsPadStub extends io.grpc.stub.AbstractStub<FsPadStub> {
    private FsPadStub(io.grpc.Channel channel) {
      super(channel);
    }

    private FsPadStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FsPadStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new FsPadStub(channel, callOptions);
    }

    /**
     * <pre>
     *PAD 登录
     * </pre>
     */
    public void padLogin(cn.iopig.feed.scale.grpc.DevInfoReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.PigstyInfoRes> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getPadLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *上料命令
     * </pre>
     */
    public void loadCmd(cn.iopig.feed.scale.grpc.LoadReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.ResHeader> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getLoadCmdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *选择猪圈
     * </pre>
     */
    public void choosePigsty(cn.iopig.feed.scale.grpc.ChoosePigstyReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.CurrentFedRes> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getChoosePigstyMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *上传当前重量
     * </pre>
     */
    public void uploadRawInfo(cn.iopig.feed.scale.grpc.ChoosePigstyReq request,
        io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.CurrentFedRes> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getUploadRawInfoMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * The greeting service definition.
   * interface between feed scale and pad  
   * </pre>
   */
  public static final class FsPadBlockingStub extends io.grpc.stub.AbstractStub<FsPadBlockingStub> {
    private FsPadBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private FsPadBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FsPadBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new FsPadBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     *PAD 登录
     * </pre>
     */
    public cn.iopig.feed.scale.grpc.PigstyInfoRes padLogin(cn.iopig.feed.scale.grpc.DevInfoReq request) {
      return blockingUnaryCall(
          getChannel(), getPadLoginMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *上料命令
     * </pre>
     */
    public cn.iopig.feed.scale.grpc.ResHeader loadCmd(cn.iopig.feed.scale.grpc.LoadReq request) {
      return blockingUnaryCall(
          getChannel(), getLoadCmdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *选择猪圈
     * </pre>
     */
    public cn.iopig.feed.scale.grpc.CurrentFedRes choosePigsty(cn.iopig.feed.scale.grpc.ChoosePigstyReq request) {
      return blockingUnaryCall(
          getChannel(), getChoosePigstyMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *上传当前重量
     * </pre>
     */
    public cn.iopig.feed.scale.grpc.CurrentFedRes uploadRawInfo(cn.iopig.feed.scale.grpc.ChoosePigstyReq request) {
      return blockingUnaryCall(
          getChannel(), getUploadRawInfoMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * The greeting service definition.
   * interface between feed scale and pad  
   * </pre>
   */
  public static final class FsPadFutureStub extends io.grpc.stub.AbstractStub<FsPadFutureStub> {
    private FsPadFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private FsPadFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FsPadFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new FsPadFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     *PAD 登录
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cn.iopig.feed.scale.grpc.PigstyInfoRes> padLogin(
        cn.iopig.feed.scale.grpc.DevInfoReq request) {
      return futureUnaryCall(
          getChannel().newCall(getPadLoginMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *上料命令
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cn.iopig.feed.scale.grpc.ResHeader> loadCmd(
        cn.iopig.feed.scale.grpc.LoadReq request) {
      return futureUnaryCall(
          getChannel().newCall(getLoadCmdMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *选择猪圈
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cn.iopig.feed.scale.grpc.CurrentFedRes> choosePigsty(
        cn.iopig.feed.scale.grpc.ChoosePigstyReq request) {
      return futureUnaryCall(
          getChannel().newCall(getChoosePigstyMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *上传当前重量
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cn.iopig.feed.scale.grpc.CurrentFedRes> uploadRawInfo(
        cn.iopig.feed.scale.grpc.ChoosePigstyReq request) {
      return futureUnaryCall(
          getChannel().newCall(getUploadRawInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_PAD_LOGIN = 0;
  private static final int METHODID_LOAD_CMD = 1;
  private static final int METHODID_CHOOSE_PIGSTY = 2;
  private static final int METHODID_UPLOAD_RAW_INFO = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final FsPadImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(FsPadImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_PAD_LOGIN:
          serviceImpl.padLogin((cn.iopig.feed.scale.grpc.DevInfoReq) request,
              (io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.PigstyInfoRes>) responseObserver);
          break;
        case METHODID_LOAD_CMD:
          serviceImpl.loadCmd((cn.iopig.feed.scale.grpc.LoadReq) request,
              (io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.ResHeader>) responseObserver);
          break;
        case METHODID_CHOOSE_PIGSTY:
          serviceImpl.choosePigsty((cn.iopig.feed.scale.grpc.ChoosePigstyReq) request,
              (io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.CurrentFedRes>) responseObserver);
          break;
        case METHODID_UPLOAD_RAW_INFO:
          serviceImpl.uploadRawInfo((cn.iopig.feed.scale.grpc.ChoosePigstyReq) request,
              (io.grpc.stub.StreamObserver<cn.iopig.feed.scale.grpc.CurrentFedRes>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (FsPadGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getPadLoginMethod())
              .addMethod(getLoadCmdMethod())
              .addMethod(getChoosePigstyMethod())
              .addMethod(getUploadRawInfoMethod())
              .build();
        }
      }
    }
    return result;
  }
}
