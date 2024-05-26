package main

import "fmt"

func main() {

	//Defer : ertelemek anlamına gelmektedir. yani biz bir işlemi erteleyebiliyoruz.
	// ertelemekten kasıt, biz defere herhangi bir fonksiyonda kullandığımız zaman, bu aslında erteleneyor, hemen çalıştırılmıyor,
	//fonksiyon bittiğinde statement, yani fonksiyon çağırımı yapılıyor.
	//programlarımızda, temizleme görevi olduğu zaman defer kullanılır.
	//veritabanı işlemi açtınız, bişeyler eklediniz, en sonunda işlem bitince veri tabanınızın kapatılması gerekmektedir.
	//sistemde ölümcül hata olursa veya, sistemde şu hata olursa sistemi kapat.
	//biz bişey açtık belli işlemlerden sonra açtığımız o şeyi kapatmak istiyorsak kullan.

	// defer fmt.Println("hello") //defere uygulandı, artık bu kodun çalışması ertelendi,. fonksiyon bittikten sonra defer çalışır.
	// defer fmt.Println("hello1")
	// //deferlerde stack yapısı kullanılır, yani ilk giren son çıkar. kural.
	// //enson ilk defer, sonra ikinci defer yazdırır. son beklediğimiz ilk yazılır

	// fmt.Println("Burası main fonksiyonu")
	// defer fmt.Println("birinci defer")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// defer fmt.Println("4")

	// var condition bool //true yada false

	// condition = true
	// defer fmt.Println("1.defer")

	// if condition == true {

	// 	return //yani programı sonlandır, retuna düştüğü zaman gerisini okumaz.
	// }

	// defer fmt.Println("2.defer")

	// a := 10
	// b := 11

	// defer fmt.Println(a)
	// defer fmt.Println(b)

	// a = 100
	// b = 200

	// fmt.Println(a)
	// fmt.Println(b)

	var condition bool
	condition = true

	//panic deferden önce çalıştırılır. uygulamamızda panic en son çalışır.

	defer cleanApp()

	if condition == true {
		//panic : uygulamanızı sonlandırır.
		panic("Uygulama hackleniyor ... ")

		//dış dünyaya kapattık kendi içimizde işlemler yapacaksak, defer kullanacağız

	}

}

func cleanApp() {
	//uygulamayı temizlesin işlemlerden sonra

	fmt.Println("Clean up çalıştırıldı.")
}
