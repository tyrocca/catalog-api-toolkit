# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from catalog.v1 import catalog_pb2 as catalog_dot_v1_dot_catalog__pb2
from catalog.v1 import company_pb2 as catalog_dot_v1_dot_company__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class CatalogServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateCompany = channel.unary_unary(
                '/catalog.v1.CatalogService/CreateCompany',
                request_serializer=catalog_dot_v1_dot_company__pb2.CreateCompanyRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_company__pb2.Company.FromString,
                )
        self.GetCompany = channel.unary_unary(
                '/catalog.v1.CatalogService/GetCompany',
                request_serializer=catalog_dot_v1_dot_company__pb2.GetCompanyRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_company__pb2.Company.FromString,
                )
        self.ListCompanies = channel.unary_unary(
                '/catalog.v1.CatalogService/ListCompanies',
                request_serializer=catalog_dot_v1_dot_company__pb2.ListCompaniesRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_company__pb2.ListCompaniesResponse.FromString,
                )
        self.UpdateCompany = channel.unary_unary(
                '/catalog.v1.CatalogService/UpdateCompany',
                request_serializer=catalog_dot_v1_dot_company__pb2.UpdateCompanyRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_company__pb2.Company.FromString,
                )
        self.DeleteCompany = channel.unary_unary(
                '/catalog.v1.CatalogService/DeleteCompany',
                request_serializer=catalog_dot_v1_dot_company__pb2.DeleteCompanyRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.CreateCatalog = channel.unary_unary(
                '/catalog.v1.CatalogService/CreateCatalog',
                request_serializer=catalog_dot_v1_dot_catalog__pb2.CreateCatalogRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_catalog__pb2.Catalog.FromString,
                )
        self.GetCatalog = channel.unary_unary(
                '/catalog.v1.CatalogService/GetCatalog',
                request_serializer=catalog_dot_v1_dot_catalog__pb2.GetCatalogRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_catalog__pb2.Catalog.FromString,
                )
        self.ListCatalogs = channel.unary_unary(
                '/catalog.v1.CatalogService/ListCatalogs',
                request_serializer=catalog_dot_v1_dot_catalog__pb2.ListCatalogsRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_catalog__pb2.ListCatalogsResponse.FromString,
                )
        self.UpdateCatalog = channel.unary_unary(
                '/catalog.v1.CatalogService/UpdateCatalog',
                request_serializer=catalog_dot_v1_dot_catalog__pb2.UpdateCatalogRequest.SerializeToString,
                response_deserializer=catalog_dot_v1_dot_catalog__pb2.Catalog.FromString,
                )
        self.DeleteCatalog = channel.unary_unary(
                '/catalog.v1.CatalogService/DeleteCatalog',
                request_serializer=catalog_dot_v1_dot_catalog__pb2.DeleteCatalogRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class CatalogServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateCompany(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCompany(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListCompanies(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateCompany(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteCompany(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateCatalog(self, request, context):
        """Catalog CRUD
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCatalog(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListCatalogs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateCatalog(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteCatalog(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CatalogServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateCompany': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCompany,
                    request_deserializer=catalog_dot_v1_dot_company__pb2.CreateCompanyRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_company__pb2.Company.SerializeToString,
            ),
            'GetCompany': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCompany,
                    request_deserializer=catalog_dot_v1_dot_company__pb2.GetCompanyRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_company__pb2.Company.SerializeToString,
            ),
            'ListCompanies': grpc.unary_unary_rpc_method_handler(
                    servicer.ListCompanies,
                    request_deserializer=catalog_dot_v1_dot_company__pb2.ListCompaniesRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_company__pb2.ListCompaniesResponse.SerializeToString,
            ),
            'UpdateCompany': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateCompany,
                    request_deserializer=catalog_dot_v1_dot_company__pb2.UpdateCompanyRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_company__pb2.Company.SerializeToString,
            ),
            'DeleteCompany': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteCompany,
                    request_deserializer=catalog_dot_v1_dot_company__pb2.DeleteCompanyRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'CreateCatalog': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCatalog,
                    request_deserializer=catalog_dot_v1_dot_catalog__pb2.CreateCatalogRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_catalog__pb2.Catalog.SerializeToString,
            ),
            'GetCatalog': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCatalog,
                    request_deserializer=catalog_dot_v1_dot_catalog__pb2.GetCatalogRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_catalog__pb2.Catalog.SerializeToString,
            ),
            'ListCatalogs': grpc.unary_unary_rpc_method_handler(
                    servicer.ListCatalogs,
                    request_deserializer=catalog_dot_v1_dot_catalog__pb2.ListCatalogsRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_catalog__pb2.ListCatalogsResponse.SerializeToString,
            ),
            'UpdateCatalog': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateCatalog,
                    request_deserializer=catalog_dot_v1_dot_catalog__pb2.UpdateCatalogRequest.FromString,
                    response_serializer=catalog_dot_v1_dot_catalog__pb2.Catalog.SerializeToString,
            ),
            'DeleteCatalog': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteCatalog,
                    request_deserializer=catalog_dot_v1_dot_catalog__pb2.DeleteCatalogRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'catalog.v1.CatalogService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CatalogService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateCompany(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/CreateCompany',
            catalog_dot_v1_dot_company__pb2.CreateCompanyRequest.SerializeToString,
            catalog_dot_v1_dot_company__pb2.Company.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCompany(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/GetCompany',
            catalog_dot_v1_dot_company__pb2.GetCompanyRequest.SerializeToString,
            catalog_dot_v1_dot_company__pb2.Company.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListCompanies(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/ListCompanies',
            catalog_dot_v1_dot_company__pb2.ListCompaniesRequest.SerializeToString,
            catalog_dot_v1_dot_company__pb2.ListCompaniesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateCompany(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/UpdateCompany',
            catalog_dot_v1_dot_company__pb2.UpdateCompanyRequest.SerializeToString,
            catalog_dot_v1_dot_company__pb2.Company.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteCompany(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/DeleteCompany',
            catalog_dot_v1_dot_company__pb2.DeleteCompanyRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateCatalog(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/CreateCatalog',
            catalog_dot_v1_dot_catalog__pb2.CreateCatalogRequest.SerializeToString,
            catalog_dot_v1_dot_catalog__pb2.Catalog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCatalog(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/GetCatalog',
            catalog_dot_v1_dot_catalog__pb2.GetCatalogRequest.SerializeToString,
            catalog_dot_v1_dot_catalog__pb2.Catalog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListCatalogs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/ListCatalogs',
            catalog_dot_v1_dot_catalog__pb2.ListCatalogsRequest.SerializeToString,
            catalog_dot_v1_dot_catalog__pb2.ListCatalogsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateCatalog(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/UpdateCatalog',
            catalog_dot_v1_dot_catalog__pb2.UpdateCatalogRequest.SerializeToString,
            catalog_dot_v1_dot_catalog__pb2.Catalog.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteCatalog(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/catalog.v1.CatalogService/DeleteCatalog',
            catalog_dot_v1_dot_catalog__pb2.DeleteCatalogRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
