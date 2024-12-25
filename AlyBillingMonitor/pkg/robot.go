package pkg

import (
	"AlyBillingMonitor/config"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// RobotSender 群机器人消息推送
func RobotSender(msg, people string) (err error) {
	data := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s","mentioned_mobile_list":["%s"]}}`, msg, people)
	if _, err := http.Post(config.OpsGroupRobot, "application/json", strings.NewReader(data)); err != nil {
		log.Println("群机器人消息推送失败，err：", err)
		return err
	}
	return
}
