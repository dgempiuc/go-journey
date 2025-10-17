1.package

go programları package'lar halinde organize edilir. 
bir package, birlikte compile edilen aynı directory içerisindeki source filelar yani functionlar, typelar, variablelar ve constantlardır.
bunlar, aynı paket içindeki diğer tüm source filelar tarafından görülebilir.


2.module 

bir repo, bir ya da daha fazla module içerebilir.
module, birlikte release edilen, ilgili go packagelarının derlemesi.
bir go reposu, tipik olarak yalnızca bir module içerir, reponun root'unda konumludur,
module, go.mod dosyasının bulunduğu directorydeki packageları ve bu directornin altındaki diğer directoryleri içerir.
ismi go.mod'dur, orada module path tanımlanır, bu da, module içindeki tüm packagelar için import path prefixidir.

mesela

hello.go (package ismi main olmalı)
go.mod (module x olsun. her zaman project rootunda olmalı)
tree/
- tree.go (package ismi tree)

tree.go dosyasını hello.go main içerisinde kullanabilmek için, import edilmeli ama package import edilmeli.

import "x/tree"

go.mod'da module tanımı a/b/c olabilir. ama fiziksel olarak a/b/c paketinin olmasına gerek yok.
ama tree'yi import ederken a/b/c/tree yazılır.

ya da go workspace kullanılır.
multi-module development için.
aynı anda birden fazla go module üzerinde çalışmayı sağlar.

ana root directoryde
go.work dosyası. içerisinde, go.mod olan tüm directoryleri yani module'leri dahil eder.

go 1.25

use (
t02/localrepo
t02/testrepo
t02
t01
)

go work sync // go.work dosyasındaki tüm module'lerin dependencysisini sync eder.

3. install/build

code'u, build etmeden önce remote repoya publish etmeye gerek yok.
bir module, local olarak bir repoya ait olmadan tanımlanabilir.
ama remote repolardaki go kodlarını da kullanmaya imkan var.

go install çalışması için, cwd'deki module'deki path verilmeli, yoksa hata. aşağıdakilerin 3ü de kabul.
go install example/user/hello
// build eder ve binary dosyası oluşturur, localde %USERPROFILE%\go\bin\ altına exe dosyası koyar.
// GOPATH ve GOBIN environemnt variableları ile, install directory control edilir.
// GOPATH altındaki bin klasörüne ya da GOBIN klasörüne.
// go env -w GOBIN=/x/y ile default value ayarlanır.
// go env -u GOBIN ile de, önceden ayarlanan unset edilir.

go install example/user/hello
go install .
go install

daha sonra da, exe dosyası çalıştırılır.
hello // hello world

install dediğin için build de yaptı ve hello.exe oluşturdu.
eğer exe olmadan hello yazsaydın exe bulamadım diyecekti.
build etmeden direkt çalıştırmak için.

go run hello.go 

4. remote paketler
go'da remote paketler de indirilip kullanılabilir.

import "github.com/gin-gonic/gin"

go get ile indirilip kullanılır.
go get github.com/gin-gonic/gin

http isteği yapar, url'deki <meta name="go-import"> tagından gerçek repo url'ini bulur ve git clone yapar.
bu sayede, repo'nun yerini değiştiresen bile, import pathler değişmez.

özetle
java maven		go modules
pom.xml			go.mod
maven central	direkt git repoları
mvn install		go get
mvn clean pakcage go build

5.kendi remote repomuzu oluşturma

githubda repo oluştur.
mesela github.com/dgempiuc/go-journey

rootda go.mod olacak.

'''
module github.com/dgempiuc/go-journey

go 1.25.1
'''

daha sonra istediğim gibi directory olusturabilirim.
remoterepo diye directory olustrudum, içerisinde stringutils.go olusturdum, Reverse functionı yazdım.
bunu pushladım

daha sonra test etmek istediğim yerde, önce bunu indiriyorum.

go get github.com/dgempiuc/go-journey

bu gidip, bulunduğu konumdaki go.mod'u günceller ve '''require github.com/dgempiuc/go-journey''' olarak ekler.

'''
import (
"fmt"
"github.com/dgempiuc/go-journey/remoterepo"
)

remoterepo.Reverse(original)
'''

olarak kullanırım.


6.local module kullanma

localrepo isminde module oluştur (denizg/repos), stringutils.go dosyasını koy. 

bu modul'ü eklemek istediğin yerde, oranın go.mod'una şunu eklersin

'''
require denizg/repo/localrepo v0.0.0
replace denizg/repo/localrepo => ../localrepo
'''

kullanmak istediğin yerde de

'''
import (
"denizg/repo/localrepo"
)

localrepo.Truncate(reversed, 4, "aaa")
'''

