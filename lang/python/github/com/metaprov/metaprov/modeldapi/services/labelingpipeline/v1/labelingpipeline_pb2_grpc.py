# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from github.com.metaprov.modeldapi.services.labelingpipeline.v1 import labelingpipeline_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2


class LabelingPipelineServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListLabelingPipelines = channel.unary_unary(
                '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/ListLabelingPipelines',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.ListLabelingPipelineRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.ListLabelingPipelineResponse.FromString,
                )
        self.CreateLabelingPipeline = channel.unary_unary(
                '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/CreateLabelingPipeline',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.CreateLabelingPipelineRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.CreateLabelingPipelineResponse.FromString,
                )
        self.GetLabelingPipeline = channel.unary_unary(
                '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/GetLabelingPipeline',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.GetLabelingPipelineRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.GetLabelingPipelineResponse.FromString,
                )
        self.UpdateLabelingPipeline = channel.unary_unary(
                '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/UpdateLabelingPipeline',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.UpdateLabelingPipelineRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.UpdateLabelingPipelineResponse.FromString,
                )
        self.DeleteLabelingPipeline = channel.unary_unary(
                '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/DeleteLabelingPipeline',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.DeleteLabelingPipelineRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.DeleteLabelingPipelineResponse.FromString,
                )


class LabelingPipelineServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListLabelingPipelines(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateLabelingPipeline(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLabelingPipeline(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateLabelingPipeline(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteLabelingPipeline(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LabelingPipelineServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListLabelingPipelines': grpc.unary_unary_rpc_method_handler(
                    servicer.ListLabelingPipelines,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.ListLabelingPipelineRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.ListLabelingPipelineResponse.SerializeToString,
            ),
            'CreateLabelingPipeline': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateLabelingPipeline,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.CreateLabelingPipelineRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.CreateLabelingPipelineResponse.SerializeToString,
            ),
            'GetLabelingPipeline': grpc.unary_unary_rpc_method_handler(
                    servicer.GetLabelingPipeline,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.GetLabelingPipelineRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.GetLabelingPipelineResponse.SerializeToString,
            ),
            'UpdateLabelingPipeline': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateLabelingPipeline,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.UpdateLabelingPipelineRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.UpdateLabelingPipelineResponse.SerializeToString,
            ),
            'DeleteLabelingPipeline': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteLabelingPipeline,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.DeleteLabelingPipelineRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.DeleteLabelingPipelineResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class LabelingPipelineService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListLabelingPipelines(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/ListLabelingPipelines',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.ListLabelingPipelineRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.ListLabelingPipelineResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateLabelingPipeline(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/CreateLabelingPipeline',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.CreateLabelingPipelineRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.CreateLabelingPipelineResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetLabelingPipeline(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/GetLabelingPipeline',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.GetLabelingPipelineRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.GetLabelingPipelineResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateLabelingPipeline(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/UpdateLabelingPipeline',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.UpdateLabelingPipelineRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.UpdateLabelingPipelineResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteLabelingPipeline(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/DeleteLabelingPipeline',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.DeleteLabelingPipelineRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_labelingpipeline_dot_v1_dot_labelingpipeline__pb2.DeleteLabelingPipelineResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
