package main

import "fmt"

func main() {
	//bir uygulamamız var, üye kayıt işlemleri yapıyoruz, üyeler kullanıcı bilgilerini giriyor ve biz veri tabanına kaydediyoruz.
	//biz, database bir saldırı olduğunda parolaları, şifrelememiz gerekiyor.
	//M1234. ŞİFRELİYİM ŞİFRELİ HALİ = MKADŞMKFŞEAŞRNVMDAVCÇW!'^+%&/&/&/&%+^'^^'%+%/%/&%++^'!ASDŞCŞÇÇ12E56
	//parolaları biz belli algoritmalar ile encode ediyoruz.
	//bizim büyük projelerimiz 1-2 senelik değil genelde 5-6 senelik oluyor.
	//biz burada şifleme işlemini md5 şifreleme yöntemi ile yaptık sonra yeni bir şifreleme yöntemi geldi onu biz denemek istiyoruz.
	//bizim şiflediğimiz, identitynumber, cart number, prola gibi 50 tane yerde şifreleme yaptığımızı düşünelim.
	//biz bir deneme yapmak için bu 50 yerde tek tek şifreleme yöntemini değiştirecekmiyiz, yoksa biz bunu kolay bir yazılımla mı yapacağız, bu yazılım interfacedir.
	//biz bu yönetimi tek bir yerden yapacağız. bunun içinde interfaceleri kullanacağız.
	//bir şifreleme algoritmasının içerisinde şifreyi çözme fonksiyonuda vardır.

	//şifreyi neden çözdürüyoruz, M1234. = mjemfmaksfı2329reropfe = M1234.

	var encoder ISifrele

	encoder = &ZEncoder{}

	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Encoder("123456")
	encoder.Decoder("eceörclascçqwdç3çedwqeds")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")
	encoder.Decoder("Kartnumarasılidmemrvlaşcöalspmpiep")

}

type XEncoder struct {
}

type YEncoder struct {
}

type ZEncoder struct {
}

type ISifrele interface {
	Encoder(value string)
	Decoder(value string)
}

func (xEncoder XEncoder) Encoder(value string) {
	fmt.Println("XEncoder ile şifrelendi.")
}

func (xEncoder XEncoder) Decoder(value string) {
	fmt.Println("XEncoder ile Şifreyi çözdü")
}

func (yEncoder YEncoder) Encoder(value string) {
	fmt.Println("YEncoder ile şifrelendi.")
}

func (yEncoder YEncoder) Decoder(value string) {
	fmt.Println("YEncoder ile Şifreyi çözdü")
}

func (zEncoder ZEncoder) Encoder(value string) {
	fmt.Println("ZEncoder ile şifrelendi.")
}

func (zEncoder ZEncoder) Decoder(value string) {
	fmt.Println("ZEncoder ile Şifreyi çözdü")
}

//projemizin 50 yerinde biz bu XCoder şifreleme algoritmasını kullandık.
//yarın birgün bu xencoderı değiştirmek isteyebiliriz.
