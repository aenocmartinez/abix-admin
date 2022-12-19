package database

var user string = DataSource().UserDB()
var pass string = DataSource().PassDB()
var host string = DataSource().HostDB()
var port string = DataSource().PortDB()
var database string = DataSource().NameDB()

func getSource() IConnectionDB {
	var source IConnectionDB
	if DataSource().DriverDB() == "mysql" {
		source = NewMySQL()
	}

	if DataSource().DriverDB() == "postgresql" {
		source = NewPostgreSQL()
	}
	return source
}
