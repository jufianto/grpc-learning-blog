package client

import (
	safemap "github.com/jufianto/state-agent/pkg"
	agentsrv_proto "github.com/jufianto/state-agent/proto"
	"google.golang.org/grpc"
)

type AgentService struct {
	Key string
	sm  *safemap.SafeMap
	agentsrv_proto.UnimplementedTaskServiceServer
}

type TaskResult struct {
	TaskID     string
	TaskResult string
	TaskStatus string
}

func NewAgentClient(key string) *AgentService {
	return &AgentService{
		Key: key,
		sm:  safemap.NewSyncMap(),
	}
}

func (a *AgentService) RegisterGW(srvGrpc *grpc.Server) {
	agentsrv_proto.RegisterTaskServiceServer(srvGrpc, a)
}
