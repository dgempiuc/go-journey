# Google Wire - Dependency Injection

## Manual Dependency Injection

**Go** community'de yaygın olan manuel DI:

```go
func main() {
    cfg := config.LoadConfig()
    db := config.NewDatabase(cfg)
    repo := repository.NewWarRepository(db)
    svc := service.NewWarService(repo)
    handler := handler.NewWarHandler(svc)
}
```

#### Javada spring varken go'da neden manuel di?

- Runtime'da ekstra yük yok (runtime overhead yok)
- Hata derleme zamanında çıkar (compile-time safe)
- Explicit - her şey açık

#### Google Wire

Eğer büyük mikroservislerde fazlaca dependency'ler varsa, Google'ın geliştirdiği **Wire** var.
Dependency değiştirmeyi kolaylaştırır.

- Spring Container'a göre avantajlı
- Spring tarzı dependency injection yapmıyor
- İşleri hala **compile-time**'da yapıyor (**compile-time dependency injection**)
- Spring Container tarzı bir şeyi olmadığı için **memory tüketimi çok düşük** (neredeyse 10 kat daha az memory)
- Wire ile aynı fayda, **sıfır overhead** ile

Micronaut da benzer, o da **compile-time DI**.
Fakat orada da annotation tabanlı tarama, component scan ve lifecycle var.

#### Wire Kullanımı

google/wire

This repository was archived by the owner on Aug 25, 2025. It is now read-only.


**go.mod** dosyasına **require** olarak ekle:

```bash
go get github.com/google/wire
```

**go mod tidy** ile yükle:

```bash
go mod tidy
```

Wire CLI'ı yükle:

```bash
go install github.com/google/wire/cmd/wire@latest
wire version  # kontrol et
```

## Wire Provider Tanımlama

Artık her **package**'da wire provider'ları tanımla.

### Repository Layer

**repository/repository.go**:

```go
type WarRepository struct{}

func NewWarRepository() WarRepository {
    return WarRepository{}
}
```

**repository/wire.go**:

Tüm **wire.go** dosyalarının başında şu olmalı: Anlamı: "Bu dosyayı sadece wire komutu çalışırken kullan!"

```go
//go:build wireinject
// +build wireinject

package repository

import "github.com/google/wire"

// ProviderSet repository layer'ı için tüm provider'ları gruplar
var ProviderSet = wire.NewSet(
    NewWarRepository,
)
```

### Provider Kavramı

**Provider** dediğin, bir şeyi sağlayan **function**.

- **NewWarRepository** **function**, **WarRepository** sağlıyor
- Wire'a diyorsun ki: "Birisi **WarRepository** isterse, **NewWarRepository** **function**'ı çağır"
- Aslında bildiğin **constructor function** - Wire context'inde ona **provider** diyoruz

**Önemli:** Aynı **Type** için birden fazla provider olamaz!

#### Hatalı Örnek

```go
// repository.go
func NewWarRepository() WarRepository {
    return WarRepository{}
}

func NewWarRepository2() WarRepository {
    return WarRepository{}
}
```

```go
// wire.go
// ❌ HATA: multiple providers for repository.WarRepository
var ProviderSet = wire.NewSet(
    NewWarRepository,   // WarRepository sağlar
    NewWarRepository2,  // WarRepository sağlar // CONFLICT!
)
```

Çünkü **WarRepository** istiyorum dediğinde, hangisini kullanacağını bilmez.

### Service Layer

**service/service.go**:

```go
type WarService struct {
    wp repository.WarRepository
}

func NewWarService(wp repository.WarRepository) WarService {
    return WarService{wp}
}
```

**service/wire.go**:

```go
package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
    NewWarService,
)
```

**NewWarService** provider'ı:
- Sağladığı: **WarService**
- İhtiyaç duyduğu: **WarRepository**

Wire'a diyorsun ki: "Birisi **WarService** isterse önce başka provider'lardan **WarRepository** bul, sonra **NewWarService(repo)** çağır"

### Handler Layer

**handler/handler.go**:

```go
type WarHandler struct {
    ws service.WarService
}

func NewWarHandler(ws service.WarService) WarHandler {
    return WarHandler{ws}
}
```

**handler/wire.go**:

```go
package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
    NewWarHandler,
)
```


#### t05'deki Main (manual dependency injection)

**main.go**:

```go
func main() {
    warRepo := repository.NewWarRepository()
    warService := service.NewWarService(warRepo)
    warHandler := handler.NewWarHandler(warService)

    router := gin.Default()
    router.GET("/wars", warHandler.GetWars)
    router.POST("/wars", warHandler.AddWar)
    router.GET("/wars/:name", warHandler.GetWarByName)

    router.Run("localhost:8080")
}
```

#### Wire Injector

**wire.go**:

```go
func InitializeApp() handler.WarHandler {
    wire.Build(
        repository.ProviderSet,  // NewWarRepository içerir
        service.ProviderSet,     // NewWarService içerir
        handler.ProviderSet,     // NewWarHandler içerir
    )

    return handler.WarHandler{}
}
```

Wire dependency graph oluşturur:

```
NewWarRepository() → WarRepository
        ↓
NewWarService(repo) → WarService
        ↓
NewWarHandler(service) → WarHandler
```

#### Yeni Main (Wire ile)

**main.go**:

```go
func main() {
    warHandler := InitializeApp()

    router := gin.Default()
    router.GET("/wars", warHandler.GetWars)
    router.POST("/wars", warHandler.AddWar)
    router.GET("/wars/:name", warHandler.GetWarByName)

    router.Run("localhost:8080")
}
```

#### Wire Code Generation

**main.go**'nun olduğu dizinde şu komut çalıştırılır:

```bash
wire
```

> **Not:** Kodda her değişiklik oldukça bu komut da çalıştırılmalı.

**wire_gen.go** dosyası oluşturur. İçeriği şu şekilde:

```go
func InitializeApp() handler.WarHandler {
    warRepository := repository.NewWarRepository()
    warService := service.NewWarService(warRepository)
    warHandler := handler.NewWarHandler(warService)
    return warHandler
}
```

**main.go**'dan silinmiş 3 kod satırı artık otomatik oluşturuldu.

## Çalıştırma

Artık `go run main.go` olarak çalıştıramazsın - **main.go**'da **InitializeApp** **function**'ı bulamadım hatası atar.

```bash
go run .
```

olmalı.
