package main

import (
	"fmt"
	"log"
	"time"
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func main() {

	message := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

	l := &Location{
		Id:        65988,
		UserId:    55,
		Lat:       -7.112213,
		Long:      -34.843408,
		Hdrop:     22111,
		Altitude:  58,
		Speed:     30,
		Message:   message,
		Timestamp: ptypes.TimestampNow(),
	}

	protom, err := proto.Marshal(l)
	if err != nil {
		log.Fatal("Marshall error: ", err)
	} else {

		time := time.Now().Unix()
		in := fmt.Sprintf("{Id: 65988, UserId: 55, Lat: -7.112213, Long: -34.843408, Hdrop: 22111, Altitude: 58, Speed: 30, Message: %s, Timestamp: %d}", message, time)
		jsonm, _ := json.Marshal(in)

		//fmt.Println([]byte(jsonm))
		fmt.Printf("json: %d\n", len([]byte(jsonm)))

		//fmt.Println([]byte(protom))
		fmt.Printf("proto: %d\n", len([]byte(protom)))
	
	}

}
