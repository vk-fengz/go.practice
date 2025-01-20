package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskBluePrintRequest struct {
	SubSystemType        string `json:"subSystemType"`
	UserContext          string `json:"userContext"`
	PlaybookID           string `json:"playbookID"`
	TasksRequireDescribe []byte `json:"tasksRequireDescribe"`
}

func main() {
	str := "YXBpVmVyc2lvbjogYXBwcy5nYWlhLmlvL3YxYWxwaGExCmtpbmQ6IERlc2NyaXB0aW9uCm1ldGFkYXRhOgogIG5hbWU6IGFsaXl1bi1haS10cmFpbi0xCiAgbmFtZXNwYWNlOiBnYWlhLXJlc2VydmVkCiAgbGFiZWxzOgogICAgYXBwcy5nYWlhLmlvL3VzZXIuaWQ6IHpoYW9mZW5nCiAgICBhcHBzLmdhaWEuaW8vbmV0LXBsYW46IGlwCiAgICBoeXBlcm5vZGUuY2x1c3Rlci5wbWwuY29tLmNuL3VzZXItY29udGV4dDogbGJuOTg3NjU0MzIxMC0xCiAgICBoeXBlcm5vZGUuY2x1c3Rlci5wbWwuY29tLmNuL3BsYXlib29rLWlkOiB4eHh4eHh4eHh4eGFsaXl1bi1haS10cmFpbi0xCnNwZWM6CiAgYXBwSUQ6IGFsaXl1bi1haS10cmFpbi0xCiAgd29ya2xvYWRDb21wb25lbnRzOgogICAgLSBjb21wb25lbnROYW1lOiBhbGl5dW4tYWktdHJhaW4tMS1jb20xCiAgICAgIG5hbWVzcGFjZTogYWkKICAgICAgc2FuZGJveDogcnVuYwogICAgICB3b3JrbG9hZFR5cGU6ICJzdGF0ZWxlc3Mtc3lzdGVtLXNlcnZpY2UiCiAgICAgIG1vZHVsZToKICAgICAgICBzcGVjOgogICAgICAgICAgY29udGFpbmVyczoKICAgICAgICAgICAgLSBuYW1lOiBncHVwb2QxCiAgICAgICAgICAgICAgaW1hZ2U6IHJlZ2lzdHJ5LXZwYy5jbi1oYW5nemhvdS5hbGl5dW5jcy5jb20vbXVjZXAvdHJhaW46bGF0ZXN0CiAgICAgICAgICAgICAgcmVzb3VyY2VzOgogICAgICAgICAgICAgICAgbGltaXRzOgogICAgICAgICAgICAgICAgICBjcHU6IDEyCiAgICAgICAgICAgICAgICAgIG1lbW9yeTogOTJHaQogICAgICAgICAgICAgICAgcmVxdWVzdHM6CiAgICAgICAgICAgICAgICAgIGNwdTogMTIKICAgICAgICAgICAgICAgICAgbWVtb3J5OiA5MkdpCiAgICAgICAgICAgICAgdm9sdW1lTW91bnRzOgogICAgICAgICAgICAgICAgLSBtb3VudFBhdGg6IC9kZXYvc2htCiAgICAgICAgICAgICAgICAgIG5hbWU6IGNhY2hlLXZvbHVtZQogICAgICAgICAgdm9sdW1lczoKICAgICAgICAgICAgLSBlbXB0eURpcjoKICAgICAgICAgICAgICAgIG1lZGl1bTogTWVtb3J5CiAgICAgICAgICAgICAgICBzaXplTGltaXQ6IDQ4R2kKICAgICAgICAgICAgICBuYW1lOiBjYWNoZS12b2x1bWUKICBkZXBsb3ltZW50Q29uZGl0aW9uOgogICAgbWFuZGF0b3J5OgogICAgICAtIHN1YmplY3Q6CiAgICAgICAgICBuYW1lOiBhbGl5dW4tYWktdHJhaW4tMS1jb20xCiAgICAgICAgICB0eXBlOiBjb21wb25lbnQKICAgICAgICBvYmplY3Q6CiAgICAgICAgICBuYW1lOiBnZW8tbG9jYXRpb24KICAgICAgICAgIHR5cGU6IGxhYmVsCiAgICAgICAgcmVsYXRpb246IEluCiAgICAgICAgZXh0ZW50OiBXeUl3TURBd01EQXRNREF3TURBMUxUWXhNREF3TUMwMk1UQXhNREF0TmpFd01URTJJbDA9CiAgICAgIC0gc3ViamVjdDoKICAgICAgICAgIG5hbWU6IGFsaXl1bi1haS10cmFpbi0xLWNvbTEKICAgICAgICAgIHR5cGU6IGNvbXBvbmVudAogICAgICAgIG9iamVjdDoKICAgICAgICAgIG5hbWU6IGdwdVJlcXVlc3QKICAgICAgICAgIHR5cGU6IHJlc291cmNlcwogICAgICAgIHJlbGF0aW9uOiBJbgogICAgICAgIGV4dGVudDoKICAgICAgICAgIFczc2lZMjl1ZEdGcGJtVnlUbUZ0WlNJNkltZHdkWEJ2WkRFaUxDSm5jSFZRY205a2RXTjBJam9pVGxaSlJFbEJMVll4TURBaUxDSm5jSFZEYjNWdWRDSTZNU3dpWjNCMVRXVnRiM0o1SWpvd0xDSnpZMmhsWkhWc1pWUjVjR1VpT2lJaUxDSm5jSFZKYzI5c1lYUmxJanAwY25WbGZWMD0KICBleHBlY3RlZFBlcmZvcm1hbmNlOgogICAgYm91bmRhcmllczoKICAgICAgaW5uZXI6CiAgICAgICAgLSBuYW1lOiBib3VuZGFyeTAKICAgICAgICAgIHN1YmplY3Q6IGFsaXl1bi1haS10cmFpbi0xLWNvbTEKICAgICAgICAgIHR5cGU6IHJlcGxpY2FzCiAgICAgICAgICB2YWx1ZTogTVE9PQ=="
	byteData := []byte(str)

	task := TaskBluePrintRequest{
		SubSystemType:        "generalComputing",
		UserContext:          "f1f71e8b11efae15e9c4fb7402f1ea418cf76d3e",
		PlaybookID:           "gc_public",
		TasksRequireDescribe: byteData,
	}

	// convert struct to json file and print it
	jsonData, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// create a new file
	file, err := os.Create("task.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// write json data to file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data written to task.json")
}
