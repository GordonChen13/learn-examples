package mongo

import (
	"fmt"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type MatchRepository struct {
	Client *mongo.Client
}

const (
	DbName = "users"
	CollectionName = "match"
)

func NewMatchRepository() (*MatchRepository, error) {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	//dbUser := viper.GetString(`database.user`)
	//dbPass := viper.GetString(`database.pass`)
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	ctxPing, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctxPing, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &MatchRepository{
		Client:client,
	}, nil
}

func (m *MatchRepository) Store(ctx context.Context, match *models.Match) error {
	collection := m.getMatchCollection()

	data := bson.M{
		"id" : match.Id,
		"name" : match.Name,
	}
	_, err := collection.InsertOne(ctx, data)

	if err != nil {
		return err
	}

	return nil
}

func (m *MatchRepository) GetById(ctx context.Context, id string) (*models.Match, error) {
	match := &models.Match{}
	collection := m.getMatchCollection()

	log.Println(id)
	intId, _ := strconv.ParseInt(id, 10, 64)
	filter := bson.M{"id": intId}
	err := collection.FindOne(ctx, filter).Decode(match)

	if err != nil {
		return nil, err
	}

	return match,nil
}

func (m *MatchRepository) GetByName(ctx context.Context, name string) (*models.Match, error) {
	match := &models.Match{}
	collection := m.getMatchCollection()

	filter := bson.M{"name":name}
	err := collection.FindOne(ctx, filter).Decode(match)

	if err != nil {
		return nil, err
	}

	return match,nil
}

func (m *MatchRepository) getMatchCollection() *mongo.Collection {
	return m.Client.Database(DbName).Collection(CollectionName)
}