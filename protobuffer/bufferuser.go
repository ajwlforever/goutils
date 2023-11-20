package protobuffer

import (
	"fmt"
	pb "github.com/ajwlforever/goutils/protobuffer/message"
	"github.com/golang/protobuf/proto"
	"log"
)

func UseStudent() {
	s := &pb.Student{
		Name:   "yqx",
		Male:   true,
		Scores: []int32{88, 99, 91},
	}

	data, err := proto.Marshal(s)
	if err != nil {
		log.Fatal("Marshal failed")
	}
	fmt.Println("data:", data)

	s1 := &pb.Student{}
	err = proto.Unmarshal(data, s1)
	if err != nil {
		log.Fatal("Unmarshal failed")
	}
	fmt.Println("s1:", s1)
	if s1.GetName() != s.GetName() {
		log.Fatal("ssss")
	}
}
