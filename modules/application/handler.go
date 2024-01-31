package application

import (
	"fmt"
	"log-sys-api/domain"
	"log-sys-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type applicationHandler struct {
	applicationService Service
}

func NewApplicationHandler(v1 *gin.RouterGroup, applicationService Service) {

	handler := &applicationHandler{applicationService}

	application := v1.Group("application")
	application.POST("", handler.Create)
}

// @Summary Create Application
// @Description Create Application
// @Accept  json
// @Param ApplicationRequest body domain.ApplicationRequest true " ApplicationRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.ApplicationData}
// @Router /api/v1/application [post]
// @Tags Application
func (h *applicationHandler) Create(c *gin.Context) {
	start := time.Now()
	applicationInput := domain.ApplicationRequest{}

	err := c.ShouldBindJSON(&applicationInput)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	application, err := h.applicationService.Store(applicationInput)
	if err != nil {
		utils.LogInit(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        application,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
