package main

import (
	"github.com/emvi/logbuch"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	staticDir       = "public/static"
	staticDirPrefix = "/static/"
	buildJsFile     = "public/dist/build.js"
	buildJsPrefix   = "/dist/build.js"
	cssFile         = "public/dist/main.css"
	cssFilePrefix   = "/dist/main.css"
	indexFile       = "public/index.html"
	rootDirPrefix   = "/"

	envPrefix = "STS_SLOTLIST_"
	pwdString = "PASSWORD" // do not log passwords!

	defaultHttpWriteTimeout = 20
	defaultHttpReadTimeout  = 20
)

var (
	buildJs      []byte
	watchBuildJs = false
)

func configureLog() {
	logbuch.Info("Configure logging...")
	logbuch.SetFormatter(logbuch.NewFieldFormatter(logbuch.StandardTimeFormat, "\t\t"))
	level := strings.ToLower(os.Getenv("STS_SLOTLIST_LOGLEVEL"))

	if level == "debug" {
		logbuch.SetLevel(logbuch.LevelDebug)
	} else if level == "info" {
		logbuch.SetLevel(logbuch.LevelInfo)
	} else {
		logbuch.SetLevel(logbuch.LevelWarning)
	}
}

func logEnvConfig() {
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, envPrefix) && !strings.Contains(e, pwdString) {
			pair := strings.Split(e, "=")
			logbuch.Info(pair[0] + "=" + pair[1])
		}
	}
}

func loadBuildJs() {
	logbuch.Info("Loading build.js...")
	watchBuildJs = os.Getenv("STS_SLOTLIST_WATCH_BUILD_JS") != ""
	content, err := ioutil.ReadFile(buildJsFile)

	if err != nil {
		logbuch.Fatal("build.js not found", logbuch.Fields{"err": err})
	}

	buildJs = content
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// static content
	router.PathPrefix(staticDirPrefix).Handler(http.StripPrefix(staticDirPrefix, http.FileServer(http.Dir(staticDir))))
	router.PathPrefix(buildJsPrefix).Handler(http.StripPrefix(buildJsPrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if watchBuildJs {
			loadBuildJs()
		}

		w.Header().Add("Content-Type", "text/javascript")

		if _, err := w.Write(buildJs); err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
	})))
	router.PathPrefix(cssFilePrefix).Handler(http.StripPrefix(cssFilePrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		http.ServeFile(w, r, cssFile)
	})))
	router.PathPrefix(rootDirPrefix).Handler(http.StripPrefix(rootDirPrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, indexFile)
	})))

	return router
}

func configureCors(router *mux.Router) http.Handler {
	logbuch.Info("Configuring CORS...")

	origins := strings.Split(os.Getenv("STS_SLOTLIST_ALLOWED_ORIGINS"), ",")
	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            strings.ToLower(os.Getenv("STS_SLOTLIST_CORS_LOGLEVEL")) == "debug",
	})
	return c.Handler(router)
}

func start(handler http.Handler) {
	logbuch.Info("Starting server...")

	writeTimeout := defaultHttpWriteTimeout
	readTimeout := defaultHttpReadTimeout
	var err error

	if os.Getenv("STS_SLOTLIST_HTTP_WRITE_TIMEOUT") != "" {
		writeTimeout, err = strconv.Atoi(os.Getenv("STS_SLOTLIST_HTTP_WRITE_TIMEOUT"))

		if err != nil {
			logbuch.Fatal(err.Error())
		}
	}

	if os.Getenv("STS_SLOTLIST_HTTP_READ_TIMEOUT") != "" {
		readTimeout, err = strconv.Atoi(os.Getenv("STS_SLOTLIST_HTTP_READ_TIMEOUT"))

		if err != nil {
			logbuch.Fatal(err.Error())
		}
	}

	logbuch.Info("Using HTTP read/write timeouts", logbuch.Fields{"write_timeout": writeTimeout, "read_timeout": readTimeout})

	server := &http.Server{
		Handler:      handler,
		Addr:         os.Getenv("STS_SLOTLIST_HOST"),
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
	}

	if strings.ToLower(os.Getenv("STS_SLOTLIST_TLS_ENABLE")) == "true" {
		logbuch.Info("TLS enabled")

		if err := server.ListenAndServeTLS(os.Getenv("STS_SLOTLIST_TLS_CERT"), os.Getenv("STS_SLOTLIST_TLS_PKEY")); err != nil {
			logbuch.Fatal(err.Error())
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			logbuch.Fatal(err.Error())
		}
	}
}

func main() {
	configureLog()
	logEnvConfig()
	loadBuildJs()
	router := setupRouter()
	corsConfig := configureCors(router)
	start(corsConfig)
}
