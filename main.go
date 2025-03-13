package main

import (
	envutils "warehouse-management-system/utils/env"
	"warehouse-management-system/utils/loggerutils"
)

func main() {
	loggerutils.SetLogger(loggerutils.NewLogrusLogger())
	envutils.LoadEnv()
}
