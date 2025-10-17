1. ilk adım

go.mod oluştur.

module journey/denizg/tutorial01

go 1.25

daha sonra hello.go dosyası oluştur.
package ismi main olacak.
hello world yazdıracak main functiona sahip.

go install komutunu çalıştırınca, build eder, exe üretir ve onu da %USERPROFILE%\go\bin\ altına install eder /yükler.
install directory, GOPATH ve GOBIN ile değiştirilir.

exe ismi, go.mod'daki en son paket ismi. yani tutorial01.exe

go install yerine go build denseydi, output dosyası o anki klasörüde oluşturulurdu ama exe oluşturulmazdı.

2. package'ları module'den import etme

morepackage olustur. altına da reverse.go oluştur.

hello.go'ya import edilir.

import "journey/denizg/tutorial01/morepackage"

morepackage.Reverse("Hello, world!")

3. package'ları remote module'lerden import etme

import path, package source code'un git gibi version kontrol sistemi kullanarak nasıl elde edeceğini açıklayabilir.
eğer url görürse, packageları otomatik olarak remote repolardan çeker.

hello dosyasına import olarak bunu ekle.

import "github.com/google/go-cmp/cmp"

cmp.Diff("Hello World", "Hello Go")

şimdi, external module'e bağımlılık/depenndency var. bu module'ü download etmeli ve go.mod file'a kaydetmeli.
go mod tidy
bu komut, import edilen package'lar için eksik module gereksinimlerini ekler,  artık kullanılmayan module gereksinimleri kaldırır.

bunu çalıştırınca, go mod dosyasına ekler.
require github.com/google/go-cmp v0.5.4

ya da tek tek bagımlılıkları eklersin.
go get go get github.com/google/go-cmp/cmp

bunlar %USER%/go altında pkg/mod altına indirilir.

4. test etme

4. test etme

go, lightweight test frameworküne sahip.
"go test" command + "testing" package

bir testi, _test.go ismiyle biten bir dosya yaratarak yazarsın.
bu, TestXXX isimli functionlar içerir ve signature, func (t *testing.T) içerir.
test framework, bu tür her bir functionı çalıştırır.

eğer function, t.Error yada t.Fail gibi failure functionları çağırırsa, test, failed oldu sayılır.

go test
// PASS
