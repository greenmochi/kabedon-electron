# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from hydro_engine.proto import youtubedl_pb2 as hydro__engine_dot_proto_dot_youtubedl__pb2


class YoutubeDLStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Ping = channel.unary_unary(
        '/youtubdl.YoutubeDL/Ping',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.PingRequest.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.PingReply.FromString,
        )
    self.AddToDownloadQueue = channel.unary_unary(
        '/youtubdl.YoutubeDL/AddToDownloadQueue',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItem.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItemResponse.FromString,
        )
    self.AddToDownloadQueueSlow = channel.unary_unary(
        '/youtubdl.YoutubeDL/AddToDownloadQueueSlow',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItem.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItemResponse.FromString,
        )
    self.AllStatus = channel.unary_unary(
        '/youtubdl.YoutubeDL/AllStatus',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllStatusRequest.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllStatusResponse.FromString,
        )


class YoutubeDLServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Ping(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AddToDownloadQueue(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AddToDownloadQueueSlow(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AllStatus(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_YoutubeDLServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Ping': grpc.unary_unary_rpc_method_handler(
          servicer.Ping,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.PingRequest.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.PingReply.SerializeToString,
      ),
      'AddToDownloadQueue': grpc.unary_unary_rpc_method_handler(
          servicer.AddToDownloadQueue,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItem.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItemResponse.SerializeToString,
      ),
      'AddToDownloadQueueSlow': grpc.unary_unary_rpc_method_handler(
          servicer.AddToDownloadQueueSlow,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItem.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItemResponse.SerializeToString,
      ),
      'AllStatus': grpc.unary_unary_rpc_method_handler(
          servicer.AllStatus,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllStatusRequest.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllStatusResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'youtubdl.YoutubeDL', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
