package g

import (
	"github.com/open-falcon/common/model"
	"log"
	"time"
)

var (
	TransferClient *SingleConnRpcClient
)

func InitRpcClients() {
	if Config().Transfer.Enabled {
		TransferClient = &SingleConnRpcClient{
			RpcServer: Config().Transfer.Addr,
			Timeout:   time.Duration(Config().Transfer.Timeout) * time.Millisecond,
		}
	}
}

func SendToTransfer(metrics []*model.MetricValue) error {
	if len(metrics) == 0 {
		return nil
	}

	debug := Config().Debug

	if debug {
		log.Printf("=> <Total=%d> %v\n", len(metrics), metrics[0])
	}

	var resp model.TransferResponse
	err := TransferClient.Call("Transfer.Update", metrics, &resp)
	if err != nil {
		log.Println("call Transfer.Update fail", err)
	} else {
		if debug {
			log.Println("<=", &resp)
		}
	}

	return err
}
