package corebanking

import (
	"context"
	entitiesServer "core-system/core/entities/server"
	"core-system/core/utils/database/postgres"
	"core-system/core/utils/getenv"
	"core-system/core/utils/logger"
	"core-system/core/utils/server"
	"core-system/service/core-banking/controller/inquiry"
	repository "core-system/service/core-banking/repository"
	usecaseInquiry "core-system/service/core-banking/usecase/inquiry"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func InitServiceCoreBanking() CoreBankingInterface {
	engine := gin.New()

	log := logger.NewLoggerService(&logger.Config{
		Level: "",
		Color: true,
	})

	server, errServer := server.InitServer(entitiesServer.ServerConfig{
		Port:         getenv.Getenv[string]("PORT"),
		Handler:      engine,
		IdleTimeout:  getenv.Getenv[int]("IDLE_TIMEOUT"),
		ReadTimeout:  getenv.Getenv[int]("READ_TIMEOUT"),
		WriteTimeout: getenv.Getenv[int]("WRITE_TIMEOUT"),
	})

	if errServer != nil {
		logs.Fatalf("[!] Error Initialize Server: %+v", errServer)
	}
	repo := repository.InitRepo(postgres.ConfigPostgres{
		Username: getenv.Getenv[string]("USER_DATABASE"),
		Password: getenv.Getenv[string]("PASS_DATABASE"),
		Host:     getenv.Getenv[string]("HOST_DATABASE"),
		Port:     getenv.Getenv[string]("PORT_DATABASE"),
		DbName:   getenv.Getenv[string]("DB_DATABASE"),
		Debug:    getenv.Getenv[string]("DEBUG_DATABASE"),
		Mode:     getenv.Getenv[string]("LOG_MODE_DATABASE"),
	})
	usecase, errUc := usecaseInquiry.InitUsecase(usecaseInquiry.UsecaseConfig{
		Repo: repo,
	})

	if errUc != nil {
		logs.Fatalf("[!] Error Initialize Usecase: %v", errUc)
	}
	inquiryCtrl, errCtrl := inquiry.InitInquiryController(inquiry.ControllerConfig{
		Usecase: usecase,
		Log:     log,
	})

	if errCtrl != nil {
		logs.Fatalf("[!] Error Initilization Controller Inquiry: %v", errCtrl)
	}

	return &coreBankingService{
		inquiry: inquiryCtrl,
		engine:  engine,
		server:  server,
	}
}

func (r *coreBankingService) InitRoutes() error {
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	inquiry := r.engine.Group("/inquiry")
	{
		inquiry.GET("/:userID", r.inquiry.GetBalanceByUserID)
		inquiry.POST("/withdraw", r.inquiry.WithdrawByUserID)
	}

	if err := r.server.StartServer(); err != nil {
		return err
	}

	return nil
}

func (r *coreBankingService) Stop(ctx context.Context) error {
	return r.server.ShutdownServer(ctx)
}
