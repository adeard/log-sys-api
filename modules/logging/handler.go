package logging

import (
	"fmt"
	"log-sys-api/domain"
	"log-sys-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type loggingHandler struct {
	loggingService Service
}

func NewLoggingHandler(v1 *gin.RouterGroup, loggingService Service) {

	handler := &loggingHandler{loggingService}

	log := v1.Group("log")
	log.GET("", handler.GetAll)
	log.GET("range", handler.GetAllByRange)
	log.POST("", handler.Create)
}

// @Summary Get All Log
// @Description Get All Log
// @Accept  json
// @Param LogFilterRequest query domain.LogFilterRequest true " LogFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.PagingResponse{data=[]domain.LogData}}
// @Router /api/v1/log [get]
// @Tags Log
func (h *loggingHandler) GetAll(c *gin.Context) {
	start := time.Now()
	var input domain.LogFilterRequest

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	logs, err := h.loggingService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        logs,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get All Log By Range
// @Description Get All Log By Range
// @Accept  json
// @Param LogFilterRequest query domain.LogFilterRequest true " LogFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.LogTotalData}
// @Router /api/v1/log/range [get]
// @Tags Log
func (h *loggingHandler) GetAllByRange(c *gin.Context) {
	start := time.Now()
	var input domain.LogFilterRequest

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	logs, err := h.loggingService.GetTotalByDate(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        logs,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
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

		utils.LogInit(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	log, err := h.loggingService.Store(logInput)
	if err != nil {
		utils.LogInit(err.Error())

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
