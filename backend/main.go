package main

import (
    "fmt"
    "log"
    
    "course-management/config"
    "course-management/handlers"
    "course-management/models"
    
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    // 加载配置
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("配置加载失败:", err)
    }

    gin.SetMode(cfg.App.GinMode)
    
    // 连接数据库
    db, err := models.NewDatabase(cfg.Database)
    if err != nil {
        log.Fatal("数据库连接失败:", err)
    }
    defer db.Close()
    
    // 初始化示例数据
    if cfg.Security.SampleDataEnabled {
        if err := db.InitializeSampleData(); err != nil {
            log.Printf("示例数据初始化失败: %v", err)
        }
    }
    
    // 创建路由器
    r := gin.Default()
    r.SetTrustedProxies([]string{})
    
    // 配置CORS
    corsMiddleware := cors.Config{
        AllowOrigins:     cfg.CORS.AllowedOrigins,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{
            "Origin", 
            "Content-Type", 
            "Accept", 
            "Authorization",
            "X-Requested-With",
        },
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: cfg.CORS.AllowCredentials,
        MaxAge:           cfg.CORS.MaxAge,
    }
    
    r.Use(cors.New(corsMiddleware))
    
    // 安全中间件
    if cfg.Security.HeadersEnabled {
        r.Use(securityHeaders(cfg.App.Environment))
    }
    
    r.Use(requestLogger(cfg.Log))
    
    // 创建API处理器并设置路由
    apiHandler := handlers.NewAPIHandler(db)
    r.Use(apiHandler.ErrorHandler())
    apiHandler.SetupRoutes(r)
    
    // 添加调试端点
    if cfg.Security.DebugRoutesEnabled {
        setupDebugRoutes(r, db)
    }
    
    // 启动服务器
    serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
    log.Printf("🚀 服务器启动: http://localhost:%d", cfg.Server.Port)
    log.Printf("📝 环境: %s", cfg.App.Environment)
    log.Printf("🌐 允许的CORS源: %v", cfg.CORS.AllowedOrigins)
    
    if err := r.Run(serverAddr); err != nil {
        log.Fatal("服务器启动失败:", err)
    }
}

func securityHeaders(environment string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 防止XSS攻击
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-Frame-Options", "DENY")
        c.Header("X-XSS-Protection", "1; mode=block")
        
        // HTTPS相关（生产环境）
        if environment == "production" {
            c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        }
        
        c.Next()
    }
}

func requestLogger(logConfig config.LogConfig) gin.HandlerFunc {
    if logConfig.Format == "json" {
        return gin.Logger()
    }
    
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
            param.ClientIP,
            param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
            param.Method,
            param.Path,
            param.Request.Proto,
            param.StatusCode,
            param.Latency,
            param.Request.UserAgent(),
            param.ErrorMessage,
        )
    })
}

func setupDebugRoutes(r *gin.Engine, db *models.Database) {
    debug := r.Group("/debug")
    {
        debug.GET("/stats", func(c *gin.Context) {
            stats, err := db.GetDataStats()
            if err != nil {
                c.JSON(500, gin.H{"error": err.Error()})
                return
            }
            c.JSON(200, gin.H{"stats": stats})
        })
        
        debug.POST("/reset-data", func(c *gin.Context) {
            if err := db.ClearAllData(); err != nil {
                c.JSON(500, gin.H{"error": err.Error()})
                return
            }
            
            if err := db.InitializeSampleData(); err != nil {
                c.JSON(500, gin.H{"error": err.Error()})
                return
            }
            
            c.JSON(200, gin.H{"message": "数据重置成功"})
        })
    }
}
