package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type ZIMU struct {
	Aliyun
}

func NewZimu(AK, AS, AppKey string) *ZIMU {
	var zimu = ZIMU{}
	zimu.AK = AK
	zimu.AS = AS
	zimu.AppKey = AppKey
	return &zimu
}
func (z *ZIMU) StartConvert(link string) {
	// 地域ID，固定值。
	const REGION_ID string = "cn-shanghai"
	const ENDPOINT_NAME string = "cn-shanghai"
	const PRODUCT string = "nls-filetrans"
	const DOMAIN string = "filetrans.cn-shanghai.aliyuncs.com"
	const API_VERSION string = "2018-08-17"
	const POST_REQUEST_ACTION string = "SubmitTask"
	const GET_REQUEST_ACTION string = "GetTaskResult"
	// 请求参数
	const KEY_APP_KEY string = "appkey"
	const KEY_FILE_LINK string = "file_link"
	const KEY_VERSION string = "version"
	const KEY_ENABLE_WORDS string = "enable_words"
	// 响应参数
	const KEY_TASK string = "Task"
	const KEY_TASK_ID string = "TaskId"
	const KEY_STATUS_TEXT string = "StatusText"
	const KEY_RESULT string = "Result"
	// 状态值
	const STATUS_SUCCESS string = "SUCCESS"
	const STATUS_RUNNING string = "RUNNING"
	const STATUS_QUEUEING string = "QUEUEING"
	var accessKeyId string = z.AK
	var accessKeySecret string = z.AS
	var appKey string = z.AppKey
	var fileLink string = link
	client, err := sdk.NewClientWithAccessKey(REGION_ID, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	postRequest := requests.NewCommonRequest()
	postRequest.Domain = DOMAIN
	postRequest.Version = API_VERSION
	postRequest.Product = PRODUCT
	postRequest.ApiName = POST_REQUEST_ACTION
	postRequest.Method = "POST"
	mapTask := make(map[string]string)
	mapTask[KEY_APP_KEY] = appKey
	mapTask[KEY_FILE_LINK] = fileLink
	// 新接入请使用4.0版本，已接入（默认2.0）如需维持现状，请注释掉该参数设置。
	mapTask[KEY_VERSION] = "4.0"
	// 设置是否输出词信息，默认为false。开启时需要设置version为4.0。
	mapTask[KEY_ENABLE_WORDS] = "false"
	task, err := json.Marshal(mapTask)
	if err != nil {
		panic(err)
	}
	postRequest.FormParams[KEY_TASK] = string(task)
	postResponse, err := client.ProcessCommonRequest(postRequest)
	if err != nil {
		panic(err)
	}
	postResponseContent := postResponse.GetHttpContentString()
	fmt.Println(postResponseContent)
	if postResponse.GetHttpStatus() != 200 {
		fmt.Println("录音文件识别请求失败，Http错误码: ", postResponse.GetHttpStatus())
		return
	}
	var postMapResult map[string]interface{}
	err = json.Unmarshal([]byte(postResponseContent), &postMapResult)
	if err != nil {
		panic(err)
	}
	var taskId string = ""
	var statusText string = ""
	statusText = postMapResult[KEY_STATUS_TEXT].(string)
	if statusText == STATUS_SUCCESS {
		fmt.Println("录音文件识别请求成功响应!")
		taskId = postMapResult[KEY_TASK_ID].(string)
	} else {
		fmt.Println("录音文件识别请求失败!")
		return
	}
	getRequest := requests.NewCommonRequest()
	getRequest.Domain = DOMAIN
	getRequest.Version = API_VERSION
	getRequest.Product = PRODUCT
	getRequest.ApiName = GET_REQUEST_ACTION
	getRequest.Method = "GET"
	getRequest.QueryParams[KEY_TASK_ID] = taskId
	statusText = ""
	for true {
		getResponse, err := client.ProcessCommonRequest(getRequest)
		if err != nil {
			panic(err)
		}
		getResponseContent := getResponse.GetHttpContentString()
		fmt.Println("识别查询结果：", getResponseContent)
		if getResponse.GetHttpStatus() != 200 {
			fmt.Println("识别结果查询请求失败，Http错误码：", getResponse.GetHttpStatus())
			break
		}
		var getMapResult ZimuResult
		err = json.Unmarshal([]byte(getResponseContent), &getMapResult)
		if err != nil {
			panic(err)
		}
		statusText = getMapResult.StatusText
		if statusText == STATUS_RUNNING || statusText == STATUS_QUEUEING {
			time.Sleep(10 * time.Second)
		} else {
			log.Println("End")
			log.Println("[*] Start translation")
			z.GetResult(getMapResult)
			break
		}
	}
	if statusText == STATUS_SUCCESS {
		fmt.Println("录音文件识别成功！")
	} else {
		fmt.Println("录音文件识别失败！")
	}
}
func (c *ZIMU) time(t int) string {
	var timeObj = time.Unix(int64(t), 0)
	return timeObj.Format("15:04:05")
}
func (z *ZIMU) GetResult(result ZimuResult) {
	var resultText = ""
	fy := Fanyi{}
	fy.AK = z.AK
	fy.AS = z.AS
	for index, row := range result.Result.Sentences {
		row.Text = strings.ReplaceAll(row.Text, " ", "")
		zh := fy.GetResult(row.Text)
		if strings.Contains(resultText, zh) {
			continue
		}
		var text = "\n%s --> %s\n%s\n\n"
		text = fmt.Sprintf(text, Int2time(row.BeginTime), Int2time(row.EndTime), zh)
		if strings.Contains(resultText, text) {
			continue
		}
		text = fmt.Sprintf("%d%s", index, text)
		if strings.Contains(resultText, fmt.Sprintf("%s --> %s", Int2time(row.BeginTime), Int2time(row.EndTime))) {
			continue
		}
		resultText += text
	}
	ioutil.WriteFile("output/video.srt", []byte(resultText), 0777)
}
