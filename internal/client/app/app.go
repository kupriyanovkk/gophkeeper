package app

import (
	"context"

	"github.com/kupriyanovkk/gophkeeper/internal/client/config"
	"github.com/kupriyanovkk/gophkeeper/internal/client/interceptor"
	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"github.com/kupriyanovkk/gophkeeper/internal/client/service"
	"github.com/kupriyanovkk/gophkeeper/internal/client/storage"
	storageMemory "github.com/kupriyanovkk/gophkeeper/internal/client/storage/memory"
	storageSync "github.com/kupriyanovkk/gophkeeper/internal/client/storage/sync"
	"github.com/kupriyanovkk/gophkeeper/pkg/cert"
	"github.com/kupriyanovkk/gophkeeper/pkg/crypt"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"

	pb "github.com/kupriyanovkk/gophkeeper/proto"
)

type App struct {
	Storage        storage.MemoryAbstractStorage
	Cancel         context.CancelFunc
	PrivateService *service.PrivateService
	UserService    *service.UserService
	Cron           *cron.Cron
	Sync           storage.SyncAbstract
}

// NewApp initializes and returns a new App instance.
//
// No parameters.
// Returns *App and error.
func NewApp() (*App, error) {
	ctx, cancel := context.WithCancel(context.Background())
	globalCtx := model.GlobalContext{Ctx: ctx, Cancel: cancel}
	config := config.NewConfig()
	crypt, err := crypt.NewCrypt()
	if err != nil {
		return nil, err
	}

	tlsCreds, err := cert.NewSSLConfigService().LoadClientCertificate(config)
	if err != nil {
		return nil, err
	}

	interceptor := interceptor.NewInterceptor(map[string]bool{
		"/proto.Private/GetPrivateDataByType": true,
		"/proto.Private/CreatePrivateData":    true,
		"/proto.Private/GetPrivateData":       true,
		"/proto.Private/DeletePrivateData":    true,
		"/proto.Private/UpdatePrivateData":    true,
	})

	conn, err := grpc.Dial(":"+config.Port,
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		return nil, err
	}

	privateClient := pb.NewPrivateClient(conn)
	userClient := pb.NewUserClient(conn)

	storage := storageMemory.NewMemoryStorage()
	sync := storageSync.NewSync(storage, privateClient, &globalCtx, crypt)

	privateService := service.NewPrivateService(storage, &globalCtx, privateClient, crypt, sync)
	userService := service.NewUserService(&globalCtx, userClient)

	cron := cron.New()
	cron.AddFunc("* * * * *", sync.SyncAll)

	return &App{
		Storage:        storage,
		Cancel:         cancel,
		PrivateService: privateService,
		UserService:    userService,
		Cron:           cron,
		Sync:           sync,
	}, nil
}
