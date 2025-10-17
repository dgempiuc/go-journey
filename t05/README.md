# Gin Web Framework ve GORM

## Database Konfigürasyonu

Root directory'de `.env` dosyası oluştur ve DB değerlerini yaz.

Daha sonra database'e bağlanmak için **database.go** dosyası oluştur.

### Config Struct

Config'leri tutmak için **struct** oluştur:

```go
type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}
```

### Environment Variables'ı Yükleme

`.env` dosyasını oku ve **struct** nesnesi oluştur:

```go
func LoadDBConfig() DBConfig {
    return DBConfig{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }
}
```

### Database Connection

Connection string oluştur ve **GORM** ile DB bağlantısı aç, **gorm.DB** nesnesi dön:

```go
func DatabaseConnection(cfg DBConfig) (*gorm.DB, error) {
    dbStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
    var dialector gorm.Dialector = postgres.Open(dbStr)
    var dbConn *gorm.DB
    var err error
    dbConn, err = gorm.Open(dialector)
    return dbConn, err
}
```

## Model Katmanı

Şimdi t04'de kullanılan **War** **struct**'ını genişletme işi.

O zaman sadece JSON işlemi vardı. Şimdi ise DB işlemleri de var.

Fakat Java'daki gibi ayrı ayrı entity ve model nesneleri oluşturulmayacak.

Onun yerine **struct** tag'e **GORM** tag'ları de eklenecek.

Model için ayrı **model.go** dosyasında yap:

```go
type War struct {
    Name      string    `json:"war-name" gorm:"primary_key"`
    DateBegin time.Time `json:"begin-date" gorm:"not null"`
    Duration  int       `json:"total-day" gorm:"not null"`
}
```

## Repository Katmanı

**repository.go** oluşturulur:

```go
type WarRepository struct {
    DB *gorm.DB
}

func (wp WarRepository) GetAllWar() []model.War {
    var wars []model.War
    wp.DB.Find(&wars)
    return wars
}

func (wp WarRepository) AddWar(war model.War) {
    wp.DB.Create(&war)
}

func (wp WarRepository) GetWarByName(name string) model.War {
    var war model.War
    wp.DB.Where("name = ?", name).First(&war)
    return war
}
```

## Service Katmanı

**service.go** oluşturulur:

```go
type WarService struct {
    WP repository.WarRepository
}

func (ws WarService) GetAllWar() []model.War {
    return ws.WP.GetAllWar()
}

func (ws WarService) AddWar(war model.War) {
    ws.WP.AddWar(war)
}

func (ws WarService) GetWarByName(name string) model.War {
    return ws.WP.GetWarByName(name)
}
```

## Handler Katmanı

**handler.go** oluşturulur:

```go
type WarHandler struct {
    WS service.WarService
}

func (wh WarHandler) GetWars(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, wh.WS.GetAllWar())
}

func (wh WarHandler) AddWar(c *gin.Context) {
    var newWar model.War
    c.BindJSON(&newWar)
    wh.WS.AddWar(newWar)
    c.IndentedJSON(http.StatusCreated, newWar)
}

func (wh WarHandler) GetWarByName(c *gin.Context) {
    name := c.Param("name")
    c.IndentedJSON(http.StatusOK, wh.WS.GetWarByName(name))
}
```

## Main Function

**main.go**'yu güncelle:

```go
func main() {
    godotenv.Load()
    var cfg config.DBConfig = config.LoadDBConfig()
    db, err := config.DatabaseConnection(cfg)
    if err != nil {
        fmt.Println("db connection sırasında hata olustu. %v", err)
    }
    err = db.AutoMigrate(model.War{})
    if err != nil {
        fmt.Println("tabloları otomatik olusturma sırasında hata olustu. %v", err)
    }
    var warRepo repository.WarRepository = repository.WarRepository{db}
    var warService service.WarService = service.WarService{warRepo}
    var warHandler handler.WarHandler = handler.WarHandler{warService}

    initDBData(warService)

    router := gin.Default()
    router.GET("/wars", warHandler.GetWars)
    router.POST("/wars", warHandler.AddWar)
    router.GET("/wars/:name", warHandler.GetWarByName)

    router.Run("localhost:8080")
}
```

## Initial Data

In-memory data ve database initialization:

```go
var InMemoryWarData = []model.War{
    {Name: "Miryokefalon", DateBegin: time.Date(1176, time.September, 17, 0, 0, 0, 0, time.UTC), Duration: 1},
    {Name: "Yassıçemen", DateBegin: time.Date(1230, time.August, 10, 0, 0, 0, 0, time.UTC), Duration: 2},
}

func initDBData(ws service.WarService) {
    fmt.Println(cap(ws.GetAllWar()))
    if len(ws.GetAllWar()) == 0 {
        for _, data := range InMemoryWarData {
            ws.AddWar(data)
        }
    }
}
```
