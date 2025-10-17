# Go ve Gin Web Framework

RESTful web service API yazmak için

## Bağımlılık Ekleme

Önce bağımlılığı ekle:

```go
import "github.com/gin-gonic/gin"
```

## Struct Tags

Model oluştururken **struct** tag'ları oluyor: `json:"title"` gibi.

Bu **struct** tag'lar, serialize/deserialize işlemlerinde kullanılıyor.

**Struct** field'ı JSON'a serialize ederken hangi JSON field'a serialize edeceğini belirtir.
Yani **struct** field'ın JSON'da görüleceği isim.

```go
type War struct {
    Name      string    `json:"war-name"`
    DateBegin time.Time `json:"begin-date"`
    Duration  int       `json:"total-day"`
}
```

### Struct Tag Seçenekleri

- Field JSON'da görünmesin: `json:"-"`
- Eğer boşsa JSON'da görünmesin: `json:"title,omitempty"`
- **float64**'ü **string** olarak encode et: `json:"price,string"`
- Birden fazla tag: `json:"artist" db:"artist_name"`

### Validation Tags (binding tag ile)

- **Required**: `json:"id" binding:"required"`
- **Min/Max**: `json:"title" binding:"required,min=1,max=100"`
- **Greater than**: `json:"price" binding:"required,gt=0"`

## In-Memory Data

```go
var InMemoryWarData = []War{
    {Name: "Miryokefalon", DateBegin: time.Date(1176, time.September, 17, 0, 0, 0, 0, time.UTC), Duration: 1},
    {Name: "Yassıçemen", DateBegin: time.Date(1230, time.August, 10, 0, 0, 0, 0, time.UTC), Duration: 2},
}
```

## GET - Tüm Kayıtları Getirme

Handler **function**'ı oluştur. Tüm savaşları kullanıcıya dönecek.

**gin.Context** parametresi:
- Request detaylarını taşır
- Validate eder
- JSON'ı serialize eder

`c.IndentedJSON` → **struct**'ı JSON'a serialize eder ve onu da response'a ekler.

```go
func getWars(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, InMemoryWarData)
}
```

### Router Konfigürasyonu

Daha sonra da bu handler **function**'ı endpoint'e assign et.

- `gin.Default()` ile Gin router initialize edilir
- Endpoint-handler **function** mapping yapılır
- `Run()` ile router'a HTTP server eklenir ve server başlatılır

```go
func main() {
    router := gin.Default()
    router.GET("/wars", getWars)

    router.Run("localhost:8080")
}
```

### Response Örneği

`localhost:8080/wars`

```json
[
    {
        "war-name": "Miryokefalon",
        "begin-date": "1176-09-17T00:00:00Z",
        "total-day": 1
    },
    {
        "war-name": "Yassıçemen",
        "begin-date": "1230-08-10T00:00:00Z",
        "total-day": 2
    }
]
```

## POST - Yeni Kayıt Ekleme

`c.BindJSON()` ile request body, **newWar** **variable**'ına bağlanır.

```go
func addWar(c *gin.Context) {
    var newWar War
    err := c.BindJSON(&newWar)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "request body is empty or invalid"})
        return
    }
    InMemoryWarData = append(InMemoryWarData, newWar)
    c.IndentedJSON(http.StatusCreated, InMemoryWarData)
}
```

### Request Örneği

`POST localhost:8080/wars`

```json
{
    "war-name": "Kösedağ",
    "begin-date": "1243-07-03T00:00:00Z",
    "total-day": 1
}
```

## GET - Parametreli İstek

`/wars/[id]`

`c.Param()` ile URL'deki path parameter alınır.

Hangi path parameter ismini alacağını, handler **function**-endpoint mapping'indeyken belirtiyorsun.

```go
func getWar(c *gin.Context) {
    name := c.Param("war-name")

    for _, a := range InMemoryWarData {
        if a.Name == name {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "war not found"})
}
```

### Router Mapping

```go
router.GET("/wars/:war-name", getWar)
```

### Request Örneği

`GET localhost:8080/wars/Miryokefalon`
