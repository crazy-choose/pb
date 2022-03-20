// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: stream.proto

#include "stream.pb.h"
#include "stream.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>

static const char* StreamService_method_names[] = {
  "/StreamService/Request",
  "/StreamService/Notify",
  "/StreamService/SendTrap",
};

std::unique_ptr< StreamService::Stub> StreamService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< StreamService::Stub> stub(new StreamService::Stub(channel, options));
  return stub;
}

StreamService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Request_(StreamService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::CLIENT_STREAMING, channel)
  , rpcmethod_Notify_(StreamService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::BIDI_STREAMING, channel)
  , rpcmethod_SendTrap_(StreamService_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::BIDI_STREAMING, channel)
  {}

::grpc::ClientWriter< ::ReqMsg>* StreamService::Stub::RequestRaw(::grpc::ClientContext* context, ::RspMsg* response) {
  return ::grpc::internal::ClientWriterFactory< ::ReqMsg>::Create(channel_.get(), rpcmethod_Request_, context, response);
}

void StreamService::Stub::async::Request(::grpc::ClientContext* context, ::RspMsg* response, ::grpc::ClientWriteReactor< ::ReqMsg>* reactor) {
  ::grpc::internal::ClientCallbackWriterFactory< ::ReqMsg>::Create(stub_->channel_.get(), stub_->rpcmethod_Request_, context, response, reactor);
}

::grpc::ClientAsyncWriter< ::ReqMsg>* StreamService::Stub::AsyncRequestRaw(::grpc::ClientContext* context, ::RspMsg* response, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncWriterFactory< ::ReqMsg>::Create(channel_.get(), cq, rpcmethod_Request_, context, response, true, tag);
}

::grpc::ClientAsyncWriter< ::ReqMsg>* StreamService::Stub::PrepareAsyncRequestRaw(::grpc::ClientContext* context, ::RspMsg* response, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncWriterFactory< ::ReqMsg>::Create(channel_.get(), cq, rpcmethod_Request_, context, response, false, nullptr);
}

::grpc::ClientReaderWriter< ::NotifyMsg, ::NetResult>* StreamService::Stub::NotifyRaw(::grpc::ClientContext* context) {
  return ::grpc::internal::ClientReaderWriterFactory< ::NotifyMsg, ::NetResult>::Create(channel_.get(), rpcmethod_Notify_, context);
}

void StreamService::Stub::async::Notify(::grpc::ClientContext* context, ::grpc::ClientBidiReactor< ::NotifyMsg,::NetResult>* reactor) {
  ::grpc::internal::ClientCallbackReaderWriterFactory< ::NotifyMsg,::NetResult>::Create(stub_->channel_.get(), stub_->rpcmethod_Notify_, context, reactor);
}

::grpc::ClientAsyncReaderWriter< ::NotifyMsg, ::NetResult>* StreamService::Stub::AsyncNotifyRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderWriterFactory< ::NotifyMsg, ::NetResult>::Create(channel_.get(), cq, rpcmethod_Notify_, context, true, tag);
}

::grpc::ClientAsyncReaderWriter< ::NotifyMsg, ::NetResult>* StreamService::Stub::PrepareAsyncNotifyRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderWriterFactory< ::NotifyMsg, ::NetResult>::Create(channel_.get(), cq, rpcmethod_Notify_, context, false, nullptr);
}

::grpc::ClientReaderWriter< ::TrapMsg, ::NetResult>* StreamService::Stub::SendTrapRaw(::grpc::ClientContext* context) {
  return ::grpc::internal::ClientReaderWriterFactory< ::TrapMsg, ::NetResult>::Create(channel_.get(), rpcmethod_SendTrap_, context);
}

void StreamService::Stub::async::SendTrap(::grpc::ClientContext* context, ::grpc::ClientBidiReactor< ::TrapMsg,::NetResult>* reactor) {
  ::grpc::internal::ClientCallbackReaderWriterFactory< ::TrapMsg,::NetResult>::Create(stub_->channel_.get(), stub_->rpcmethod_SendTrap_, context, reactor);
}

::grpc::ClientAsyncReaderWriter< ::TrapMsg, ::NetResult>* StreamService::Stub::AsyncSendTrapRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderWriterFactory< ::TrapMsg, ::NetResult>::Create(channel_.get(), cq, rpcmethod_SendTrap_, context, true, tag);
}

::grpc::ClientAsyncReaderWriter< ::TrapMsg, ::NetResult>* StreamService::Stub::PrepareAsyncSendTrapRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderWriterFactory< ::TrapMsg, ::NetResult>::Create(channel_.get(), cq, rpcmethod_SendTrap_, context, false, nullptr);
}

StreamService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      StreamService_method_names[0],
      ::grpc::internal::RpcMethod::CLIENT_STREAMING,
      new ::grpc::internal::ClientStreamingHandler< StreamService::Service, ::ReqMsg, ::RspMsg>(
          [](StreamService::Service* service,
             ::grpc::ServerContext* ctx,
             ::grpc::ServerReader<::ReqMsg>* reader,
             ::RspMsg* resp) {
               return service->Request(ctx, reader, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      StreamService_method_names[1],
      ::grpc::internal::RpcMethod::BIDI_STREAMING,
      new ::grpc::internal::BidiStreamingHandler< StreamService::Service, ::NotifyMsg, ::NetResult>(
          [](StreamService::Service* service,
             ::grpc::ServerContext* ctx,
             ::grpc::ServerReaderWriter<::NetResult,
             ::NotifyMsg>* stream) {
               return service->Notify(ctx, stream);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      StreamService_method_names[2],
      ::grpc::internal::RpcMethod::BIDI_STREAMING,
      new ::grpc::internal::BidiStreamingHandler< StreamService::Service, ::TrapMsg, ::NetResult>(
          [](StreamService::Service* service,
             ::grpc::ServerContext* ctx,
             ::grpc::ServerReaderWriter<::NetResult,
             ::TrapMsg>* stream) {
               return service->SendTrap(ctx, stream);
             }, this)));
}

StreamService::Service::~Service() {
}

::grpc::Status StreamService::Service::Request(::grpc::ServerContext* context, ::grpc::ServerReader< ::ReqMsg>* reader, ::RspMsg* response) {
  (void) context;
  (void) reader;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status StreamService::Service::Notify(::grpc::ServerContext* context, ::grpc::ServerReaderWriter< ::NetResult, ::NotifyMsg>* stream) {
  (void) context;
  (void) stream;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status StreamService::Service::SendTrap(::grpc::ServerContext* context, ::grpc::ServerReaderWriter< ::NetResult, ::TrapMsg>* stream) {
  (void) context;
  (void) stream;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

