package config

import "github.com/Jacobbrewer1/chess-boards/src/dataaccess/connection"

var Cfg Root

type (
	Root struct {
		Setup     *Setup     `json:"setup,omitempty" yaml:"setup,omitempty"`
		Databases *Databases `json:"databases,omitempty" yaml:"databases,omitempty"`
	}

	Setup struct {
		ListeningPort string `json:"listeningPort,omitempty" yaml:"listeningPort,omitempty"`
		CertPath      string `json:"certificatePath,omitempty" yaml:"certPath,omitempty"`
		KeyPath       string `json:"keyPath,omitempty" yaml:"keyPath,omitempty"`
	}

	Databases struct {
		MysqlCredentials   *connection.MySql   `json:"mysql,omitempty" yaml:"mysql,omitempty"`
		MongoDbCredentials *connection.MongoDB `json:"mongoDb,omitempty" yaml:"mongoDb,omitempty"`
	}
)
