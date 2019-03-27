package datebase

import "fmt"
import "gopkg.in/redis.v5"

var client *redis.Client
func factory(name string) *redis.Client {
	//conf := "master"
	//host := "192.168.99.100:6379"
	host := "127.0.0.1:6379"
	port := 6379
	password := ""
	fmt.Printf("conf-redis: %s:%s - %s\r\n", host, port, password)

	//address := fmt.Sprintf("%s:%s", host, port)
	if client ==nil {
		client = redis.NewClient(&redis.Options{
			Addr:        host,
			Password:    password,
			DB:          0,
			PoolSize:    5,
		})
	}

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

/**
 * 获取连接
 */
func getRedis(name string) *redis.Client {
	return factory(name)
}

/**
 * 获取master连接
 */
func Master() *redis.Client {
	return getRedis("master")
}

/**
 * 获取slave连接
 */
func Slave() *redis.Client {
	return getRedis("slave")
}
