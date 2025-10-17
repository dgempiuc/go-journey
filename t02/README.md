# Tutorial 02 - Go Modules Deep Dive

## 1. Package

**Go** programları **package**'lar halinde organize edilir.

Bir **package**, birlikte compile edilen aynı directory içerisindeki source file'lar yani **function**'lar, **type**'lar, **variable**'lar ve **constant**'lardır.

Bunlar, aynı **package** içindeki diğer tüm source file'lar tarafından görülebilir.

## 2. Module

Bir repo, bir ya da daha fazla **module** içerebilir.

**Module**, birlikte release edilen, ilgili **Go** **package**'larının derlemesi.

Bir **Go** reposu, tipik olarak yalnızca bir **module** içerir, repo'nun root'unda konumludur.

**Module**, **go.mod** dosyasının bulunduğu directory'deki **package**'ları ve bu directory'nin altındaki diğer directory'leri içerir.

İsmi **go.mod**'dur, orada **module** path tanımlanır. Bu da, **module** içindeki tüm **package**'lar için **import** path prefix'idir.

### Örnek Yapı

```
hello.go          (package ismi main olmalı)
go.mod            (module x olsun. her zaman project root'unda olmalı)
tree/
  tree.go         (package ismi tree)
```

**tree.go** dosyasını **hello.go** **main** içerisinde kullanabilmek için **import** edilmeli:

```go
import "x/tree"
```

**go.mod**'da **module** tanımı `a/b/c` olabilir. Ama fiziksel olarak `a/b/c` **package**'inin olmasına gerek yok.

Ama **tree**'yi **import** ederken `a/b/c/tree` yazılır.

## 3. Go Workspace

**Go** workspace, multi-module development için kullanılır.

Aynı anda birden fazla **Go** **module** üzerinde çalışmayı sağlar.

Ana root directory'de **go.work** dosyası. İçerisinde, **go.mod** olan tüm directory'leri yani **module**'leri dahil eder.

```go
go 1.25

use (
  t02/localrepo
  t02/testrepo
  t02
  t01
)
```

```bash
go work sync  # go.work dosyasındaki tüm module'lerin dependency'sini sync eder
```

## 4. Install/Build

Code'u build etmeden önce remote repo'ya publish etmeye gerek yok.

Bir **module**, local olarak bir repo'ya ait olmadan tanımlanabilir.

Ama remote repo'lardaki **Go** kodlarını da kullanmaya imkan var.

### Go Install

`**go** install` çalışması için, cwd'deki **module**'deki path verilmeli, yoksa hata. Aşağıdakilerin 3'ü de kabul:

```bash
go install example/user/hello
go install .
go install
```

`**go** install`:
- Build eder ve binary dosyası oluşturur
- Local'de `%USERPROFILE%\go\bin\` altına exe dosyası koyar
- `GOPATH` ve `GOBIN` environment variable'ları ile install directory control edilir
- `GOPATH` altındaki **bin** klasörüne ya da `GOBIN` klasörüne

```bash
go env -w GOBIN=/x/y  # default value ayarlanır
go env -u GOBIN       # önceden ayarlanan unset edilir
```

Daha sonra da exe dosyası çalıştırılır:

```bash
hello  # hello world
```

**install** dediğin için build de yaptı ve **hello.exe** oluşturdu.

Eğer exe olmadan **hello** yazsaydın exe bulamadım diyecekti.

Build etmeden direkt çalıştırmak için:

```bash
go run hello.go
```

## 5. Remote Paketler

**Go**'da remote **package**'ler de indirilip kullanılabilir:

```go
import "github.com/gin-gonic/gin"
```

`**go** get` ile indirilip kullanılır:

```bash
go get github.com/gin-gonic/gin
```

HTTP isteği yapar, URL'deki `<meta name="go-import">` tag'inden gerçek repo URL'ini bulur ve **git clone** yapar.

Bu sayede, repo'nun yerini değiştiresen bile **import** path'ler değişmez.

### Karşılaştırma

| Java Maven | Go Modules |
|------------|------------|
| pom.xml | **go.mod** |
| Maven Central | Direkt **git** repoları |
| mvn install | **go get** |
| mvn clean package | **go build** |

## 6. Kendi Remote Repo'muzu Oluşturma

GitHub'da repo oluştur:
- Mesela `github.com/dgempiuc/go-journey`

Root'da **go.mod** olacak:

```go
module github.com/dgempiuc/go-journey

go 1.25.1
```

Daha sonra istediğim gibi directory oluşturabilirim.

**remoterepo** diye directory oluşturdum, içerisinde **stringutils.go** oluşturdum, **Reverse** **function**'ı yazdım.

Bunu push'ladım.

Daha sonra test etmek istediğim yerde, önce bunu indiriyorum:

```bash
go get github.com/dgempiuc/go-journey
```

Bu gidip, bulunduğu konumdaki **go.mod**'u günceller ve şunu ekler:

```go
require github.com/dgempiuc/go-journey
```

Kullanım:

```go
import (
  "fmt"
  "github.com/dgempiuc/go-journey/remoterepo"
)

remoterepo.Reverse(original)
```

## 7. Local Module Kullanma

**localrepo** isminde **module** oluştur (`denizg/repos`), **stringutils.go** dosyasını koy.

Bu **module**'ü eklemek istediğin yerde, oranın **go.mod**'una şunu eklersin:

```go
require denizg/repo/localrepo v0.0.0
replace denizg/repo/localrepo => ../localrepo
```

Kullanmak istediğin yerde de:

```go
import (
  "denizg/repo/localrepo"
)

localrepo.Truncate(reversed, 4, "aaa")
```
