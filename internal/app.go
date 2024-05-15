package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wildegor/kaspi-rest/internal/configs"
	"github.com/wildegor/kaspi-rest/internal/db/postgres"
	get_people_info_handler "github.com/wildegor/kaspi-rest/internal/handlers/add_people_info"
	find_people_info_by_iin_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_iin"
	find_people_info_by_phone_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_key"
	health_check_handler "github.com/wildegor/kaspi-rest/internal/handlers/health_check"
	iin_check_handler "github.com/wildegor/kaspi-rest/internal/handlers/iin_check"
	users_repository "github.com/wildegor/kaspi-rest/internal/repositories/users"
	"github.com/wildegor/kaspi-rest/internal/routers"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Server struct {
	App    *http.Server
	DbConn *postgres.PostgresConnection
}

func (srv *Server) Run(ctx context.Context) {
	slog.Info("server running!")

	log.Fatal(srv.App.ListenAndServe())
}

func (srv *Server) Shutdown(ctx context.Context) {
	if err := srv.App.Shutdown(ctx); err != nil {
		slog.Error("fail shutdown server")
	}
}

func NewServer() (*Server, error) {
	// Init configs
	c := configs.NewConfigurator()
	ac := configs.NewAppConfig(c)
	lc := configs.NewLoggerConfig()
	pc := configs.NewPostgresConfig(c)

	// Setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: lc.Level,
	}))
	if lc.Format == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: lc.Level,
		}))
	}
	slog.SetDefault(logger)

	// Setup db
	dbConn := postgres.NewPostgresConnection(pc)

	// Setup repos
	ur := users_repository.NewUsersRepository(dbConn)

	// Setup handlers
	ich := iin_check_handler.NewCheckIINHandler()
	gph := get_people_info_handler.NewAddPeopleHandler(ur)
	fpph := find_people_info_by_phone_handler.NewFindPeopleInfoByKeyHandler(ur)
	fpih := find_people_info_by_iin_handler.NewFindPeopleInfoByIINHandler(ur)
	hh := health_check_handler.NewHealthCheckHandler()

	// Setup routers
	r := mux.NewRouter()
	hcr := routers.NewHealthRouter(hh)
	hcr.Setup(r)
	ccr := routers.NewCheckRouter(ich)
	ccr.Setup(r)
	pr := routers.NewPeopleRouter(gph, fpih, fpph)
	pr.Setup(r)
	sr := routers.NewSwaggerRouter()
	sr.Setup(r)

	// Setup server
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%d", ac.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Server{
		App:    srv,
		DbConn: dbConn,
	}, nil
}
