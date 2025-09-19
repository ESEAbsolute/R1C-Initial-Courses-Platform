package models

import (
    "database/sql"
    "fmt"
)

func (db *Database) GetAllStudents() ([]Student, error) {
    query := `
        SELECT id, email, username, created_at
        FROM students
        ORDER BY username
    `
    
    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query students: %w", err)
    }
    defer rows.Close()
    
    var students []Student
    for rows.Next() {
        var student Student
        err := rows.Scan(&student.ID, &student.Email, &student.Username, &student.CreatedAt)
        if err != nil {
            return nil, fmt.Errorf("failed to scan student: %w", err)
        }
        students = append(students, student)
    }
    
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }
    
    return students, nil
}

func (db *Database) GetStudentByID(studentID int) (*Student, error) {
    query := `
        SELECT id, email, username, created_at
        FROM students
        WHERE id = $1
    `
    
    var student Student
    err := db.DB.QueryRow(query, studentID).Scan(
        &student.ID, &student.Email, &student.Username, &student.CreatedAt)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // 学生不存在
        }
        return nil, fmt.Errorf("failed to get student: %w", err)
    }
    
    return &student, nil
}

func (db *Database) AddStudent(email, username string) (*Student, error) {
    query := `
        INSERT INTO students (email, username)
        VALUES ($1, $2)
        RETURNING id, email, username, created_at
    `
    
    var student Student
    err := db.DB.QueryRow(query, email, username).Scan(
        &student.ID, &student.Email, &student.Username, &student.CreatedAt)
    
    if err != nil {
        return nil, fmt.Errorf("failed to add student: %w", err)
    }
    
    return &student, nil
}

func (db *Database) StudentExists(studentID int) (bool, error) {
    query := `SELECT COUNT(*) > 0 FROM students WHERE id = $1`
    
    var exists bool
    err := db.DB.QueryRow(query, studentID).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("failed to check if student exists: %w", err)
    }
    
    return exists, nil
}