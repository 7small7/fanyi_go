package cmd

import (
	"flag"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func CommandArgs() (string, string, string) {
	translationContentPlatform := flag.String("t", "baidu", "the platform used for translation has optional values of bd (Baidu), tx (Tencent), and al (Alibaba).")
	translationContent := flag.String("c", "", "the content that needs to be translated.")
	translationlanguage := flag.String("l", "en", "the language of the content to be translated has optional values of en (English), zh (Chinese), ja (Japanese)")

	flag.Parse()

	definePlatform := []string{"bd", "tx", "al"}

	if *translationlanguage == "" {
		fmt.Println("Please enter the language type for translation")
		return "", "", ""
	}

	if *translationContentPlatform != "" {
		mapString := mapset.NewSet[string]()
		for i := 0; i < len(definePlatform); i++ {
			mapString.Add(definePlatform[i])
		}
		if !mapString.Contains(*translationContentPlatform) {
			fmt.Println("please enter the correct platform name or use --help for prompt operation.")
			return "", "", ""
		}
	} else {
		*translationContentPlatform = definePlatform[0]
	}

	if *translationContent == "" {
		fmt.Println("please enter the content to be translated, or use -- help to view the explanation of command parameters.")
		return "", "", ""
	}

	return *translationContentPlatform, *translationContent, *translationlanguage
}
