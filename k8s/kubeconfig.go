package main

import (
	"context"
	"fmt"

	gaiaclientset "github.com/lmxia/gaia/pkg/generated/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog/v2"
)

func getBootstrapKubeConfigForParentCluster() (*rest.Config, error) {
	// get bootstrap kubeconfig from token
	clientConfig, err := GenerateKubeConfigFromToken("https://172.24.33.8:6443", "2b1j4b.ys7qfar6y4fszjd9",
		nil, 1)
	if err != nil {
		return nil, fmt.Errorf("error while creating kubeconfig: %v", err)
	}

	return clientConfig, nil
}

// GenerateKubeConfigFromToken composes a kubeconfig from token
func GenerateKubeConfigFromToken(serverURL, token string, caCert []byte, flowRate int) (*rest.Config, error) {
	clientConfig := CreateKubeConfigWithToken(serverURL, token, caCert)
	config, err := clientcmd.NewDefaultClientConfig(*clientConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("error while creating kubeconfig: %v", err)
	}

	return applyDefaultRateLimiter(config, flowRate), nil
}

func applyDefaultRateLimiter(config *rest.Config, flowRate int) *rest.Config {
	if flowRate < 0 {
		flowRate = 1
	}

	// here we magnify the default qps and burst in client-go
	config.QPS = rest.DefaultQPS * float32(flowRate)
	config.Burst = rest.DefaultBurst * flowRate

	return config
}

// CreateKubeConfigWithToken creates a KubeConfig object with access to the API server with a token
func CreateKubeConfigWithToken(serverURL, token string, caCert []byte) *clientcmdapi.Config {
	userName := "test"
	clusterName := "test"
	config := createBasicKubeConfig(serverURL, clusterName, userName, caCert)
	config.AuthInfos[userName] = &clientcmdapi.AuthInfo{
		Token: token,
	}
	return config
}

// createBasicKubeConfig creates a basic, general KubeConfig object that then can be extended
func createBasicKubeConfig(serverURL, clusterName, userName string, caCert []byte) *clientcmdapi.Config {
	// Use the cluster and the username as the context name
	contextName := fmt.Sprintf("%s@%s", userName, clusterName)

	var insecureSkipTLSVerify bool
	if caCert == nil {
		insecureSkipTLSVerify = true
	}

	return &clientcmdapi.Config{
		Clusters: map[string]*clientcmdapi.Cluster{
			clusterName: {
				Server:                   serverURL,
				InsecureSkipTLSVerify:    insecureSkipTLSVerify,
				CertificateAuthorityData: caCert,
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			contextName: {
				Cluster:  clusterName,
				AuthInfo: userName,
			},
		},
		AuthInfos:      map[string]*clientcmdapi.AuthInfo{},
		CurrentContext: contextName,
	}
}

func main() {
	clientConfig, err := getBootstrapKubeConfigForParentCluster()
	if err != nil {
		fmt.Println("errorororor")
	}
	// create ClusterRegistrationRequest
	client := gaiaclientset.NewForConfigOrDie(clientConfig)
	crr, err := client.PlatformV1alpha1().ClusterRegistrationRequests().Get(context.TODO(), "cluster4-d0585737-fa09-4123-94f7-e0588c2d984b", metav1.GetOptions{})
	if err != nil {
		klog.Errorf("failed to get ClusterRegistrationRequest: %v", err)
		return
	} else {
		fmt.Printf("%s", crr.Name)
	}
}
