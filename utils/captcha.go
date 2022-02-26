package utils

import (
	"math/rand"
	"time"
)

func GenerateCaptchaValues() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func UpdateCaptcha(captcha *[2]int) {
	for {
		captcha[0], captcha[1] = GenerateCaptchaValues()
		time.Sleep(5 * time.Minute)
	}
}
