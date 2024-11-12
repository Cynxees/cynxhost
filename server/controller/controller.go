package controller

import (
	"context"
	"mchost-spot-instance/server/api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
  pb "mchost-spot-instance/server/pb"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupHandlers(router *gin.Engine, server *api.Server) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.LoadHTMLFiles("index.html")

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/api/launch", func(c *gin.Context) {
		launchSpotFleetHandler(c, server)
	})

	router.POST("/api/get", func(c *gin.Context) {
		getSpotFleetHandler(c, server)
	})

	router.POST("/api/create", func(c *gin.Context) {
		createSpotFleetHandler(c, server)
	})

	router.POST("/api/stop", func(c *gin.Context) {
		stopSpotFleetHandler(c, server)
	})

	router.GET("/api/ping", Helloworld)
	router.POST("/api/pong", Pong)
}

func Pong(g *gin.Context) {

	currentTime := time.Now()
	_, err := g.GetRawData()
	time.Sleep(1 * time.Second)
	duration := time.Since(currentTime)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"time": duration.Seconds()})
}


// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping [get]
func Helloworld(g *gin.Context) {
	time.Sleep(1 * time.Second)
	g.JSON(http.StatusOK, "helloworld")
}

// LaunchSpotFleetTemplate godoc
// @Summary Launch Spot Fleet Instances
// @Tags Spot
// @Accept json
// @Produce json
// @Param			requestBody	body		pb.LaunchTemplateRequest	true	"Request Body"
// @Success 200 {string} Helloworld
// @Router /launch [post]
func launchSpotFleetHandler(c *gin.Context, server *api.Server) {
	ctx := context.Background()
	
	var request pb.LaunchTemplateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	
	resp, err := server.LaunchSpotFleet(ctx, &request);
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// CreateSpotFleetTemplate godoc
// @Summary Create Spot Fleet Template
// @Tags Spot
// @Accept json
// @Produce json
// @Param			requestBody	body		pb.CreateTemplateRequest	true	"Request Body"
// @Success 200 {string} Helloworld
// @Router /create [post]
func createSpotFleetHandler(c *gin.Context, server *api.Server) {
	ctx := context.Background()

	var request pb.CreateTemplateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := server.CreateTemplate(ctx, &request);
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// StopSpotFleetTemplate godoc
// @Summary Stop Spot Fleet Instances
// @Tags Spot
// @Accept json
// @Produce json
// @Param			requestBody	body		pb.StopTemplateRequest	true	"Request Body"
// @Success 200 {string} Helloworld
// @Router /stop [post]
func stopSpotFleetHandler(c *gin.Context, server *api.Server) {
	ctx := context.Background()

	var request pb.StopTemplateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := server.StopTemplate(ctx, &request);
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// GetSpotFleetTemplate godoc
// @Summary Get Spot Fleet Instances
// @Tags Spot
// @Accept json
// @Produce json
// @Param			requestBody	body		pb.GetTemplateRequest	true	"Request Body"
// @Success 200 {string} Helloworld
// @Router /get [post]
func getSpotFleetHandler(c *gin.Context, server *api.Server) {
	ctx := context.Background()

	var request pb.GetTemplateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := server.GetTemplate(ctx, &request);
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}