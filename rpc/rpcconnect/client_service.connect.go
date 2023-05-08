// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: client_service.proto

package rpcconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	rpc "github.com/gokch/ipfs_mount/rpc"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ClientServiceName is the fully-qualified name of the ClientService service.
	ClientServiceName = "proto.ClientService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClientServiceConnectProcedure is the fully-qualified name of the ClientService's Connect RPC.
	ClientServiceConnectProcedure = "/proto.ClientService/Connect"
	// ClientServiceDisconnectProcedure is the fully-qualified name of the ClientService's Disconnect
	// RPC.
	ClientServiceDisconnectProcedure = "/proto.ClientService/Disconnect"
	// ClientServiceIsConnectProcedure is the fully-qualified name of the ClientService's IsConnect RPC.
	ClientServiceIsConnectProcedure = "/proto.ClientService/IsConnect"
	// ClientServiceUploadProcedure is the fully-qualified name of the ClientService's Upload RPC.
	ClientServiceUploadProcedure = "/proto.ClientService/Upload"
	// ClientServiceDownloadProcedure is the fully-qualified name of the ClientService's Download RPC.
	ClientServiceDownloadProcedure = "/proto.ClientService/Download"
)

// ClientServiceClient is a client for the proto.ClientService service.
type ClientServiceClient interface {
	Connect(context.Context, *connect_go.Request[rpc.ConnectRequest]) (*connect_go.Response[rpc.ConnectResponse], error)
	Disconnect(context.Context, *connect_go.Request[rpc.DisconnectRequest]) (*connect_go.Response[rpc.DisconnectResponse], error)
	IsConnect(context.Context, *connect_go.Request[rpc.IsConnectRequest]) (*connect_go.Response[rpc.IsConnectResponse], error)
	Upload(context.Context, *connect_go.Request[rpc.UploadRequest]) (*connect_go.Response[rpc.UploadResponse], error)
	Download(context.Context, *connect_go.Request[rpc.DownloadRequest]) (*connect_go.Response[rpc.DownloadResponse], error)
}

// NewClientServiceClient constructs a client for the proto.ClientService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClientServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ClientServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clientServiceClient{
		connect: connect_go.NewClient[rpc.ConnectRequest, rpc.ConnectResponse](
			httpClient,
			baseURL+ClientServiceConnectProcedure,
			opts...,
		),
		disconnect: connect_go.NewClient[rpc.DisconnectRequest, rpc.DisconnectResponse](
			httpClient,
			baseURL+ClientServiceDisconnectProcedure,
			opts...,
		),
		isConnect: connect_go.NewClient[rpc.IsConnectRequest, rpc.IsConnectResponse](
			httpClient,
			baseURL+ClientServiceIsConnectProcedure,
			opts...,
		),
		upload: connect_go.NewClient[rpc.UploadRequest, rpc.UploadResponse](
			httpClient,
			baseURL+ClientServiceUploadProcedure,
			opts...,
		),
		download: connect_go.NewClient[rpc.DownloadRequest, rpc.DownloadResponse](
			httpClient,
			baseURL+ClientServiceDownloadProcedure,
			opts...,
		),
	}
}

// clientServiceClient implements ClientServiceClient.
type clientServiceClient struct {
	connect    *connect_go.Client[rpc.ConnectRequest, rpc.ConnectResponse]
	disconnect *connect_go.Client[rpc.DisconnectRequest, rpc.DisconnectResponse]
	isConnect  *connect_go.Client[rpc.IsConnectRequest, rpc.IsConnectResponse]
	upload     *connect_go.Client[rpc.UploadRequest, rpc.UploadResponse]
	download   *connect_go.Client[rpc.DownloadRequest, rpc.DownloadResponse]
}

// Connect calls proto.ClientService.Connect.
func (c *clientServiceClient) Connect(ctx context.Context, req *connect_go.Request[rpc.ConnectRequest]) (*connect_go.Response[rpc.ConnectResponse], error) {
	return c.connect.CallUnary(ctx, req)
}

// Disconnect calls proto.ClientService.Disconnect.
func (c *clientServiceClient) Disconnect(ctx context.Context, req *connect_go.Request[rpc.DisconnectRequest]) (*connect_go.Response[rpc.DisconnectResponse], error) {
	return c.disconnect.CallUnary(ctx, req)
}

// IsConnect calls proto.ClientService.IsConnect.
func (c *clientServiceClient) IsConnect(ctx context.Context, req *connect_go.Request[rpc.IsConnectRequest]) (*connect_go.Response[rpc.IsConnectResponse], error) {
	return c.isConnect.CallUnary(ctx, req)
}

// Upload calls proto.ClientService.Upload.
func (c *clientServiceClient) Upload(ctx context.Context, req *connect_go.Request[rpc.UploadRequest]) (*connect_go.Response[rpc.UploadResponse], error) {
	return c.upload.CallUnary(ctx, req)
}

// Download calls proto.ClientService.Download.
func (c *clientServiceClient) Download(ctx context.Context, req *connect_go.Request[rpc.DownloadRequest]) (*connect_go.Response[rpc.DownloadResponse], error) {
	return c.download.CallUnary(ctx, req)
}

// ClientServiceHandler is an implementation of the proto.ClientService service.
type ClientServiceHandler interface {
	Connect(context.Context, *connect_go.Request[rpc.ConnectRequest]) (*connect_go.Response[rpc.ConnectResponse], error)
	Disconnect(context.Context, *connect_go.Request[rpc.DisconnectRequest]) (*connect_go.Response[rpc.DisconnectResponse], error)
	IsConnect(context.Context, *connect_go.Request[rpc.IsConnectRequest]) (*connect_go.Response[rpc.IsConnectResponse], error)
	Upload(context.Context, *connect_go.Request[rpc.UploadRequest]) (*connect_go.Response[rpc.UploadResponse], error)
	Download(context.Context, *connect_go.Request[rpc.DownloadRequest]) (*connect_go.Response[rpc.DownloadResponse], error)
}

// NewClientServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClientServiceHandler(svc ClientServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ClientServiceConnectProcedure, connect_go.NewUnaryHandler(
		ClientServiceConnectProcedure,
		svc.Connect,
		opts...,
	))
	mux.Handle(ClientServiceDisconnectProcedure, connect_go.NewUnaryHandler(
		ClientServiceDisconnectProcedure,
		svc.Disconnect,
		opts...,
	))
	mux.Handle(ClientServiceIsConnectProcedure, connect_go.NewUnaryHandler(
		ClientServiceIsConnectProcedure,
		svc.IsConnect,
		opts...,
	))
	mux.Handle(ClientServiceUploadProcedure, connect_go.NewUnaryHandler(
		ClientServiceUploadProcedure,
		svc.Upload,
		opts...,
	))
	mux.Handle(ClientServiceDownloadProcedure, connect_go.NewUnaryHandler(
		ClientServiceDownloadProcedure,
		svc.Download,
		opts...,
	))
	return "/proto.ClientService/", mux
}

// UnimplementedClientServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClientServiceHandler struct{}

func (UnimplementedClientServiceHandler) Connect(context.Context, *connect_go.Request[rpc.ConnectRequest]) (*connect_go.Response[rpc.ConnectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.ClientService.Connect is not implemented"))
}

func (UnimplementedClientServiceHandler) Disconnect(context.Context, *connect_go.Request[rpc.DisconnectRequest]) (*connect_go.Response[rpc.DisconnectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.ClientService.Disconnect is not implemented"))
}

func (UnimplementedClientServiceHandler) IsConnect(context.Context, *connect_go.Request[rpc.IsConnectRequest]) (*connect_go.Response[rpc.IsConnectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.ClientService.IsConnect is not implemented"))
}

func (UnimplementedClientServiceHandler) Upload(context.Context, *connect_go.Request[rpc.UploadRequest]) (*connect_go.Response[rpc.UploadResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.ClientService.Upload is not implemented"))
}

func (UnimplementedClientServiceHandler) Download(context.Context, *connect_go.Request[rpc.DownloadRequest]) (*connect_go.Response[rpc.DownloadResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.ClientService.Download is not implemented"))
}
