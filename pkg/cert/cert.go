package cert

import (
	"crypto/tls"

	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"

	clientConfig "github.com/kupriyanovkk/gophkeeper/internal/client/config"
	serverConfig "github.com/kupriyanovkk/gophkeeper/internal/server/config"
	"github.com/kupriyanovkk/gophkeeper/pkg/logger"
)

type SSLConfigLoaderService interface {
	LoadClientCertificate(config clientConfig.Config) (credentials.TransportCredentials, error)
	LoadServerCertificate(config serverConfig.Config) (*tls.Config, error)
}

type sslConfigService struct {
	logger *zap.Logger
}

// NewSSLConfigService initializes a new SSL configuration service.
//
// No parameters.
// Returns a pointer to sslConfigService.
func NewSSLConfigService() *sslConfigService {
	return &sslConfigService{
		logger: logger.NewLogger(),
	}
}

// LoadServerCertificate loads the server certificate for the SSL configuration.
//
// It takes a serverConfig.Config parameter and returns a *tls.Config and an error.
func (s sslConfigService) LoadServerCertificate(config serverConfig.Config) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(config.SSLCertPath, config.SSLKeyPath)
	if err != nil {
		s.logger.Error(err.Error())

		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

// LoadClientCertificate loads the client certificate for SSL configuration.
//
// It takes a clientConfig.Config parameter and returns credentials.TransportCredentials and error.
func (s sslConfigService) LoadClientCertificate(config clientConfig.Config) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(config.SSLCertPath, config.SSLKeyPath)
	if err != nil {
		s.logger.Error(err.Error())

		return nil, err
	}

	return credentials.NewTLS(
		&tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	), nil
}
