package types

// ==================== API响应结构体 ====================
// 严格按照Swagger文档定义，并添加管理功能需要的结构体

// 错误响应结构体
type ErrorResponse struct {
    Error string `json:"error" example:"查询失败"`
}

// 成功响应结构体
type SuccessResponse struct {
    Message string `json:"message" example:"操作成功"`
}

// 学生信息结构体 - API版本
type Student struct {
    ID    int    `json:"id" example:"1"`
    Name  string `json:"name" example:"张三"`
    Email string `json:"email" example:"zhangsan@connect.hku.hk"`
}

// 课程信息结构体 - API版本 (简化版，用于列表显示)
type Course struct {
    ID         int    `json:"id" example:"1"`
    CourseCode string `json:"course_code" example:"COMP1117"`
    CourseName string `json:"course_name" example:"Computer programming"`
}

// 课程详细信息结构体 - 用于获取单个课程的完整信息
type CourseDetail struct {
    ID                int    `json:"id" example:"1"`
    CourseCode        string `json:"course_code" example:"COMP1117"`
    CourseName        string `json:"course_name" example:"Computer programming"`
    CourseDescription string `json:"course_description" example:"学习计算机程序设计基础"`
    Credits           int    `json:"credits" example:"3"`
    Instructor        string `json:"instructor" example:"张教授"`
    Semester          string `json:"semester" example:"2024春"`
    TimeSlot          string `json:"time_slot" example:"周一3-4节, 周三5-6节"`
    CourseLocation    string `json:"course_location" example:"教学楼A101"`
}

// 学生选课信息结构体
type StudentCourse struct {
    CourseID   int    `json:"course_id" example:"1"`
    CourseCode string `json:"course_code" example:"COMP1117"`
    CourseName string `json:"course_name" example:"Computer programming"`
}

// ==================== API响应结构体 ====================

// 课程列表响应
type CoursesResponse struct {
    Courses []Course `json:"courses"`
}

// 课程详情响应
type CourseDetailResponse struct {
    Course CourseDetail `json:"course"`
}

// 学生列表响应
type StudentsResponse struct {
    Students []Student `json:"students"`
}

// 学生选课响应
type StudentCoursesResponse struct {
    Student    Student        `json:"student"`
    Courses    []StudentCourse `json:"courses"`
    TotalCount int            `json:"total_count" example:"2"`
}

// ==================== 请求结构体 ====================

// 添加课程请求
type AddCourseRequest struct {
    CourseCode        string `json:"course_code" binding:"required" example:"COMP1117"`
    CourseName        string `json:"course_name" binding:"required" example:"Computer programming"`
    CourseDescription string `json:"course_description" example:"学习计算机程序设计基础"`
    Credits           int    `json:"credits" example:"3"`
    Instructor        string `json:"instructor" example:"张教授"`
    Semester          string `json:"semester" example:"2024春"`
    TimeSlot          string `json:"time_slot" example:"周一3-4节, 周三5-6节"`
    CourseLocation    string `json:"course_location" example:"教学楼A101"`
}

// 添加课程响应
type AddCourseResponse struct {
    Course  Course `json:"course"`
    Message string `json:"message" example:"课程添加成功"`
}

// 添加学生请求
type AddStudentRequest struct {
    Name  string `json:"name" binding:"required" example:"张三"`
    Email string `json:"email" binding:"required,email" example:"zhangsan@connect.hku.hk"`
}

// 添加学生响应
type AddStudentResponse struct {
    Student Student `json:"student"`
    Message string  `json:"message" example:"学生添加成功"`
}