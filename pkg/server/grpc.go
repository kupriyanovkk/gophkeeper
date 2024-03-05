package server

import (
	"context"
	"crypto/tls"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
	"github.com/kupriyanovkk/gophkeeper/pkg/cert"
)

type Service interface {
	RegisterService(grpc.ServiceRegistrar)
}

type GRPCServer struct {
	logger             *zap.Logger
	config             config.Config
	server             *grpc.Server
	services           []Service
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

type GRPCServerOption func(*GRPCServer)

// UseConfig returns a GRPCServerOption that sets the configuration for the server.
//
// It takes a config.Config parameter and returns a GRPCServerOption.
func UseConfig(config config.Config) GRPCServerOption {
	return func(s *GRPCServer) {
		s.config = config
	}
}

// UseLogger sets the logger for the GRPC server.
//
// It takes a pointer to a zap.Logger as a parameter and returns a GRPCServerOption.
func UseLogger(logger *zap.Logger) GRPCServerOption {
	return func(s *GRPCServer) {
		s.logger = logger
	}
}

// UseServices creates a GRPCServerOption with the provided services.
//
// It takes a variadic parameter of type Service and returns a GRPCServerOption.
func UseServices(services ...Service) GRPCServerOption {
	return func(s *GRPCServer) {
		s.services = services
	}
}

// UseStreamInterceptors returns a GRPCServerOption function that adds the provided stream server interceptors to the server.
//
// The function takes a variadic parameter of type grpc.StreamServerInterceptor.
// It returns a GRPCServerOption.
func UseStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) GRPCServerOption {
	return func(server *GRPCServer) {
		server.streamInterceptors = append(server.streamInterceptors, interceptors...)
	}
}

// UseUnaryInterceptors creates a GRPC server option to set the unary server interceptors.
//
// The function takes a variadic parameter of type grpc.UnaryServerInterceptor to be used as unary server interceptors.
// It returns a GRPCServerOption.
func UseUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) GRPCServerOption {
	return func(server *GRPCServer) {
		server.unaryInterceptors = append(server.unaryInterceptors, interceptors...)
	}
}

// RegisterServices registers the given services to the GRPCServer.
//
// services ...Service
func (s *GRPCServer) RegisterServices(services ...Service) {
	for _, service := range services {
		service.RegisterService(s.server)
	}
}

// Start starts the gRPC server.
//
// It takes a cancel function as a parameter and does not return anything.
func (s *GRPCServer) Start(cancel context.CancelFunc) {
	sslConfig, err := cert.NewSSLConfigService().LoadServerCertificate(s.config)
	conn, errListen := tls.Listen("tcp", ":"+s.config.Port, sslConfig)
	if errListen != nil {
		s.logger.Error(errListen.Error())
	}

	s.server = grpc.NewServer(
		grpc.ChainStreamInterceptor(s.streamInterceptors...),
		grpc.ChainUnaryInterceptor(s.unaryInterceptors...),
	)

	s.RegisterServices(s.services...)
	go func() {
		err = s.server.Serve(conn)
		if err != nil {
			s.logger.Error(err.Error())
			cancel()
		}
	}()

	s.logger.Sugar().Infof("gRPC server is running on %s port", s.config.Port)
}

// Stop gracefully stops the gRPC server.
//
// No parameters.
// No return type.
func (s *GRPCServer) Stop() {
	s.logger.Info("Gracefully stopping gRPC server")

	s.server.GracefulStop()

	s.logger.Info("gRPC server stopped")
}

// NewGRPCServer creates a new GRPCServer with the given options.
//
// The function takes a variadic number of GRPCServerOption(s) as its parameter(s) and returns a pointer to GRPCServer.
func NewGRPCServer(opts ...GRPCServerOption) *GRPCServer {
	s := &GRPCServer{}

	for _, option := range opts {
		option(s)
	}

	return s
}
