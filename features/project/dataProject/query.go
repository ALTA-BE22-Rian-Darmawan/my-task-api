package dataProject

import (
	"my-task-app/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) project.DataProjectInterface {
	return &projectQuery{
		db: db,
	}
}

// Delete implements project.DataProjectInterface.
func (p *projectQuery) Delete(id uint) error {
	tx := p.db.Delete(&Project{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectById implements project.DataProjectInterface.
func (p *projectQuery) SelectById(id uint) (*project.ProjectEntity, error) {
	var projectGorm Project
	tx := p.db.Preload("Tasks").First(&projectGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var projectcore = project.ProjectEntity{
		ID:          projectGorm.ID,
		UserID:      projectGorm.UserID,
		ProjectName: projectGorm.ProjectName,
		Description: projectGorm.Description,
		Tasks:       projectGorm.Tasks,
		CreatedAt:   projectGorm.CreatedAt,
		UpdatedAt:   projectGorm.UpdatedAt,
	}

	return &projectcore, nil
}

// Insert implements project.DataProjectInterface.
func (p *projectQuery) Insert(project project.ProjectEntity) error {
	//var projectGorm Project
	projectGorm := Project{
		UserID:      project.UserID,
		ProjectName: project.ProjectName,
		Description: project.Description,
	}
	tx := p.db.Create(&projectGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Update implements project.DataProjectInterface.
func (p *projectQuery) Update(id uint, project project.ProjectEntity) error {
	var projectGorm Project
	tx := p.db.First(&projectGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	projectGorm.ProjectName = project.ProjectName
	projectGorm.Description = project.Description

	tx = p.db.Save(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *projectQuery) SelectByUserId(id uint) ([]project.ProjectEntity, error) {
	var allProject []Project // var penampung data yg dibaca dari db
	tx := p.db.Where("user_id =?", id).Find(&allProject)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allProjectCore []project.ProjectEntity
	for _, v := range allProject {
		allProjectCore = append(allProjectCore, project.ProjectEntity{
			ID:          v.ID,
			UserID:      v.UserID,
			ProjectName: v.ProjectName,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return allProjectCore, nil
}
