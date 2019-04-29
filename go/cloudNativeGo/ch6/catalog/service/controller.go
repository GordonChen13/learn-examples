package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"strconv"
	"time"
)

type Controller struct {
	DB *mongo.Client
}
func NewController() *Controller {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("connect to mongodb err: %s", err)
	}

	return &Controller{
		DB:client,
	}
}

func (c *Controller) Create(ctx *gin.Context) (res *Catalog, err error) {
	sku := ctx.PostForm("sku")
	shipWithIn, err := strconv.Atoi(ctx.PostForm("shipswithin"))
	quantilyInStock, err := strconv.Atoi(ctx.PostForm("stock"))
	id := NewCatalogId()

	catalog := &Catalog{
		Id:id,
		SKU:sku,
		ShipsWithin:shipWithIn,
		QuantityInStock:quantilyInStock,
	}

	_, err = c.DB.Database("users").Collection("catalog").InsertOne(ctx, catalog)

	return catalog, err
}

func (c *Controller) Get(ctx *gin.Context) (res *Catalog, err error) {
	sku := ctx.Query("sku")
	if sku != "" {
		res, err = c.GetBySku(ctx, sku)
	}

	return
}

func (c *Controller) GetBySku(ctx *gin.Context, sku string) (*Catalog, error) {
	catalog := &Catalog{}
	filter := bson.M{"sku": sku}
	err := c.DB.Database("users").Collection("catalog").FindOne(ctx, filter).Decode(&catalog)

	return catalog,err
}
