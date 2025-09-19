# R1C 课程管理系统

基于 React + Go + PostgreSQL 的现代化课程选课平台，支持学生选课、课程管理、用户认证等功能。

## 项目概述

本项目是为 RIC 2025-2026 年度执委招募技术考核开发的模拟课程管理系统。

### 核心功能

- 课程管理：
  - 显示所有课程
  - 按 课程名称 / 课程代码 / 教师名称 搜索课程
  - 添加新课程（没有引入管理员账号鉴权功能，不安全）
  - 删除一个课程中的所有学生（仅课程被废弃用；并没有禁止学生再次选择该课程；没有引入管理员账号鉴权功能，不安全）
- 学生管理：
  - 下拉菜单查看学生列表
  - 输入全新的 姓名 + 邮箱 自动注册新学生
- 选课管理：
  - 查看指定学生选课信息
  - 学生使用 **姓名 + 邮箱 模拟**登陆（没有使用密码和邮箱验证，不安全）
  - 已登录学生的 选课 / 退课 功能

## 技术栈

- 前端: React
- 后端: Go
- 数据库: PostgreSQL

## 分支说明

- `main`：主分支

- `feat-frontend`：前端开发分支
- `feat-backend`：后端开发分支
- `feat-database`：数据库开发分支

## 项目结构

```text
r1c-course-management/
├── frontend/                # 前端代码
│   └── course-management/   # React 应用目录
│       ├── public/          # 静态资源
│       ├── src/
│       │   └── App.js       # 主应用组件
│       ├── package.json
│       └── .env.example
├── backend/                 # 后端代码
│   ├── main.go              # 主程序入口
│   ├── config/              # 配置管理
│   │   └── config.go
│   ├── handlers/            # API处理器
│   │   └── api_handler.go
│   ├── models/              # 数据模型
│   │   ├── database.go
│   │   ├── courses.go
│   │   ├── student.go
│   │   ├── enrollment.go
│   │   └── sample_data.go
│   ├── types/               # 类型定义
│   │   └── responses.go
│   ├── go.mod
│   ├── go.sum
│   ├── .env.development.example
│   ├── .env.production.example
│   └── .env.test.example
├── database/                # 数据库相关
│   ├── init.sql             # 数据库初始化脚本
│   └── setup.md             # 数据库设置说明
├── .gitignore
└── README.md
```

## 快速开始

### 前置需求

前端使用的是 React，所以你需要安装 [Node.js](https://nodejs.org/en/download)。

后端使用的是 [Go](https://go.dev/dl/)。

数据库使用的是 [PostgreSQL](https://www.postgresql.org/download/)。

### 克隆项目

```bash
git clone https://github.com/ESEAbsolute/R1C-Initial-Courses-Platform.git
cd r1c-course-management
```

### 数据库设置

#### 连接到 PostgreSQL

```bash
psql -U postgres
```

#### 创建数据库和用户

```sql
CREATE DATABASE course_management;
CREATE USER course_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE course_management TO course_user;
\q
```

#### 初始化数据表

```bash
psql -U postgres -d course_management -f database/init.sql
```

> 提示：数据库初始化脚本会自动创建所需的表结构和索引，后端启动时会根据配置文件选择是否自动插入示例数据。

### 后端设置

1. 切换到后端目录（`cd backend`）；
2. 配置环境变量，将 .env.\*.example 中三者选择其一复制到 .env 文件，并编辑 .env 文件；
3. 切换到主目录，双击 `run_backend.cmd`。

### 前端设置

1. 切换到前端目录（`cd frontend/course-management`）；
2. 使用 `npm install` 安装依赖；
3. 配置环境变量，将 .env*.example 复制到 .env 文件，并编辑 .env 文件；
4. 切换到主目录，双击 `run_frontend.cmd`。

### License

MIT License

```text
The MIT License (MIT)

Copyright (c) 2025 ESEAbsolute

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

