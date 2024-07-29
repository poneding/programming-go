package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/client"
)

func main() {
	// 创建 Docker 客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	// 定义镜像名称
	imageName := "registry.cn-hangzhou.aliyuncs.com/pding/nginx:latest"

	// 获取镜像的 inspect 信息
	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		log.Fatalf("Error inspecting image: %v", err)
	}

	// 打印 inspect 信息
	// fmt.Printf("Image ID: %s\n", imageInspect.ID)
	// fmt.Printf("RepoTags: %v\n", imageInspect.RepoTags)
	// fmt.Printf("Size: %d\n", imageInspect.Size)
	// fmt.Printf("Created: %s\n", imageInspect.Created)
	// fmt.Printf("Architecture: %s\n", imageInspect.Architecture)
	// fmt.Printf("Os: %s\n", imageInspect.Os)
	fmt.Printf("Os: %s\n", imageInspect.Config.Env)
	// 其他需要的信息也可以在 imageInspect 中找到
}
