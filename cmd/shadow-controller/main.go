package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	docker "github.com/fsouza/go-dockerclient"

	"github.com/wwqdrh/ds-connect/pkg/swarm"
	"github.com/wwqdrh/logger"
)

var client *docker.Client

func init() {
	var err error
	client, err = docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
}

func updateNetwork(ctx context.Context) {
	t := time.NewTicker(2 * time.Minute)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("do a update...")
			if err := swarm.ServiceUpdateNet(client, "ds-connect-shadow", swarm.GetServiceDiff(client, "ds-connect-shadow")); err != nil {
				fmt.Println(err.Error())
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	_, err := swarm.ServiceInspect(client, "ds-connect-shadow")
	if err != nil {
		envs := []string{}
		id, err := swarm.ServiceCreate(client, "ds-connect-shadow", "wwqdrh/ds-connect-shadow:dev", envs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("创建成功: " + id)
	} else {
		fmt.Println("已经存在")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go updateNetwork(ctx)

	mux := http.NewServeMux()
	registerService(mux)

	srv := http.Server{
		Addr:         ":18080",
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}
	go func() {
		logger.DefaultLogger.Info("Starting httpserver at :8080")
		if err := srv.ListenAndServe(); err != nil {
			logger.DefaultLogger.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-quit
	srv.Close()
}

func registerService(engine *http.ServeMux) {
	engine.HandleFunc("/service/query", func(w http.ResponseWriter, r *http.Request) {
		// 获取服务名的ip地址
		name := r.URL.Query().Get("name")
		servers, err := swarm.ServiceIPMap(client)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		if val, ok := servers[name]; !ok {
			w.WriteHeader(404)
			w.Write([]byte("not found"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte(val))
		}
	})

	engine.HandleFunc("/service/dnsmap", func(w http.ResponseWriter, r *http.Request) {
		servers, err := swarm.ServiceIPMap(client)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		} else {
			res := []string{}
			for _, val := range servers {
				res = append(res, val)
			}

			data, err := json.Marshal(res)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(200)
				w.Write(data)
			}
		}
	})
}
