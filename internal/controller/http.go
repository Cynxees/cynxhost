package controller

import (
	"cynxhost/internal/app"
	"cynxhost/internal/controller/persistentnodecontroller"
	"cynxhost/internal/controller/servertemplatecontroller"
	"cynxhost/internal/controller/usercontroller"
	"cynxhost/internal/middleware"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.elastic.co/apm/module/apmhttp/v2"
)

type HttpServer struct {
	*http.Server
}

func NewHttpServer(app *app.App) (*HttpServer, error) {

	config := app.Dependencies.Config

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{config.Security.CORS.Origin},               // replace with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // allowed methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"},           // allowed headers
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	routerPath := app.Dependencies.Config.Router.Default
	debug := app.Dependencies.Config.App.Debug

	handleRouterFunc := func(path string, handler middleware.HandlerFuncWithHelper, requireAuth bool) *mux.Route {
		wrappedHandler := middleware.WrapHandler(handler, debug)

		if requireAuth && !debug {
			wrappedHandler = middleware.AuthMiddleware(app.Dependencies.JWTManager, wrappedHandler, debug)
		}
		return r.HandleFunc(routerPath+path, wrappedHandler).Methods("POST", "GET")
	}

	userController := usercontroller.New(app.Usecases.UserUseCase, app.Dependencies.Validator, app.Dependencies.Config)
	serverTemplateController := servertemplatecontroller.New(app.Usecases.ServerTemplateUseCase, app.Dependencies.Validator)
	persistentNodeController := persistentnodecontroller.New(app.Usecases.PersistentNodeUseCase, app.Dependencies.Validator)

	// User
	handleRouterFunc("user/register", userController.RegisterUser, false)
	handleRouterFunc("user/login", userController.LoginUser, false)
	handleRouterFunc("user/logout", userController.LogoutUser, false)
	handleRouterFunc("user/check-username", userController.CheckUsername, false)
	handleRouterFunc("user/paginate", userController.PaginateUser, true)
	handleRouterFunc("user/profile", userController.GetProfile, true)

	// Server Template
	handleRouterFunc("server-template/paginate", serverTemplateController.PaginateServerTemplate, true)
	handleRouterFunc("server-template/detail", serverTemplateController.GetServerTemplate, true)
	handleRouterFunc("server-template/categories", serverTemplateController.GetServerTemplateCategories, true)

	// Persistent Node
	handleRouterFunc("persistent-node/show-owned", persistentNodeController.GetAllPersistentNodesFromUser, true)
	handleRouterFunc("persistent-node/detail", persistentNodeController.GetPersistentNodeDetail, true)
	handleRouterFunc("persistent-node/create", persistentNodeController.CreatePersistentNode, true)
	handleRouterFunc("persistent-node/force-shutdown", persistentNodeController.ForceShutdownPersistentNode, true)

	// Callbacks
	handleRouterFunc("persistent-node/callback/launch", persistentNodeController.LaunchCallbackPersistentNode, false)
	handleRouterFunc("persistent-node/callback/shutdown", persistentNodeController.ShutdownCallbackPersistentNode, false)
	handleRouterFunc("persistent-node/callback/update-status", persistentNodeController.StatusCallbackPersistentNode, false)

	// Persistent Node Dashboard
	handleRouterFunc("persistent-node/dashboard/send-command", persistentNodeController.SendCommandPersistentNode, true)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			app.Dependencies.Logger.Errorf("Failed to write response: %v", err)
		}
	})

	corsHandler := c.Handler(r)

	address := app.Dependencies.Config.App.Address + ":" + strconv.Itoa(app.Dependencies.Config.App.Port)
	app.Dependencies.Logger.Infof("Starting http server on %s\n", address)

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      apmhttp.Wrap(corsHandler),
	}

	return &HttpServer{srv}, nil
}

func (s *HttpServer) Start() error {
	return s.ListenAndServe()
}

func (s *HttpServer) Stop() error {
	return errors.New("http stop not implemented")
}
