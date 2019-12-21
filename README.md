# Lark Bot: 飞书机器人🤖️

------

本仓库为整合木犀团队飞书内机器人服务的仓库，目前想要完成的feature：

+ Github每日Push量统计推送👌🏻
+ Github每日Star量最高仓库推送👌🏻
+ 木犀工作台进度统计推送👊🏻

------

### Quick Start

本仓库内置了时刻表（TimeTable），能够帮助您快速搭建自己想要的飞书机器人功能。

Github相关的BOT需要满足接口：`githubbot.GithubBot`:

```golang
type GithubBot interface {
	Do()
	Run(duration time.Duration)
}
```

在自己的BOT中，将业务逻辑写在`Do()`方法中，使用`Run(time.Duration)`方法来在一段时间后调用`Do()`方法。

然后只需要构建一个`TimeTable`的实例，然后将期望运行时间点（整小时）和`BOT`的实例一起插入，运行`TimeTable`就可以了（记得运行`Close`方法）

```golang
githubBotTimeTable.Append(BOT, 10) // 在每天的上午十点整运行

githubBotTimeTable.Run()
defer githubBotTimeTable.Close()
```