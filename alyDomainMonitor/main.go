package main

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	// 企业微信机器人
	OpsGroupRobot = os.Getenv("OPS_GROUP_ROBOT")

	contentType     = "application/json"
	accessKeyId     = os.Getenv("ACCESS_KEY_ID")
	accessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
	prefixMsg       = "以下域名即将过期，请确认是否续费：\n"
)

// RobotSender 群机器人消息推送
func robotSender(msg, people string) (err error) {

	data := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s","mentioned_mobile_list":["%s"]}}`, msg, people)

	if _, err := http.Post(OpsGroupRobot, contentType, strings.NewReader(data)); err != nil {
		log.Println("群机器人消息推送失败，err：", err)
		return err
	}
	return
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *domain20180129.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("domain.aliyuncs.com")
	_result = &domain20180129.Client{}
	_result, _err = domain20180129.NewClient(config)
	return _result, _err
}

func run() {
	client, err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		log.Println("初始化客户端失败：", err)
		return
	}

	queryDomainListRequest := &domain20180129.QueryDomainListRequest{
		PageSize: tea.Int32(1000),
		PageNum:  tea.Int32(1),
	}

	runtime := &util.RuntimeOptions{}

	data, err := client.QueryDomainListWithOptions(queryDomainListRequest, runtime)

	if *data.StatusCode != int32(200) {
		log.Println("请求异常：", err)
		return
	}

	if *data.Body.TotalItemNum == 0 {
		log.Println("该账号未查询到域名")
		return
	}

	msg := prefixMsg

	for _, domainName := range data.Body.Data.Domain {
		if *domainName.DomainStatus == strconv.Itoa(1) {
			msg += "*" + "域名：" + *domainName.DomainName + "，过期时间：" + *domainName.ExpirationDate + "\n"
		}
	}

	if msg != prefixMsg {
		robotSender(msg, "@all")
	}

}

func main() {
	log.Println("程序已启动")
	for true {
		log.Println("开始查询阿里云域名过期情况")
		run()
		log.Println("结束查询阿里云域名过期情况，休眠中...")
		time.Sleep(5 * 24 * time.Hour)
	}
}
