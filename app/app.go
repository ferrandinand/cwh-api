package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/domain"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/service"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/logger"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
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

	vulnerabilityRepositoryDb := domain.NewVulnerabilityRepositoryDb(dbClient)
	vh := VulnerabilityHandler{service.NewVulnerabilityService(vulnerabilityRepositoryDb)}

	costRepositoryDb := domain.NewCostRepositoryDb(dbClient)
	ch := CostHandler{service.NewCostService(costRepositoryDb)}

	client := &http.Client{}
	statusAdapter := domain.NewStatusAdapter(client, os.Getenv("STATUS_API"))
	sh := StatusHandler{service.NewStatusService(statusAdapter)}

	// define routes
	router.
		HandleFunc("/vulnerabilities/{project_id}", vh.GetVulnerability).
		Methods(http.MethodGet).
		Name("GetVulnerability")
	router.
		HandleFunc("/resources/{project_id}", ch.GetResource).
		Methods(http.MethodGet).
		Name("GetResource")
	router.
		HandleFunc("/status/{project_id}", sh.GetStatus).
		Methods(http.MethodGet).
		Name("GetStatus")

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
