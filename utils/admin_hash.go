package utils

import (
	"bufio"
	"os"
)

func LoadAdminHash(path string) [32]byte {
	file, err := os.Open(path)
	if err != nil {
		panic("Could not find DSN for database...")
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	// The dsn should be the first line of the file
	if reader.Scan() {
		var buff [32]byte
		for i, b := range reader.Bytes() {
			buff[i] = b
		}
		return buff
	}
	panic(path + " is empty...")
}
