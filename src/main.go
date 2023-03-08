package main

import (
	"flag"
	"fmt"
	"github.com/Jacobbrewer1/chess-boards/src/config"
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"github.com/Jacobbrewer1/chess-boards/src/dataaccess/dal"
	"github.com/Jacobbrewer1/chess-boards/src/entities"
	"github.com/Jacobbrewer1/chess-boards/src/session"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func init() {
	initFlags()
	initLogging()
	initConfig()
	initDatabases()
	initTemplates()
}

func initFlags() {
	configPath := flag.String(config.FlagName, "./config/config.yml", "Provides the location of the config file (required)")

	flag.Parse()

	config.Location = *configPath
}

func isFlagProvided(flagName string) (isProvided bool) {
	isProvided = false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			isProvided = true
		}
	})

	return isProvided
}

func initLogging() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func initConfig() {
	if !isFlagProvided(config.FlagName) {
		panic("no config flag provided")
	}

	if err := config.CreateConfig(); err != nil {
		panic(err)
	}
}

// confirm database connections are valid
func initDatabases() {
	log.Println("testing mysql database connection")
	config.Cfg.Databases.MysqlCredentials.Ping()
	dal.Connections.SetChessMysql(config.Cfg.Databases.MysqlCredentials)
	log.Println("mysql database connected")

	log.Println("testing mongodb connection")
	config.Cfg.Databases.MongoDbCredentials.Ping()
	dal.Connections.SetChessMongoDB(config.Cfg.Databases.MongoDbCredentials)
	log.Println("mongodb connected")

	dal.Connections.SetRedisDb(config.Cfg.Databases.RedisDbCredentials)
}

func initTemplates() {
	log.Println("parsing templates")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	templates = template.Must(template.New("").Option("missingkey=error").Funcs(template.FuncMap{
		// Web methods
	}).ParseGlob("./assets/templates/*.gohtml"))

	log.Println("templates parsed successfully")
}

func init() {
	key := uuid.NewString()
	n := time.Now()

	db := dal.NewSessionDal(0)
	err := db.SaveSession(session.Session{
		Key: key,
		User: entities.User{
			Id:        1,
			FirstName: "temp",
			Surname:   "test",
			Email:     "test@mail.co.uk",
			Password:  "test",
			LastLogin: (custom.Datetime)(n),
		},
		Expiry: (custom.Datetime)(n),
	})
	if err != nil {
		log.Println(err)
		return
	}

	got, err := db.GetSession(key)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(got)

	//fmt.Println("Go Redis Tutorial")
	//
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379",
	//	Password: "",
	//	DB:       0,
	//})
	//
	//now := time.Now()
	//exp := time.Now().Add(time.Hour * 6)
	//s := session.Session{
	//	User: entities.User{
	//		Id:        1,
	//		FirstName: "Test",
	//		Surname:   "Session",
	//		Email:     "test.session@session.com",
	//		Password:  "test",
	//		LastLogin: (custom.Datetime)(now),
	//	},
	//	Expiry: (custom.Datetime)(exp),
	//}
	//
	//data, err := json.Marshal(s)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = client.Set("id1234", data, 0).Err()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//val, err := client.Get("id1234").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(val)
}

func main() {
	defer func() {
		if recovery := recover(); recovery != nil {
			log.Println("main recovery:", recovery)
			log.Println("stacktrace:", string(debug.Stack()))
		}
	}()

	r := mux.NewRouter()
	//secure := r.PathPrefix("/secure")
	//ajax := r.PathPrefix("/ajax").Subrouter()
	//secureAjax := secure.PathPrefix("/ajax").Subrouter()

	r.HandleFunc("/login", login).Methods(http.MethodGet)

	http.Handle("/", r)

	log.Println("listening...")
	if err := http.ListenAndServeTLS(
		fmt.Sprintf(":%s", config.Cfg.Setup.ListeningPort),
		config.Cfg.Setup.CertPath,
		config.Cfg.Setup.KeyPath,
		nil); err != nil {
		panic(err)
	}
}
