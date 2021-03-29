# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from github.com.metaprov.modeldapi.services.cronprediction.v1 import cronprediction_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2


class CronPredictionServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListCronPredictions = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/ListCronPredictions',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.ListCronPredictionsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.ListCronPredictionsResponse.FromString,
                )
        self.CreateCronPrediction = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/CreateCronPrediction',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.CreateCronPredictionRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.CreateCronPredictionResponse.FromString,
                )
        self.GetCronPrediction = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/GetCronPrediction',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.GetCronPredictionRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.GetCronPredictionResponse.FromString,
                )
        self.UpdateCronPrediction = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/UpdateCronPrediction',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.UpdateCronPredictionRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.UpdateCronPredictionResponse.FromString,
                )
        self.DeleteCronPrediction = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/DeleteCronPrediction',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DeleteCronPredictionRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DeleteCronPredictionResponse.FromString,
                )
        self.Download = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/Download',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DownloadCronPredictionRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DownloadCronPredictionResponse.FromString,
                )


class CronPredictionServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListCronPredictions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateCronPrediction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCronPrediction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateCronPrediction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteCronPrediction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Download(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CronPredictionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListCronPredictions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListCronPredictions,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.ListCronPredictionsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.ListCronPredictionsResponse.SerializeToString,
            ),
            'CreateCronPrediction': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCronPrediction,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.CreateCronPredictionRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.CreateCronPredictionResponse.SerializeToString,
            ),
            'GetCronPrediction': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCronPrediction,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.GetCronPredictionRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.GetCronPredictionResponse.SerializeToString,
            ),
            'UpdateCronPrediction': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateCronPrediction,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.UpdateCronPredictionRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.UpdateCronPredictionResponse.SerializeToString,
            ),
            'DeleteCronPrediction': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteCronPrediction,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DeleteCronPredictionRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DeleteCronPredictionResponse.SerializeToString,
            ),
            'Download': grpc.unary_unary_rpc_method_handler(
                    servicer.Download,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DownloadCronPredictionRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DownloadCronPredictionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CronPredictionService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListCronPredictions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/ListCronPredictions',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.ListCronPredictionsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.ListCronPredictionsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateCronPrediction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/CreateCronPrediction',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.CreateCronPredictionRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.CreateCronPredictionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCronPrediction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/GetCronPrediction',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.GetCronPredictionRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.GetCronPredictionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateCronPrediction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/UpdateCronPrediction',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.UpdateCronPredictionRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.UpdateCronPredictionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteCronPrediction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/DeleteCronPrediction',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DeleteCronPredictionRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DeleteCronPredictionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Download(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.cronprediction.v1.CronPredictionService/Download',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DownloadCronPredictionRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_cronprediction_dot_v1_dot_cronprediction__pb2.DownloadCronPredictionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)