package hash

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type BikeInfo struct {
	Model string `redis:"model"`
	Brand string `redis:"brand"`
	Type  string `redis:"type"`
	Price int    `redis:"price"`
}

func HashFunction(ctx context.Context, rdb *redis.Client) {
	hashFields := []string{
		"model", "Deimos",
		"brand", "Ergonom",
		"type", "Enduro bikes",
		"price", "4972",
	}

	res1, err := rdb.HSet(ctx, "bike:1", hashFields).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res1) // >>> 4

	res2, err := rdb.HGet(ctx, "bike:1", "model").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res2) // >>> Deimos

	res3, err := rdb.HGet(ctx, "bike:1", "price").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res3) // >>> 4972

	res4, err := rdb.HGetAll(ctx, "bike:1").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res4)
	// >>> map[brand:Ergonom model:Deimos price:4972 type:Enduro bikes]

	var res4a BikeInfo
	err = rdb.HGetAll(ctx, "bike:1").Scan(&res4a)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Model: %v, Brand: %v, Type: %v, Price: $%v\n",
		res4a.Model, res4a.Brand, res4a.Type, res4a.Price)

}
