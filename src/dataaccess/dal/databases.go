package dal

import "github.com/Jacobbrewer1/chess-boards/src/dataaccess/connection"

var Connections ConnectionSet

type ConnectionSet struct {
	chessMysql   *connection.MySql
	chessMongoDB *connection.MongoDB
}

func (c *ConnectionSet) BadmintonManagerMysql() *connection.MySql {
	return c.chessMysql
}

func (c *ConnectionSet) SetBadmintonManagerMysql(chessMysql *connection.MySql) {
	c.chessMysql = chessMysql
}

func (c *ConnectionSet) BadmintonManagerMongoDB() *connection.MongoDB {
	return c.chessMongoDB
}

func (c *ConnectionSet) SetBadmintonManagerMongoDB(chessMongoDB *connection.MongoDB) {
	c.chessMongoDB = chessMongoDB
}
