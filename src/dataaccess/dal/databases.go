package dal

import "github.com/Jacobbrewer1/chess-boards/src/dataaccess/connection"

var Connections ConnectionSet

type ConnectionSet struct {
	chessMysql   *connection.MySql
	chessMongoDB *connection.MongoDB
}

func (c *ConnectionSet) ChessMysql() *connection.MySql {
	return c.chessMysql
}

func (c *ConnectionSet) SetChessMysql(chessMysql *connection.MySql) {
	c.chessMysql = chessMysql
}

func (c *ConnectionSet) ChessMongoDB() *connection.MongoDB {
	return c.chessMongoDB
}

func (c *ConnectionSet) SetChessMongoDB(chessMongoDB *connection.MongoDB) {
	c.chessMongoDB = chessMongoDB
}
