package appserver

import (
	"encoding/json"
	"net/http"

	"k8s.io/client-go/rest"

	appsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	"github.com/redhat-developer/app-service/appserver/topology"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HandleTopology returns the handler function for the /status endpoint
func (srv *AppServer) HandleTopology() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var bytes []byte

		openshiftAPIConfig := rest.Config{
			Host:        "https://api.rohit15.devcluster.openshift.com:6443",
			BearerToken: "-",
			TLSClientConfig: rest.TLSClientConfig{
				Insecure: true,
			},
		}
		client, err := appsv1.NewForConfig(&openshiftAPIConfig)
		if err != nil {
			return
		}

		namespace := "d-proj"
		_, err = client.DeploymentConfigs(namespace).List(metav1.ListOptions{})

		response := topology.GetSampleTopology()
		bytes, err = json.Marshal(&response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
