package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"Construction-API/graph/model"
	"Construction-API/graph/utils"
	"context"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	task := utils.ToTask(input, r.DB)
	err := r.DB.Create(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id int, input model.TaskInput) (*model.Task, error) {
	task := utils.ToTask(input, r.DB)
	task.ID = id

	err := r.DB.Save(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id int) (bool, error) {
	err := r.DB.Where("id = ?", id).Delete(&model.Task{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, name string) (*model.Project, error) {
	project := model.Project{
		Name: name,
	}

	err := r.DB.Create(&project).Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id int, name string) (*model.Project, error) {
	project := model.Project{
		ID:   id,
		Name: name,
	}

	err := r.DB.Save(&project).Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (bool, error) {
	err := r.DB.Where("id = ?", id).Delete(&model.Project{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, name string) (*model.Location, error) {
	location := model.Location{
		Name: name,
	}

	err := r.DB.Create(&location).Error

	if err != nil {
		return nil, err
	}

	return &location, nil
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, id int, name string) (*model.Location, error) {
	location := model.Location{
		ID:   id,
		Name: name,
	}

	err := r.DB.Save(&location).Error

	if err != nil {
		return nil, err
	}

	return &location, nil
}

// DeleteLocation is the resolver for the deleteLocation field.
func (r *mutationResolver) DeleteLocation(ctx context.Context, id int) (bool, error) {
	err := r.DB.Where("id = ?", id).Delete(&model.Location{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateStaff is the resolver for the createStaff field.
func (r *mutationResolver) CreateStaff(ctx context.Context, input model.StaffInput) (*model.Staff, error) {
	staff := model.Staff{
		Name:         input.Name,
		DepartmentID: input.DepartmentID,
		Role:         input.Role,
	}

	err := r.DB.Create(&staff).Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

// UpdateStaff is the resolver for the updateStaff field.
func (r *mutationResolver) UpdateStaff(ctx context.Context, id int, input model.StaffInput) (*model.Staff, error) {
	staff := model.Staff{
		Name:         input.Name,
		DepartmentID: input.DepartmentID,
		Role:         input.Role,
	}

	err := r.DB.Model(&model.Staff{}).Where("id = ?", id).Updates(staff).Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

// DeleteStaff is the resolver for the deleteStaff field.
func (r *mutationResolver) DeleteStaff(ctx context.Context, id int) (bool, error) {
	err := r.DB.Where("id = ?", id).Delete(&model.Staff{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateDepartment is the resolver for the createDepartment field.
func (r *mutationResolver) CreateDepartment(ctx context.Context, name string) (*model.Department, error) {
	department := model.Department{
		Name: name,
	}

	err := r.DB.Create(&department).Error

	if err != nil {
		return nil, err
	}

	return &department, nil
}

// UpdateDepartment is the resolver for the updateDepartment field.
func (r *mutationResolver) UpdateDepartment(ctx context.Context, id int, name string) (*model.Department, error) {
	department := model.Department{
		ID:   id,
		Name: name,
	}

	err := r.DB.Save(&department).Error

	if err != nil {
		return nil, err
	}

	return &department, nil
}

// DeleteDepartment is the resolver for the deleteDepartment field.
func (r *mutationResolver) DeleteDepartment(ctx context.Context, id int) (bool, error) {
	err := r.DB.Where("id = ?", id).Delete(&model.Department{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.DB.Set("gorm:auto_preload", true).Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id int) (*model.Task, error) {
	var task *model.Task
	err := r.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&task).Error

	if err != nil {
		return nil, err
	}

	return task, nil
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project
	err := r.DB.Set("gorm:auto_preload", true).Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, id int) (*model.Project, error) {
	var project *model.Project
	err := r.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&project).Error

	if err != nil {
		return nil, err
	}

	return project, nil
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	var locations []*model.Location
	err := r.DB.Set("gorm:auto_preload", true).Find(&locations).Error

	if err != nil {
		return nil, err
	}

	return locations, nil
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, id int) (*model.Location, error) {
	var location *model.Location
	err := r.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&location).Error

	if err != nil {
		return nil, err
	}

	return location, nil
}

// Staff is the resolver for the staff field.
func (r *queryResolver) Staff(ctx context.Context) ([]*model.Staff, error) {
	var staff []*model.Staff
	err := r.DB.Set("gorm:auto_preload", true).Find(&staff).Error

	if err != nil {
		return nil, err
	}

	return staff, nil
}

// StaffMember is the resolver for the staffMember field.
func (r *queryResolver) StaffMember(ctx context.Context, id int) (*model.Staff, error) {
	var staffMember *model.Staff
	err := r.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&staffMember).Error

	if err != nil {
		return nil, err
	}

	return staffMember, nil
}

// Departments is the resolver for the departments field.
func (r *queryResolver) Departments(ctx context.Context) ([]*model.Department, error) {
	var departments []*model.Department
	err := r.DB.Set("gorm:auto_preload", true).Find(&departments).Error

	if err != nil {
		return nil, err
	}

	return departments, nil
}

// Department is the resolver for the department field.
func (r *queryResolver) Department(ctx context.Context, id int) (*model.Department, error) {
	var department *model.Department
	err := r.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&department).Error

	if err != nil {
		return nil, err
	}

	return department, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
