package ccancel

import (
	"CraneFrontEnd/generated/protos"
	"CraneFrontEnd/internal/util"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	stub protos.CraneCtldClient
)

func CancelTask(taskId uint32) {
	req := &protos.CancelTaskRequest{OperatorUid: uint32(os.Getuid()), TaskId: uint32(taskId)}

	reply, err := stub.CancelTask(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send TerminateTask gRPC: %s", err.Error())
	}

	if reply.Ok {
		fmt.Printf("Task #%d is terminating...\n", taskId)
	} else {
		fmt.Printf("Failed to terminating task #%d: %s\n", taskId, reply.Reason)
	}
}
func Init() {
	config := util.ParseConfig()

	serverAddr := fmt.Sprintf("%s:%s", config.ControlMachine, config.CraneCtldListenPort)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Cannot connect to CraneCtld: " + err.Error())
	}

	stub = protos.NewCraneCtldClient(conn)
}
