# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from github.com.metaprov.modeldapi.services.catalog.v1 import catalog_pb2 as github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2


class CatalogServiceStub(object):
    """////////////////////////////////////////////////// Service ///////////////////////////////////

    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListAlgorithm = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListAlgorithm',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListAlgorithmsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListAlgorithmsResponse.FromString,
                )
        self.GetAlgorithm = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetAlgorithm',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetAlgorithmRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetAlgorithmResponse.FromString,
                )
        self.ListMLFrameworks = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListMLFrameworks',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListMLFrameworksRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListMLFrameworksResponse.FromString,
                )
        self.GetMLFramework = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetMLFramework',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetMLFrameworkRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetMLFrameworkResponse.FromString,
                )
        self.ListClouds = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListClouds',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListCloudsRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListCloudsResponse.FromString,
                )
        self.GetCloud = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetCloud',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetCloudRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetCloudResponse.FromString,
                )
        self.ListWorkloadClasses = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListWorkloadClasses',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListWorkloadClassesRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListWorkloadClassesResponse.FromString,
                )
        self.GetWorkloadClass = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetWorkloadClass',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetWorkloadClassRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetWorkloadClassResponse.FromString,
                )
        self.ListUserRoleClasses = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListUserRoleClasses',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListUserRoleClassesRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListUserRoleClassesResponse.FromString,
                )
        self.GetUserRoleClass = channel.unary_unary(
                '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetUserRoleClass',
                request_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetUserRoleClassRequest.SerializeToString,
                response_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetUserRoleClassResponse.FromString,
                )


class CatalogServiceServicer(object):
    """////////////////////////////////////////////////// Service ///////////////////////////////////

    """

    def ListAlgorithm(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAlgorithm(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListMLFrameworks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMLFramework(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListClouds(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCloud(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListWorkloadClasses(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetWorkloadClass(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUserRoleClasses(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserRoleClass(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CatalogServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListAlgorithm': grpc.unary_unary_rpc_method_handler(
                    servicer.ListAlgorithm,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListAlgorithmsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListAlgorithmsResponse.SerializeToString,
            ),
            'GetAlgorithm': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAlgorithm,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetAlgorithmRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetAlgorithmResponse.SerializeToString,
            ),
            'ListMLFrameworks': grpc.unary_unary_rpc_method_handler(
                    servicer.ListMLFrameworks,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListMLFrameworksRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListMLFrameworksResponse.SerializeToString,
            ),
            'GetMLFramework': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMLFramework,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetMLFrameworkRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetMLFrameworkResponse.SerializeToString,
            ),
            'ListClouds': grpc.unary_unary_rpc_method_handler(
                    servicer.ListClouds,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListCloudsRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListCloudsResponse.SerializeToString,
            ),
            'GetCloud': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCloud,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetCloudRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetCloudResponse.SerializeToString,
            ),
            'ListWorkloadClasses': grpc.unary_unary_rpc_method_handler(
                    servicer.ListWorkloadClasses,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListWorkloadClassesRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListWorkloadClassesResponse.SerializeToString,
            ),
            'GetWorkloadClass': grpc.unary_unary_rpc_method_handler(
                    servicer.GetWorkloadClass,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetWorkloadClassRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetWorkloadClassResponse.SerializeToString,
            ),
            'ListUserRoleClasses': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUserRoleClasses,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListUserRoleClassesRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListUserRoleClassesResponse.SerializeToString,
            ),
            'GetUserRoleClass': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUserRoleClass,
                    request_deserializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetUserRoleClassRequest.FromString,
                    response_serializer=github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetUserRoleClassResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'github.com.metaprov.modeld.services.catalog.v1.CatalogService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CatalogService(object):
    """////////////////////////////////////////////////// Service ///////////////////////////////////

    """

    @staticmethod
    def ListAlgorithm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListAlgorithm',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListAlgorithmsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListAlgorithmsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAlgorithm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetAlgorithm',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetAlgorithmRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetAlgorithmResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListMLFrameworks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListMLFrameworks',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListMLFrameworksRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListMLFrameworksResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMLFramework(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetMLFramework',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetMLFrameworkRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetMLFrameworkResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListClouds(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListClouds',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListCloudsRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListCloudsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCloud(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetCloud',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetCloudRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetCloudResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListWorkloadClasses(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListWorkloadClasses',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListWorkloadClassesRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListWorkloadClassesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetWorkloadClass(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetWorkloadClass',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetWorkloadClassRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetWorkloadClassResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListUserRoleClasses(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/ListUserRoleClasses',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListUserRoleClassesRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.ListUserRoleClassesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUserRoleClass(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/github.com.metaprov.modeld.services.catalog.v1.CatalogService/GetUserRoleClass',
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetUserRoleClassRequest.SerializeToString,
            github_dot_com_dot_metaprov_dot_modeldapi_dot_services_dot_catalog_dot_v1_dot_catalog__pb2.GetUserRoleClassResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
