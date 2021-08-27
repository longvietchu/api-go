package service

import (
	"api-go/configs"
	_ "api-go/docs"
	courseDto "api-go/dto"
	"api-go/entity"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

// GetAll godoc
// @Summary Get courses.
// @Description Get all the courses.
// @Tags courses
// @Produce json
// @Success 200 {array} entity.Course
// @Router /courses [get]
func GetAll(c echo.Context) error {
	limit, limitError := strconv.Atoi(c.QueryParam("limit"))
	page, pageError := strconv.Atoi(c.QueryParam("page"))
	if limitError != nil {
		limit = 10
	}
	if pageError != nil {
		page = 0
	}
	var courses []entity.Course
	var totalRecord int64
	configs.DBConn.Limit(limit).Offset(limit * page).Find(&courses).Limit(-1).Offset(-1).Count(&totalRecord)
	totalPage := math.Ceil(float64(totalRecord / int64(limit)))
	return c.JSON(http.StatusOK, echo.Map{
		"data": courses,
		"meta": echo.Map{
			"page_index":   page,
			"page_size":    limit,
			"total_record": totalRecord,
			"total_page":   totalPage,
		},
	})
}

// GetOne godoc
// @Summary Retrieves course based on given ID.
// @Tags courses
// @Produce json
// @Param id path integer true "Course ID"
// @Success 200 {object} entity.Course
// @Router /courses/{id} [get]
func GetOne(c echo.Context) error {
	courseId, _ := strconv.Atoi(c.Param("id"))
	course := entity.Course{}
	configs.DBConn.First(&course, courseId)
	return c.JSON(http.StatusOK, echo.Map{
		"data": course,
	})
}

// Create godoc
// @Summary Add a course
// @Description Add new course
// @Tags courses
// @Accept  json
// @Produce  json
// @Param course body courseDto.CreateCourseRequest true "Add course"
// @Success 200 {object} courseDto.CreateCourseRequest
// @Router /courses [post]
func Create(c echo.Context) error {
	// Validate courseRequest
	var courseRequest courseDto.CreateCourseRequest
	if err := courseRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	course, _ := entity.NewCourse(courseRequest.Title, courseRequest.Description, courseRequest.Creator, courseRequest.Level, courseRequest.URL, courseRequest.Language, courseRequest.Commitment, courseRequest.Rating)
	configs.DBConn.Create(&course)
	return c.JSON(http.StatusOK, echo.Map{
		"data": courseRequest,
	})
}

// Update godoc
// @Summary Update a course
// @Description Update existing courses
// @Tags courses
// @Accept  json
// @Produce  json
// @Param course body courseDto.CreateCourseRequest true "Update course"
// @Param id path integer true "Course ID"
// @Success 200 {object} courseDto.CreateCourseRequest
// @Router /courses/{id} [put]
func Update(c echo.Context) error {
	// Get model if exist
	course := entity.Course{}
	if err := configs.DBConn.First(&course, c.Param("id")).Error; err != nil {
		return err
	}

	// Validate input
	var courseRequest courseDto.CreateCourseRequest
	if err := courseRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	configs.DBConn.Model(&course).Updates(courseRequest)
	return c.JSON(http.StatusOK, echo.Map{
		"data": course,
	})
}

// Delete godoc
// @Summary Delete a course
// @Description Delete existing courses
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id path integer true "Course ID"
// @Success 200 {object} entity.Course
// @Router /courses/{id} [delete]
func Delete(c echo.Context) error {
	// Get model if exist
	course := entity.Course{}
	if err := configs.DBConn.First(&course, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	configs.DBConn.Delete(&course)
	return c.JSON(http.StatusOK, echo.Map{
		"data": course,
	})
}
