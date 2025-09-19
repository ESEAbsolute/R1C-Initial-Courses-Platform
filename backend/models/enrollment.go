package models

import (
    "fmt"
)

func (db *Database) GetStudentCourses(studentID int) ([]Course, error) {
    query := `
        SELECT c.id, c.course_code, c.course_name, c.course_description,
               c.credits, c.instructor, c.semester, c.time_slot, c.course_location, c.created_at
        FROM courses c
        JOIN student_courses sc ON c.id = sc.course_id
        WHERE sc.student_id = $1
        ORDER BY c.course_code, c.semester
    `
    
    rows, err := db.DB.Query(query, studentID)
    if err != nil {
        return nil, fmt.Errorf("failed to query student courses: %w", err)
    }
    defer rows.Close()
    
    var courses []Course
    for rows.Next() {
        var course Course
        err := rows.Scan(
            &course.ID, &course.CourseCode, &course.CourseName, &course.CourseDescription,
            &course.Credits, &course.Instructor, &course.Semester, &course.TimeSlot,
            &course.CourseLocation, &course.CreatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan student course: %w", err)
        }
        courses = append(courses, course)
    }
    
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }
    
    return courses, nil
}

func (db *Database) EnrollStudentInCourse(studentID, courseID int) error {
    // 首先检查学生和课程是否存在
    studentExists, err := db.StudentExists(studentID)
    if err != nil {
        return fmt.Errorf("failed to check student existence: %w", err)
    }
    if !studentExists {
        return fmt.Errorf("student with ID %d does not exist", studentID)
    }
    
    courseExists, err := db.CourseExists(courseID)
    if err != nil {
        return fmt.Errorf("failed to check course existence: %w", err)
    }
    if !courseExists {
        return fmt.Errorf("course with ID %d does not exist", courseID)
    }
    
    // 检查是否已经选过这门课
    enrolled, err := db.isStudentEnrolled(studentID, courseID)
    if err != nil {
        return fmt.Errorf("failed to check enrollment status: %w", err)
    }
    if enrolled {
        return fmt.Errorf("student is already enrolled in this course")
    }
    
    // 执行选课
    query := `
        INSERT INTO student_courses (student_id, course_id)
        VALUES ($1, $2)
    `
    
    _, err = db.DB.Exec(query, studentID, courseID)
    if err != nil {
        return fmt.Errorf("failed to enroll student in course: %w", err)
    }
    
    return nil
}

func (db *Database) UnenrollStudentFromCourse(studentID, courseID int) error {
    query := `
        DELETE FROM student_courses
        WHERE student_id = $1 AND course_id = $2
    `
    
    result, err := db.DB.Exec(query, studentID, courseID)
    if err != nil {
        return fmt.Errorf("failed to unenroll student from course: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("student is not enrolled in this course")
    }
    
    return nil
}

func (db *Database) ClearCourseEnrollments(courseID int) error {
    query := `DELETE FROM student_courses WHERE course_id = $1`
    
    _, err := db.DB.Exec(query, courseID)
    if err != nil {
        return fmt.Errorf("failed to clear course enrollments: %w", err)
    }
    
    return nil
}

// 私有辅助方法，检查学生是否已选课
func (db *Database) isStudentEnrolled(studentID, courseID int) (bool, error) {
    query := `
        SELECT COUNT(*) > 0
        FROM student_courses
        WHERE student_id = $1 AND course_id = $2
    `
    
    var enrolled bool
    err := db.DB.QueryRow(query, studentID, courseID).Scan(&enrolled)
    if err != nil {
        return false, fmt.Errorf("failed to check enrollment: %w", err)
    }
    
    return enrolled, nil
}