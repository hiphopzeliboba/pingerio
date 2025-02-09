package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"log"
	"net/http"
	"time"

	"github.com/docker/docker/client"
)

type ContainerInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	//IP       string    `json:"ip"`
	Status   string    `json:"status"`
	Created  time.Time `json:"created"`
	PingTime time.Time `json:"ping_time"`
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Ошибка при создании Docker клиента: %v", err)
	}

	http.HandleFunc("/containers", func(w http.ResponseWriter, r *http.Request) {
		containers, err := getContainers(cli)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(containers)
	})

	fmt.Println("Сервер запущен на порту 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func getContainers(cli *client.Client) ([]ContainerInfo, error) {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var containersInfo []ContainerInfo
	for _, container := range containers {
		info := ContainerInfo{
			ID:       container.ID[:12],
			Name:     container.Names[0][1:], // Удаляем начальный слэш
			Image:    container.Image,
			Status:   container.Status,
			Created:  time.Unix(container.Created, 0),
			PingTime: time.Now(),
		}
		containersInfo = append(containersInfo, info)
	}

	return containersInfo, nil
}
