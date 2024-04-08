package service

import (
	"admin-api/utils"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 使用store
var store = utils.RedisStore{}

// 生成验证码
func CaptMake() (id, b64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, &store)
	lid, lb64s, _, _ := captcha.Generate()
	return lid, lb64s
}
