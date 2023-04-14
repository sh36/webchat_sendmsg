# wechatbot
项目基于[openwechat](https://github.com/eatmoreapple/openwechat)开发
### 目前实现了以下功能
 + 定时发送消息

### 启动项目
1. ```nohup sh startup.sh &```进行项目后台启动
2. 使用vim打开 nohup.log ，找到最新的日志，即为登录链接(在vim命令模式下输入'G'可快速跳转到最后一行)
3. 点击链接使用浏览器打开，微信扫码进行登陆

### 关闭项目
1. ```ps -aux | grep main.go```找到项目对应的进程id
2. ``` kill -9 进程id ```
