package main

import "fmt"

func main() {
	//Genelde banka, e ticaret uygulamaları gibi komplex uygulamaları parçalara bölerek yazarız, bu parçalar struct yardımı ile yapmaktayız.
	//örneğin bir sipariş kendi içerisinde birden fazla küçük yapılar olabilir herbir yapı için ayrı bir struct kullanmaktayız.
	//c# ve javada classlar vardırher bir classta bir veya birden fazla method vardır. Biz golangde bu methodlara struc deriz.

	fmt.Println(SiparişVer{})
}

type SiparişVer struct {
	//ödeme alındı
}

type SiparişBeklemede struct {
	//notification gönder
	//mesaj gönder
	//kargoya bildirim ver
	X int
	Y int
}

type sipraişKrgoyabildirimGönderir struct {
	//işlemleri yazacağız.
}
