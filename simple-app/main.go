package main

import (
	"fmt"
	"log"

	"gitlab.com/bfelipe/go-grpc/simple-app/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	app := pb.Application{
		Name:      "my-app",
		Id:        1,
		IsBackend: true,
		Type:      pb.ApplicationType_API,
		Team: &pb.Team{
			Name:    "Awesome team",
			Members: []string{"Erick", "Sonia", "Matt"},
		},
		CreatedAt: timestamppb.Now(),
	}

	fmt.Printf("Proto Application %s \n", &app)

	// convert to bytes for transmition over network
	data, err := proto.Marshal(&app)
	if err != nil {
		log.Fatalf("Err %s", err)
	}

	fmt.Printf("Marshalled data: %v\n", data)

	// convert to application struct variable
	var unmarshalledApp pb.Application
	if err := proto.Unmarshal(data, &unmarshalledApp); err != nil {
		log.Fatalf("Err %s", err)
	}

	fmt.Printf("Proto Application %s \n", &unmarshalledApp)

	// enconding/json does not serialize proto time properly, instead use protojson
	jsonData, err := protojson.Marshal(&app)
	if err != nil {
		log.Fatalf("Err %s", err)
	}

	fmt.Printf("Json Application %s \n", jsonData)

	fmt.Printf("Proto size data %d \n", len(data))
	fmt.Printf("Json size data %d \n", len(jsonData))

	// Something I noticed when using enconding/json and enconding/protojson
	// is that enconding/json marshal results into a slightly bigger size
	// compared to enconding/protojson marshal. Although enconding/protojson
	// is slightly slow compared to enconding/json
}
