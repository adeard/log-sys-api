package logging

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.GET("top", handler.GetTopError)
	log.GET("range", handler.GetAllByRange)
	log.GET(":log_id", handler.GetDetail)
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

// @Summary Get Detail Log
// @Description Get Detail Log
// @Accept  json
// @Param log_id path string true " Log Id "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.LogData}
// @Router /api/v1/log/{log_id} [get]
// @Tags Log
func (h *loggingHandler) GetDetail(c *gin.Context) {
	start := time.Now()
	logId := c.Param("log_id")

	logData, err := h.loggingService.GetDetail(logId)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	logDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", logData.CreatedAt)
	logDate := logDateParse.Format("2006-01-02 15:04:05")

	logUpdateDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", logData.UpdatedAt)
	logUpdateDate := logUpdateDateParse.Format("2006-01-02 15:04:05")

	logData.CreatedAt = logDate
	logData.UpdatedAt = logUpdateDate

	c.JSON(http.StatusOK, domain.Response{
		Data:        logData,
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

	for i, logData := range logs {

		logDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", logData.LogDate)
		logDate := logDateParse.Format("2006-01-02")
		if logDate == "1900-01-01" {
			logDate = ""
		}

		logs[i].LogDate = logDate
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        logs,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get Top Error Log
// @Description Get Top Error Log
// @Accept  json
// @Param LogFilterRequest query domain.LogFilterRequest true " LogFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.LogTopErrorData}
// @Router /api/v1/log/top [get]
// @Tags Log
func (h *loggingHandler) GetTopError(c *gin.Context) {
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

	logs, err := h.loggingService.GetTopError(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	for i, logData := range logs {
		logDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", logData.LastCreatedAt)
		logDate := logDateParse.Format("2006-01-02 15:04:05")

		logs[i].LastCreatedAt = logDate
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

		utils.LogInit("error", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	logData, err := h.loggingService.Store(logInput)
	if err != nil {
		utils.LogInit("error", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal([]byte(logInput.Response), &data)
	if err != nil {
		var dataObject map[string]interface{}
		err = json.Unmarshal([]byte(logInput.Response), &dataObject)
		if err == nil {
			prettyJson, err := json.MarshalIndent(dataObject, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			logInput.Response = string(prettyJson)
		}
	} else {
		prettyJson, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		logInput.Response = string(prettyJson)
	}

	utils.LogInit2(logInput.Source, logInput)

	c.JSON(http.StatusOK, domain.Response{
		Data:        logData,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
