package service

import (
	"context"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Repository struct {
	Client *mongo.Client
}

func NewRepository() (*Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27010"))
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
	return &Repository{
		Client:client,
	}, nil
}

func (r *Repository) UpdateLastTelemetryEvent(event common.TelemetryUpdatedEvent) (err error)  {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := r.Client.Database("user").Collection("telemetry_event")
	uo := options.Update()
	uo.SetUpsert(true)

	_, err = collection.UpdateOne(ctx, bson.M{"drone_id": event.DroneID}, bson.M{
		"drone_id": event.DroneID,
		"core_temp": event.CoreTemp,
		"up_time" : event.Uptime,
		"battery": event.RemainingBattery,
		"received_on": event.ReceiveOn,
	}, uo)
	return err
}

func (r *Repository) UpdateLastAlertEvent(event common.AlertSignalledEvent) (err error)  {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := r.Client.Database("user").Collection("alert_event")
	uo := options.Update()
	uo.SetUpsert(true)

	_, err = collection.UpdateOne(ctx, bson.M{"drone_id": event.DroneID}, bson.M{
		"drone_id": event.DroneID,
		"description": event.Description,
		"fault_code" : event.FaultCode,
		"received_on": event.ReceiveOn,
	}, uo)
	return err
}

func (r *Repository) UpdateLastPositionEvent(event common.PositionChangedEvent) (err error)  {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := r.Client.Database("user").Collection("alert_event")
	uo := options.Update()
	uo.SetUpsert(true)

	_, err = collection.UpdateOne(ctx, bson.M{"drone_id": event.DroneID}, bson.M{
		"drone_id": event.DroneID,
		"altitude": event.Altitude,
		"latitude" : event.Latitude,
		"longitude" : event.Longitude,
		"description": event.CurrentSpeed,
		"heading_cardinal" : event.HeadingCardinal,
		"received_on": event.ReceiveOn,
	}, uo)
	return err
}
