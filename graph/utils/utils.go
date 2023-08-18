package utils

import (
	"Construction-API/graph/model"
	"errors"
	"strconv"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
		staffMember, _ := ToStaff(staffID, r)
		staff = append(staff, staffMember)
	}
	return staff
}

func ToTask(data model.TaskData, r *gorm.DB) model.Task {
	user, _ := ToStaff(data.UserID, r)
	project := ToProject(data.ProjectID, r)
	location := ToLocation(data.LocationID, r)
	staff := MapStaffFromInput(arrayToSlice(data.StaffIDs), r)

	task := model.Task{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Detail:      data.Detail,
		User:        user,
		Status:      data.Status,
		Project:     project,
		Location:    location,
		Staff:       staff,
	}
	return task
}

func ToTaskData(input model.TaskInput, r *gorm.DB) model.TaskData {
	var staffIDs []int64

	for _, value := range input.StaffIDs {
		staffIDs = append(staffIDs, int64(value))
	}

	data := model.TaskData{
		Name:        input.Name,
		Description: input.Description,
		Detail:      input.Detail,
		UserID:      input.UserID,
		Status:      input.Status,
		ProjectID:   input.ProjectID,
		LocationID:  input.LocationID,
		StaffIDs:    pq.Int64Array(staffIDs),
	}
	return data
}

func ToStaff(id int, r *gorm.DB) (*model.Staff, error) {
	var data *model.StaffData
	err := r.Set("gorm:auto_preload", true).Where("id = ?", id).First(&data).Error

	if err != nil {
		return nil, err
	}

	var department *model.Department
	err = r.Set("gorm:auto_preload", true).Where("id = ?", data.DepartmentID).First(&department).Error

	if err != nil {
		return nil, err
	}

	staffMember := model.Staff{
		ID:         data.ID,
		Name:       data.Name,
		Department: department,
		Role:       data.Role,
		Password:   data.Password,
		Token:      data.Token,
	}

	return &staffMember, nil
}

func arrayToSlice(array pq.Int64Array) []int {
	var intSlice []int
	for _, value := range array {
		intSlice = append(intSlice, int(value))
	}
	return intSlice
}

func DataToTask(data model.TaskData, r *gorm.DB) model.Task {
	user, _ := ToStaff(data.ID, r)
	project := ToProject(data.ProjectID, r)
	location := ToLocation(data.LocationID, r)
	staff := MapStaffFromInput(arrayToSlice(data.StaffIDs), r)

	task := model.Task{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Detail:      data.Detail,
		User:        user,
		Status:      data.Status,
		Project:     project,
		Location:    location,
		Staff:       staff,
	}
	return task
}

func CheckTask(input model.TaskInput, r *gorm.DB) (bool, error) {
	user, err := ToStaff(input.UserID, r)
	if user == nil || err != nil {
		return false, errors.New("User not found! UserID: " + strconv.Itoa(input.UserID))
	}

	project := ToProject(input.ProjectID, r)
	if project == nil {
		return false, errors.New("Project not found! ProjectID: " + strconv.Itoa(input.ProjectID))
	}

	location := ToLocation(input.LocationID, r)
	if location == nil {
		return false, errors.New("Location not found! LocationID: " + strconv.Itoa(input.LocationID))
	}

	if input.StaffIDs == nil {
		return false, errors.New("StaffIDs not found")
	}

	for _, staffID := range input.StaffIDs {
		staffMember, err := ToStaff(staffID, r)
		if staffMember == nil || err != nil {
			return false, errors.New("StaffID not found! staffID: " + strconv.Itoa(staffID))
		}
	}

	return true, nil
}

func DataToStaff(data model.StaffData, r *gorm.DB) model.Staff {
	var department *model.Department
	err := r.Set("gorm:auto_preload", true).Where("id = ?", data.DepartmentID).First(&department).Error

	if err != nil {
		return model.Staff{}
	}

	staffMember := model.Staff{
		ID:         data.ID,
		Name:       data.Name,
		Department: department,
		Role:       data.Role,
		Password:   data.Password,
		Token:      data.Token,
	}

	return staffMember
}

func CheckUserPass(username, password string, checkPass bool, r *gorm.DB) (bool, error) {
	var data *model.StaffData
	var err error
	if checkPass {
		err = r.Where("name = ?", username).Where("password = ?", password).First(&data).Error
	} else {
		err = r.Where("name = ?", username).First(&data).Error
	}

	if err != gorm.ErrRecordNotFound && err != nil {
		return false, err
	}

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}

	return false, nil
}

func GenerateToken(name string) (string, error) {
	userPassword := []byte(name)

	hashedPassword, err := bcrypt.GenerateFromPassword(userPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
