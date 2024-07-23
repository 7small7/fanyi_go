package main

import (
	"flag"
	"fmt"
	"github.com/7samll7/handle/driver"
	mapset "github.com/deckarep/golang-set/v2"
)

// +----------------------------------------------------------------------
// | 兔兔答题考试系统
// +----------------------------------------------------------------------
// | 感谢使用兔兔答题考试系统
// | 本系统经过商业授权，不能转售、开源或者以其他方式进行使用，只可在购买协议范围内容使用
// | 若发现未被版权协议，将追究相应的公司、团队、企业等
// | 访问官网：https://www.tutudati.com
// | 开发者：Mandy
// | 创建时间：2024/7/23 23:30
// | 兔兔答题考试系统开发者版权所有 拥有最终解释权
// +----------------------------------------------------------------------

func main() {
	// define command-line parameters
	translationContentPlatform := flag.String("t", "baidu", "the platform used for translation has optional values of bd (Baidu), tx (Tencent), and al (Alibaba)")
	translationContent := flag.String("c", "", "the content that needs to be translated")

	flag.Parse()

	// define available translation platforms
	definePlatform := []string{"bd", "tx", "al"}

	// process command-line parameters. if no platform is set, a platform is set by default.
	// if it is set, check if the platform supports it
	if *translationContentPlatform != "" {
		mapString := mapset.NewSet[string]()
		for i := 0; i < len(definePlatform); i++ {
			mapString.Add(definePlatform[i])
		}
		if !mapString.Contains(*translationContentPlatform) {
			fmt.Println("please enter the correct platform name or use --help for prompt operation.")
			return
		}
	} else {
		*translationContentPlatform = definePlatform[0]
	}

	// verify translation content parameters
	if *translationContent == "" {
		fmt.Println("please enter the content to be translated, or use -- help to view the explanation of command parameters")
		return
	}

	// after verification, determine the platform used and select the corresponding translation interface based on the platform
	if *translationContentPlatform == "bd" {
		baidu := &driver.BaiDu{}
		baidu.Input(*translationContent)
		fmt.Println(baidu.Output())
	}
}
