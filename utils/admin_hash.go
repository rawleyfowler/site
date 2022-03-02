package utils

import (
	"bufio"
	"os"
)

func LoadAdminHash(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic("Could not find DSN for database...")
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	// The dsn should be the first line of the file
	if reader.Scan() {
		return reader.Text()
	}
	panic(path + " is empty...")
}
