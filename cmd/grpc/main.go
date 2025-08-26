package main

import (
	"log"
	"net"

	"inventory-service/helpers/database"
	slogLogger "inventory-service/helpers/logger"
	"inventory-service/helpers/utils/converter"
	"inventory-service/helpers/xvalidator"
	_middelware "inventory-service/module/auth/delivery/middleware_grpc"

	"inventory-service/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/grpc"

	_userGrpc "inventory-service/module/user/delivery/grpc"
	_userRepository "inventory-service/module/user/repository/postgres"
	_userUseCase "inventory-service/module/user/usecase"

	_authRepository "inventory-service/module/auth/repository/postgres"
	_authUseCase "inventory-service/module/auth/usecase"
)

func main() {
	validate, _ := xvalidator.NewValidator()
	conf := config.InitConfig(validate)
	slogLogger.SetupLogger(&slogLogger.Config{
		CurrentEnv: conf.AppEnv.CurrentEnv,
		LogPath:    conf.AppEnv.LogFilePath,
	})
	env := conf.AppEnv.CurrentEnv
	if env == "" {
		log.Println("Service RUN on DEBUG mode")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Recover())
	dbGorm := database.NewDatabase("postgres", &database.Config{
		DbHost: conf.Database.Pghost,
		DbUser: conf.Database.Pguser,
		DbPass: conf.Database.Pgpassword,
		DbName: conf.Database.Pgdatabase,
		DbPort: converter.ToString(conf.Database.Pgport),
	})

	userRepository := _userRepository.NewUserRepository()
	userUseCase := _userUseCase.NewUserUseCase(dbGorm.GetDB(), userRepository)

	ordersRepository := _ordersRepository.NewOrdersRepository()
	ordersUseCase := _ordersUseCase.NewOrdersUseCase(dbGorm.GetDB(), ordersRepository)

	inventoryRepository := _inventoryRepository.NewInventoryRepository()
	inventoryUseCase := _inventoryUseCase.NewInventoryUseCase(dbGorm.GetDB(), ordersRepository)

	authRepository := _authRepository.NewAuthRepository()
	authUseCase := _authUseCase.NewAuthUseCase(dbGorm.GetDB(), authRepository, userRepository)

	middlewareJWT := _middelware.NewAuthenticationJWT(authUseCase, map[string][]string{
		"AuthService": {"LoginUser", "RegisterUser"}, // ← ini yang benar
	})
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middlewareJWT.JwtInterceptor),
	)

	_userGrpc.NewUserService(grpcServer, userUseCase)
	_ordersGrpc.NewOrdersService(grpcServer, ordersUseCase)
	_inventoryGrpc.NewInventoryService(grpcServer, inventoryUseCase)

	httpPort := conf.AppEnv.HttpPort

	if httpPort == "" {
		httpPort = "9090"
	}

	lis, err := net.Listen("tcp", ":"+conf.AppEnv.HttpPort)
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}

	log.Fatal(grpcServer.Serve(lis))

}
