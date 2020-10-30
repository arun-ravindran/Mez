// Consumer client (Analytics application)
package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Ann-Geo/Mez/api/edgenode"
	"github.com/Ann-Geo/Mez/api/edgeserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ConsumerClient struct {
	Auth           Authentication
	ConnConsClient *grpc.ClientConn
	Cl             edgeserver.PubSubClient
	Cancel         context.CancelFunc
	Ctx            context.Context
}

var NumImRcvdUnsubTest uint64

func NewConsumerClient(login, password string) *ConsumerClient {
	return &ConsumerClient{Auth: Authentication{login: login, password: password}}
}

func (cc *ConsumerClient) Connect(url, userAddress string) error {

	var err error

	creds, err := credentials.NewClientTLSFromFile("../../cert/server.crt", "")
	if err != nil {
		return err
	}

	// Dial to Mez
	cc.ConnConsClient, err = grpc.Dial(url, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&cc.Auth))
	if err != nil {
		return err
	}

	cc.Cl = edgeserver.NewPubSubClient(cc.ConnConsClient)

	cc.Ctx, cc.Cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//Connect with Mez
	connReq := &edgeserver.Url{
		Address: userAddress,
	}
	_, connErr := cc.Cl.Connect(cc.Ctx, connReq)
	if connErr != nil {
		return connErr
	}

	return nil
}

func (cc *ConsumerClient) Retry(url, userAddress string, numRetries int) (err error) {

	for i := 0; i < numRetries; i++ {
		err = cc.Connect(url, userAddress)
		if err == nil {

			return nil
		}

	}
	return err
}

func (cc *ConsumerClient) SubscribeImage(client edgeserver.PubSubClient, tbegin time.Time) error {
	// Server side streaming
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := tbegin.Format(customTimeformat)
	stop := tbegin.Add(5 * time.Second).Format(customTimeformat)

	imPars := &edgeserver.ImageStreamParameters{Camid: "cam1", Latency: "100", Accuracy: "100",
		Start: start, Stop: stop}
	// Open a stream to gRPC server
	stream, err := client.Subscribe(ctx, imPars)
	fmt.Println(stream)
	if err != nil {
		return err
	}

	numImagesRecvd := 0
	for {
		im, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		numImagesRecvd++
		log.Printf("Consumer client: Number of images received %d, of size %d, and timestamp %s", numImagesRecvd, len(im.GetImage()), im.GetTimestamp())

		// {Here would be application code to process image received}

	}
	return nil
}

func (cc *ConsumerClient) Unsubscribe(client edgeserver.PubSubClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	appInfo := &edgeserver.AppInfo{Appid: "app1", Camid: "cam1"}
	_, err := client.Unsubscribe(ctx, appInfo)
	if err != nil {
		log.Fatalf("Unsubscribe failed")
	}
}


