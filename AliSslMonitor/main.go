package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	cas20200407 "github.com/alibabacloud-go/cas-20200407/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	prefixMsg       = "【ssl证书过期提醒】\n"
	AccessKeyId     = os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
	OpsGroupRobot   = os.Getenv("OPS_GROUP_ROBOT")
)

func main() {
	log.Println("程序已启动")
	for {
		log.Println("开始查询阿里云证书过期情况")
		run()
		log.Println("结束查询阿里云证书过期情况，休眠中...")
		time.Sleep(3 * 24 * time.Hour)
	}
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *cas20200407.Client, _err error) {
	c := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	c.Endpoint = tea.String("cas.aliyuncs.com")
	_result = &cas20200407.Client{}
	_result, _err = cas20200407.NewClient(c)
	return _result, _err
}

func run() {
	client, err := CreateClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
	if err != nil {
		log.Println("初始化客户端错误：", err)
	}
	listUserCertificateOrderRequest := &cas20200407.ListUserCertificateOrderRequest{
		ShowSize: tea.Int64(1000000),
	}
	runtime := &util.RuntimeOptions{}

	resp, err := client.ListUserCertificateOrderWithOptions(listUserCertificateOrderRequest, runtime)
	if err != nil {
		log.Println("请求数据错误：", err)
	}
	if *(resp.StatusCode) > 399 {
		log.Println("请求异常：", resp.String())
	}

	msg := prefixMsg
	for _, ssl := range resp.Body.CertificateOrderList {
		instanceId := *(ssl.InstanceId)
		domain := *(ssl.Domain)

		endTimestamp := *(ssl.CertEndTime) / 1000
		EndTime := time.Unix(endTimestamp, 0)
		certEndTime := EndTime.Format("2006-01-02 15:04:05")

		timeNowTimestamp := time.Now().Unix()

		if endTimestamp-timeNowTimestamp < 7*24*60*60 {
			data := fmt.Sprintf("实例ID为：%s的ssl证书%s将于%s到期\n", instanceId, domain, certEndTime)
			msg += data
		}
	}

	if msg != prefixMsg {
		log.Println(msg)
		RobotSender(msg, "@all")
	}
}

// RobotSender 群机器人消息推送
func RobotSender(msg, people string) (err error) {

	data := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s","mentioned_mobile_list":["%s"]}}`, msg, people)

	if _, err := http.Post(OpsGroupRobot, "application/json", strings.NewReader(data)); err != nil {
		log.Println("群机器人消息推送失败，err：", err)
		return err
	}
	return
}
