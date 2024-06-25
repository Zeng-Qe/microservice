package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"time"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var REDIS *redis.Client

func RedisInit() *redis.Client {
	conf, err := loadConfig()
	if err != nil {
		log.Panic("加载redis配置失败:", err)
		return nil
	}

	client := connectRedis(conf)

	if checkRedisClient(client) != nil {
		log.Panic("连接redis失败:", err)
		return nil
	}
	REDIS = client
	return client
}

func connectRedis(conf *RedisConfig) *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	// 如果返回nil，就创建这个DB

	return redisClient
}

func checkRedisClient(redisClient *redis.Client) error {
	// 通过 client.Ping() 来检查是否成功连接到了 redis 服务器
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	} else {
		log.Println("连接redis成完成...")
		return nil
	}
}

func loadConfig() (*RedisConfig, error) {
	redisConfig := &RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}

	if redisConfig.Addr == "" {
		redisConfig.Addr = "localhost:6379"
	}
	return redisConfig, nil
}

// Set 存储字符串
func Set(ctx context.Context, key string, value any, duration time.Duration) error {
	err := REDIS.Set(ctx, key, value, duration).Err()
	if err != nil {
		fmt.Sprintf("err:%v", err)
	}
	return err
}

// Get 获取字符串
func Get(ctx context.Context, key string) (any, error) {
	data, err := REDIS.Get(ctx, key).Result()
	if err != nil {
		fmt.Sprintf("err:%v", err)
	}
	return data, err
}

var rdb = REDIS

// hash 使用场景：存储结构化的对象数据，如用户信息。
func hash() {

	ctx := context.Background()

	// 存储用户信息到哈希
	err := rdb.HSet(ctx, "user:1000", map[string]interface{}{
		"name":  "kimi",
		"email": "kimi@example.com",
		"age":   20,
	}).Err()
	if err != nil {
		panic(err)
	}

	// 获取哈希中的字段
	age, err := rdb.HGet(ctx, "user:1000", "age").Int64()
	if err != nil {
		panic(err)
	}
	fmt.Println("Age:", age)

}

// list 使用场景：实现消息队列、任务队列。
func list() {
	ctx := context.Background()

	// LPUSH将元素添加到列表头部，用作入队
	err := rdb.LPush(ctx, "task_queue", "task1").Err()
	if err != nil {
		panic(err)
	}

	// BRPOP从列表尾部弹出元素，用作出队
	task, err := rdb.BRPop(ctx, 0, "task_queue").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Task:", task[1])
}

// set 使用场景：存储唯一性的数据集，如标签集合、好友列表。
func set() {
	ctx := context.Background()

	// 将元素添加到集合
	err := rdb.SAdd(ctx, "friends", "alice", "bob", "charlie").Err()
	if err != nil {
		panic(err)
	}

	// 获取集合中的元素数量
	count, err := rdb.SCard(ctx, "friends").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Friends count:", count)
}

// zset 使用场景：存储具有分数或权重的数据，如排行榜、优先级队列。
//func zset() {
//	ctx := context.Background()
//
//	// 将元素和分数添加到有序集合
//	err := rdb.ZAdd(ctx, "leaderboard", &redis.Z{Score: 83.5, Member: "kimi"}, &redis.Z{Score: 88.0, Member: "alice"}).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 获取有序集合中的元素排名
//	rank, err := rdb.ZRank(ctx, "leaderboard", "kimi").Int64()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Kimi's rank:", rank)
//}

//
//Redis的`Hash`是一种字符串类型的键值集合，非常适合用于存储对象或者缓存数据库中行的数据。在Redis中，每个`Hash`可以存储多达40亿个字段。以下是一些常用的`Hash`命令以及它们的使用场景：
//
//1. **HSET key field value**：
//- 为指定的`key`设置指定的`field`和`value`。如果`key`不存在，一个新的`Hash`被创建。如果`field`已经存在于`Hash`中，那么它的`value`将被更新。
//
//2. **HGET key field**：
//- 获取存储在`key`的`Hash`对象中的指定`field`的值。
//
//3. **HSETNX key field value**：
//- 仅当`field`不存在时，设置`field`的值。这可以用来避免覆盖已经存在的`field`。
//
//4. **HGETALL key**：
//- 返回`key`对应的`Hash`中的所有字段和值。返回值是一个列表，列表中的每个项都是一个字段名和值对。
//
//5. **HMSET key field1 value1 [field2 value2 ...]**：
//- 同时设置一个或多个`Hash`的字段。如果`key`不存在，一个新的`Hash`被创建。
//
//6. **HMGET key field1 [field2 ...]**：
//- 获取所有给定字段的值。
//
//7. **HINCRBY key field increment**：
//- 为`Hash`对象中的`field`增加指定的数值。这通常用于计数器。
//
//8. **HEXISTS key field**：
//- 检查给定`field`是否存在于`Hash`对象中。
//
//9. **HDEL key field [field ...]**：
//- 删除`Hash`中的一个或多个字段。
//
//10. **HLEN key**：
//- 返回`Hash`对象中字段的数量。
//
//11. **HKEYS key**：
//- 返回`Hash`对象中的所有字段列表。
//
//12. **HVALS key**：
//- 返回`Hash`对象中的所有值列表。
//
//13. **HSCAN key cursor [MATCH pattern] [COUNT count]**：
//- 迭代`Hash`中的字段和值。`cursor`参数是一个游标，`MATCH`和`COUNT`参数用于迭代控制。
//
//以下是使用Go语言和`go-redis/redis`库展示如何使用`Hash`命令的示例代码：
//
//```go
//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/go-redis/redis/v8"
//)
//
//func main() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "",
//		DB:       0,
//	})
//
//	ctx := context.Background()
//
//	key := "user:1000"
//	fieldName := "name"
//	value := "kimi"
//
//	// 设置单个字段的值
//	err := rdb.HSet(ctx, key, fieldName, value).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 获取单个字段的值
//	result, err := rdb.HGet(ctx, key, fieldName).Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Value:", result)
//
//	// 获取所有字段和值
//	allFieldsAndValues, err := rdb.HGetAll(ctx, key).Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("All fields and values:", allFieldsAndValues)
//
//	// 增加数值
//	err = rdb.HIncrBy(ctx, key, "age", 1).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 删除字段
//	err = rdb.HDel(ctx, key, "email").Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 获取字段数量
//	fieldCount, err := rdb.HLen(ctx, key).Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Field count:", fieldCount)
//}
//```
//
//在实际应用中，`Hash`命令可以用来存储用户信息、配置参数、商品详情等结构化数据。通过`Hash`命令，可以灵活地对这些数据进行增删改查操作。
