# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from github.com.metaprov.modeldapi.services.model.v1 import model_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2


class ModelServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListModels = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/ListModels',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ListModelsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ListModelsResponse.FromString,
                )
        self.CreateModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CreateModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelResponse.FromString,
                )
        self.GetModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelResponse.FromString,
                )
        self.UpdateModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/UpdateModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.UpdateModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.UpdateModelResponse.FromString,
                )
        self.DeleteModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/DeleteModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeleteModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeleteModelResponse.FromString,
                )
        self.DeployModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/DeployModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeployModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeployModelResponse.FromString,
                )
        self.PublishModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/PublishModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PublishModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PublishModelResponse.FromString,
                )
        self.CreateModelProfile = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CreateModelProfile',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelProfileRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelProfileResponse.FromString,
                )
        self.GetModelProfile = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModelProfile',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelProfileRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelProfileResponse.FromString,
                )
        self.GetModelMisclass = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModelMisclass',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetMisclassRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetMisclassResponse.FromString,
                )
        self.GetModelLogs = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModelLogs',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelLogsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelLogsResponse.FromString,
                )
        self.AbortModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/AbortModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.AbortModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.AbortModelResponse.FromString,
                )
        self.PauseModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/PauseModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PauseModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PauseModelResponse.FromString,
                )
        self.ResumeModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/ResumeModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ResumeModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ResumeModelResponse.FromString,
                )
        self.CompareModels = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CompareModels',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompareModelsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompareModelsResponse.FromString,
                )
        self.CompileModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CompileModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompileModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompileModelResponse.FromString,
                )
        self.DownloadModel = channel.unary_unary(
                '/github.com.metaprov.modeldapi.services.model.v1.ModelService/DownloadModel',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DownloadModelRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DownloadModelResponse.FromString,
                )


class ModelServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListModels(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeployModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PublishModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateModelProfile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetModelProfile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetModelMisclass(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetModelLogs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AbortModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PauseModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ResumeModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CompareModels(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CompileModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DownloadModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ModelServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListModels': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModels,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ListModelsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ListModelsResponse.SerializeToString,
            ),
            'CreateModel': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelResponse.SerializeToString,
            ),
            'GetModel': grpc.unary_unary_rpc_method_handler(
                    servicer.GetModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelResponse.SerializeToString,
            ),
            'UpdateModel': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.UpdateModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.UpdateModelResponse.SerializeToString,
            ),
            'DeleteModel': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeleteModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeleteModelResponse.SerializeToString,
            ),
            'DeployModel': grpc.unary_unary_rpc_method_handler(
                    servicer.DeployModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeployModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeployModelResponse.SerializeToString,
            ),
            'PublishModel': grpc.unary_unary_rpc_method_handler(
                    servicer.PublishModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PublishModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PublishModelResponse.SerializeToString,
            ),
            'CreateModelProfile': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateModelProfile,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelProfileRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelProfileResponse.SerializeToString,
            ),
            'GetModelProfile': grpc.unary_unary_rpc_method_handler(
                    servicer.GetModelProfile,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelProfileRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelProfileResponse.SerializeToString,
            ),
            'GetModelMisclass': grpc.unary_unary_rpc_method_handler(
                    servicer.GetModelMisclass,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetMisclassRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetMisclassResponse.SerializeToString,
            ),
            'GetModelLogs': grpc.unary_unary_rpc_method_handler(
                    servicer.GetModelLogs,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelLogsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelLogsResponse.SerializeToString,
            ),
            'AbortModel': grpc.unary_unary_rpc_method_handler(
                    servicer.AbortModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.AbortModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.AbortModelResponse.SerializeToString,
            ),
            'PauseModel': grpc.unary_unary_rpc_method_handler(
                    servicer.PauseModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PauseModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PauseModelResponse.SerializeToString,
            ),
            'ResumeModel': grpc.unary_unary_rpc_method_handler(
                    servicer.ResumeModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ResumeModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ResumeModelResponse.SerializeToString,
            ),
            'CompareModels': grpc.unary_unary_rpc_method_handler(
                    servicer.CompareModels,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompareModelsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompareModelsResponse.SerializeToString,
            ),
            'CompileModel': grpc.unary_unary_rpc_method_handler(
                    servicer.CompileModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompileModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompileModelResponse.SerializeToString,
            ),
            'DownloadModel': grpc.unary_unary_rpc_method_handler(
                    servicer.DownloadModel,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DownloadModelRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DownloadModelResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'github.com.metaprov.modeldapi.services.model.v1.ModelService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ModelService(object):
    """Missing associated documentation comment in .proto file."""

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
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/ListModels',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ListModelsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ListModelsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CreateModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/UpdateModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.UpdateModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.UpdateModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/DeleteModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeleteModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeleteModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeployModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/DeployModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeployModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DeployModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PublishModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/PublishModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PublishModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PublishModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateModelProfile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CreateModelProfile',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelProfileRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CreateModelProfileResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetModelProfile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModelProfile',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelProfileRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelProfileResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetModelMisclass(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModelMisclass',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetMisclassRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetMisclassResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetModelLogs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/GetModelLogs',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelLogsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.GetModelLogsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AbortModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/AbortModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.AbortModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.AbortModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PauseModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/PauseModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PauseModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.PauseModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ResumeModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/ResumeModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ResumeModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.ResumeModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CompareModels(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CompareModels',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompareModelsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompareModelsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CompileModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/CompileModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompileModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.CompileModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DownloadModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeldapi.services.model.v1.ModelService/DownloadModel',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DownloadModelRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_model_dot_v1_dot_model__pb2.DownloadModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
