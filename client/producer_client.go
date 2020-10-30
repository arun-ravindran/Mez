// Producer client (camera & edgenode functionality)
package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Ann-Geo/Mez/api/edgenode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var customTimeformat string = "Monday, 02-Jan-06 15:04:05.00000 MST"

type ProducerClient struct {
	Auth           Authentication
	ConnProdClient *grpc.ClientConn
	Cl             edgenode.PubSubClient
	Cancel         context.CancelFunc
	Ctx            context.Context
}

func NewProducerClient(login, password string) *ProducerClient {
	return &ProducerClient{Auth: Authentication{login: login, password: password}}

}

func (pc *ProducerClient) Connect(url, userAddress string) error {
	var err error

	creds, err := credentials.NewClientTLSFromFile("../../cert/server.crt", "")
	if err != nil {
		return err
	}

	// Dial to Mez
	pc.ConnProdClient, err = grpc.Dial(url, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&pc.Auth))
	if err != nil {
		return err
	}

	pc.Cl = edgenode.NewPubSubClient(pc.ConnProdClient)

	pc.Ctx, pc.Cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//Connect with Mez
	connReq := &edgenode.Url{
		Address: userAddress,
	}
	_, connErr := pc.Cl.Connect(pc.Ctx, connReq)
	if connErr != nil {
		return connErr
	}

	return nil
}

func (pc *ProducerClient) Retry(url, userAddress string, numRetries int) (err error) {

	for i := 0; i < numRetries; i++ {
		err = pc.Connect(url, userAddress)
		if err == nil {

			return nil
		}

	}
	return err
}

func (pc *ProducerClient) PublishImage(client edgenode.PubSubClient) error {
	// Client side streaming
	var (
		writing        = true
		buf            []byte
		chunkSize      int = 1 << 20
		imageFilesPath     = "../test_images"
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Open a stream to gRPC server
	stream, err := client.Publish(ctx)
	if err != nil {
		return fmt.Errorf("%v.Publish(_) = _, %v", client, err)
	}
	defer stream.CloseSend()

	// Read image file names
	var files []string
	absImageFilePath, _ := filepath.Abs(imageFilesPath)
	err = filepath.Walk(absImageFilePath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return err
	}
	// Copy image file to buffer
	numSend := 0
	for i := 1; i < len(files); i++ {
		writing = true
		fd, err := os.Open(files[i])
		if err != nil {
			return err
		}
		defer fd.Close()
		buf = make([]byte, chunkSize)
		for writing {
			n, err := fd.Read(buf)
			if err != nil {
				if err == io.EOF {
					writing = false
					err = nil
					continue
				}
				return fmt.Errorf("errored while copying from file to buf")
			}
			time.Sleep(1 * time.Second)
			ts, _ := time.Parse(customTimeformat, time.Now().Format(customTimeformat))
			err = stream.Send(&edgenode.Image{
				Image:     buf[:n],                     // Needed if n < chunksize
				Timestamp: ts.Format(customTimeformat), // Timestamp the image
			})
			numSend++
			log.Println("ProducerClient: Image size sent kB", n/1024)
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("failed to send chunk via stream")
			}
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Println("From TestBroker:PublishImage", reply)
	return nil
}


