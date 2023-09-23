package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

/*
The init functions will be run before any functions in the package
we can also have multiple init functions in a single package, and they will
be run respectively as they are defined, The init functions can not be invoked directly
*/
func init() {
	/*
		Although it is a very bad mistake to create database connections in an init function
		Here it's just for a test so don't follow it ;(
	*/
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

/*
This will run in after the previous init function
*/
func init() {
	pingStatus := rdb.Ping(ctx)
	if pingStatus.Val() != "PONG" {
		fmt.Printf("Redis server is unavailable")
	}
}

func startTheApp() {
	statusCmd := rdb.Set(ctx, "key", "Maziar", 0)
	fmt.Println("Status Command  is : ", statusCmd)
	fmt.Println("value is : ", rdb.Get(ctx, "key").Val())
}
