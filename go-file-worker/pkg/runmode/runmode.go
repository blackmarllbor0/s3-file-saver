package runmode

import "os"

const (
	DEV  = "dev"
	PROD = "prod"
	TEST = "test"
)

func GetAppRunMode() (runMode string) {
	runMode = os.Getenv("RUN_MODE")
	if runMode == "" {
		runMode = DEV
	}

	return runMode
}
