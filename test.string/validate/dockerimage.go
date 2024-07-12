package main

import (
	"fmt"
	"regexp"
)

// validateDockerImage checks if the given Docker image string is valid.
func validateDockerImage(image string) bool {
	// Docker image regex pattern
	// This pattern aims to cover common formats but might not be exhaustive for all valid Docker image names.
	dockerImagePattern := `^([a-z0-9]+([._-][a-z0-9]+)*(:[a-z0-9]+([-._][a-z0-9]+)*)?\/)?([a-z0-9]+([-._][a-z0-9]+)*)+(:[a-z0-9]+([-._][a-z0-9]+)*)?$`

	// Compile the regex and check if the image name matches the pattern
	matched, err := regexp.MatchString(dockerImagePattern, image)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return false
	}

	return matched
}

// isValidDockerImage 检查给定的字符串是否符合 Docker 镜像地址的格式
func isValidDockerImage(image string) bool {
	// Docker 镜像地址的正则表达式
	// 参考 Docker 官方文档关于镜像命名的部分
	// https://docs.docker.com/engine/reference/commandline/tag/
	re := regexp.MustCompile(`^((?:[a-zA-Z0-9]+(?:[._-][a-zA-Z0-9]+)*)?/)?(?:([a-zA-Z0-9]+(?:[._-][a-zA-Z0-9]+)*)/)*([a-z0-9]+(?:[-._][a-z0-9]+)*)?(?::([a-zA-Z0-9][a-zA-Z0-9_.-]*))?$`)
	return re.MatchString(image)
}
// ^((?:[a-zA-Z0-9]+(?:[._-][a-zA-Z0-9]+)*)?/)?(?:([a-zA-Z0-9]+(?:[._-][a-zA-Z0-9]+)*)/)*([a-z0-9]+(?:[-._][a-z0-9]+)*)?(?::([a-zA-Z0-9][a-zA-Z0-9_.-]*))?$


func isValidDockerImage2(imageURL string) bool {
	// Docker镜像地址的正则表达式匹配规则
	pattern := `^([a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*/)?[a-z0-9]+(?:(?:[._]|__|[-]*)[a-z0-9]+)*(?:(?:[:-])[.]{1}[.]*[.]{1})?(?:/[^/]+)+(?::[A-Za-z0-9_./-]+)?(?:@[A-Fa-f0-9]{64})?$`
	match, err := regexp.MatchString(pattern, imageURL)
	if err != nil {
		fmt.Println("正则表达式匹配出错：", err)
		return false
	}
	return match
}

func main() {
	images := []string{
		"nginx",
		"nginx:latest",
		"my-username/my-image",
		"my-username/my-image:tag",
		"docker.io/library/nginx",
		"docker.io/library/nginx:stable",
		"invalid_image_name",
		"also-invalid:image",
		":tag-only",
		"121.40.102.76:30080/zzac",
		"121.40.102.76:30080/zzac/zzacpython01_7c76ab6645_zzacpython01_demo_20240529183743",
		"121.40.102.76:30080/zzacpython01_7c76ab6645_zzacpython01_demo_20240529183743",
		"121.40.102.76:30080/zzac/zzacpython01_7c76ab6645_zzacpython01_demo_20240529183743:SecureContainer",
	}

	for _, image := range images {
		fmt.Printf("%s: %v\n", image, isValidDockerImage2(image))

		//fmt.Printf("%s: %s\n", image, map[bool]string{true: "Valid", false: "Invalid"}[validateDockerImage(image)])
	}
}
