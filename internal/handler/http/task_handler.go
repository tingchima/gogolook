// Package http provides
package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tingchima/gogolook/internal/application"
	"github.com/tingchima/gogolook/internal/domain"
	"github.com/tingchima/gogolook/internal/domain/common"
	"gopkg.in/guregu/null.v4"
)

// TaskResponse .
type TaskResponse struct {
	// 任務ID
	ID int64 `json:"id"`
	// 任務名稱
	Name string `json:"name"`
	// 任務狀態
	Status bool `json:"status"`
}

// @Summary 取得任務列表
// @Router /tasks [GET]
// @Produce json
// @Tags Task
// @Success 200 {object} List{data=[]http.TaskResponse} "任務列表"
// @Failure 400 {object} ErrResponse "{"code":"400400","message":"Wrong parameter format or invalid"}" "參數錯誤"
// @Failure 500 {object} ErrResponse "{"code":"500000","message":"Internal server error"}" "伺服器內部錯誤"
func ListTasks(app *application.Application) func(c *gin.Context) {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		tasks, err := app.TaskService.ListTasks(ctx, domain.TaskParam{})
		if err != nil {
			responseWithError(c, err)
			return
		}

		response := make([]TaskResponse, len(tasks))

		for i := range tasks {
			response[i] = TaskResponse{
				ID:     tasks[i].ID,
				Name:   tasks[i].Name,
				Status: tasks[i].Status,
			}
		}

		responseWithJSON(c, http.StatusOK, response)
	}
}

// @Summary 建立任務
// @Router /task [POST]
// @Produce json
// @Tags Task
// @Success 201 {object} http.TaskResponse "任務內容"
// @Failure 400 {object} ErrResponse "{"code":"400400","message":"Wrong parameter format or invalid"}" "參數錯誤"
// @Failure 500 {object} ErrResponse "{"code":"500000","message":"Internal server error"}" "伺服器內部錯誤"
func CreateTask(app *application.Application) func(c *gin.Context) {

	// Request .
	type Request struct {
		// 任務名稱
		Name string `form:"name" json:"name" binding:"required"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var req Request
		err := c.ShouldBind(&req)
		if err != nil {
			responseWithError(c, common.NewError(common.ErrCodeInvalidParameter, err, common.WithMsg(err.Error())))
			return
		}

		createdTask, err := app.TaskService.CreateTask(ctx, domain.Task{Name: req.Name})
		if err != nil {
			fmt.Println(err.Error())
			responseWithError(c, err)
			return
		}

		response := TaskResponse{
			ID:     createdTask.ID,
			Name:   createdTask.Name,
			Status: createdTask.Status,
		}

		responseWithJSON(c, http.StatusOK, response)
	}
}

// @Summary 聊天室訊息輸入中
// @Router /task/:id [PUT]
// @Produce json
// @Tags Task
// @Param id path int true "任務ID"
// @Success 200 {object} http.TaskResponse "任務內容"
// @Failure 400 {object} ErrResponse "{"code":"400400","message":"Wrong parameter format or invalid"}" "參數錯誤"
// @Failure 404 {object} ErrResponse "{"code":"400404","message":"User not found"}" "找不到此資源"
// @Failure 500 {object} ErrResponse "{"code":"500000","message":"Internal server error"}" "伺服器內部錯誤"
func UpdateTask(app *application.Application) func(c *gin.Context) {

	// Request .
	type Request struct {
		// 任務ID
		ID int64 `form:"id" json:"id" binding:"required"`
		// 任務名稱
		Name string `form:"name" json:"name" binding:"required"`
		// 任務狀態
		Status null.Bool `form:"status" json:"status" binding:"required"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var req Request
		err := c.ShouldBind(&req)
		if err != nil {
			responseWithError(c, common.NewError(common.ErrCodeInvalidParameter, err, common.WithMsg(err.Error())))
			return
		}

		taskID, err := GetPathInt(c, "id")
		if err != nil {
			responseWithError(c, err)
			return
		}

		updatedTask, err := app.TaskService.UpdateTask(ctx, domain.Task{
			ID:     int64(taskID),
			Name:   req.Name,
			Status: req.Status.Bool,
		})
		if err != nil {
			responseWithError(c, err)
			return
		}

		response := TaskResponse{
			ID:     updatedTask.ID,
			Name:   updatedTask.Name,
			Status: updatedTask.Status,
		}

		responseWithJSON(c, http.StatusOK, response)
	}
}

// @Summary 聊天室訊息輸入中
// @Router /task/:id [DELETE]
// @Produce json
// @Tags Task
// @Param id path int true "任務ID"
// @Success 200 {string} string "" No Content
// @Failure 400 {object} ErrResponse "{"code":"400400","message":"Wrong parameter format or invalid"}" "參數錯誤"
// @Failure 404 {object} ErrResponse "{"code":"400404","message":"User not found"}" "找不到此資源"
// @Failure 500 {object} ErrResponse "{"code":"500000","message":"Internal server error"}" "伺服器內部錯誤"
func DeleteTask(app *application.Application) func(c *gin.Context) {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		taskID, err := GetPathInt(c, "id")
		if err != nil {
			responseWithError(c, err)
			return
		}

		err = app.TaskService.DeleteTaskByID(ctx, int64(taskID))
		if err != nil {
			responseWithError(c, err)
			return
		}

		responseWithNoContent(c, http.StatusOK)
	}
}
