# Tutorial 01 - Go Basics

## Go Commands

```bash
go <command> [arguments]

go build ...     # compile packages and dependencies
go install ...   # compile and install packages and dependencies
go get ...       # add dependencies to current module and install them
go test ...      # test packages
```

## 1. İlk Adım

**go.mod** oluştur:

```go
module journey/denizg/tutorial01

go 1.25
```

Daha sonra **hello.go** dosyası oluştur.
- **package** ismi **main** olacak
- Hello world yazdıracak **main** **function**'a sahip

### Build & Install

```bash
go install
```

`**go** install` komutunu çalıştırınca:
- Build eder
- exe üretir
- Onu da `%USERPROFILE%\go\bin\` altına install eder/yükler

**Install directory**, `GOPATH` ve `GOBIN` ile değiştirilir.

**exe** ismi, **go.mod**'daki en son **package** ismi. Yani `tutorial01.exe`

> `**go** install` yerine `**go** build` denseydi, output dosyası o anki klasörde oluşturulurdu ama exe oluşturulmazdı.

## 2. Package'ları Module'den Import Etme

**morepackage** oluştur. Altına da **reverse.go** oluştur.

**hello.go**'ya **import** edilir:

```go
import "journey/denizg/tutorial01/morepackage"

morepackage.Reverse("Hello, world!")
```

## 3. Package'ları Remote Module'lerden Import Etme

**import** path, **package** source code'un **git** gibi version kontrol sistemi kullanarak nasıl elde edeceğini açıklayabilir.

Eğer URL görürse, **package**'ları otomatik olarak remote repolardan çeker.

**hello** dosyasına **import** olarak bunu ekle:

```go
import "github.com/google/go-cmp/cmp"

cmp.Diff("Hello World", "Hello Go")
```

Şimdi, external **module**'e bağımlılık/dependency var. Bu **module**'ü download etmeli ve **go.mod** file'a kaydetmeli.

```bash
go mod tidy
```

Bu komut:
- **import** edilen **package**'lar için eksik **module** gereksinimlerini ekler
- Artık kullanılmayan **module** gereksinimleri kaldırır

Her zaman en güncel halini getirmeyebilir, çünkü **module** cache kullanır.
- Yeni değişiklikleri hemen çekmez
- **Module**'i ilk indirdiğinde cache'e koyar (`pkg/mod`)
- `**go** mod tidy` desen bile cache'ten alır

```bash
go clean -modcache  # module cache'i temizler
```

Bunu çalıştırınca, **go.mod** dosyasına ekler:

```go
require github.com/google/go-cmp v0.5.4
```

Ya da tek tek bağımlılıkları eklersin:

```bash
go get github.com/google/go-cmp/cmp
```

Bunlar `%USER%/go` altında `pkg/mod` altına indirilir.

## 4. Test Etme

**Go**, lightweight test framework'üne sahip:
- `**go** test` command + `testing` **package**

Bir testi, `_test.go` ismiyle biten bir dosya yaratarak yazarsın.

Bu, `TestXXX` isimli **function**'lar içerir ve signature:

```go
func (t *testing.T)
```

Test framework, bu tür her bir **function**'ı çalıştırır.

Eğer **function**, `t.Error` yada `t.Fail` gibi failure **function**'ları çağırırsa, test failed oldu sayılır.

```bash
go test
# PASS
```
