package logging

import (
	"fmt"
	"log-sys-api/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type loggingHandler struct {
	loggingService Service
}

func NewLoggingHandler(v1 *gin.RouterGroup, loggingService Service) {

	handler := &loggingHandler{loggingService}

	log := v1.Group("log")
	log.POST("", handler.Create)
}

// @Summary Create Log
// @Description Create Log
// @Accept  json
// @Param LogRequest body domain.LogRequest true " LogRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.LogData}
// @Router /api/v1/log [post]
// @Tags Log
func (h *loggingHandler) Create(c *gin.Context) {
	start := time.Now()
	logInput := domain.LogRequest{}

	err := c.ShouldBindJSON(&logInput)
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

	log, err := h.loggingService.Store(logInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        log,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
