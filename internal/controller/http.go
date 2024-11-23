package controller

import (
	"cynxhost/internal/app"
	"cynxhost/internal/controller/registerusercontroller"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmhttp/v2"
)

type HttpServer struct {
	*http.Server
}

func NewHttpServer(app *app.App) (*HttpServer, error) {

	r := mux.NewRouter()

	registerUserController := registerusercontroller.New(app.Usecases.RegisterUserUseCase, app.Dependencies.Validator)

	routerPath := app.Dependencies.Config.Router.Default
	fmt.Println("routerPath: ", routerPath)
	r.HandleFunc(routerPath + "/register", registerUserController.RegisterUser).Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			app.Dependencies.Logger.Errorf("Failed to write response: %v", err)
		}
	})

	address := app.Dependencies.Config.App.Address + ":" + strconv.Itoa(app.Dependencies.Config.App.Port)
	app.Dependencies.Logger.Infof("Starting http server on %s\n", address)

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      apmhttp.Wrap(r),
	}

	return &HttpServer{srv}, nil
}

func (s *HttpServer) Start() error {
	return s.ListenAndServe()
}

func (s *HttpServer) Stop() error {
	return errors.New("http stop not implemented")
}
