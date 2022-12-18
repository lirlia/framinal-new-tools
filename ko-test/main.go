package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(GetEnv("GCP_PROJECT_ID", "dummy"))
	fmt.Println(GetEnv("SPANNER_INSTANCE", "dummy"))
	fmt.Println(GetEnv("SPANNER_DATABASE", "dummy"))
	fmt.Println(GetEnv("MASTER_DATA_PATH", "dummy"))
	fmt.Println(GetEnv("USER_IDS_LIST_PATH", "dummy"))

}

func GetEnv(key string, def string) string {
	var ret string
	if ret = os.Getenv(key); ret != "" {
		return ret
	}
	return def
}
