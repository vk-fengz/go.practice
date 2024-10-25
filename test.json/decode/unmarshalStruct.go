package main

import (
	"encoding/json"
	"fmt"
)

type addr struct {
	Province string `json:"province"` //转化后对应字段的json名为province
	City     string `json:"city"`
}
type stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr addr   `json:"addr"`
}

type gpuResource struct {
	ContainerName string `json:"containerName"` // 容器名称
	GPUProduct    string `json:"gpuProduct"`    // GPU产品型号
	GPUCount      int    `json:"gpuCount"`      // GPU数量
	GPUMemory     int    `json:"gpuMemory"`     // GPU显存大小
	ScheduleType  string `json:"scheduleType"`  // 调度类型：共享或独占  shared/exclusive(默认)
	GpuIsolate    bool   `json:"gpuIsolate"`    // 调度类型为共享时，算力是否隔离：true/false
}

func main() {
	var xm = stu{Name: "xiaoming", Age: 18, Addr: addr{Province: "Hunan", City: "ChangSha"}}
	js, err := json.Marshal(xm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(js))
	var xxm stu
	err = json.Unmarshal(js, &xxm)
	fmt.Println(xxm)
	fmt.Println(xxm.Age)

	var gpu1 = gpuResource{ContainerName: "gpupod1", GPUProduct: "Tesla-P4", GPUCount: 0, GPUMemory: 0, ScheduleType: "", GpuIsolate: false}
	var gpu2 = gpuResource{ContainerName: "gpupod2", GPUProduct: "Tesla-V100", GPUCount: 0, GPUMemory: 0, ScheduleType: "", GpuIsolate: false}
	gpuSli := []gpuResource{gpu1, gpu2}
	gpuJson, err := json.Marshal(gpuSli)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(gpuJson))

}