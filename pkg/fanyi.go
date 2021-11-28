package pkg

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
)

type Fanyi struct {
	Aliyun
}

func (f *Fanyi) GetResult(text string) string {
	client, err := alimt.NewClientWithAccessKey("cn-qingdao", f.AK, f.AS)
	request := alimt.CreateTranslateGeneralRequest()
	request.Scheme = "https"

	request.SourceLanguage = "ja"
	request.TargetLanguage = "zh"
	request.SourceText = text
	request.FormatType = "text"

	response, err := client.TranslateGeneral(request)
	if err != nil {
		fmt.Print(err.Error())
		return ""
	}
	return response.Data.Translated
}
