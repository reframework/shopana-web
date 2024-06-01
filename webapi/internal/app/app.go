package application

import (
	"fmt"
	"log"

	graphqlHandler "webapi/internal/http/graphql"
	"webapi/internal/http/server/middlewares"
	"webapi/internal/service"
	postgres "webapi/pkg/postgres/connect"

	app "webapi/internal/app/interface"
	server "webapi/internal/http/server"
	repository "webapi/internal/storage"
	isql "webapi/pkg/postgres/interface"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

type App struct {
	logger    *slog.Logger
	db        *sqlx.DB
	server    *fiber.App
	service   app.Service
	storage   app.Storage
	validator app.Validator
}

func NewApp() *App {
	app := &App{}

	db := postgres.CreatePostgreSQL(&postgres.Input{
		DSN:             viper.GetString("DB"),
		SetMaxIdeConns:  viper.GetInt("DB_MAX_IDLE_CONNS"),
		SetMaxOpenConns: viper.GetInt("DB_MAX_OPEN_CONNS"),
	})

	logger := CreateLogger()
	server := server.NewServer()
	storage := repository.NewStorage(app)
	service := service.NewService(app)
	validator := validator.New()

	app.db = db
	app.logger = logger
	app.server = server
	app.validator = validator
	app.storage = storage
	app.service = service

	server.Use(middlewares.ReqIdMiddleware())
	graphqlHandler.InitRouter(server, app)

	return app
}

func (a *App) Run() {
	PORT := viper.GetString("PORT")

	func() {
		if err := a.server.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", PORT)); err != nil {
			log.Fatalf("Error occurred while running HTTP server: %s", err.Error())
		}
	}()
}

func (a App) Service() app.Service {
	return a.service
}

func (a App) Storage() app.Storage {
	return a.storage
}

func (a App) Logger() *slog.Logger {
	return a.logger
}

func (a App) DB() isql.Runner {
	return a.db
}

func (a App) Validator() app.Validator {
	return a.validator
}
