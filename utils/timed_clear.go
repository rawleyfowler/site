package utils

import (
	"strconv"
	"time"
)

func TimedClearMap(m *map[string]uint, timer uint) {
	for {
		t, err := time.ParseDuration(strconv.FormatUint(uint64(timer), 10) + "ms")
		if err != nil {
			panic("Something went wrong initializing time!!")
		}
		time.Sleep(t)
		for key := range *m {
			delete(*m, key)
		}
	}
}
