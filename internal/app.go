package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wildegor/kaspi-rest/internal/configs"
	"github.com/wildegor/kaspi-rest/internal/db/postgres"
	find_people_info_by_iin_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_iin"
	find_people_info_by_phone_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_phone"
	get_people_info_handler "github.com/wildegor/kaspi-rest/internal/handlers/get_people_info"
	iin_check_handler "github.com/wildegor/kaspi-rest/internal/handlers/iin_check"
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

	// Setup handlers
	ich := iin_check_handler.NewCheckIINHandler()
	gph := get_people_info_handler.NewGetPeopleHandler()
	fpph := find_people_info_by_phone_handler.NewFindPeopleInfoByPhoneHandler()
	fpih := find_people_info_by_iin_handler.NewFindPeopleInfoByIINHandler()

	// Setup routers
	r := mux.NewRouter()
	api := r.PathPrefix("/api")
	hcr := routers.NewHealthRouter()
	hcr.Setup(api)
	ccr := routers.NewCheckRouter(ich)
	ccr.Setup(api)
	pr := routers.NewPeopleRouter(gph, fpih, fpph)
	pr.Setup(api)

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
