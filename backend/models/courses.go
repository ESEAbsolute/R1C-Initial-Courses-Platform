package models

import (
    "database/sql"
    "fmt"
)

func (db *Database) GetAllCourses() ([]Course, error) {
    query := `
        SELECT id, course_code, course_name, course_description,
               credits, instructor, semester, time_slot, course_location, created_at
        FROM courses
        ORDER BY course_code, semester
    `
    
    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query courses: %w", err)
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
            return nil, fmt.Errorf("failed to scan course: %w", err)
        }
        courses = append(courses, course)
    }
    
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }
    
    return courses, nil
}

func (db *Database) GetCourseByID(courseID int) (*Course, error) {
    query := `
        SELECT id, course_code, course_name, course_description,
               credits, instructor, semester, time_slot, course_location, created_at
        FROM courses
        WHERE id = $1
    `
    
    var course Course
    err := db.DB.QueryRow(query, courseID).Scan(
        &course.ID, &course.CourseCode, &course.CourseName, &course.CourseDescription,
        &course.Credits, &course.Instructor, &course.Semester, &course.TimeSlot,
        &course.CourseLocation, &course.CreatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // 课程不存在
        }
        return nil, fmt.Errorf("failed to get course: %w", err)
    }
    
    return &course, nil
}

func (db *Database) AddCourse(courseCode, courseName, courseDescription string, 
                             credits int, instructor, semester, timeSlot, courseLocation string) (*Course, error) {
    query := `
        INSERT INTO courses (course_code, course_name, course_description, credits, 
                           instructor, semester, time_slot, course_location)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, course_code, course_name, course_description, credits,
                  instructor, semester, time_slot, course_location, created_at
    `
    
    var course Course
    err := db.DB.QueryRow(query, courseCode, courseName, courseDescription, credits,
                         instructor, semester, timeSlot, courseLocation).Scan(
        &course.ID, &course.CourseCode, &course.CourseName, &course.CourseDescription,
        &course.Credits, &course.Instructor, &course.Semester, &course.TimeSlot,
        &course.CourseLocation, &course.CreatedAt,
    )
    
    if err != nil {
        return nil, fmt.Errorf("failed to add course: %w", err)
    }
    
    return &course, nil
}

func (db *Database) SearchCourses(keyword string) ([]Course, error) {
    query := `
        SELECT id, course_code, course_name, course_description,
               credits, instructor, semester, time_slot, course_location, created_at
        FROM courses
        WHERE course_name ILIKE '%' || $1 || '%'
        OR course_code ILIKE '%' || $1 || '%'
        OR instructor ILIKE '%' || $1 || '%'
        ORDER BY course_code
    `
    
    rows, err := db.DB.Query(query, keyword)
    if err != nil {
        return nil, fmt.Errorf("failed to search courses: %w", err)
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
            return nil, fmt.Errorf("failed to scan course: %w", err)
        }
        courses = append(courses, course)
    }
    
    return courses, nil
}

func (db *Database) CourseExists(courseID int) (bool, error) {
    query := `SELECT COUNT(*) > 0 FROM courses WHERE id = $1`
    
    var exists bool
    err := db.DB.QueryRow(query, courseID).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("failed to check if course exists: %w", err)
    }
    
    return exists, nil
}