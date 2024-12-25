package config

import (
	util "github.com/alibabacloud-go/tea-utils/service"
	"os"
)

var (
	Endpoint = "business.aliyuncs.com"
	Runtime  = &util.RuntimeOptions{}

	OpsGroupRobot = os.Getenv("OPS_GROUP_ROBOT")

	AKs = []map[string]string{
		{
			"accountName":     os.Getenv("ACCOUNT"),
			"MainAccountId":   os.Getenv("MAIN_ACCOUNT_ID"),
			"AccessKeyId":     os.Getenv("ACCESS_KEY_ID"),
			"AccessKeySecret": os.Getenv("ACCESS_KEY_SECRET"),
		},
	}
)
