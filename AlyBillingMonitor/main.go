package main

import (
	"AlyBillingMonitor/config"
	"AlyBillingMonitor/pkg"
	"fmt"
	"log"
	"time"

	"github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	client2 "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

func main() {
	log.Println("程序已启动")
	for {
		fmt.Println("开始查询阿里云费用余额")
		runner(time.Now())
		fmt.Println("结束查询阿里云费用余额")
		time.Sleep(1 * time.Hour)
	}
}

func runner(nowTime time.Time) {
	for _, akInfo := range config.AKs {
		conf := &client2.Config{
			AccessKeyId:     tea.String(akInfo["AccessKeyId"]),
			AccessKeySecret: tea.String(akInfo["AccessKeySecret"]),
			Endpoint:        tea.String(config.Endpoint),
		}

		cli, err := client.NewClient(conf)
		if err != nil {
			log.Fatal("客户端初始化失败：", err)
		}
		// 当前账号余额
		availableAmount := pkg.AvailableAmount(cli)

		firstDate, lastDate := pkg.GetAZDay(nowTime)
		// 当月充值次数，充值金额
		paymentAccount, payment := pkg.Transactions(cli, firstDate+"T00:00:00Z", lastDate+"T23:59:59Z", "Payment")

		msg := fmt.Sprintf("【阿里云财务监控】\n账号名称：%s\n当月充值次数：%d,\n当月充值金额/元：%.2f,\n当前账号余额/元：%.2f,\n请考虑充值", akInfo["accountName"], paymentAccount, payment, availableAmount)

		log.Println(msg)

		if availableAmount < 3000.0 {
			pkg.RobotSender(msg, "")
		}
	}
}
