package factory

import "github.com/larrygf02/factory/connection"

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.Mysql{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
