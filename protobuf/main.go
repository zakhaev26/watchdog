package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/protozakhaev26/model"
)

func main() {
	message := &model.KibanaMessage{
		CpuUsage:  123.213,
		Time:      "test",
		Timestamp: "asdsasdf",
	}

	data, err := proto.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	var m model.KibanaMessage

	err = proto.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.CpuUsage)
}
