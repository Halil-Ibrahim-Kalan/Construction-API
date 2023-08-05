package utils

import (
	"github.com/Halil-Ibrahim-Kalan/Construction-API/graph/model"
	"github.com/jinzhu/gorm"
)

func ToStaff(userID int, r *gorm.DB) *model.Staff {
	staff := model.Staff{}
	err := r.First(&staff, "id = ?", userID).Error
	if err != nil {
		return nil
	}
	return &staff
}

func ToProject(userID int, r *gorm.DB) *model.Project {
	project := model.Project{}
	err := r.First(&project, "id = ?", userID).Error
	if err != nil {
		return nil
	}
	return &project
}

func ToLocation(userID int, r *gorm.DB) *model.Location {
	location := model.Location{}
	err := r.First(&location, "id = ?", userID).Error
	if err != nil {
		return nil
	}
	return &location
}

func MapStaffFromInput(staffIDs []*int, r *gorm.DB) []*model.Staff {
	var staff []*model.Staff
	for _, staffID := range staffIDs {
		staffMember := ToStaff(*staffID, r)
		staff = append(staff, staffMember)
	}
	return staff
}

func ToTask(input model.TaskInput, r *gorm.DB) model.Task {
	user := ToStaff(*input.UserID, r)
	project := ToProject(*input.ProjectID, r)
	location := ToLocation(*input.LocationID, r)
	staff := MapStaffFromInput(input.StaffIDs, r)

	task := model.Task{
		Name:        input.Name,
		Description: input.Description,
		Detail:      input.Detail,
		User:        user,
		Status:      input.Status,
		Project:     project,
		Location:    location,
		Staff:       staff,
	}
	return task
}
