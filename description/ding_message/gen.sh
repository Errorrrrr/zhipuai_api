#!/bin/bash

API_FILE="ding_message.api"

# 执行 goctl 命令生成代码
goctl api go -api $API_FILE -dir .
