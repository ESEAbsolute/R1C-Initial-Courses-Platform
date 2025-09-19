package models

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    _ "github.com/lib/pq"
)

type Database struct {
    DB *sql.DB
}

type Student struct {
    ID            int       `json:"id"`
    Email         string    `json:"email"`
    Username      string    `json:"username"`
    CreatedAt     time.Time `json:"created_at"`
}

type Course struct {
    ID                int       `json:"id"`
    CourseCode        string    `json:"course_code"`
    CourseName        string    `json:"course_name"`
    CourseDescription string    `json:"course_description"`
    Credits           int       `json:"credits"`
    Instructor        string    `json:"instructor"`
    Semester          string    `json:"semester"`
    TimeSlot          string    `json:"time_slot"`
    CourseLocation    string    `json:"course_location"`
    CreatedAt         time.Time `json:"created_at"`
}

type StudentCourse struct {
    ID         int       `json:"id"`
    StudentID  int       `json:"student_id"`
    CourseID   int       `json:"course_id"`
    EnrolledAt time.Time `json:"enrolled_at"`
}

type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// 连接数据库
func NewDatabase(config DBConfig) (*Database, error) {
    connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
    
    // 连接数据库并测试
    db, err := sql.Open("postgres", connection)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }
    
    // 设置连接池参数
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    log.Println("数据库连接成功")
    
    database := &Database{DB: db}
    
    return database, nil
}

func (db *Database) Close() error {
    return db.DB.Close()
}