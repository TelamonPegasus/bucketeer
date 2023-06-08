# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.auth import service_pb2 as proto_dot_auth_dot_service__pb2


class AuthServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetAuthCodeURL = channel.unary_unary(
                '/bucketeer.auth.AuthService/GetAuthCodeURL',
                request_serializer=proto_dot_auth_dot_service__pb2.GetAuthCodeURLRequest.SerializeToString,
                response_deserializer=proto_dot_auth_dot_service__pb2.GetAuthCodeURLResponse.FromString,
                )
        self.ExchangeToken = channel.unary_unary(
                '/bucketeer.auth.AuthService/ExchangeToken',
                request_serializer=proto_dot_auth_dot_service__pb2.ExchangeTokenRequest.SerializeToString,
                response_deserializer=proto_dot_auth_dot_service__pb2.ExchangeTokenResponse.FromString,
                )
        self.RefreshToken = channel.unary_unary(
                '/bucketeer.auth.AuthService/RefreshToken',
                request_serializer=proto_dot_auth_dot_service__pb2.RefreshTokenRequest.SerializeToString,
                response_deserializer=proto_dot_auth_dot_service__pb2.RefreshTokenResponse.FromString,
                )


class AuthServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetAuthCodeURL(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ExchangeToken(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RefreshToken(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetAuthCodeURL': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAuthCodeURL,
                    request_deserializer=proto_dot_auth_dot_service__pb2.GetAuthCodeURLRequest.FromString,
                    response_serializer=proto_dot_auth_dot_service__pb2.GetAuthCodeURLResponse.SerializeToString,
            ),
            'ExchangeToken': grpc.unary_unary_rpc_method_handler(
                    servicer.ExchangeToken,
                    request_deserializer=proto_dot_auth_dot_service__pb2.ExchangeTokenRequest.FromString,
                    response_serializer=proto_dot_auth_dot_service__pb2.ExchangeTokenResponse.SerializeToString,
            ),
            'RefreshToken': grpc.unary_unary_rpc_method_handler(
                    servicer.RefreshToken,
                    request_deserializer=proto_dot_auth_dot_service__pb2.RefreshTokenRequest.FromString,
                    response_serializer=proto_dot_auth_dot_service__pb2.RefreshTokenResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'bucketeer.auth.AuthService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AuthService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetAuthCodeURL(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/bucketeer.auth.AuthService/GetAuthCodeURL',
            proto_dot_auth_dot_service__pb2.GetAuthCodeURLRequest.SerializeToString,
            proto_dot_auth_dot_service__pb2.GetAuthCodeURLResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ExchangeToken(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/bucketeer.auth.AuthService/ExchangeToken',
            proto_dot_auth_dot_service__pb2.ExchangeTokenRequest.SerializeToString,
            proto_dot_auth_dot_service__pb2.ExchangeTokenResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RefreshToken(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/bucketeer.auth.AuthService/RefreshToken',
            proto_dot_auth_dot_service__pb2.RefreshTokenRequest.SerializeToString,
            proto_dot_auth_dot_service__pb2.RefreshTokenResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)