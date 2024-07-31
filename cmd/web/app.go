package main

import (
	"context"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5"
	"github.com/malpania/beerproj/pkg/config"
	"github.com/malpania/beerproj/pkg/handlers"
	"github.com/malpania/beerproj/pkg/helpers"
	"github.com/malpania/beerproj/pkg/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := pgx.Connect(context.Background(), "postgres://amazinguser:perfectpassword@postgres-service:5432/awesomedb")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)

	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Cannot ping db")
	}

	log.Fatal("Success ping db")

	defer conn.Close(context.Background())

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	server := &http.Server{Addr: portNumber, Handler: routes(&app)}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	//	http.ListenAndServe(portNumber, nil)
}

func run() error {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.ReadFolderCache()
	if err != nil {
		log.Fatal("Error loading cache folder")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)

	render.InitializeApp(&app)
	//	http.HandleFunc("/", handlers.Repo.Home)
	//	http.HandleFunc("/about", handlers.Repo.About)
	return nil
}
