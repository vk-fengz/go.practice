// ref: https://github.com/kubernetes/klog/tree/main/examples

package main

import (
	"flag"

	klog "k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "5")
	flag.Parse()

	var a int32 = 3
	var b int32 = 6
	klog.V(3).InfoS("3 cron start updated replicas of deployment", "\"from replicas\"", a, "to replicas", b)
	klog.V(5).InfoS("5 cron start updated replicas of deployment", "from replicas", a, "to replicas", b)

	klog.Flush()
}

// --- out
// I0725 10:04:44.981831  198746 infos.go:16] "3 cron start updated replicas of deployment" "from replicas"=3 to replicas=6
// I0725 10:04:44.981906  198746 infos.go:17] "5 cron start updated replicas of deployment" from replicas=3 to replicas=6
