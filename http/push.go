package http

import (
	"encoding/json"
	"github.com/open-falcon/common/model"
	"github.com/open-falcon/forwarder/g"
	"net/http"
)

func configPushRoutes() {
	http.HandleFunc("/v1/push", func(w http.ResponseWriter, req *http.Request) {
		if req.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(req.Body)
		var metrics []*model.MetricValue
		err := decoder.Decode(&metrics)
		if err != nil {
			http.Error(w, "connot decode body", http.StatusBadRequest)
			return
		}

		tags := g.Config().AttachTags
		if tags != "" {
			count := len(metrics)
			for i := 0; i < count; i++ {
				if metrics[i].Tags == "" {
					metrics[i].Tags = tags
				} else {
					metrics[i].Tags = metrics[i].Tags + "," + tags
				}
			}
		}

		err = g.SendToTransfer(metrics)
		msg := "success"
		if err != nil {
			msg = err.Error()
		}
		w.Write([]byte(msg))
	})
}
