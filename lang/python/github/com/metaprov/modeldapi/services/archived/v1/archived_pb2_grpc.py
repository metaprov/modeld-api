# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from github.com.metaprov.modeldapi.services.archived.v1 import archived_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class ArchivedServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RecordObject = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/RecordObject',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.RecordObjectRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.GetObject = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/GetObject',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.GetObjectRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.GetObjectResponse.FromString,
                )
        self.ListObjects = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListObjects',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListObjectRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListObjectResponse.FromString,
                )
        self.DeleteObject = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/DeleteObject',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.DeleteObjectRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ListDatasets = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListDatasets',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListDatasetRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListDatasetResponse.FromString,
                )
        self.ListFeaturePipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListFeaturePipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListFeaturePipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListFeaturePipelineRunResponse.FromString,
                )
        self.ListLabelingPipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListLabelingPipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListLabelingPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListLabelingPipelineRunResponse.FromString,
                )
        self.ListRecipeRun = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListRecipeRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
                )
        self.ListModelAutoBuilders = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListModelAutoBuilders',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListModelAutoBuilderRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListModelAutoBuilderResponse.FromString,
                )
        self.ListModels = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListModels',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
                )
        self.ListModelPipelineRuns = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListModelPipelineRuns',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
                )
        self.ListNotebookRuns = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListNotebookRuns',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
                )
        self.ListReports = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListReports',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListReportRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListReportResponse.FromString,
                )
        self.ListStudies = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListStudies',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListStudyRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListStudyResponse.FromString,
                )
        self.ListPredictionPipelineRun = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListPredictionPipelineRun',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListPredictionPipelineRunRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListPredictionPipelineRunResponse.FromString,
                )


class ArchivedServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RecordObject(self, request, context):
        """Record all objects in a blob
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetObject(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListObjects(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteObject(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListDatasets(self, request, context):
        """records the operational objects
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListFeaturePipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListLabelingPipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListRecipeRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListModelAutoBuilders(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListModels(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListModelPipelineRuns(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListNotebookRuns(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListReports(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListStudies(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListPredictionPipelineRun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ArchivedServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RecordObject': grpc.unary_unary_rpc_method_handler(
                    servicer.RecordObject,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.RecordObjectRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetObject': grpc.unary_unary_rpc_method_handler(
                    servicer.GetObject,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.GetObjectRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.GetObjectResponse.SerializeToString,
            ),
            'ListObjects': grpc.unary_unary_rpc_method_handler(
                    servicer.ListObjects,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListObjectRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListObjectResponse.SerializeToString,
            ),
            'DeleteObject': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteObject,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.DeleteObjectRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ListDatasets': grpc.unary_unary_rpc_method_handler(
                    servicer.ListDatasets,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListDatasetRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListDatasetResponse.SerializeToString,
            ),
            'ListFeaturePipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.ListFeaturePipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListFeaturePipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListFeaturePipelineRunResponse.SerializeToString,
            ),
            'ListLabelingPipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.ListLabelingPipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListLabelingPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListLabelingPipelineRunResponse.SerializeToString,
            ),
            'ListRecipeRun': grpc.unary_unary_rpc_method_handler(
                    servicer.ListRecipeRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.SerializeToString,
            ),
            'ListModelAutoBuilders': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModelAutoBuilders,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListModelAutoBuilderRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListModelAutoBuilderResponse.SerializeToString,
            ),
            'ListModels': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModels,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.SerializeToString,
            ),
            'ListModelPipelineRuns': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModelPipelineRuns,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.SerializeToString,
            ),
            'ListNotebookRuns': grpc.unary_unary_rpc_method_handler(
                    servicer.ListNotebookRuns,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.SerializeToString,
            ),
            'ListReports': grpc.unary_unary_rpc_method_handler(
                    servicer.ListReports,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListReportRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListReportResponse.SerializeToString,
            ),
            'ListStudies': grpc.unary_unary_rpc_method_handler(
                    servicer.ListStudies,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListStudyRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListStudyResponse.SerializeToString,
            ),
            'ListPredictionPipelineRun': grpc.unary_unary_rpc_method_handler(
                    servicer.ListPredictionPipelineRun,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListPredictionPipelineRunRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListPredictionPipelineRunResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'github.com.metaprov.modeldapi.services.archived.v1.ArchivedService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ArchivedService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RecordObject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/RecordObject',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.RecordObjectRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetObject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/GetObject',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.GetObjectRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.GetObjectResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListObjects(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListObjects',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListObjectRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListObjectResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteObject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/DeleteObject',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.DeleteObjectRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListDatasets(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListDatasets',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListDatasetRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListDatasetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListFeaturePipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListFeaturePipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListFeaturePipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListFeaturePipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListLabelingPipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListLabelingPipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListLabelingPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListLabelingPipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListRecipeRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListRecipeRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListModelAutoBuilders(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListModelAutoBuilders',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListModelAutoBuilderRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListModelAutoBuilderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListModels(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListModels',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListModelPipelineRuns(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListModelPipelineRuns',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListNotebookRuns(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListNotebookRuns',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListRecipeRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListReports(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListReports',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListReportRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListReportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListStudies(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListStudies',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListStudyRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListStudyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListPredictionPipelineRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.archived.v1.ArchivedService/ListPredictionPipelineRun',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListPredictionPipelineRunRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_archived_dot_v1_dot_archived__pb2.ListPredictionPipelineRunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
