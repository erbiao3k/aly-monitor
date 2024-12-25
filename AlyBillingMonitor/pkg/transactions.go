package pkg

import (
	"AlyBillingMonitor/config"
	"strconv"

	"github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	bssopenapi20171214 "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// Transactions 查询阿里云充值金额
func Transactions(client *client.Client, transactionsStartTime, transactionsEndTime, transactionType string) (count int, money float64) {

	var pageSize int32 = 300

	queryAccountTransactionsRequest := &bssopenapi20171214.QueryAccountTransactionsRequest{
		PageSize:        tea.Int32(pageSize),
		CreateTimeStart: tea.String(transactionsStartTime),
		CreateTimeEnd:   tea.String(transactionsEndTime),
		//CreateTimeStart: tea.String("2022-09-01T00:00:00Z"),
		//CreateTimeEnd:   tea.String("2022-09-31T23:59:59Z"),
		TransactionType: tea.String(transactionType),
	}

	accountTransactionsData, _ := client.QueryAccountTransactionsWithOptions(queryAccountTransactionsRequest, config.Runtime)

	//totalCount := *accountTransactionsData.Body.Data.TotalCount

	//forCount := (totalCount / pageSize) + 1

	accountTransactionsList := accountTransactionsData.Body.Data.AccountTransactionsList.AccountTransactionsList

	count = len(accountTransactionsList)

	// 充值记录总数为0则判定充值金额为0
	if count == 0 {
		return count, 0
	}

	for _, list := range accountTransactionsList {
		amount, _ := strconv.ParseFloat(*list.Amount, 64)
		money += amount
	}
	return
}
