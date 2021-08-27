package repository

import (
	"api-go/entity"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(user *entity.Course) (*entity.Course, error)
	GetByID(id int) (*entity.Course, error)
	GetAll(limit, page int) (*entity.Courses, error)
	Update(user *entity.Course) (*entity.Course, error)
	Delete(user *entity.Course) error
}

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func (courseRepositoryImpl CourseRepositoryImpl) GetAll(limit, page int) (*entity.Courses, error) {
	courses := entity.Courses{}
	courseRepositoryImpl.db.Limit(limit).Offset(limit * page).Find(&courses)
	coursesEntiy := new(entity.Courses)
	err := copier.Copy(&coursesEntiy, &courses)
	if err != nil {
		return nil, err
	}

	return coursesEntiy, nil
}
