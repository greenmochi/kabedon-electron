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
        '/youtubedl.YoutubeDL/Ping',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.PingRequest.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.PingReply.FromString,
        )
    self.AddToQueue = channel.unary_unary(
        '/youtubedl.YoutubeDL/AddToQueue',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItem.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItemResponse.FromString,
        )
    self.RemoveFromQueue = channel.unary_unary(
        '/youtubedl.YoutubeDL/RemoveFromQueue',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadID.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadRemoveResponse.FromString,
        )
    self.GetAllDownloads = channel.unary_unary(
        '/youtubedl.YoutubeDL/GetAllDownloads',
        request_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllDownloadsRequest.SerializeToString,
        response_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllDownloads.FromString,
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

  def AddToQueue(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RemoveFromQueue(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetAllDownloads(self, request, context):
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
      'AddToQueue': grpc.unary_unary_rpc_method_handler(
          servicer.AddToQueue,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItem.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadItemResponse.SerializeToString,
      ),
      'RemoveFromQueue': grpc.unary_unary_rpc_method_handler(
          servicer.RemoveFromQueue,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadID.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.DownloadRemoveResponse.SerializeToString,
      ),
      'GetAllDownloads': grpc.unary_unary_rpc_method_handler(
          servicer.GetAllDownloads,
          request_deserializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllDownloadsRequest.FromString,
          response_serializer=hydro__engine_dot_proto_dot_youtubedl__pb2.AllDownloads.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'youtubedl.YoutubeDL', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
