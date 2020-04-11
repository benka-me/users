package rpc

//this is your entry point server, it will not be generated again.
import (
	"fmt"
	"github.com/benka-me/laruche/go-pkg/discover"
	"github.com/benka-me/users/go-pkg/config"
	"github.com/benka-me/users/go-pkg/db"
	"github.com/benka-me/users/go-pkg/users"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// This structure will be passed to your handlers. Add everything you need inside.
type App struct {
	Clients
	Engine *discover.Engine
	Config *config.Config
	DB     *gorm.DB
}

// your server port, don't change it unless you update the service on the hub.

var grpcServer *grpc.Server

func Server_2_0(engine discover.Engine) {
	var err error
	port, err := engine.ThisPort("benka-me/users")
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		Clients: InitClients(engine, grpc.WithInsecure()), // Init clients of dependencies services
		Engine:  &engine,
		Config:  config.Init(engine.Dev),
	}
	app.DB = db.Init(app.Config, engine.Dev)

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		users.RegisterUsersServer(grpcServer, app) // Register your service server.
		reflection.Register(grpcServer)
	}

	fmt.Println("benka-me/users service running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
