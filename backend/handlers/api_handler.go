package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"course-management/models"
	"course-management/types"

	"github.com/gin-gonic/gin"
)

// API处理器结构体
type APIHandler struct {
    DB *models.Database
}

// 创建新的API处理器
func NewAPIHandler(db *models.Database) *APIHandler {
    return &APIHandler{DB: db}
}

// ==================== 课程相关API ====================

// 获取课程列表
func (h *APIHandler) GetCourses(c *gin.Context) {
    courses, err := h.DB.GetAllCourses()
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "获取课程列表失败",
        })
        return
    }
    
    // 转换为API响应格式
    apiCourses := make([]types.Course, len(courses))
    for i, course := range courses {
        apiCourses[i] = types.Course{
            ID:         course.ID,
            CourseCode: course.CourseCode,
            CourseName: course.CourseName,
        }
    }
    
    c.JSON(http.StatusOK, types.CoursesResponse{
        Courses: apiCourses,
    })
}

// 获取课程详细信息
func (h *APIHandler) GetCourseByID(c *gin.Context) {
    courseID, err := strconv.Atoi(c.Param("id"))
    if err != nil || courseID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的课程ID",
        })
        return
    }
    
    course, err := h.DB.GetCourseByID(courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "查询课程信息失败",
        })
        return
    }
    
    if course == nil {
        c.JSON(http.StatusNotFound, types.ErrorResponse{
            Error: "课程不存在",
        })
        return
    }
    
    // 返回完整课程信息
    apiCourse := types.CourseDetail{
        ID:                course.ID,
        CourseCode:        course.CourseCode,
        CourseName:        course.CourseName,
        CourseDescription: course.CourseDescription,
        Credits:           course.Credits,
        Instructor:        course.Instructor,
        Semester:          course.Semester,
        TimeSlot:          course.TimeSlot,
        CourseLocation:    course.CourseLocation,
    }
    
    c.JSON(http.StatusOK, types.CourseDetailResponse{
        Course: apiCourse,
    })
}

// 搜索课程
func (h *APIHandler) SearchCourses(c *gin.Context) {
    keyword := c.Query("keyword")
    if keyword == "" {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "搜索关键词不能为空",
        })
        return
    }
    
    // 去除多余空格
    keyword = strings.TrimSpace(keyword)
    
    courses, err := h.DB.SearchCourses(keyword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "搜索课程失败",
        })
        return
    }
    
    // 转换为API响应格式
    apiCourses := make([]types.Course, len(courses))
    for i, course := range courses {
        apiCourses[i] = types.Course{
            ID:         course.ID,
            CourseCode: course.CourseCode,
            CourseName: course.CourseName,
        }
    }
    
    c.JSON(http.StatusOK, types.CoursesResponse{
        Courses: apiCourses,
    })
}

// 添加课程 (管理员功能)
func (h *APIHandler) AddCourse(c *gin.Context) {
    var req types.AddCourseRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "请求参数格式错误",
        })
        return
    }
    
    // 数据验证
    if req.CourseCode == "" || req.CourseName == "" {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "课程代码和课程名称不能为空",
        })
        return
    }
    
    course, err := h.DB.AddCourse(
        req.CourseCode, req.CourseName, req.CourseDescription,
        req.Credits, req.Instructor, req.Semester, 
        req.TimeSlot, req.CourseLocation,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "添加课程失败",
        })
        return
    }
    
    // 返回添加的课程信息
    apiCourse := types.Course{
        ID:         course.ID,
        CourseCode: course.CourseCode,
        CourseName: course.CourseName,
    }
    
    c.JSON(http.StatusCreated, types.AddCourseResponse{
        Course:  apiCourse,
        Message: "课程添加成功",
    })
}

// ==================== 学生相关API ====================

// 获取学生列表
func (h *APIHandler) GetStudents(c *gin.Context) {
    students, err := h.DB.GetAllStudents()
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "获取学生列表失败",
        })
        return
    }
    
    // 转换为API响应格式
    apiStudents := make([]types.Student, len(students))
    for i, student := range students {
        apiStudents[i] = types.Student{
            ID:    student.ID,
            Name:  student.Username,
            Email: student.Email,
        }
    }
    
    c.JSON(http.StatusOK, types.StudentsResponse{
        Students: apiStudents,
    })
}

// 添加学生 (管理员功能)
func (h *APIHandler) AddStudent(c *gin.Context) {
    var req types.AddStudentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "请求参数格式错误",
        })
        return
    }
    
    // 数据验证
    if req.Name == "" || req.Email == "" {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "姓名和邮箱不能为空",
        })
        return
    }
    
    student, err := h.DB.AddStudent(req.Email, req.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "添加学生失败",
        })
        return
    }
    
    apiStudent := types.Student{
        ID:    student.ID,
        Name:  student.Username,
        Email: student.Email,
    }
    
    c.JSON(http.StatusCreated, types.AddStudentResponse{
        Student: apiStudent,
        Message: "学生添加成功",
    })
}

// ==================== 选课相关API ====================

// 获取学生选课信息
func (h *APIHandler) GetStudentCourses(c *gin.Context) {
    studentID, err := strconv.Atoi(c.Param("id"))
    if err != nil || studentID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的学生ID",
        })
        return
    }
    
    // 获取学生基本信息
    student, err := h.DB.GetStudentByID(studentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "查询学生信息失败",
        })
        return
    }
    
    if student == nil {
        c.JSON(http.StatusNotFound, types.ErrorResponse{
            Error: "学生不存在",
        })
        return
    }
    
    // 获取学生选课信息
    courses, err := h.DB.GetStudentCourses(studentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "查询学生选课信息失败",
        })
        return
    }
    
    // 转换格式
    apiStudent := types.Student{
        ID:    student.ID,
        Name:  student.Username,
        Email: student.Email,
    }
    
    apiCourses := make([]types.StudentCourse, len(courses))
    for i, course := range courses {
        apiCourses[i] = types.StudentCourse{
            CourseID:   course.ID,
            CourseCode: course.CourseCode,
            CourseName: course.CourseName,
        }
    }
    
    c.JSON(http.StatusOK, types.StudentCoursesResponse{
        Student:    apiStudent,
        Courses:    apiCourses,
        TotalCount: len(apiCourses),
    })
}

// 学生选课
func (h *APIHandler) EnrollStudentInCourse(c *gin.Context) {
    studentID, err := strconv.Atoi(c.Param("studentId"))
    if err != nil || studentID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的学生ID",
        })
        return
    }
    
    courseID, err := strconv.Atoi(c.Param("courseId"))
    if err != nil || courseID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的课程ID",
        })
        return
    }
    
    err = h.DB.EnrollStudentInCourse(studentID, courseID)
    if err != nil {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, types.SuccessResponse{
        Message: "选课成功",
    })
}

// 学生退课
func (h *APIHandler) UnenrollStudentFromCourse(c *gin.Context) {
    studentID, err := strconv.Atoi(c.Param("studentId"))
    if err != nil || studentID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的学生ID",
        })
        return
    }
    
    courseID, err := strconv.Atoi(c.Param("courseId"))
    if err != nil || courseID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的课程ID",
        })
        return
    }
    
    err = h.DB.UnenrollStudentFromCourse(studentID, courseID)
    if err != nil {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, types.SuccessResponse{
        Message: "退课成功",
    })
}

// 批量将学生从指定课程移除 (管理员功能 - 课程deprecated时使用)
func (h *APIHandler) RemoveAllStudentsFromCourse(c *gin.Context) {
    courseID, err := strconv.Atoi(c.Param("courseId"))
    if err != nil || courseID <= 0 {
        c.JSON(http.StatusBadRequest, types.ErrorResponse{
            Error: "无效的课程ID",
        })
        return
    }
    
    // 检查课程是否存在
    exists, err := h.DB.CourseExists(courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "检查课程失败",
        })
        return
    }
    
    if !exists {
        c.JSON(http.StatusNotFound, types.ErrorResponse{
            Error: "课程不存在",
        })
        return
    }
    
    // 清空课程的所有选课记录
    err = h.DB.ClearCourseEnrollments(courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "清空课程选课记录失败",
        })
        return
    }
    
    c.JSON(http.StatusOK, types.SuccessResponse{
        Message: "已成功将所有学生从该课程中移除",
    })
}

// ==================== 路由设置 ====================

// 设置所有路由
func (h *APIHandler) SetupRoutes(r *gin.Engine) {
    // 文档要求的基础API
    r.GET("/courses", h.GetCourses)
    r.GET("/students", h.GetStudents)
    r.GET("/student/:id", h.GetStudentCourses)
    
    // 扩展的管理API
    r.POST("/courses", h.AddCourse)                              // 添加课程
    r.GET("/course/:id", h.GetCourseByID)                       // 获取课程详情
    r.GET("/course/search", h.SearchCourses)                    // 搜索课程
    
    r.POST("/students", h.AddStudent)                            // 添加学生
    
    r.POST("/students/:studentId/courses/:courseId", h.EnrollStudentInCourse)      // 学生选课
    r.DELETE("/students/:studentId/courses/:courseId", h.UnenrollStudentFromCourse) // 学生退课
    
    r.DELETE("/courses/:courseId/students", h.RemoveAllStudentsFromCourse) // 批量移除学生(课程deprecated)
}

// 错误处理中间件
func (h *APIHandler) ErrorHandler() gin.HandlerFunc {
    return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
        c.JSON(http.StatusInternalServerError, types.ErrorResponse{
            Error: "服务器内部错误",
        })
    })
}