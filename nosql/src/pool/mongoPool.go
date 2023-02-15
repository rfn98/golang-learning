package Pool

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)


type MongoPool struct {
	Pool        chan *mongo.Client
	Timeout     time.Duration
	Uri         string
	Connections int
	PoolSize    int
}


func (mp *MongoPool) getContextTimeOut() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), mp.Timeout)
	return ctx
}

func (mp *MongoPool) createToChan() {
	var conn *mongo.Client
	conn, e := mongo.NewClient(options.Client().ApplyURI(mp.Uri))
	if e != nil {
		log.Fatalf("Create the Pool failed，err=%v", e)
	}
	e = conn.Connect(mp.getContextTimeOut())
	if e != nil {
		log.Fatalf("Create the Pool failed，err=%v", e)
	}
	mp.Pool <- conn
	mp.Connections++
}

func (mp *MongoPool) CloseConnection(conn *mongo.Client) error {
	select {
	case mp.Pool <- conn:
		return nil
	default:
		if err := conn.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Close the Pool failed，err=%v", err)
			return err
		}
		mp.Connections--
		return nil
	}
}

func (mp *MongoPool) GetConnection() (*mongo.Client, error) {
	for {
		select {
		case conn := <-mp.Pool:
			err := conn.Ping(mp.getContextTimeOut(), readpref.Primary())
			if err != nil {
				log.Fatalf("ERROR GET CONNECTION，err=%v", err)
				return nil, err
			}
			return conn, nil
		default:
			if mp.Connections < mp.PoolSize {
				mp.createToChan()
			}
		}
	}
}

func GetCollection(conn *mongo.Client, dbname, collection string) *mongo.Collection {
	return conn.Database(dbname).Collection(collection)
}