package handle

import (
	"fmt"

	"github.com/7samll7/handle/driver"
)

type InputInterface interface {
	Output() (string, error)
}

func TxtTrans(platform, txt, language string) {
	var err error
	var content string
	switch platform {
	case "bd":
		content, err = (&driver.BaiDu{Content: txt, Language: language}).Output()
	case "tx":
		content, err = (&driver.Tencent{Content: txt, Language: language}).Output()
	}
	if err != nil {
		fmt.Println("request error:", err.Error())
		return
	}
	fmt.Println(content)
}
