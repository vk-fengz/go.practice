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
	GPUProduct   string `json:"gpuProduct"`
	GPUCount     int    `json:"gpuCount"`
	GPUMemory    int    `json:"gpuMemory"`
	ScheduleType string `json:"scheduleType"`
	GpuIsolate   bool   `json:"gpuIsolate"`
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

	var gpu1 = gpuResource{GPUProduct: "Tesla-P4", GPUCount: 0, GPUMemory: 0, ScheduleType: "", GpuIsolate: false}
	var gpu2 = gpuResource{GPUProduct: "Tesla-V100", GPUCount: 0, GPUMemory: 0, ScheduleType: "", GpuIsolate: false}
	gpuSli := []gpuResource{gpu1, gpu2}
	gpuJson, err := json.Marshal(gpuSli)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(gpuJson))

}
