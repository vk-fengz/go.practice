// go 实现发送post请求的两种方法
// https://blog.csdn.net/YMY_mine/article/details/98496009
// 可以用 apifox 等直接生成

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func main() {

	// 需要post的json文件
	resultSchemaSet := "{\n    \"appID\": \"desc01\",\n    \"components\": [\n        {\n            \"componentName\": \"component1\",\n            \"appType\": \"serverless\"\n        },\n        {\n            \"componentName\": \"component2\",\n            \"appType\": \"deployment\",\n            \"replics\": 5\n        }\n    ],\n    \"spec\": [\n        {\n            \"components\": [\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 1\n                    },\n                    \"clusterName\": \"gaia-cluster1-dnglj\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 1\n                    },\n                    \"clusterName\": \"gaia-cluster2-qjvw7\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster3-6lxx4\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 3\n                    },\n                    \"clusterName\": \"gaia-cluster4-gl4j4\"\n                }\n            ]\n        },\n        {\n            \"components\": [\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster1-dnglj\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 2\n                    },\n                    \"clusterName\": \"gaia-cluster2-qjvw7\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster3-6lxx4\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 3\n                    },\n                    \"clusterName\": \"gaia-cluster4-gl4j4\"\n                }\n            ]\n        },\n        {\n            \"components\": [\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster1-dnglj\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 5\n                    },\n                    \"clusterName\": \"gaia-cluster2-qjvw7\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster3-6lxx4\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster4-gl4j4\"\n                }\n            ]\n        },\n        {\n            \"components\": [\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 2\n                    },\n                    \"clusterName\": \"gaia-cluster1-dnglj\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 3\n                    },\n                    \"clusterName\": \"gaia-cluster2-qjvw7\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster3-6lxx4\"\n                },\n                {\n                    \"replicas\": {\n                        \"component1\": 1,\n                        \"component2\": 0\n                    },\n                    \"clusterName\": \"gaia-cluster4-gl4j4\"\n                }\n            ]\n        }\n    ]\n}"

	postBody, err := json.Marshal(resultSchemaSet)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := http.Post("http://127.0.0.1:8080",
		"application/json;charset=utf-8", bytes.NewReader(postBody))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
	str := (*string)(unsafe.Pointer(&content)) // 转化为string,优化内存
	fmt.Println(*str)
}
