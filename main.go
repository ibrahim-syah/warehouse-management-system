package main

import (
	"warehouse-management-system/app"
	envutils "warehouse-management-system/utils/env"
	"warehouse-management-system/utils/loggerutils"
)

func main() {
	loggerutils.SetLogger(loggerutils.NewLogrusLogger())
	envutils.LoadEnv()
	app.StartApp()
}
