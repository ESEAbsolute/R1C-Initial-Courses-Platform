@echo off
chcp 65001 > nul

REM 检查 .env 文件是否存在
if not exist .env (
    echo ❌ .env 文件不存在！
    echo 💡 请复制 .env.example 为 .env 并配置您的数据库信息：
    echo    copy .env.example .env
    echo    然后编辑 .env 文件中的 DB_PASSWORD 等配置
    pause
    exit /b 1
)

REM 加载环境变量
for /f "usebackq tokens=1,2 delims==" %%a in (".env") do (
    if "%%a" neq "" if "%%b" neq "" (
        set %%a=%%b
    )
)

echo 🔧 使用配置 
echo    数据库: %DB_USER%@%DB_HOST%:%DB_PORT%/%DB_NAME%
echo    服务器: %SERVER_HOST%:%SERVER_PORT%
echo.

REM 启动应用
echo 🚀 启动 R1C 选课系统...
go run main.go

pause