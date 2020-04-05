package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/openvino/openvino-api/src/app/handler"
	"github.com/openvino/openvino-api/src/app/model"
	"github.com/openvino/openvino-api/src/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
		config.Database.Charset)

	db, err := gorm.Open(config.Database.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database", err)
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects

	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello world! This is the openvino api.")) })
	
	a.Get("/sensor_data", a.handleRequest(handler.GetSensorDataDay)).Queries("day", "{[1-31]*?}").Queries("month", "{[0-12]*?}").Queries("year", "{[0-2030]*?}")
	a.Get("/sensor_data", a.handleRequest(handler.GetSensorDataMonth)).Queries("month", "{[1-12]*?}").Queries("year", "{[0-2030]*?}")
	a.Get("/sensor_data", a.handleRequest(handler.GetSensorDataYear)).Queries("year", "{[2019-2030]*?}")
	a.Post("/sensor_data", a.handleRequest(handler.CreateSensorData))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request))  *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request))  *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
