package models

import (
	"database/sql"
	"fmt"
	"log"
)

// 检查并插入示例数据
func (db *Database) InitializeSampleData() error {
    log.Println("检查数据库是否需要初始化示例数据...")
    
    // 检查是否已有数据
    needsData, err := db.needsSampleData()
    if err != nil {
        return fmt.Errorf("检查数据库状态失败: %w", err)
    }
    
    if !needsData {
        log.Println("数据库中已有数据，跳过示例数据插入")
        return nil
    }
    
    log.Println("数据库为空，开始插入示例数据...")
    
    // 开始事务
    tx, err := db.DB.Begin()
    if err != nil {
        return fmt.Errorf("开始事务失败: %w", err)
    }
    defer tx.Rollback()
    
    // 插入示例数据
    if err := db.insertSampleStudents(tx); err != nil {
        return fmt.Errorf("插入示例学生失败: %w", err)
    }
    
    if err := db.insertSampleCourses(tx); err != nil {
        return fmt.Errorf("插入示例课程失败: %w", err)
    }
    
    if err := db.insertSampleEnrollments(tx); err != nil {
        return fmt.Errorf("插入示例选课记录失败: %w", err)
    }
    
    // 提交事务
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("提交事务失败: %w", err)
    }
    
    log.Println("✅ 示例数据插入成功！")
    log.Println("   - 6名示例学生")
    log.Println("   - 8门示例课程")
    log.Println("   - 多条选课记录")
    log.Println("   💡 您可以随时通过API添加或删除数据")
    
    return nil
}

// 检查是否需要插入示例数据
func (db *Database) needsSampleData() (bool, error) {
    var studentCount, courseCount int
    
    err := db.DB.QueryRow("SELECT COUNT(*) FROM students").Scan(&studentCount)
    if err != nil {
        return false, err
    }
    
    err = db.DB.QueryRow("SELECT COUNT(*) FROM courses").Scan(&courseCount)
    if err != nil {
        return false, err
    }
    
    // 如果学生或课程表为空，则需要插入示例数据
    return studentCount == 0 || courseCount == 0, nil
}

// 插入示例学生
func (db *Database) insertSampleStudents(tx *sql.Tx) error {
    students := []struct {
        email    string
        username string
    }{
        {"zhang.san@connect.hku.hk", "张三"},
        {"li.si@connect.hku.hk", "李四"},
        {"wang.wu@connect.hku.hk", "王五"},
        {"zhao.liu@connect.hku.hk", "赵六"},
        {"qian.qi@connect.hku.hk", "钱七"},
        {"sun.ba@connect.hku.hk", "孙八"},
    }
    
    query := `INSERT INTO students (email, username) VALUES ($1, $2)`
    
    for _, student := range students {
        _, err := tx.Exec(query, student.email, student.username)
        if err != nil {
            return fmt.Errorf("插入学生 %s 失败: %w", student.username, err)
        }
    }
    
    log.Printf("✅ 插入了 %d 名示例学生", len(students))
    return nil
}

// 插入示例课程
func (db *Database) insertSampleCourses(tx *sql.Tx) error {
    courses := []struct {
        courseCode        string
        courseName        string
        courseDescription string
        credits           int
        instructor        string
        semester          string
        timeSlot          string
        courseLocation    string
    }{
        {
            "COMP1117", "Computer Programming",
            "Introduction to computer programming using Python", 3,
            "Prof. Chen", "2024 Spring", "Mon 9:00-12:00", "CYC LT1",
        },
        {
            "COMP2119", "Data Structures and Algorithms",
            "Fundamental data structures and algorithms", 4,
            "Prof. Li", "2024 Spring", "Wed 14:00-17:00", "CYC LT2",
        },
        {
            "COMP3234", "Database Systems",
            "Principles of database design and implementation", 3,
            "Prof. Wang", "2024 Spring", "Fri 10:00-13:00", "CYC LT3",
        },
        {
            "COMP3278", "Web Development",
            "Full-stack web development with modern technologies", 3,
            "Prof. Zhang", "2024 Spring", "Tue 14:00-17:00", "Lab 1",
        },
        {
            "COMP4331", "Machine Learning",
            "Introduction to machine learning algorithms", 4,
            "Prof. Liu", "2024 Spring", "Thu 9:00-12:00", "CYC LT4",
        },
        {
            "COMP3322", "Software Engineering",
            "Software development lifecycle and methodologies", 3,
            "Prof. Zhao", "2024 Spring", "Mon 14:00-17:00", "CYC LT5",
        },
        {
            "COMP3297", "Computer Networks",
            "Network protocols and distributed systems", 3,
            "Prof. Wu", "2024 Spring", "Wed 10:00-13:00", "CYC LT6",
        },
        {
            "MATH1013", "Calculus and Linear Algebra",
            "Mathematical foundations for computer science", 4,
            "Prof. Yang", "2024 Spring", "Fri 9:00-12:00", "Math Building LT1",
        },
    }
    
    query := `
        INSERT INTO courses (course_code, course_name, course_description, credits, 
                           instructor, semester, time_slot, course_location)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `
    
    for _, course := range courses {
        _, err := tx.Exec(query,
            course.courseCode, course.courseName, course.courseDescription,
            course.credits, course.instructor, course.semester,
            course.timeSlot, course.courseLocation)
        if err != nil {
            return fmt.Errorf("插入课程 %s 失败: %w", course.courseName, err)
        }
    }
    
    log.Printf("✅ 插入了 %d 门示例课程", len(courses))
    return nil
}

// 插入示例选课记录
func (db *Database) insertSampleEnrollments(tx *sql.Tx) error {
    // 预定义的选课关系 (student_id, course_id)
    enrollments := []struct {
        studentID int
        courseID  int
    }{
        // 张三 (ID: 1) 选了5门课
        {1, 1}, {1, 2}, {1, 3}, {1, 7}, {1, 8},
        
        // 李四 (ID: 2) 选了4门课
        {2, 1}, {2, 4}, {2, 5}, {2, 8},
        
        // 王五 (ID: 3) 选了6门课
        {3, 2}, {3, 3}, {3, 4}, {3, 6}, {3, 7}, {3, 8},
        
        // 赵六 (ID: 4) 选了3门课
        {4, 1}, {4, 4}, {4, 8},
        
        // 钱七 (ID: 5) 选了5门课
        {5, 2}, {5, 5}, {5, 6}, {5, 7}, {5, 8},
        
        // 孙八 (ID: 6) 选了4门课
        {6, 3}, {6, 5}, {6, 6}, {6, 7},
    }
    
    query := `INSERT INTO student_courses (student_id, course_id) VALUES ($1, $2)`
    
    for _, enrollment := range enrollments {
        _, err := tx.Exec(query, enrollment.studentID, enrollment.courseID)
        if err != nil {
            return fmt.Errorf("插入选课记录 (学生ID:%d, 课程ID:%d) 失败: %w", 
                             enrollment.studentID, enrollment.courseID, err)
        }
    }
    
    log.Printf("✅ 插入了 %d 条示例选课记录", len(enrollments))
    return nil
}

// 清空所有数据 (可选功能，用于重置数据库)
func (db *Database) ClearAllData() error {
    log.Println("⚠️  正在清空所有数据...")
    
    // 开始事务
    tx, err := db.DB.Begin()
    if err != nil {
        return fmt.Errorf("开始事务失败: %w", err)
    }
    defer tx.Rollback()
    
    // 按依赖关系顺序删除数据
    queries := []string{
        "DELETE FROM student_courses",
        "DELETE FROM students",
        "DELETE FROM courses",
    }
    
    for _, query := range queries {
        _, err := tx.Exec(query)
        if err != nil {
            return fmt.Errorf("清空数据失败 (%s): %w", query, err)
        }
    }
    
    // 重置序列
    resetQueries := []string{
        "ALTER SEQUENCE students_id_seq RESTART WITH 1",
        "ALTER SEQUENCE courses_id_seq RESTART WITH 1", 
        "ALTER SEQUENCE student_courses_id_seq RESTART WITH 1",
    }
    
    for _, query := range resetQueries {
        _, err := tx.Exec(query)
        if err != nil {
            return fmt.Errorf("重置序列失败 (%s): %w", query, err)
        }
    }
    
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("提交清空操作失败: %w", err)
    }
    
    log.Println("✅ 所有数据已清空，ID序列已重置")
    return nil
}

// 获取数据库统计信息
func (db *Database) GetDataStats() (map[string]int, error) {
    stats := make(map[string]int)
    
    // 获取各表数据量
    queries := map[string]string{
        "students":         "SELECT COUNT(*) FROM students",
        "courses":          "SELECT COUNT(*) FROM courses", 
        "student_courses":  "SELECT COUNT(*) FROM student_courses",
    }
    
    for name, query := range queries {
        var count int
        err := db.DB.QueryRow(query).Scan(&count)
        if err != nil {
            return nil, fmt.Errorf("查询%s统计失败: %w", name, err)
        }
        stats[name] = count
    }
    
    return stats, nil
}