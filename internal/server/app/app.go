package app

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
	authMiddleware "github.com/kupriyanovkk/gophkeeper/internal/server/middleware/auth"
	"github.com/kupriyanovkk/gophkeeper/internal/server/service"
	"github.com/kupriyanovkk/gophkeeper/internal/server/storage/mongodb"
	"github.com/kupriyanovkk/gophkeeper/pkg/crypt"
	"github.com/kupriyanovkk/gophkeeper/pkg/jwt"
	"github.com/kupriyanovkk/gophkeeper/pkg/logger"
	"github.com/kupriyanovkk/gophkeeper/pkg/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type App struct {
	DB     *mongo.Client
	Server *server.GRPCServer
	Logger *zap.Logger
}

// InterceptorLogger creates a logging interceptor.
//
// It takes a zap logger as input and returns a logging.Logger.
func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

// NewApp initializes and returns a new App instance.
//
// ctx Context parameter. Returns *App and error.
func NewApp(ctx context.Context) (*App, error) {
	logger := logger.NewLogger()
	config := config.NewConfig()

	crypt, errCrypt := crypt.NewCrypt()
	if errCrypt != nil {
		return nil, fmt.Errorf("crypt.NewCrypt error: %w", errCrypt)
	}

	jwtManager, errJwt := jwt.NewJWT(config.JWTSecret, config.JWTExp)
	if errJwt != nil {
		return nil, fmt.Errorf("jwt.NewJWT error: %w", errJwt)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.DatabaseDSN).SetServerAPIOptions(serverAPI)
	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect error: %w", err)
	}

	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err = mongoClient.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("mongo.Ping error: %w", err)
	}

	privateStorage := mongodb.NewPrivateStore(mongoClient)
	privateService := service.NewPrivateService(privateStorage)

	userStorage := mongodb.NewUserStore(mongoClient)
	userService := service.NewUserService(userStorage, jwtManager, crypt)

	authMiddleware := authMiddleware.NewAuthMiddleware(jwtManager, crypt).Auth

	gRPCServer := server.NewGRPCServer(
		server.UseConfig(config),
		server.UseLogger(logger),
		server.UseServices(privateService, userService),
		server.UseStreamInterceptors(
			logging.StreamServerInterceptor(InterceptorLogger(logger)),
			auth.StreamServerInterceptor(authMiddleware),
			recovery.StreamServerInterceptor(),
		),
		server.UseUnaryInterceptors(
			logging.UnaryServerInterceptor(InterceptorLogger(logger)),
			auth.UnaryServerInterceptor(authMiddleware),
			recovery.UnaryServerInterceptor(),
		),
	)

	return &App{
		DB:     mongoClient,
		Server: gRPCServer,
		Logger: logger,
	}, nil
}
