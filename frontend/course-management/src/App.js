import { useState, useEffect } from 'react';
import { Search, User, X, Settings, Eye } from 'lucide-react';

// CSS样式定义
const styles = {
    container: {
        minHeight: '100vh',
        backgroundColor: '#f8fafc',
        fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif'
    },
    fixedHeader: {
        position: 'fixed',
        top: 0,
        left: 0,
        right: 0,
        backgroundColor: 'white',
        borderBottom: '1px solid #e2e8f0',
        boxShadow: '0 1px 3px 0 rgba(0, 0, 0, 0.1)',
        zIndex: 40
    },
    headerContent: {
        maxWidth: '1280px',
        margin: '0 auto',
        padding: '0 1rem',
        height: '4rem',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between'
    },
    logo: {
        fontSize: '1.5rem',
        fontWeight: 'bold',
        color: '#2563eb'
    },
    button: {
        display: 'flex',
        alignItems: 'center',
        gap: '0.5rem',
        padding: '0.5rem 1rem',
        backgroundColor: '#2563eb',
        color: 'white',
        border: 'none',
        borderRadius: '0.375rem',
        cursor: 'pointer',
        fontSize: '0.875rem',
        transition: 'background-color 0.2s'
    },
    buttonHover: {
        backgroundColor: '#1d4ed8'
    },
    mainContent: {
        paddingTop: '4rem'
    },
    searchSection: {
        backgroundColor: 'white',
        borderBottom: '1px solid #e2e8f0',
        padding: '1rem 0'
    },
    searchContainer: {
        maxWidth: '1280px',
        margin: '0 auto',
        padding: '0 1rem',
        display: 'flex',
        gap: '1rem'
    },
    selectContainer: {
        width: '16rem'
    },
    label: {
        display: 'block',
        fontSize: '0.875rem',
        fontWeight: '500',
        marginBottom: '0.5rem',
        color: '#374151'
    },
    select: {
        width: '100%',
        border: '1px solid #d1d5db',
        borderRadius: '0.375rem',
        padding: '0.5rem 0.75rem',
        fontSize: '0.875rem',
        outline: 'none'
    },
    searchInputContainer: {
        flex: 1
    },
    searchInputWrapper: {
        position: 'relative'
    },
    searchInput: {
        width: '94%',
        paddingLeft: '2.5rem',
        paddingRight: '1rem',
        paddingTop: '0.5rem',
        paddingBottom: '0.5rem',
        border: '1px solid #d1d5db',
        borderRadius: '0.375rem',
        fontSize: '0.875rem',
        outline: 'none'
    },
    searchIcon: {
        position: 'absolute',
        left: '0.75rem',
        top: '50%',
        transform: 'translateY(-50%)',
        color: '#9ca3af'
    },
    coursesSection: {
        maxWidth: '1280px',
        margin: '0 auto',
        padding: '1.5rem 1rem'
    },
    coursesGrid: {
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))',
        gap: '1rem'
    },
    courseCard: {
        backgroundColor: 'white',
        borderRadius: '0.5rem',
        border: '1px solid #93c5fd',
        borderColor: '#93c5fd',
        padding: '1rem',
        cursor: 'pointer',
        transition: 'all 0.2s',
        boxShadow: '0 1px 3px 0 rgba(0, 0, 0, 0.1)'
    },
    courseCardHover: {
        borderColor: '#000000',
        boxShadow: '0 4px 6px -1px rgba(0, 0, 0, 0.1)'
    },
    courseCode: {
        fontSize: '0.875rem',
        fontWeight: '500',
        color: '#2563eb',
        marginBottom: '0.5rem'
    },
    courseName: {
        color: '#111827',
        fontWeight: '500',
        lineHeight: '1.25'
    },
    floatingButtons: {
        position: 'fixed',
        right: '1.5rem',
        bottom: '1.5rem',
        display: 'flex',
        flexDirection: 'column',
        gap: '0.75rem'
    },
    floatingButton: {
        padding: '0.75rem',
        borderRadius: '50%',
        border: 'none',
        cursor: 'pointer',
        boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1)',
        transition: 'all 0.2s',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
    },
    floatingButtonGreen: {
        backgroundColor: '#10b981',
        color: 'white'
    },
    floatingButtonOrange: {
        backgroundColor: '#f59e0b',
        color: 'white'
    },
    modal: {
        position: 'fixed',
        inset: 0,
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        zIndex: 50
    },
    modalContent: {
        backgroundColor: 'white',
        borderRadius: '0.5rem',
        padding: '1.5rem',
        width: '24rem',
        maxWidth: '90vw',
        maxHeight: '80vh',
        overflowY: 'auto'
    },
    modalContentLarge: {
        width: '37.5rem'
    },
    modalHeader: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: '1rem'
    },
    modalTitle: {
        fontSize: '1.25rem',
        fontWeight: '600',
        color: '#111827'
    },
    closeButton: {
        padding: '0.25rem',
        border: 'none',
        backgroundColor: 'transparent',
        cursor: 'pointer',
        borderRadius: '0.25rem',
        transition: 'background-color 0.2s'
    },
    closeButtonHover: {
        backgroundColor: '#f3f4f6'
    },
    input: {
        width: '90%',
        border: '1px solid #d1d5db',
        borderRadius: '0.375rem',
        padding: '0.5rem 0.75rem',
        fontSize: '0.875rem',
        outline: 'none',
        transition: 'border-color 0.2s'
    },
    input_long: {
        width: '95%',
        border: '1px solid #d1d5db',
        borderRadius: '0.375rem',
        padding: '0.5rem 0.75rem',
        fontSize: '0.875rem',
        outline: 'none',
        transition: 'border-color 0.2s'
    },
    inputFocus: {
        borderColor: '#2563eb'
    },
    textarea: {
        width: '95%',
        border: '1px solid #d1d5db',
        borderRadius: '0.375rem',
        padding: '0.5rem 0.75rem',
        fontSize: '0.875rem',
        outline: 'none',
        minHeight: '5rem',
        resize: 'vertical'
    },
    formGroup: {
        marginBottom: '1rem'
    },
    grid: {
        display: 'grid',
        gap: '1rem'
    },
    gridCols2: {
        display: 'grid',
        gap: '1rem',
        gridTemplateColumns: 'repeat(2, 1fr)'
    },
    errorAlert: {
        backgroundColor: '#fef2f2',
        border: '1px solid #fecaca',
        color: '#dc2626',
        padding: '0.75rem',
        borderRadius: '0.375rem',
        marginBottom: '1rem'
    },
    loadingText: {
        textAlign: 'center',
        padding: '3rem 0',
        color: '#6b7280'
    },
    emptyText: {
        textAlign: 'center',
        padding: '3rem 0',
        color: '#6b7280'
    },
    courseDetailGrid: {
        display: 'grid',
        gridTemplateColumns: 'repeat(2, 1fr)',
        gap: '1rem',
        fontSize: '0.875rem',
        marginBottom: '1rem'
    },
    courseDescription: {
        marginBottom: '1rem'
    },
    buttonGroup: {
        display: 'flex',
        gap: '0.5rem',
        paddingTop: '1rem'
    },
    buttonPrimary: {
        backgroundColor: '#2563eb',
        color: 'white',
        padding: '0.5rem 1.5rem',
        border: 'none',
        borderRadius: '0.375rem',
        cursor: 'pointer',
        fontSize: '0.875rem',
        transition: 'background-color 0.2s'
    },
    buttonSuccess: {
        backgroundColor: '#10b981',
        color: 'white',
        padding: '0.5rem 1rem',
        border: 'none',
        borderRadius: '0.375rem',
        cursor: 'pointer',
        fontSize: '0.875rem',
        transition: 'background-color 0.2s'
    },
    buttonDanger: {
        backgroundColor: '#ef4444',
        color: 'white',
        padding: '0.5rem 1.5rem',
        border: 'none',
        borderRadius: '0.375rem',
        cursor: 'pointer',
        fontSize: '0.875rem',
        transition: 'background-color 0.2s'
    },
    buttonDisabled: {
        opacity: 0.5,
        cursor: 'not-allowed'
    },
    tabContainer: {
        display: 'flex',
        marginBottom: '1rem'
    },
    tab: {
        padding: '0.5rem 1rem',
        border: 'none',
        cursor: 'pointer',
        backgroundColor: '#e5e7eb',
        fontSize: '0.875rem',
        transition: 'background-color 0.2s'
    },
    tabActive: {
        backgroundColor: '#2563eb',
        color: 'white'
    },
    tabLeft: {
        borderRadius: '0.375rem 0 0 0.375rem'
    },
    tabRight: {
        borderRadius: '0 0.375rem 0.375rem 0'
    },
    listItem: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '0.75rem',
        border: '1px solid #e5e7eb',
        borderRadius: '0.375rem',
        marginBottom: '0.75rem'
    },
    itemInfo: {
        flex: 1
    },
    itemTitle: {
        fontWeight: '500',
        marginBottom: '0.25rem'
    },
    itemSubtitle: {
        fontSize: '0.875rem',
        color: '#6b7280'
    }
};

// 模拟API基础URL
const API_BASE = process.env.REACT_APP_API_BASE || 'http://localhost:8080';

// API调用函数
const api = {
    getCourses: () => fetch(`${API_BASE}/courses`).then(r => r.json()),
    getStudents: () => fetch(`${API_BASE}/students`).then(r => r.json()),
    getCourseById: (id) => fetch(`${API_BASE}/course/${id}`).then(r => r.json()),
    getStudentCourses: (id) => fetch(`${API_BASE}/student/${id}`).then(r => r.json()),
    searchCourses: (keyword) => fetch(`${API_BASE}/course/search?keyword=${keyword}`).then(r => r.json()),
    addStudent: (data) => fetch(`${API_BASE}/students`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    }).then(r => r.json()),
    addCourse: (data) => fetch(`${API_BASE}/courses`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    }).then(r => r.json()),
    enrollCourse: (studentId, courseId) => fetch(`${API_BASE}/students/${studentId}/courses/${courseId}`, {
        method: 'POST'
    }).then(r => r.json()),
    unenrollCourse: (studentId, courseId) => fetch(`${API_BASE}/students/${studentId}/courses/${courseId}`, {
        method: 'DELETE'
    }).then(r => r.json()),
    removeAllStudentsFromCourse: (courseId) => fetch(`${API_BASE}/courses/${courseId}/students`, {
        method: 'DELETE'
    }).then(r => r.json())
};

// 登录/注册弹窗组件
function LoginModal({ isOpen, onClose, onLogin }) {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!name || !email) {
            setError('姓名和邮箱不能为空');
            return;
        }

        setLoading(true);
        setError('');

        try {
            const studentsData = await api.getStudents();
            const students = studentsData.students || [];

            const existingStudent = students.find(s => s.name === name || s.email === email);

            if (existingStudent) {
                if (existingStudent.name === name && existingStudent.email === email) {
                    onLogin(existingStudent);
                    onClose();
                } else {
                    setError('姓名或邮箱已存在但不匹配，请检查输入');
                }
            } else {
                const newStudent = await api.addStudent({ name, email });
                onLogin(newStudent.student);
                onClose();
            }
        } catch (err) {
            setError('操作失败，请稍后重试');
        }
        setLoading(false);
    };

    if (!isOpen) return null;

    return (
        <div style={styles.modal}>
            <div style={styles.modalContent}>
                <div style={styles.modalHeader}>
                    <h2 style={styles.modalTitle}>注册/登录</h2>
                    <button onClick={onClose} style={styles.closeButton}>
                        <X size={20} />
                    </button>
                </div>

                {error && (
                    <div style={styles.errorAlert}>
                        {error}
                    </div>
                )}

                <form onSubmit={handleSubmit}>
                    <div style={styles.formGroup}>
                        <label style={styles.label}>学生姓名</label>
                        <input
                            type="text"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            style={styles.input}
                            placeholder="请输入姓名"
                        />
                    </div>

                    <div style={styles.formGroup}>
                        <label style={styles.label}>学生邮箱</label>
                        <input
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            style={styles.input}
                            placeholder="请输入邮箱"
                        />
                    </div>

                    <button
                        type="submit"
                        disabled={loading}
                        style={{
                            ...styles.button,
                            ...styles.buttonPrimary,
                            width: '97%',
                            justifyContent: 'center',
                            ...(loading ? styles.buttonDisabled : {})
                        }}
                    >
                        {loading ? '处理中...' : '确认'}
                    </button>
                </form>
            </div>
        </div>
    );
}

// 课程详情弹窗组件
function CourseDetailModal({ courseId, isOpen, onClose, currentUser, onEnroll }) {
    const [course, setCourse] = useState(null);
    const [loading, setLoading] = useState(false);
    const [enrolling, setEnrolling] = useState(false);
    const [isEnrolled, setIsEnrolled] = useState(false);
    const [checkingEnrollment, setCheckingEnrollment] = useState(false);

    useEffect(() => {
        if (isOpen && courseId) {
            setLoading(true);
            api.getCourseById(courseId)
                .then(data => setCourse(data.course))
                .catch(err => console.error('获取课程详情失败:', err))
                .finally(() => setLoading(false));
            
            // 检查用户是否已选择该课程
            if (currentUser) {
                setCheckingEnrollment(true);
                api.getStudentCourses(currentUser.id)
                    .then(data => {
                        const studentCourses = data.courses || [];
                        const isCurrentlyEnrolled = studentCourses.some(c => c.course_id === parseInt(courseId));
                        setIsEnrolled(isCurrentlyEnrolled);
                    })
                    .catch(err => {
                        console.error('检查选课状态失败:', err);
                        setIsEnrolled(false);
                    })
                    .finally(() => setCheckingEnrollment(false));
            } else {
                setIsEnrolled(false);
            }
        }
    }, [isOpen, courseId, currentUser]);

    const handleUnenroll = async () => {
        if (!currentUser || !course || !isEnrolled) return;

        setEnrolling(true);
        try {
            await api.unenrollCourse(currentUser.id, course.id);
            setIsEnrolled(false); // 立即更新本地状态
            onEnroll(); // 刷新外部数据
        } catch (err) {
            alert('退课失败，请稍后重试');
        }
        setEnrolling(false);
    };

    const handleEnroll = async () => {
        if (!currentUser || !course || isEnrolled) return;

        setEnrolling(true);
        try {
            await api.enrollCourse(currentUser.id, course.id);
            setIsEnrolled(true); // 立即更新本地状态
            onEnroll();
            // 不关闭弹窗，让用户看到状态变化
            // onClose();
        } catch (err) {
            alert('选课失败，可能已经选择过该课程');
        }
        setEnrolling(false);
    };

    // 按钮样式定义
    const getButtonStyle = () => {
        if (!currentUser) {
            return {
                ...styles.buttonPrimary,
                ...styles.buttonDisabled
            };
        }
        
        if (isEnrolled) {
            return {
                ...styles.buttonPrimary,
                backgroundColor: '#93c5fd', // 淡蓝色
                cursor: 'not-allowed',
                opacity: 0.8
            };
        }
        
        if (enrolling || checkingEnrollment) {
            return {
                ...styles.buttonPrimary,
                ...styles.buttonDisabled
            };
        }
        
        return styles.buttonPrimary;
    };

    const getButtonText = () => {
        if (!currentUser) return '请先登录';
        if (checkingEnrollment) return '检查中...';
        if (isEnrolled) return '已选';
        if (enrolling) return '选课中...';
        return '选课';
    };

    if (!isOpen) return null;

    return (
        <div style={styles.modal}>
            <div style={{ ...styles.modalContent, ...styles.modalContentLarge }}>
                <div style={styles.modalHeader}>
                    <h2 style={styles.modalTitle}>课程详情</h2>
                    <button onClick={onClose} style={styles.closeButton}>
                        <X size={20} />
                    </button>
                </div>

                {loading ? (
                    <div style={styles.loadingText}>加载中...</div>
                ) : course ? (
                    <div>
                        <div style={styles.formGroup}>
                            <h3 style={{ fontSize: '1.125rem', fontWeight: '600', marginBottom: '0.5rem' }}>{course.course_name}</h3>
                            <p style={{ color: '#6b7280' }}>{course.course_code}</p>
                        </div>

                        <div style={styles.courseDetailGrid}>
                            <div><span style={{ fontWeight: '500' }}>学分：</span>{course.credits}</div>
                            <div><span style={{ fontWeight: '500' }}>教师：</span>{course.instructor}</div>
                            <div><span style={{ fontWeight: '500' }}>学期：</span>{course.semester}</div>
                            <div><span style={{ fontWeight: '500' }}>时间：</span>{course.time_slot}</div>
                            <div style={{ gridColumn: 'span 2' }}><span style={{ fontWeight: '500' }}>地点：</span>{course.course_location}</div>
                        </div>

                        <div style={styles.courseDescription}>
                            <h4 style={{ fontWeight: '500', marginBottom: '0.5rem' }}>课程描述</h4>
                            <p style={{ color: '#374151' }}>{course.course_description}</p>
                        </div>

                        <div style={{ display: 'flex', justifyContent: 'space-between', paddingTop: '1rem' }}>
                            {/* 左侧：退课按钮（仅在已选时显示） */}
                            <div>
                                {currentUser && isEnrolled && !checkingEnrollment && (
                                    <button
                                        onClick={handleUnenroll}
                                        disabled={enrolling}
                                        style={{
                                            ...styles.buttonDanger,
                                            ...(enrolling ? styles.buttonDisabled : {})
                                        }}
                                    >
                                        {enrolling ? '退课中...' : '退课'}
                                    </button>
                                )}
                            </div>
                            
                            {/* 右侧：选课/已选按钮 */}
                            <div>
                                <button
                                    onClick={handleEnroll}
                                    disabled={!currentUser || enrolling || checkingEnrollment || isEnrolled}
                                    style={getButtonStyle()}
                                >
                                    {getButtonText()}
                                </button>
                            </div>
                        </div>
                    </div>
                ) : (
                    <div style={styles.emptyText}>加载失败</div>
                )}
            </div>
        </div>
    );
}

// 管理员工具箱组件
function AdminToolbox({ isOpen, onClose, onRefresh }) {
    const [activeTab, setActiveTab] = useState('addCourse');
    const [courseForm, setCourseForm] = useState({
        course_code: '',
        course_name: '',
        course_description: '',
        credits: 3,
        instructor: '',
        semester: '',
        time_slot: '',
        course_location: ''
    });
    const [courses, setCourses] = useState([]);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        if (isOpen && activeTab === 'manageCourse') {
            api.getCourses().then(data => setCourses(data.courses || []));
        }
    }, [isOpen, activeTab]);

    const handleAddCourse = async (e) => {
        e.preventDefault();
        setLoading(true);
        try {
            await api.addCourse(courseForm);
            alert('课程添加成功');
            setCourseForm({
                course_code: '',
                course_name: '',
                course_description: '',
                credits: 3,
                instructor: '',
                semester: '',
                time_slot: '',
                course_location: ''
            });
            onRefresh();
        } catch (err) {
            alert('添加失败');
        }
        setLoading(false);
    };

    const handleRemoveAllStudents = async (courseId) => {
        if (window.confirm('确认要从该课程中移除所有学生吗？')) {
            try {
                await api.removeAllStudentsFromCourse(courseId);
                alert('操作成功');
            } catch (err) {
                alert('操作失败');
            }
        }
    };

    if (!isOpen) return null;

    return (
        <div style={styles.modal}>
            <div style={{ ...styles.modalContent, ...styles.modalContentLarge }}>
                <div style={styles.modalHeader}>
                    <h2 style={styles.modalTitle}>管理员工具箱</h2>
                    <button onClick={onClose} style={styles.closeButton}>
                        <X size={20} />
                    </button>
                </div>

                <div style={styles.tabContainer}>
                    <button
                        onClick={() => setActiveTab('addCourse')}
                        style={{
                            ...styles.tab,
                            ...styles.tabLeft,
                            ...(activeTab === 'addCourse' ? styles.tabActive : {})
                        }}
                    >
                        添加课程
                    </button>
                    <button
                        onClick={() => setActiveTab('manageCourse')}
                        style={{
                            ...styles.tab,
                            ...styles.tabRight,
                            ...(activeTab === 'manageCourse' ? styles.tabActive : {})
                        }}
                    >
                        管理课程
                    </button>
                </div>

                {activeTab === 'addCourse' && (
                    <form onSubmit={handleAddCourse}>
                        <div style={{ ...styles.grid, ...styles.gridCols2 }}>
                            <div>
                                <label style={styles.label}>课程代码</label>
                                <input
                                    required
                                    value={courseForm.course_code}
                                    onChange={(e) => setCourseForm({ ...courseForm, course_code: e.target.value })}
                                    style={styles.input}
                                    placeholder="COMP1117"
                                />
                            </div>
                            <div>
                                <label style={styles.label}>课程名称</label>
                                <input
                                    required
                                    value={courseForm.course_name}
                                    onChange={(e) => setCourseForm({ ...courseForm, course_name: e.target.value })}
                                    style={styles.input}
                                    placeholder="Computer Programming"
                                />
                            </div>
                        </div>

                        <div style={styles.formGroup}>
                            <label style={styles.label}>课程描述</label>
                            <textarea
                                value={courseForm.course_description}
                                onChange={(e) => setCourseForm({ ...courseForm, course_description: e.target.value })}
                                style={styles.textarea}
                                placeholder="课程描述"
                            />
                        </div>

                        <div style={{ ...styles.grid, ...styles.gridCols2 }}>
                            <div>
                                <label style={styles.label}>学分</label>
                                <input
                                    type="number"
                                    value={courseForm.credits}
                                    onChange={(e) => setCourseForm({ ...courseForm, credits: parseInt(e.target.value) })}
                                    style={styles.input}
                                    min="1"
                                    max="8"
                                />
                            </div>
                            <div>
                                <label style={styles.label}>教师</label>
                                <input
                                    value={courseForm.instructor}
                                    onChange={(e) => setCourseForm({ ...courseForm, instructor: e.target.value })}
                                    style={styles.input}
                                    placeholder="Prof. Chen"
                                />
                            </div>
                        </div>

                        <div style={{ ...styles.grid, ...styles.gridCols2 }}>
                            <div>
                                <label style={styles.label}>学期</label>
                                <input
                                    value={courseForm.semester}
                                    onChange={(e) => setCourseForm({ ...courseForm, semester: e.target.value })}
                                    style={styles.input}
                                    placeholder="2024 Spring"
                                />
                            </div>
                            <div>
                                <label style={styles.label}>时间</label>
                                <input
                                    value={courseForm.time_slot}
                                    onChange={(e) => setCourseForm({ ...courseForm, time_slot: e.target.value })}
                                    style={styles.input}
                                    placeholder="Mon 9:00-12:00"
                                />
                            </div>
                        </div>

                        <div style={styles.formGroup}>
                            <label style={styles.label}>地点</label>
                            <input
                                value={courseForm.course_location}
                                onChange={(e) => setCourseForm({ ...courseForm, course_location: e.target.value })}
                                style={styles.input_long}
                                placeholder="CYC LT1"
                            />
                        </div>

                        <button
                            type="submit"
                            disabled={loading}
                            style={{
                                ...styles.buttonSuccess,
                                width: '100%',
                                justifyContent: 'center',
                                ...(loading ? styles.buttonDisabled : {})
                            }}
                        >
                            {loading ? '添加中...' : '添加课程'}
                        </button>
                    </form>
                )}

                {activeTab === 'manageCourse' && (
                    <div>
                        <h3 style={{ fontWeight: '600', marginBottom: '1rem' }}>课程管理</h3>
                        {courses.map(course => (
                            <div key={course.id} style={styles.listItem}>
                                <div style={styles.itemInfo}>
                                    <div style={styles.itemTitle}>{course.course_name}</div>
                                    <div style={styles.itemSubtitle}>{course.course_code}</div>
                                </div>
                                <button
                                    onClick={() => handleRemoveAllStudents(course.id)}
                                    style={styles.buttonDanger}
                                >
                                    移除所有学生
                                </button>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
}

// 已选课程弹窗组件
function MyCoursesModal({ isOpen, onClose, currentUser, onUnenroll }) {
    const [myCourses, setMyCourses] = useState([]);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        if (isOpen && currentUser) {
            setLoading(true);
            api.getStudentCourses(currentUser.id)
                .then(data => setMyCourses(data.courses || []))
                .catch(err => console.error('获取已选课程失败:', err))
                .finally(() => setLoading(false));
        }
    }, [isOpen, currentUser]);

    const handleUnenroll = async (courseId) => {
        try {
            await api.unenrollCourse(currentUser.id, courseId);
            setMyCourses(myCourses.filter(c => c.course_id !== courseId));
            onUnenroll();
        } catch (err) {
            alert('退课失败，请稍后重试');
        }
    };

    if (!isOpen) return null;

    return (
        <div style={styles.modal}>
            <div style={{ ...styles.modalContent, ...styles.modalContentLarge }}>
                <div style={styles.modalHeader}>
                    <h2 style={styles.modalTitle}>我的课程</h2>
                    <button onClick={onClose} style={styles.closeButton}>
                        <X size={20} />
                    </button>
                </div>

                {!currentUser ? (
                    <div style={styles.emptyText}>请先登录</div>
                ) : loading ? (
                    <div style={styles.loadingText}>加载中...</div>
                ) : myCourses.length === 0 ? (
                    <div style={styles.emptyText}>还没有选择任何课程</div>
                ) : (
                    <div>
                        {myCourses.map(course => (
                            <div key={course.course_id} style={styles.listItem}>
                                <div style={styles.itemInfo}>
                                    <div style={styles.itemTitle}>{course.course_name}</div>
                                    <div style={styles.itemSubtitle}>{course.course_code}</div>
                                </div>
                                <button
                                    onClick={() => handleUnenroll(course.course_id)}
                                    style={styles.buttonDanger}
                                >
                                    退课
                                </button>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
}

// 主应用组件
export default function CourseSelectionPlatform() {
    const [currentUser, setCurrentUser] = useState(null);
    const [courses, setCourses] = useState([]);
    const [students, setStudents] = useState([]);
    const [selectedStudent, setSelectedStudent] = useState('all');
    const [searchKeyword, setSearchKeyword] = useState('');
    const [filteredCourses, setFilteredCourses] = useState([]);
    const [loading, setLoading] = useState(true);
    const [userLoading, setUserLoading] = useState(true); // 新增：用户状态加载中

    // 弹窗状态
    const [showLogin, setShowLogin] = useState(false);
    const [showCourseDetail, setShowCourseDetail] = useState(false);
    const [selectedCourseId, setSelectedCourseId] = useState(null);
    const [showAdminToolbox, setShowAdminToolbox] = useState(false);
    const [showMyCourses, setShowMyCourses] = useState(false);
    const [showLogoutConfirm, setShowLogoutConfirm] = useState(false);

    // 悬停状态
    const [hoveredCard, setHoveredCard] = useState(null);

    // 页面加载时从localStorage恢复用户状态
    useEffect(() => {
        const restoreUserState = async () => {
            try {
                const savedUser = localStorage.getItem('courseSelectionUser');
                if (savedUser) {
                    const userData = JSON.parse(savedUser);

                    // 验证用户是否仍然有效（可选）
                    const studentsData = await api.getStudents();
                    const students = studentsData.students || [];
                    const validUser = students.find(s => s.id === userData.id);

                    if (validUser) {
                        setCurrentUser(validUser);
                    } else {
                        // 用户不存在了，清除本地存储
                        localStorage.removeItem('courseSelectionUser');
                    }
                }
            } catch (err) {
                console.error('恢复用户状态失败:', err);
                localStorage.removeItem('courseSelectionUser');
            }
            setUserLoading(false);
        };

        restoreUserState();
    }, []);

    // 初始化数据
    useEffect(() => {
        loadData();
    }, []);

    // 搜索和筛选
    useEffect(() => {
        if (searchKeyword) {
            api.searchCourses(searchKeyword)
                .then(data => setFilteredCourses(data.courses || []))
                .catch(err => console.error('搜索失败:', err));
        } else if (selectedStudent !== 'all') {
            api.getStudentCourses(selectedStudent)
                .then(data => {
                    const studentCourses = data.courses || [];
                    setFilteredCourses(studentCourses.map(sc => ({
                        id: sc.course_id,
                        course_code: sc.course_code,
                        course_name: sc.course_name
                    })));
                })
                .catch(err => console.error('获取学生课程失败:', err));
        } else {
            setFilteredCourses(courses);
        }
    }, [searchKeyword, selectedStudent, courses]);

    const loadData = async () => {
        setLoading(true);
        try {
            const [coursesData, studentsData] = await Promise.all([
                api.getCourses(),
                api.getStudents()
            ]);
            setCourses(coursesData.courses || []);
            setStudents(studentsData.students || []);
        } catch (err) {
            console.error('数据加载失败:', err);
        }
        setLoading(false);
    };

    const handleLogin = (user) => {
        setCurrentUser(user);
        // 保存到localStorage
        localStorage.setItem('courseSelectionUser', JSON.stringify(user));
    };

    const handleLogout = () => {
        setCurrentUser(null);
        setShowLogoutConfirm(false);
        // 从localStorage清除
        localStorage.removeItem('courseSelectionUser');
    };

    const openCourseDetail = (courseId) => {
        setSelectedCourseId(courseId);
        setShowCourseDetail(true);
    };

    const refreshData = () => {
        loadData();
    };

    // 如果还在加载用户状态，显示加载界面
    if (userLoading) {
        return (
            <div style={{
                ...styles.container,
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                height: '100vh'
            }}>
                <div style={{ textAlign: 'center' }}>
                    <div style={{ fontSize: '1.125rem', color: '#6b7280' }}>加载中...</div>
                </div>
            </div>
        );
    }

    return (
        <div style={styles.container}>
            {/* 固定顶部栏 */}
            <div style={styles.fixedHeader}>
                <div style={styles.headerContent}>
                    <h1 style={styles.logo}>R1C 选课平台</h1>

                    <div style={{ position: 'relative' }}>
                        {currentUser ? (
                            <>
                                <button
                                    onClick={() => setShowLogoutConfirm(true)}
                                    style={styles.button}
                                >
                                    <User size={18} />
                                    {currentUser.name}
                                </button>

                                {showLogoutConfirm && (
                                    <div style={{
                                        position: 'absolute',
                                        right: 0,
                                        top: '100%',
                                        marginTop: '0.5rem',
                                        backgroundColor: 'white',
                                        border: '1px solid #e5e7eb',
                                        borderRadius: '0.5rem',
                                        boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1)',
                                        padding: '1rem',
                                        width: '12rem'
                                    }}>
                                        <p style={{ fontSize: '0.875rem', marginBottom: '0.75rem', color: '#374151' }}>确认要登出吗？</p>
                                        <div style={{ display: 'flex', gap: '0.5rem' }}>
                                            <button
                                                onClick={handleLogout}
                                                style={{
                                                    ...styles.buttonDanger,
                                                    flex: 1,
                                                    fontSize: '0.75rem',
                                                    padding: '0.25rem 0.75rem'
                                                }}
                                            >
                                                登出
                                            </button>
                                            <button
                                                onClick={() => setShowLogoutConfirm(false)}
                                                style={{
                                                    flex: 1,
                                                    backgroundColor: '#d1d5db',
                                                    color: '#374151',
                                                    border: 'none',
                                                    borderRadius: '0.375rem',
                                                    fontSize: '0.75rem',
                                                    padding: '0.25rem 0.75rem',
                                                    cursor: 'pointer'
                                                }}
                                            >
                                                取消
                                            </button>
                                        </div>
                                    </div>
                                )}
                            </>
                        ) : (
                            <button
                                onClick={() => setShowLogin(true)}
                                style={styles.button}
                            >
                                <User size={18} />
                                注册/登录
                            </button>
                        )}
                    </div>
                </div>
            </div>

            {/* 主内容区域 */}
            <div style={styles.mainContent}>
                {/* 搜索和筛选栏 */}
                <div style={styles.searchSection}>
                    <div style={styles.searchContainer}>
                        {/* 左侧：学生选择 */}
                        <div style={styles.selectContainer}>
                            <label style={styles.label}>选择学生</label>
                            <select
                                value={selectedStudent}
                                onChange={(e) => setSelectedStudent(e.target.value)}
                                style={styles.select}
                            >
                                <option value="all">查看全部课程</option>
                                {students.map(student => (
                                    <option key={student.id} value={student.id}>
                                        {student.name} ({student.email})
                                    </option>
                                ))}
                            </select>
                        </div>

                        {/* 右侧：搜索框 */}
                        <div style={styles.searchInputContainer}>
                            <label style={styles.label}>搜索课程</label>
                            <div style={styles.searchInputWrapper}>
                                <Search style={styles.searchIcon} size={18} />
                                <input
                                    type="text"
                                    value={searchKeyword}
                                    onChange={(e) => setSearchKeyword(e.target.value)}
                                    style={styles.searchInput}
                                    placeholder="输入课程名称、课程代码或教师姓名"
                                />
                            </div>
                        </div>
                    </div>
                </div>

                {/* 课程网格 */}
                <div style={styles.coursesSection}>
                    {loading ? (
                        <div style={styles.loadingText}>
                            <div>加载中...</div>
                        </div>
                    ) : filteredCourses.length === 0 ? (
                        <div style={styles.emptyText}>
                            <div>没有找到相关课程</div>
                        </div>
                    ) : (
                        <div style={styles.coursesGrid}>
                            {filteredCourses.map(course => (
                                <div
                                    key={course.id}
                                    onClick={() => openCourseDetail(course.id)}
                                    onMouseEnter={() => setHoveredCard(course.id)}
                                    onMouseLeave={() => setHoveredCard(null)}
                                    style={{
                                        ...styles.courseCard,
                                        ...(hoveredCard === course.id ? styles.courseCardHover : {})
                                    }}
                                >
                                    <div style={styles.courseCode}>
                                        {course.course_code}
                                    </div>
                                    <div style={styles.courseName}>
                                        {course.course_name}
                                    </div>
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </div>

            {/* 浮动按钮组 */}
            <div style={styles.floatingButtons}>
                {/* 查看已选课程按钮 */}
                <button
                    onClick={() => setShowMyCourses(true)}
                    style={{
                        ...styles.floatingButton,
                        ...styles.floatingButtonGreen
                    }}
                    title="查看已选课程"
                >
                    <Eye size={20} />
                </button>

                {/* 管理员工具箱按钮 */}
                <button
                    onClick={() => setShowAdminToolbox(true)}
                    style={{
                        ...styles.floatingButton,
                        ...styles.floatingButtonOrange
                    }}
                    title="管理员工具箱"
                >
                    <Settings size={20} />
                </button>
            </div>

            {/* 各种弹窗 */}
            <LoginModal
                isOpen={showLogin}
                onClose={() => setShowLogin(false)}
                onLogin={handleLogin}
            />

            <CourseDetailModal
                courseId={selectedCourseId}
                isOpen={showCourseDetail}
                onClose={() => setShowCourseDetail(false)}
                currentUser={currentUser}
                onEnroll={refreshData}
            />

            <AdminToolbox
                isOpen={showAdminToolbox}
                onClose={() => setShowAdminToolbox(false)}
                onRefresh={refreshData}
            />

            <MyCoursesModal
                isOpen={showMyCourses}
                onClose={() => setShowMyCourses(false)}
                currentUser={currentUser}
                onUnenroll={refreshData}
            />
        </div>
    );
}