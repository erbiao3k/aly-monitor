# 阿里云SSL证书过期提醒程序

## 概述
这是一个Go语言编写的程序，用于定期检查阿里云SSL证书的过期情况，并通过群机器人发送即将过期的证书提醒。

## 功能
- 定时查询阿里云SSL证书的过期时间。
- 如果证书将在7天内过期，通过群机器人发送提醒消息。

## 环境变量
- `ACCESS_KEY_ID`: 阿里云AccessKey ID。
- `ACCESS_KEY_SECRET`: 阿里云AccessKey Secret。
- `OPS_GROUP_ROBOT`:  企业微信群机器人的Webhook地址。
