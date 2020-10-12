// TODO: Write the controlleres for all the services.

package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"./services"
)

var (
	clusterRepository 	repository.clusterRepository = nil
	clusterService 		services.clusterService 	 = services.NewClusterService(clusterRepository)
	clusterController 	controller.clusterController = controller.clusterController(clusterService)
	httpRouter      	router.Router             	 = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is up and running...")
	})

	httpRouter.GET("/api/clusters", clusterController.GetClusters)

	httpRouter.SERVE(port)
}

type namespace struct{}

type cluster struct{}

type node struct{}

type pod struct{}

// Return the namespaces in the clustes.
func (c *namespace) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    switch r.Method {
	case "GET":
		// Get all the K8S namespace
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET all namespaces"}`))
		
	case "POST":
		// Create a new K8S namespace
        w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Create a new namespace"}`))
		
	case "PUT":
		// Edit a namespace with a specific ID {Get the needed data to edit in JSON}
        w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "Edit a namespace"}`))
		
	case "DELETE":
		// Delete a K8S namespace with a given name
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete a namespace"}`))
		
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}

// Describe all the clusters [return the nodes, and their namespaces, and pods]
func (c *cluster) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    switch r.Method {
	case "GET":
		// Get all the K8S cluster
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET all cluster"}`))
		
	case "POST":
		// Create a new K8S cluster
        w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Create a new cluster"}`))
		
	case "PUT":
		// Edit a cluster with a specific ID {Get the needed data to edit in JSON}
        w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "Edit a cluster"}`))
		
	case "DELETE":
		// Delete a K8S cluster with a given name
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete a cluster"}`))
		
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}

// To join a cluster
func (c *node) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    switch r.Method {
	case "GET":
		// Get all the K8S nodes
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET all nodes"}`))
		
	case "POST":
		// Create a new K8S nodes
        w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Join to a desired cluster"}`))
		
	case "PUT":
		// Edit a node with a specific ID {Get the needed data to edit in JSON}
        w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "Edit a node"}`))
		
	case "DELETE":
		// Delete a K8S node with a given name
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete a node"}`))
		
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}

// To get and delete some pods in the k8s cluster
func (c *pod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    switch r.Method {
	case "GET":
		// Get all the K8S pods in all the clusters
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET all pods"}`))
		
	case "DELETE":
		// Delete a K8S node with a given name
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete a node"}`))
		
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}

// func main() {
// 	namespaceObj := &namespace{}
// 	clusterObj := &cluster{}
// 	nodeObj := &node{}
// 	podObj := &pod{}

// 	http.Handle("/namespaces", namespaceObj)
// 	http.Handle("/clusters", clusterObj)
// 	http.Handle("/nodes", nodeObj)
// 	http.Handle("/pods", podObj)

//     log.Fatal(http.ListenAndServe(":8000", nil))
// }