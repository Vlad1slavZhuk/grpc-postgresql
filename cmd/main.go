package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pbShop "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/shop/v1"
	pbUser "github.com/Vlad1slavZhuk/grpc-postgresql/api/gen/user/v1"
	"github.com/Vlad1slavZhuk/grpc-postgresql/configs"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/auth"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/databases"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/handler"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/hash"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/log"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/repository"
	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := log.GetLoggerInstance()
	config := configs.ReadConfig()

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// postgresql

	db, err := databases.NewPostgreSQL(config.DSN())
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
		return
	}

	// token

	tokenManager, err := auth.NewManager("test")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
		return
	}

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repos:        repos,
		Hasher:       hash.NewSHA1Hasher("test"),
		TokenManager: tokenManager,
	})

	handler := handler.NewHandler(services, tokenManager)

	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(nil),
			)),
	)

	pbShop.RegisterShopServiceServer(gRPCServer, handler)
	pbUser.RegisterUserServiceServer(gRPCServer, handler)
	reflection.Register(gRPCServer)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		s := <-signalCh
		log.Info().Msgf("got signal %v, attempting graceful shutdown", s)
		cancel()

		gRPCServer.GracefulStop()
		wg.Done()
	}()

	gRPCAddr := fmt.Sprintf("%v:%v", config.GRPCSettings.Host, config.GRPCSettings.Port)
	lis, err := net.Listen("tcp", gRPCAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	log.Info().Msgf("starting gRPC server - %s", gRPCAddr)
	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatal().Err(err).Msg("could not start gRPC server:")
	}

	wg.Wait()
	log.Info().Msg("shutdown is done")
}
