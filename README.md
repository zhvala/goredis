# goredis

goredis 是一个go 语言的redis-client 实现，既支持redis 单点模式也支持cluster 集群模式。基于[redigo](https://github.com/garyburd/redigo) 和 [redis-go-cluster ](https://github.com/chasex/redis-go-cluster)开发，由于[redis-go-cluster ](https://github.com/chasex/redis-go-cluster) 原作者近些年已经不进行维护更新了，所以fork 单独进行维护，修复了 [redis-go-cluster](https://github.com/chasex/redis-go-cluster) 的一些bug 并添加了部分新的特性。

### 修复BUG

- 添加了对redis 单点模式也支持cluster 集群模式的支持（[redigo](https://github.com/garyburd/redigo) 不支持集群模式，[redis-go-cluster ](https://github.com/chasex/redis-go-cluster) 不支持单点模式）；
- 修复redis-server 进程挂掉-恢复或cluster 节点主从切换后，redis-client 的不能回复的bug；
- 添加对cluster 集群密码的支持；
- 修复redis数据中含有特殊字符导致的bug；

### 特性

**支持：**

- Most commands of keys, strings, lists, sets, sorted sets, hashes.
- MGET/MSET
- Pipelining

**不支持（todo）：**

- Cluster commands
- Pub/Sub
- Transaction
- Lua script

## 文档

[文档地址](https://godoc.org/github.com/zhvala/goredis)

## 安装

```shell
go get -u github.com/zhvala/goredis
```

## 例子

- 自动识别是否cluster

```go
import "https://github.com/zhvala/goredis"

// cluster 集群模式
conn, err := redis.NewConn(
    &redis.Options{
    // 至少提供一个redis cluster 节点
	StartNodes: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
	ConnTimeout: 50 * time.Millisecond,
	ReadTimeout: 50 * time.Millisecond,
	WriteTimeout: 50 * time.Millisecond,
	KeepAlive: 16,
	AliveTime: 60 * time.Second,
    })
// single-node 模式
conn, err := redis.NewConn(
    &redis.Options{
	StartNodes: []string{"127.0.0.1:6379"},
	ConnTimeout: 50 * time.Millisecond,
	ReadTimeout: 50 * time.Millisecond,
	WriteTimeout: 50 * time.Millisecond,
	KeepAlive: 16,
	AliveTime: 60 * time.Second,
    })
```

## 联系方式

```shell
zhvala@foxmail.com
```

## License

redis-go-cluster is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).