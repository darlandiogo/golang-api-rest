package database

var Conn Connection;

func GetConnection () Connection {

	Conn.OpenConnection()

	return Conn;
}
