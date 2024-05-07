package handler

import (
	"errors"
	"my-task-app/app/middlewares"
	"my-task-app/features/project"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ServiceProjectInterface
}

func New(pr project.ServiceProjectInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: pr,
	}
}

func (ph *ProjectHandler) CreateProject(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	// membaca data dari request body
	newProject := ProjectRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind project: " + errBind.Error(),
		})
	}

	// mapping  dari request ke project
	inputProject := project.ProjectEntity{
		UserID:      uint(userID),
		ProjectName: newProject.ProjectName,
		Description: newProject.Description,
	}
	// memanggil/mengirimkan data ke method service layer
	errInsert := ph.projectService.Create(inputProject)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error insert project :" + errInsert.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success add project",
	})
}

func (ph *ProjectHandler) GetAllProject(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	result, err := ph.projectService.GetByUserId(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": "error read project",
		})
	}
	var allProjectsResponse []ProjectResponse
	for _, value := range result {
		allProjectsResponse = append(allProjectsResponse, ProjectResponse{
			UserID:      value.UserID,
			ID:          value.ID,
			ProjectName: value.ProjectName,
			Description: value.Description,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read project",
		"results": allProjectsResponse,
	})
}

func (ph *ProjectHandler) GetByIdProject(c echo.Context) error {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error converting id: " + err.Error(),
		})
	}

	project, err := ph.projectService.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "failed",
			"message": "error read project: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "successfully retrieved project",
		"data":    project,
	})
}

func (ph *ProjectHandler) DeleteProject(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}
	err := ph.projectService.Delete(uint(idConv), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete project " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete project",
	})
}

func (ph *ProjectHandler) UpdateProject(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error converting id: " + errConv.Error(),
		})
	}

	var updateData ProjectRequest
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error binding project: " + err.Error(),
		})
	}

	inputProject := project.ProjectEntity{
		ProjectName: updateData.ProjectName,
		Description: updateData.Description,
	}

	if err := ph.projectService.Update(uint(idConv), uint(userID), inputProject); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "error updating project: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "successfully updated project",
	})
}
