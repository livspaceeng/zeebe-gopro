package gateway

import (
	"context"
	"log"

	"github.com/zeebe-io/zeebe/clients/go/pkg/pb"
	"github.com/zeebe-io/zeebe/clients/go/pkg/zbc"
)

type GatewayServerImpl struct {
	// Implementation of pb.GatewayServer interface
}

var zbClient pb.GatewayClient
var ctx = context.Background()

func Init(gatewayAddr string) error {
	var err error
	zbClient, err = NewClient(&zbc.ClientConfig{
		GatewayAddress:         gatewayAddr,
		UsePlaintextConnection: true,
	})
	return err
}

func (*GatewayServerImpl) ActivateJobs(req *pb.ActivateJobsRequest, srv pb.Gateway_ActivateJobsServer) error {
	stream, err := zbClient.ActivateJobs(ctx, req)
	if err != nil {
		log.Println("Error while activating jobs:", err)
		return err
	}
	for {
		response, err := stream.Recv()
		if err != nil { // err == io.EOF
			break
		}
		srv.Send(response)
	}
	return err
}
func (*GatewayServerImpl) CancelWorkflowInstance(ctx context.Context, req *pb.CancelWorkflowInstanceRequest) (*pb.CancelWorkflowInstanceResponse, error) {
	response, err := zbClient.CancelWorkflowInstance(ctx, req)
	log.Println("Cancelled workflow instance: ", req)
	return response, err
}
func (*GatewayServerImpl) CompleteJob(ctx context.Context, req *pb.CompleteJobRequest) (*pb.CompleteJobResponse, error) {
	response, err := zbClient.CompleteJob(ctx, req)
	log.Println("Completed job: ", req)
	return response, err
}
func (*GatewayServerImpl) CreateWorkflowInstance(ctx context.Context, req *pb.CreateWorkflowInstanceRequest) (*pb.CreateWorkflowInstanceResponse, error) {
	response, err := zbClient.CreateWorkflowInstance(ctx, req)
	log.Println("Created workflow instance: ", req)
	return response, err
}
func (*GatewayServerImpl) CreateWorkflowInstanceWithResult(ctx context.Context, req *pb.CreateWorkflowInstanceWithResultRequest) (*pb.CreateWorkflowInstanceWithResultResponse, error) {
	response, err := zbClient.CreateWorkflowInstanceWithResult(ctx, req)
	return response, err
}
func (*GatewayServerImpl) DeployWorkflow(ctx context.Context, req *pb.DeployWorkflowRequest) (*pb.DeployWorkflowResponse, error) {
	response, err := zbClient.DeployWorkflow(ctx, req)
	log.Println("Deployed workflow.")
	return response, err
}
func (*GatewayServerImpl) FailJob(ctx context.Context, req *pb.FailJobRequest) (*pb.FailJobResponse, error) {
	response, err := zbClient.FailJob(ctx, req)
	log.Println("Failed job: ", req)
	return response, err
}
func (*GatewayServerImpl) ThrowError(ctx context.Context, req *pb.ThrowErrorRequest) (*pb.ThrowErrorResponse, error) {
	response, err := zbClient.ThrowError(ctx, req)
	return response, err
}
func (*GatewayServerImpl) PublishMessage(ctx context.Context, req *pb.PublishMessageRequest) (*pb.PublishMessageResponse, error) {
	response, err := zbClient.PublishMessage(ctx, req)
	return response, err
}
func (*GatewayServerImpl) ResolveIncident(ctx context.Context, req *pb.ResolveIncidentRequest) (*pb.ResolveIncidentResponse, error) {
	response, err := zbClient.ResolveIncident(ctx, req)
	log.Println("Resolved incident: ", req)
	return response, err
}
func (*GatewayServerImpl) SetVariables(ctx context.Context, req *pb.SetVariablesRequest) (*pb.SetVariablesResponse, error) {
	response, err := zbClient.SetVariables(ctx, req)
	return response, err
}
func (s *GatewayServerImpl) Topology(context.Context, *pb.TopologyRequest) (*pb.TopologyResponse, error) {
	topology, err := zbClient.Topology(ctx, &pb.TopologyRequest{})
	log.Println("Fetched topology: ", topology)
	return topology, err
}
func (*GatewayServerImpl) UpdateJobRetries(ctx context.Context, req *pb.UpdateJobRetriesRequest) (*pb.UpdateJobRetriesResponse, error) {
	response, err := zbClient.UpdateJobRetries(ctx, req)
	return response, err
}
