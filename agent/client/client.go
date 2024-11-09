package client

import (
	"sync"

	agentsrv_proto "github.com/jufianto/state-agent/proto"
	"google.golang.org/grpc"
)

type AgentService struct {
	Key string
	sm  *safeMap
	agentsrv_proto.UnimplementedTaskServiceServer
}

type TaskResult struct {
	TaskID     string
	TaskResult string
	TaskStatus string
}

type safeMap struct {
	mu sync.Mutex
	m  map[string]TaskResult
}

func (sm *safeMap) Get(key string) (TaskResult, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	val, ok := sm.m[key]
	return val, ok
}

func (sm *safeMap) Set(key string, val TaskResult) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.m[key] = val
}

func NewSyncMap() *safeMap {
	return &safeMap{m: make(map[string]TaskResult, 5)}
}

func NewAgentClient(key string) *AgentService {
	return &AgentService{
		Key: key,
		sm:  NewSyncMap(),
	}
}

func (a *AgentService) RegisterGW(srvGrpc *grpc.Server) {
	agentsrv_proto.RegisterTaskServiceServer(srvGrpc, a)
}
