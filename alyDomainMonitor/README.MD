# 阿里云域名过期监控脚本

监控阿里云域名的过期情况，并将即将过期的域名信息通过企业微信机器人推送消息。

## 功能描述

- 使用阿里云开放 API 查询账号下所有域名的过期状态。
- 筛选出即将过期的域名，并构造消息内容。
- 通过企业微信机器人推送即将过期的域名信息。

## 配置说明

在运行脚本前，需要设置以下环境变量：

- `OPS_GROUP_ROBOT`：企业微信机器人的 webhook 地址。
- `ACCESS_KEY_ID`：阿里云账号的 AccessKey ID。需要有域名资源的只读权限
- `ACCESS_KEY_SECRET`：阿里云账号的 AccessKey Secret。
