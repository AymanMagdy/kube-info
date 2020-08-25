package controller

import (
	"encoding/json"
	"net/http"
	"../entity"
	"../errors"
	"../service"
)

type controller struct {}

var clusterService service.ClusterService

type clusterController interface {
	GetClusters(response http.ResponseWriter, request *http.Request)
	
}