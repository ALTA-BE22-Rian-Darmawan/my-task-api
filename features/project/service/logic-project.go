package service

import (
	"errors"
	"my-task-app/features/project"
)

type projectService struct {
	projectData project.DataProjectInterface
}

func New(pr project.DataProjectInterface) project.ServiceProjectInterface {
	return &projectService{
		projectData: pr,
	}

}

// Create implements project.ServiceProjectInterface.
func (pr *projectService) Create(project project.ProjectEntity) error {
	if project.ProjectName == "" {
		return errors.New("project name cannot be empty")
	}
	err := pr.projectData.Insert(project)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements project.ServiceProjectInterface.
func (pr *projectService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New("invalid project ID")
	}
	cekuserid, err := pr.projectData.SelectById(id)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id not match, cannot delete project")
	}

	return pr.projectData.Delete(id)
}

// GetById implements project.ServiceProjectInterface.
func (pr *projectService) GetById(id uint) (project *project.ProjectEntity, err error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}
	return pr.projectData.SelectById(id)
}

// Update implements project.ServiceProjectInterface.
func (pr *projectService) Update(id uint, userid uint, project project.ProjectEntity) error {
	if id == 0 {
		return errors.New("invalid project ID")
	}
	if project.ProjectName == "" {
		return errors.New("project name cannot be empty")
	}

	cekuserid, err := pr.projectData.SelectById(id)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id not match, cannot update project")
	}

	return pr.projectData.Update(id, project)
}

// GetByUserId implements project.ServiceProjectInterface.
func (pr *projectService) GetByUserId(id uint) ([]project.ProjectEntity, error) {
	return pr.projectData.SelectByUserId(id)
}
