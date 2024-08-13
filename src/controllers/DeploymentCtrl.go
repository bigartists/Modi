package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	corev1 "k8s.io/api/core/v1"
	"log"
	"modi/client"
	"modi/src/result"
	"modi/src/service"
	"net/http"
)

type DeploymentController struct {
}

func NewDeploymentHandler() *DeploymentController {
	return &DeploymentController{}
}

func (this *DeploymentController) Build(r *gin.RouterGroup) {
	r.GET("/deployments", deploymentList) // /modi/v1/deployments?ns=infra
	r.GET("/deployment", deploymentDetail)
	r.POST("/deployment/update/scale", incrReplicas)
	r.GET("/pod/json", podJson)
	r.GET("/pods", pods)
	r.GET("/pod", podDetail)
	r.DELETE("/pod", deletePod)
	r.GET("/ns", namespaces)
	r.GET("/pod/logs", podLogs)
	r.POST("/pod/log/stream", streamLogs)
	r.GET("/pod/containers", GetPodContainer)
}

func namespaces(c *gin.Context) {
	//ResultWrapper(c)(service.DeploymentServiceGetter.GetNs().Unwrap(), "")(OK)

	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetNs().Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func streamLogs(c *gin.Context) {

	params := &struct {
		Namespace string `json:"ns"  required:"true"`
		Pod       string `json:"pname" required:"true"`
		Container string `json:"cname" required:"true""`
	}{}

	log.Println("params:", params, params.Namespace, params.Pod, params.Container)
	result.Result(c.ShouldBindJSON(params)).Unwrap()

	req := client.K8sClient.CoreV1().Pods(params.Namespace).GetLogs(params.Pod, &corev1.PodLogOptions{
		Follow:    true,
		Container: params.Container,
	})

	reader, err := req.Stream(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer reader.Close()

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.(http.Flusher).Flush()

	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Stream finished, sending finish event")
				c.Writer.Write([]byte("event: finish\n"))
				c.Writer.Write([]byte("data: Stream finished\n\n"))
				c.Writer.(http.Flusher).Flush()
				break
			}
			log.Println("Error reading from stream:", err)
			break
		}
		if n > 0 {
			log.Println("Sending log event")
			c.Writer.Write([]byte("event: log\n"))
			c.Writer.Write([]byte("data: " + string(buf[:n]) + "\n\n"))
			//c.Writer.(http.Flusher).Flush()
		}
	}

	// 下面的代码没执行；
	log.Println("Exiting loop, sending finish event")
	c.Writer.Write([]byte("event: finish\n"))
	c.Writer.Write([]byte("data: Stream finished\n\n"))
	c.Writer.(http.Flusher).Flush()

}

func podLogs(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pname" binding:"required"`
		Container string `form:"cname" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetPodLogs(c, req.Namespace, req.Pod, req.Container).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func GetPodContainer(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pname" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetPodContainer(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func deletePod(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pod" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.DeletePod(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func podJson(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pod" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	//ResultWrapper(c)(service.DeploymentServiceGetter.GetPodJson(req.Namespace, req.Pod).Unwrap(), "")(OK)
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetPodJson(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func deploymentList(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetDeploymentsByNs(namespace.Namespace).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func pods(c *gin.Context) {
	req := &struct {
		Namespace  string `form:"ns"`
		Deployment string `form:"deployment"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetPods(req.Namespace, req.Deployment).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func podDetail(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pod" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetPodDetail(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func deploymentDetail(c *gin.Context) {
	req := &struct {
		Namespace  string `form:"ns" binding:"required"`
		Deployment string `form:"deployment" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.GetDeploymentDetailByNsDName(req.Namespace, req.Deployment).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func incrReplicas(c *gin.Context) {
	req := &struct {
		Namespace  string `json:"ns" binding:"required,min=1"`
		Deployment string `json:"deployment" binding:"required,min=1"`
		Dec        bool   `json:"dec"` //是否减少一个副本
	}{}
	result.Result(c.ShouldBindJSON(req)).Unwrap()
	ret := ResultWrapper(c)(service.DeploymentServiceGetter.IncrReplicas(req.Namespace, req.Deployment, req.Dec), "")(OK)
	c.JSON(http.StatusOK, ret)
}
