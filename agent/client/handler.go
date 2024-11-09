package client

import (
	"context"
	"fmt"
	"log"
	"time"

	task "github.com/jufianto/state-agent/agent/usecase"
	agentsrv_proto "github.com/jufianto/state-agent/proto"
)

func (a *AgentService) CreateTask(ctx context.Context, req *agentsrv_proto.TaskRequest) (*agentsrv_proto.TaskResponse, error) {
	fmt.Printf("got request %+v \n", req)

	if true {
		// taskVal, ok := a.sm.Get(req.TaskId)
		// if ok {
		// 	log.Printf("duplicate taskid %s, the task status is: %s \n", taskVal.TaskID, taskVal.TaskStatus)
		// 	return &agentsrv_proto.TaskResponse{
		// 		TaskId:     req.TaskId,
		// 		TaskResult: "task already processing",
		// 		TaskStatus: agentsrv_proto.TaskStatus_PROCESSING,
		// 	}, nil
		// }

		go func(taskID string) {
			log.Println("run task concurrently", taskID)
			ctxTimeout, _ := context.WithTimeout(context.Background(), 3*time.Minute)
			// defer cancel()
			tp, err := task.NewTaskProcessing(ctxTimeout, req.GetTaskUrl())
			if err != nil {
				log.Printf("error on create task %v \n", err)
				return
			}
			a.sm.Set(req.GetTaskId(), TaskResult{TaskID: taskID, TaskStatus: "processing"})
			res, err := tp.DoTaskScanner()
			if err != nil {
				log.Printf("error on result task %v \n", err)
				return
			}
			a.sm.Set(req.GetTaskId(), TaskResult{
				TaskID:     taskID,
				TaskResult: res,
				TaskStatus: "done",
			})
			// fmt.Printf("done with task %s: %v", taskID, res)
		}(req.TaskId)
	}

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
