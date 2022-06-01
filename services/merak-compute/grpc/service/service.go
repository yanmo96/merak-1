package service

import (
	"context"
	"flag"
	"log"

	pb "github.com/futurewei-cloud/merak/api/proto/v1/merak"
	constants "github.com/futurewei-cloud/merak/services/common"
	"github.com/futurewei-cloud/merak/services/merak-compute/workflows/vm"
	"go.temporal.io/sdk/client"
)

var (
	Port            = flag.Int("port", constants.COMPUTE_GRPC_SERVER_PORT, "The server port")
	workflowOptions client.StartWorkflowOptions
)

type Server struct {
	pb.UnimplementedMerakComputeServiceServer
}

var TemporalClient client.Client

func (s *Server) ComputeHandler(ctx context.Context, in *pb.InternalComputeConfigInfo) (*pb.ReturnMessage, error) {
	log.Printf("Received on ComputeHandler %s", in)
	// Parse input

	switch op := in.OperationType; op {
	case pb.OperationType_INFO:
		workflowOptions = client.StartWorkflowOptions{
			ID:        "InfoWorkflow",
			TaskQueue: "Info",
		}
	case pb.OperationType_CREATE:
		workflowOptions = client.StartWorkflowOptions{
			ID:        "CreateWorkflow",
			TaskQueue: "Create",
		}
	case pb.OperationType_UPDATE:
		workflowOptions = client.StartWorkflowOptions{
			ID:        "UpdateWorkflow",
			TaskQueue: "Update",
		}
	case pb.OperationType_DELETE:
		workflowOptions = client.StartWorkflowOptions{
			ID:        "DeleteWorkflow",
			TaskQueue: "Delete",
		}
	default:
		log.Println("Unknown Operation")
	}

	we, err := TemporalClient.ExecuteWorkflow(context.Background(), workflowOptions, vm.Create, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	return in, nil
}

func (s *Server) TestHandler(ctx context.Context, in *pb.InternalComputeConfigInfo) (*pb.ReturnMessage, error) {
	log.Printf("Received on TestHandler %s", in)
	return in, nil
}
