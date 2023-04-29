package test

import "golang-web-testing/config"

func DBCleanup() {
	config.DB.Exec("TRUNCATE users RESTART IDENTITY;")
}
