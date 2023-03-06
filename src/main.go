package main

import (
	"flag"
	"github.com/Jacobbrewer1/chess-boards/src/config"
	"github.com/Jacobbrewer1/chess-boards/src/dataaccess/dal"
	"html/template"
	"log"
	"net/http"
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
	dal.Connections.SetBadmintonManagerMysql(config.Cfg.Databases.MysqlCredentials)
	log.Println("mysql database connected")

	log.Println("testing mongodb connection")
	config.Cfg.Databases.MongoDbCredentials.Ping()
	dal.Connections.SetBadmintonManagerMongoDB(config.Cfg.Databases.MongoDbCredentials)
	log.Println("mongodb connected")
}

func initTemplates() {
	log.Println("parsing templates")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	templates = template.Must(template.New("").Option("missingkey=error").Funcs(template.FuncMap{
		// Web methods
	}).ParseGlob("./assets/templates/*.gohtml"))

	log.Println("templates parsed successfully")
}

func main() {
	log.Println(config.Cfg.Setup.ListeningPort)
}
