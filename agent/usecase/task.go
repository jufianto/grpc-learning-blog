package task

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Ullaakut/nmap/v3"
)

type TaskProcessing struct {
	nmapCl *nmap.Scanner
}

func NewTaskProcessing(ctx context.Context, targetHost string) (*TaskProcessing, error) {
	nmapClient, err := nmap.NewScanner(ctx, nmap.WithTargets(targetHost), nmap.WithPorts("80"))
	if err != nil {
		return nil, fmt.Errorf("failed to init nmap client: %w", err)
	}

	return &TaskProcessing{
		nmapCl: nmapClient,
	}, nil
}

func (t *TaskProcessing) DoTaskScanner() (string, error) {
	log.Println("run nmap")

	res, warning, err := t.nmapCl.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run nmap: %w", err)
	}

	if len(*warning) > 0 {
		log.Println("have warnings", *warning)
	}

	var bufRes = make([]byte, 10)
	var resultB []byte

	resReader := res.ToReader()

	for {
		n, err := resReader.Read(bufRes)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("has an error: %v \n", err)
			break
		}
		resultB = append(resultB, bufRes[:n]...)

	}

	for _, host := range res.Hosts {
		fmt.Println("host", host.Addresses, host.Ports, host.Status)
	}

	return string(resultB), nil
}
