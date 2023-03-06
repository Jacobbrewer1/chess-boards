package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

type MySql struct {
	db               *sql.DB
	connectionString *string
	User             string  `json:"user,omitempty" yaml:"user,omitempty"`
	Password         string  `json:"password,omitempty" yaml:"password,omitempty"`
	Method           string  `json:"method,omitempty" yaml:"method,omitempty"`
	Host             string  `json:"host,omitempty" yaml:"host,omitempty"`
	Port             string  `json:"port,omitempty" yaml:"port,omitempty"`
	Schema           string  `json:"schema,omitempty" yaml:"schema,omitempty"`
	Query            *string `json:"query,omitempty" yaml:"query,omitempty"`
	sync.RWMutex
}

func (c *MySql) Db() *sql.DB {
	if c.db == nil {
		c.connect()
	}

	return c.db
}

func (c *MySql) Ping() {
	if c.db == nil {
		c.connect()
	}

	if err := c.db.Ping(); err != nil {
		c.db = nil
		panic(err)
	}
}

func (c *MySql) connect() {
	if c.connectionString == nil {
		c.generateConnectionString()
	}

	c.RLock()
	defer c.RUnlock()

	db, err := sql.Open("mysql", *c.connectionString)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		if err := db.Close(); err != nil {
			log.Println("closing db error:", err)
		}

		panic(err)
	}

	c.db = db
}

func (c *MySql) generateConnectionString() {
	if c.User == "" || c.Password == "" || c.Method == "" && c.Host == "" ||
		c.Port == "" && c.Schema == "" {
		panic("invalid connection")
	}

	// user:pssword@method(destination:port)/schema?query
	// Example: root:password@tcp(127.0.0.1:3306)/schema?timeout=2s&parseTime=true
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", c.User, c.Password, c.Method, c.Host, c.Port, c.Schema)

	if c.Query != nil {
		connectionString = fmt.Sprintf("%s?%s", connectionString, *c.Query)
	}

	c.connectionString = &connectionString
}
