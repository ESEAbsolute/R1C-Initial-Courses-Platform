package config

import (
    "os"
    "strconv"
    "strings"
    "time"

	"course-management/models"
    
    "github.com/joho/godotenv"
)

type Config struct {
    App      AppConfig       `json:"app"`
    Server   ServerConfig    `json:"server"`
    Database models.DBConfig `json:"database"`
    CORS     CORSConfig      `json:"cors"`
    Security SecurityConfig  `json:"security"`
    Log      LogConfig       `json:"log"`
}

type AppConfig struct {
    Environment string `json:"environment"`
    GinMode     string `json:"gin_mode"`
}

type ServerConfig struct {
    Port int `json:"port"`
}

type CORSConfig struct {
    AllowedOrigins   []string      `json:"allowed_origins"`
    AllowCredentials bool          `json:"allow_credentials"`
    MaxAge           time.Duration `json:"max_age"`
}

type SecurityConfig struct {
    HeadersEnabled     bool `json:"headers_enabled"`
    DebugRoutesEnabled bool `json:"debug_routes_enabled"`
    SampleDataEnabled  bool `json:"sample_data_enabled"`
}

type LogConfig struct {
    Level  string `json:"level"`
    Format string `json:"format"`
}

func LoadConfig() (*Config, error) {
    // 加载.env文件（如果存在）
    if err := godotenv.Load(); err != nil {
        // .env文件不存在也不是错误，继续使用环境变量
    }
    
    config := &Config{
        App: AppConfig{
            Environment: getEnvWithDefault("APP_ENV", "development"),
            GinMode:     getEnvWithDefault("GIN_MODE", "debug"),
        },
        Server: ServerConfig{
            Port: getIntEnvWithDefault("SERVER_PORT", 8080),
        },
        Database: models.DBConfig{
            Host:     getEnvWithDefault("DB_HOST", "localhost"),
            Port:     getIntEnvWithDefault("DB_PORT", 5432),
            User:     getEnvWithDefault("DB_USER", "postgres"),
            Password: getEnvWithDefault("DB_PASSWORD", ""),
            DBName:   getEnvWithDefault("DB_NAME", "course_management"),
            SSLMode:  getEnvWithDefault("DB_SSLMODE", "disable"),
        },
        CORS: CORSConfig{
            AllowedOrigins:   parseOrigins(getEnvWithDefault("CORS_ALLOWED_ORIGINS", "http://localhost:3000")),
            AllowCredentials: getBoolEnvWithDefault("CORS_ALLOW_CREDENTIALS", true),
            MaxAge:           time.Duration(getIntEnvWithDefault("CORS_MAX_AGE", 43200)) * time.Second,
        },
        Security: SecurityConfig{
            HeadersEnabled:     getBoolEnvWithDefault("SECURITY_HEADERS_ENABLED", true),
            DebugRoutesEnabled: getBoolEnvWithDefault("DEBUG_ROUTES_ENABLED", true),
            SampleDataEnabled:  getBoolEnvWithDefault("SAMPLE_DATA_ENABLED", true),
        },
        Log: LogConfig{
            Level:  getEnvWithDefault("LOG_LEVEL", "info"),
            Format: getEnvWithDefault("LOG_FORMAT", "text"),
        },
    }
    
    return config, nil
}

// 辅助函数：获取环境变量，如果不存在则使用默认值
func getEnvWithDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

// 辅助函数：获取整数环境变量
func getIntEnvWithDefault(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}

// 辅助函数：获取布尔环境变量
func getBoolEnvWithDefault(key string, defaultValue bool) bool {
    if value := os.Getenv(key); value != "" {
        if boolValue, err := strconv.ParseBool(value); err == nil {
            return boolValue
        }
    }
    return defaultValue
}

// 辅助函数：解析CORS源列表
func parseOrigins(originsStr string) []string {
    if originsStr == "" {
        return []string{}
    }
    
    origins := strings.Split(originsStr, ",")
    result := make([]string, 0, len(origins))
    
    for _, origin := range origins {
        origin = strings.TrimSpace(origin)
        if origin != "" {
            result = append(result, origin)
        }
    }
    
    return result
}

