package client

import (
	"context"
	"fmt"
	"log"

	agentsrv_proto "github.com/jufianto/state-agent/proto"
)

func (a *AgentService) CreateTask(ctx context.Context, req *agentsrv_proto.TaskRequest) (*agentsrv_proto.TaskResponse, error) {
	fmt.Printf("got request %+v \n", req)

	// if true {
	// taskVal, ok := a.sm.Get(req.TaskId)
	// if ok {
	// 	log.Printf("duplicate taskid %s, the task status is: %s \n", taskVal.TaskID, taskVal.TaskStatus)
	// 	return &agentsrv_proto.TaskResponse{
	// 		TaskId:     req.TaskId,
	// 		TaskResult: "task already processing",
	// 		TaskStatus: agentsrv_proto.TaskStatus_PROCESSING,
	// 	}, nil
	// }

	return &agentsrv_proto.TaskResponse{
		TaskId:     req.TaskId,
		TaskResult: "waiting for processing",
		TaskStatus: agentsrv_proto.TaskStatus_PROCESSING,
	}, nil
}

func (a *AgentService) ListTask(ctx context.Context, req *agentsrv_proto.TaskListRequest) (*agentsrv_proto.TaskListResponse, error) {

	log.Println("list tasks", req.GetTasksId())

	return nil, fmt.Errorf("not yet implemented")

}

func (a *AgentService) StatusTask(ctx context.Context, req *agentsrv_proto.TaskStatusRequest) (*agentsrv_proto.TasksStatusResponse, error) {
	log.Println("status tasks", req.GetTaskId())

	return nil, fmt.Errorf("not yet implemented")
}
