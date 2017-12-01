package client

import (
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

// Install creates a new MariaDB chart release
func Install(id string) error {
    db_host := os.Getenv("MARIADB_HOST")
    db_port := os.Getenv("MARIADB_PORT")
    db_user := os.Getenv("MARIADB_USER")
    db_pass := os.Getenv("MARIADB_PASS")

    db, err := sql.Open("mysql", db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/")
	if err != nil {
		return err
	}
    defer db.Close()

	glog.Infof("Create database: %s\n", id)
    _, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + id + "`")
	if err != nil {
		panic(err)
	}

	glog.Infof("Create user: %s\n", id)
    _, err = db.Exec("GRANT ALL ON `" + id + "`.* TO `" + id + "`@'%' IDENTIFIED BY '" + id + "'")
	if err != nil {
		panic(err)
	}

	return nil
}

// Delete deletes a MariaDB chart release
func Delete(id string) error {
    db_host := os.Getenv("MARIADB_HOST")
    db_port := os.Getenv("MARIADB_PORT")
    db_user := os.Getenv("MARIADB_USER")
    db_pass := os.Getenv("MARIADB_PASS")

    db, err := sql.Open("mysql", db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/")
	if err != nil {
		return err
	}
    defer db.Close()

	glog.Infof("Drop database: %s\n", id)
    _, err = db.Exec("DROP DATABASE IF EXISTS `" + id + "`")
	if err != nil {
		return err
	}

	glog.Infof("Drop user: %s\n", id)
    _, err = db.Exec("DROP USER IF EXISTS `" + id + "`")
	if err != nil {
		return err
	}

	return nil
}
