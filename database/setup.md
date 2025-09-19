连接到 PostgreSQL

```bash
psql -U postgres
```

创建数据库，创建用户

```sql
CREATE DATABASE course_management;
CREATE USER course_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE course_management TO course_user;
```

初始化脚本

```bash
psql -U postgres -d course_management -f init.sql
```

