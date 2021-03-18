package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/logger"
	"github.com/ferrandinand/cwh-api/service"

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
	ch := UserHandlers{service.NewUserService(userRepositoryDb)}
	ah := ProjectHandler{service.NewProjectService(projectRepositoryDb)}

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
		HandleFunc("/project/new", ah.NewProject).
		Methods(http.MethodPost).
		Name("NewProject")
	router.
		HandleFunc("/project/{project_id}", ah.GetProject).
		Methods(http.MethodGet).
		Name("GetProject")
	router.
		HandleFunc("/projects/", ah.GetAllProject).
		Methods(http.MethodGet).
		Name("GetAllProject")
	//router.
	//	HandleFunc("/users/{user_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
	//	Methods(http.MethodPost).
	//	Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository(authClient)}
	router.Use(am.authorizationHandler())
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
	// See "Important settings" section.
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
