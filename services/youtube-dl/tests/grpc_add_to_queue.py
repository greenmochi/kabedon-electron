import grpc
import time

from hydro_engine.proto import youtubedl_pb2
from hydro_engine.proto import youtubedl_pb2_grpc

with grpc.insecure_channel("localhost:8123") as channel:
    stub = youtubedl_pb2_grpc.YoutubeDLStub(channel)
    stub.AddToDownloadQueue(youtubedl_pb2.DownloadItem(url="https://www.youtube.com/watch?v=vDEbdmXR4Jc"))

    time.sleep(3)

    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    status = stub.AllStatus(youtubedl_pb2.AllStatusRequest())
    print(status)
    
    stub.AddToDownloadQueue(youtubedl_pb2.DownloadItem(url="https://www.youtube.com/watch?v=ajVH72iiIdM"))