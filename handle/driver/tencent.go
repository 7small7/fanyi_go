package driver

import (
	"encoding/json"
	"fmt"
	"github.com/7samll7/config"
	"github.com/atotto/clipboard"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

type Tencent struct {
	Content string
}

func (receiver *Tencent) Output() (string, error) {
	return receiver.trans(receiver.Content)
}

func (receiver *Tencent) trans(content string) (string, error) {
	tencentConfig, _ := config.TencentConfig()
	credential := common.NewCredential(
		tencentConfig.Key,
		tencentConfig.Secret,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tmt.tencentcloudapi.com"
	client, _ := tmt.NewClient(credential, "ap-chengdu", cpf)

	request := tmt.NewTextTranslateRequest()
	request.SourceText = common.StringPtr(content)
	request.Source = common.StringPtr("zh")
	request.Target = common.StringPtr("en")
	request.ProjectId = common.Int64Ptr(0)

	response, err := client.TextTranslate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", err
	}
	if err != nil {
		return "", err
	}
	var responseMessage map[string]map[string]interface{}
	err = json.Unmarshal([]byte(response.ToJsonString()), &responseMessage)
	if err != nil {
		return "", err
	}
	err = clipboard.WriteAll(responseMessage["Response"]["TargetText"].(string))
	if err != nil {
		fmt.Println(err)
	}
	return responseMessage["Response"]["TargetText"].(string), err
}
