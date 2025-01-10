package controllers

import (
	"github.com/bigartists/Modi/client"
	"github.com/bigartists/Modi/src/handler"
	"github.com/bigartists/Modi/src/result"
	"github.com/bigartists/Modi/src/service"
	"github.com/gin-gonic/gin"
	"io"
	corev1 "k8s.io/api/core/v1"
	"log"
	"net/http"
)

type DeploymentController struct {
	deploymentService *service.DeploymentService
}

func ProviderDeploymentController(deploymentService *service.DeploymentService) *DeploymentController {
	return &DeploymentController{deploymentService: deploymentService}
}

func (this *DeploymentController) Build(r *gin.RouterGroup) {
	//r.GET("/deployments", deploymentList) // /modi/v1/deployments?ns=infra
	r.GET("/deployments", this.deploymentList2) // /modi/v1/deployments?ns=infra
	r.GET("/deployment", this.deploymentDetail)
	r.POST("/deployment/update/scale", this.incrReplicas)
	r.GET("/pod/json", this.podJson)
	r.GET("/pods", this.pods)
	r.GET("/pod", this.podDetail)
	r.DELETE("/pod", this.deletePod)
	r.GET("/ns", this.namespaces)
	r.GET("/pod/logs", this.podLogs)
	r.POST("/pod/log/stream", this.streamLogs)
	r.GET("/pod/containers", this.GetPodContainer)
}

func (this *DeploymentController) namespaces(c *gin.Context) {
	//ResultWrapper(c)(service.DeploymentServiceGetter.GetNs().Unwrap(), "")(OK)

	ret := ResultWrapper(c)(this.deploymentService.GetNs().Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) streamLogs(c *gin.Context) {

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

func (this *DeploymentController) podLogs(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pname" binding:"required"`
		Container string `form:"cname" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.GetPodLogs(c, req.Namespace, req.Pod, req.Container).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) GetPodContainer(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pname" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.GetPodContainer(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) deletePod(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pod" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.DeletePod(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) podJson(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pod" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	//ResultWrapper(c)(service.DeploymentServiceGetter.GetPodJson(req.Namespace, req.Pod).Unwrap(), "")(OK)
	ret := ResultWrapper(c)(this.deploymentService.GetPodJson(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) deploymentList(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	result.Result(c.ShouldBindQuery(namespace)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.GetDeploymentsByNs(namespace.Namespace).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) deploymentList2(c *gin.Context) {
	namespace := &struct {
		Namespace string `form:"ns"`
	}{}
	if handler.BindAndCheck(c, namespace) {
		return
	}
	ret, err := this.deploymentService.GetDeploymentsByNs2(namespace.Namespace)
	handler.HandleResponse(c, err, ret)
}

func (this *DeploymentController) pods(c *gin.Context) {
	req := &struct {
		Namespace  string `form:"ns"`
		Deployment string `form:"deployment"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.GetPods(req.Namespace, req.Deployment).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) podDetail(c *gin.Context) {
	req := &struct {
		Namespace string `form:"ns" binding:"required"`
		Pod       string `form:"pod" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.GetPodDetail(req.Namespace, req.Pod).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) deploymentDetail(c *gin.Context) {
	req := &struct {
		Namespace  string `form:"ns" binding:"required"`
		Deployment string `form:"deployment" binding:"required"`
	}{}
	result.Result(c.ShouldBindQuery(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.GetDeploymentDetailByNsDName(req.Namespace, req.Deployment).Unwrap(), "")(OK)
	c.JSON(http.StatusOK, ret)
}

func (this *DeploymentController) incrReplicas(c *gin.Context) {
	req := &struct {
		Namespace  string `json:"ns" binding:"required,min=1"`
		Deployment string `json:"deployment" binding:"required,min=1"`
		Dec        bool   `json:"dec"` //是否减少一个副本
	}{}
	result.Result(c.ShouldBindJSON(req)).Unwrap()
	ret := ResultWrapper(c)(this.deploymentService.IncrReplicas(req.Namespace, req.Deployment, req.Dec), "")(OK)
	c.JSON(http.StatusOK, ret)
}
