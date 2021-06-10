package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/service"

	"github.com/ferrandinand/cwh-lib/logger"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"AUTH_ADDRESS",
		"AUTH_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	//wiring
	dbClient := getDbClient()
	authClient := getAuthURL()
	userRepositoryDb := domain.NewUserRepositoryDb(dbClient)
	projectRepositoryDb := domain.NewProjectRepositoryDb(dbClient)
	environmentRepositoryDb := domain.NewEnvironmentRepositoryDb(dbClient)
	serviceOrderRepositoryDb := domain.NewServiceOrderRepositoryDb(dbClient)
	ch := UserHandlers{service.NewUserService(userRepositoryDb)}
	ph := ProjectHandler{service.NewProjectService(projectRepositoryDb)}
	eh := EnvironmentHandler{service.NewEnvironmentService(environmentRepositoryDb)}
	sh := ServiceOrderHandler{service.NewServiceOrderService(serviceOrderRepositoryDb)}

	// define routes
	router.
		HandleFunc("/users", ch.getAllUsers).
		Methods(http.MethodGet).
		Name("GetAllUsers")
	router.
		HandleFunc("/users/{user_id:[0-9]+}", ch.getUser).
		Methods(http.MethodGet).
		Name("GetUser")
	router.
		HandleFunc("/user", ch.getCurrentUser).
		Methods(http.MethodGet).
		Name("GetCurrentUser")
	router.
		HandleFunc("/users/new", ch.NewUser).
		Methods(http.MethodPost).
		Name("NewUser")
	router.
		HandleFunc("/users/{user_id:[0-9]+}", ch.UpdateUser).
		Methods(http.MethodPatch).
		Name("UpdateUser")
	router.
		HandleFunc("/users/{user_id:[0-9]+}", ch.DeleteUser).
		Methods(http.MethodDelete).
		Name("DeleteUser")
	router.
		HandleFunc("/project/new", ph.NewProject).
		Methods(http.MethodPost).
		Name("NewProject")
	router.
		HandleFunc("/project/{project_id}", ph.GetProject).
		Methods(http.MethodGet).
		Name("GetProject")
	router.
		HandleFunc("/project", ph.GetAllProject).
		Methods(http.MethodGet).
		Name("GetAllProject")
	router.
		HandleFunc("/project/{project_id}/environments/new", eh.NewEnvironment).
		Methods(http.MethodPost).
		Name("NewEnvironment")
	router.
		HandleFunc("/project/{project_id}/environments", eh.GetAllEnvironment).
		Methods(http.MethodGet).
		Name("GetAllEnvironment")
	router.
		HandleFunc("/project/{project_id:[0-9]+}/environments/{environment_id:[0-9]+}/services", sh.GetEnvironmentServiceOrders).
		Methods(http.MethodGet).
		Name("GetEnvironmentServiceOrders")
	router.
		HandleFunc("/project/{project_id}/environments/{environment_id}/services/{service_order_id}", sh.GetServiceOrder).
		Methods(http.MethodGet).
		Name("GetServiceOrder")
	router.
		HandleFunc("/project/{project_id:[0-9]+}/environments/{environment_id:[0-9]+}/services", sh.NewServiceOrder).
		Methods(http.MethodPost).
		Name("NewServiceOrder")

	//User auth implementation
	am := AuthMiddleware{domain.NewAuthRepository(authClient)}
	router.Use(am.authorizationHandler())

	//Implement pagination pageId
	pm := PaginationMiddleware{}
	router.Use(pm.paginationHandler())

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// "DB settings".
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func getAuthURL() string {
	authPort := os.Getenv("AUTH_PORT")
	authURL := os.Getenv("AUTH_ADDRESS")

	URL := fmt.Sprintf("%s:%s", authURL, authPort)
	return URL
}
