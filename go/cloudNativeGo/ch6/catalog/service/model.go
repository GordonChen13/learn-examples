package service

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

type Catalog struct {
	Id int64
	SKU string `form:"sku" binding:"required"`
	ShipsWithin int `form:"shipswithin" binding:"required"`
	QuantityInStock int `form:"stock" binding:"required"`
}

func NewCatalogId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	id := node.Generate()
	return int64(id)
}
