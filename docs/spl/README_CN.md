# SPL参考手册

nexus使用spl来分析日志数据，spl语言采用类似于unix管道符的形式来表达数据分析的过程，通过$|$来拆分步骤，每一步是一个特定的命令，用于完成

相应的分析，spl语法如下:

```sql
<spl commands> = | <spl command> | <spl command> | ...
```

比如:

```sql
| from dataset:test | search starttime="2006-01-02T15:04:05Z07:00" endtime="2007-01-02T15:04:05Z07:00" 
```

nexus支持的命令如下: (未来会支持更多的命令)

| spl命令 | 描述                                                         |
| ------- | ------------------------------------------------------------ |
| dedup   | 去重指定字段                                                 |
| eval    | 根据输入字段进行计算得到新的字段，并将新字段和输入字段一起输出 |
| fields  | 删除或者保留输入字段                                         |
| join    | join                                                         |
| limit   | 保留输入结果的前几行                                         |
| search  | 搜索命令                                                     |
| sort    | 根据指定字段排序                                             |
| stats   | 对日志搜索结果做统计分析(分组聚合)                           |
| where   | 根据布尔表达式过滤输入数据                                   |

### dedup

dedup的语法如下:

```sql
dedup <field-list>
```

例:

``` sql
... | dedup host
... | dedup host, host
```

### eval

eval的语法如下:

```sql
eval <field>=<expression>["," <field>=<expression>]...
```

例:

```sql
... | eval d = a+b
```

### fields

fields的语法如下:

```sql
fields [+|-] <field-list>
```

例:

```sql
... | eval a --保留a字段
... | eval +a --保留a字段
... | eval -b --移除a字段
```

### join

join的语法如下:

```sql
join <join-options> <field-list> <subsearch>
<join-options> = type=inner|left|diff -- 目前只支持inner
<subsearch> = "[" <sql commands> "]"
```

例:

```sql
... | join type=inner uid [| from dataset:test]
```

### limit

limit的语法如下:

```sql
limit <int>
```

例:

```sql
... | limit 10
```

### search

search的语法如下:

```sql
search [<time-opts>] [<index-expression>]
<time-opts> = startime=<string>|endtime=<string>
<index-expression> = <string> -- 全文搜索
```

例:

```sql
search starttime="2006-01-02T15:04:05Z07:00" endtime="2007-01-02T15:04:05Z07:00" 
```

### sort

sort的语法如下:

```sql
sort [<int>] <sort-field-list>
<sort-field-list> = <sort-field> [desc|asc], ["," <sort-field> [desc|asc]]...
```

例:

```sql
... | sort by uid
... | sort 10 by uid, date
```

### stats

stats的语法如下:

```sql
stats <stats-function(field)> [as field] ... [by field-list]
```

stats支持的统计函数如下:(未来会支持更多的函数)

* count
* sum
* max
* min
* avg

例:

```sql
... | stats count(host) as cnt by ip
```

### where

where的语法如下:

```sql
where <eval-expression>
```

例:

```sql
... | where like(ip, "198.*")
... | where a = 1
```

### 函数

nexus支持以下函数:

| 函数名                | 函数描述                                                     | 例子                            |
| --------------------- | ------------------------------------------------------------ | ------------------------------- |
| if(x, y, z)           | 如果x为true返回y，否则返回z                                  | if(status==200, "ok", "error")  |
| in(field, value-list) | 如果field在value-list中可以找到，那么返回true，否则返回false | in(status, "404", "500", "503") |
| like(text, pattern)   | 模式匹配                                                     | like(uid, "x.*")                |
| isnull(x)             | 判断x是否为null                                              | isnull(uid)                     |
| nullif(x, y)          | 如果x=y，返回null，否则返回x                                 | nullif(uid, gid)                |
| isnotnull(x)          | 判断x是否不等于null                                          | isnotnull(uid)                  |

### 运算符

nexus支持以下运算符:

```sql
+, -, *, /, <, <=, =, >=, <>, %, and, or, not
```



## Dataset

datasets由多个dataset构成，dataset包含以下分类:

* source dataset - 用于直接接受日志数据
* derived dataset - 通过spl命令派生的dataset，不支持直接插入数据，通过从其他的dataset获取数据来增量更新

一个dataset由多个字段构成，如果不指定那么一个dataset只有系统字段。



### 系统字段

nexus的系统字段如下:

| 字段名字   | 类型   | 说明                             |
| ---------- | ------ | -------------------------------- |
| _raw       | string | 数据文本内容                     |
| _time      | long   | 事件发生的事件                   |
| host       | string | 事件来源机器的hostname或者ip地址 |
| linecount  | Long   | 事件的行数                       |
| source     | string | 事件的来源的信息                 |
| sourcetype | string | 事件来源的输入格式               |



## 数据类型

nexus支持以下数据类型: 

* boolean - openspl的boolean为小写的true和false
* int - 4字节
* long - 8字节
* float - 4字节
* double - 8字节
* string - 所有的字符串都用双引号
