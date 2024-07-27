package cmd

import (
	"flag"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func CommandArgs() (string, string) {
	translationContentPlatform := flag.String("t", "baidu", "the platform used for translation has optional values of bd (Baidu), tx (Tencent), and al (Alibaba).")
	translationContent := flag.String("c", "", "the content that needs to be translated.")

	flag.Parse()

	definePlatform := []string{"bd", "tx", "al"}

	if *translationContentPlatform != "" {
		mapString := mapset.NewSet[string]()
		for i := 0; i < len(definePlatform); i++ {
			mapString.Add(definePlatform[i])
		}
		if !mapString.Contains(*translationContentPlatform) {
			fmt.Println("please enter the correct platform name or use --help for prompt operation.")
			return "", ""
		}
	} else {
		*translationContentPlatform = definePlatform[0]
	}

	if *translationContent == "" {
		fmt.Println("please enter the content to be translated, or use -- help to view the explanation of command parameters.")
		return "", ""
	}

	return *translationContentPlatform, *translationContent
}
