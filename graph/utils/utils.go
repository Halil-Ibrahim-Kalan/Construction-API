package utils

import (
	"Construction-API/graph/model"

	"github.com/jinzhu/gorm"
)

func ToStaff(id int, r *gorm.DB) *model.Staff {
	staff := model.Staff{}
	err := r.First(&staff, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &staff
}

func ToProject(id int, r *gorm.DB) *model.Project {
	project := model.Project{}
	err := r.First(&project, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &project
}

func ToLocation(id int, r *gorm.DB) *model.Location {
	location := model.Location{}
	err := r.First(&location, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &location
}

func ToDepartment(id int, r *gorm.DB) *model.Department {
	department := model.Department{}
	err := r.First(&department, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &department
}

func MapStaffFromInput(staffIDs []int, r *gorm.DB) []*model.Staff {
	var staff []*model.Staff
	for _, staffID := range staffIDs {
		staffMember := ToStaff(staffID, r)
		staff = append(staff, staffMember)
	}
	return staff
}

func ToTask(input model.TaskInput, r *gorm.DB) model.Task {
	user := ToStaff(input.UserID, r)
	project := ToProject(input.ProjectID, r)
	location := ToLocation(input.LocationID, r)
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
