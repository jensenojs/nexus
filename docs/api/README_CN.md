# API参考手册

nexus通过http工作，接下来阐述nexus的api。

## Dataset管理

Dataset管理是nexus创建，删除和查询dataset的接口。

### 创建dataset

通过spl或者schema信息创建dataset。

例:

```shell
// request
curl -X POST --data '{"method": "create_dataset", "params":{"name":"test","attributes":[{"name":"uid","type":"string"}]} }'
//result
{
	"errcode": 0, // success
	"errmsg": ""
}
// request
curl -X POST --data '{"method": "create_dataset", "params":{"name":"test","spl": "| from dataset:users" }}'
//result
{
	"errcode": 0, // success
	"errmsg": ""
}
```

### 删除dataset

根据dataset的名字删除dataset。

例:

```shell
// request
curl -X POST --data '{"method": "delete_dataset", "params":{"name":"test"}}'
//result
{
	"errcode": 0, // success
	"errmsg": ""
}
```

### 查询dataset信息

通过dataaset名字查询dataset信息。

例:

```shell
// request
curl -X POST --data '{"method": "get_dataset", "params":{"name":"test"}}'
//result
{
	"errcode": 0, // success
	"errmsg": "",
	"result":{"type":"source","name":"test","spl":"","attributes":[{"name":"uid","type":"string"}]}
}
```

### 查询dataset列表

列出系统中的所有dataset。

例:

```shell
// request
curl -X POST --data '{"method": "list_dataset", "params":{}}'
//result
{
	"errcode": 0, // success
	"errmsg": "",
	"result":[{"type":"source","name":"test","spl":"","attributes":[{"name":"uid","type":"string"}]}]
}
```

## 数据打点

该接口用于将日志数据导入source dataset，需要注意的是nexus并不会做日志的采集和预处理。对于nexus，生成系统字段是必要的预处理。

例:

```shell
// request
curl -X POST --data '{"method": "insert", "params":{"name":"test","type":"csv","Data":"1, 2, 3\n2,3,4\n"}}'
//result
{
	"errcode": 0, // success
	"errmsg": "",
}
```

## 数据查询

该接口用于用户执行spl语句。

例：

```shell
// request
curl -X POST --data '{"method": "query", "params":"|dataset:test"}'
//result
{
	"errcode": 0, // success
	"errmsg": "",
	"result": {
        "fields": [
           "_raw"
        ],
        "rows": "1, 2, 3\n2,3,4\n"
    }
}
```
