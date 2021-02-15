# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from github.com.metaprov.modeldapi.services.predictionpipelinerun.v1 import predictionpipelinerun_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2


class PredictionPipelineRunServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListPredictionPipelineRuns = channel.unary_unary(
                '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/ListPredictionPipelineRuns',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.ListPredictionPipelineRunsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.ListPredictionPipelineRunsResponse.FromString,
                )
        self.CreatePredictionPipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/CreatePredictionPipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.CreatePredictionPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.CreatePredictionPipelineRunResponse.FromString,
                )
        self.GetPredictionPipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/GetPredictionPipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.GetPredictionPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.GetPredictionPipelineRunResponse.FromString,
                )
        self.UpdatePredictionPipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/UpdatePredictionPipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.UpdatePredictionPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.UpdatePredictionPipelineRunResponse.FromString,
                )
        self.DeletePredictionPipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/DeletePredictionPipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DeletePredictionPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DeletePredictionPipelineRunResponse.FromString,
                )
        self.Download = channel.unary_unary(
                '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/Download',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DownloadPredictionPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DownloadPredictionPipelineRunResponse.FromString,
                )


class PredictionPipelineRunServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListPredictionPipelineRuns(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreatePredictionPipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPredictionPipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdatePredictionPipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeletePredictionPipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Download(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PredictionPipelineRunServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListPredictionPipelineRuns': grpc.unary_unary_rpc_method_handler(
                    servicer.ListPredictionPipelineRuns,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.ListPredictionPipelineRunsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.ListPredictionPipelineRunsResponse.SerializeToString,
            ),
            'CreatePredictionPipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.CreatePredictionPipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.CreatePredictionPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.CreatePredictionPipelineRunResponse.SerializeToString,
            ),
            'GetPredictionPipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPredictionPipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.GetPredictionPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.GetPredictionPipelineRunResponse.SerializeToString,
            ),
            'UpdatePredictionPipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdatePredictionPipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.UpdatePredictionPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.UpdatePredictionPipelineRunResponse.SerializeToString,
            ),
            'DeletePredictionPipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.DeletePredictionPipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DeletePredictionPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DeletePredictionPipelineRunResponse.SerializeToString,
            ),
            'Download': grpc.unary_unary_rpc_method_handler(
                    servicer.Download,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DownloadPredictionPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DownloadPredictionPipelineRunResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PredictionPipelineRunService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListPredictionPipelineRuns(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/ListPredictionPipelineRuns',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.ListPredictionPipelineRunsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.ListPredictionPipelineRunsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreatePredictionPipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/CreatePredictionPipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.CreatePredictionPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.CreatePredictionPipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPredictionPipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/GetPredictionPipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.GetPredictionPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.GetPredictionPipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdatePredictionPipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/UpdatePredictionPipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.UpdatePredictionPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.UpdatePredictionPipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeletePredictionPipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/DeletePredictionPipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DeletePredictionPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DeletePredictionPipelineRunResponse.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.predictionpipelinerun.v1.PredictionPipelineRunService/Download',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DownloadPredictionPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_predictionpipelinerun_dot_v1_dot_predictionpipelinerun__pb2.DownloadPredictionPipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
