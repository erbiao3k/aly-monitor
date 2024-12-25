package pkg

import (
	"AlyBillingMonitor/config"
	"fmt"
	"github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"log"
	"strconv"
	"strings"
)

func ParseFloatByLocale(str string) string {
	return strings.ReplaceAll(str, ",", "")
}

// AvailableAmount 查询阿里云可用余额
func AvailableAmount(client *client.Client) (availableAmount float64) {
	billData, err := client.QueryAccountBalanceWithOptions(config.Runtime)

	if err != nil {
		log.Fatal("请求失败：", err)
	}
	if !*billData.Body.Success {
		log.Fatal("返回数据异常：", billData)
	}

	availableAmount, err = strconv.ParseFloat(ParseFloatByLocale(*(billData.Body.Data.AvailableAmount)), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(availableAmount)
	}

	return
}
