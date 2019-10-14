# cronTask
钉钉定时提醒

### 使用方法 ###
    编译
    make build
    
    运行
    make run

### 配置 ###

####    example:
    tasks:
      - {schedule: "*/1 * * * *",message: "发日报了",atAll: true,dingding: "",title: "日报提醒1"}
      - {schedule: "*/2 * * * *",message: "发周报了",atPhone: "111",dingding: "",title: "周报提醒"}
#### 配置属性
| 属性  | 类型 | 值 | 备注             |   |
|-------|------|----|------------------|---|
| tasks | 数组 |    | 对应要执行的任务 |   |

#### task属性
| 属性     | 类型   | 值 | 备注                                    |   |
|----------|--------|----|-----------------------------------------|---|
| schedule | string |    | 运行周期 参考：https://en.wikipedia.org/wiki/Cron                               |   |
| message  | string |    | 要发送钉钉的消息                        |   |
| dingding | string |    | 钉钉的webhook地址                       |   |
| title    | string |    | 任务名称：唯一 用来标识任务             |   |
| atAll    | bool   |    | 是否@所有人                             |   |
| atPhone  | string |    | 要@的用户的手机号，多个用户使用逗号分隔 |   |
| skipChineseHoliday| bool | | 是否跳过中国节假日 |