/*
E-Ticaret Sepeti ve İndirim Uygulaması

Bir e-ticaret sitesi için alışveriş sepetindeki ürünlere ve
kullanıcının indirim durumuna göre toplam fiyatı hesaplayan
bir program yazmalısınız. Aşağıdaki gereksinimleri dikkate alın:

Her bir ürünün fiyatı ve adedi kullanıcı tarafından girilecektir.
Kullanıcı, sepete eklediği ürünlerin toplam fiyatını görmek
 isteyecektir.
Kullanıcıya bir indirim kodu sorulacaktır. Kullanıcı herhangi
 bir indirim kodu girebilir veya indirim kodu girmeden devam edebilir

Eğer kullanıcı "INDIRIM10" kodunu girerse, toplam fiyattan
%10 indirim yapılacaktır.
Eğer kullanıcı "BEDAVA" kodunu girerse ve toplam ürün adedi
3 veya daha fazlaysa, en ucuz ürün bedava olacaktır.
Eğer herhangi bir indirim kodu girilmezse veya geçersiz
bir kod girilirse, indirim uygulanmayacaktır.
Yukarıdaki gereksinimlere göre, kullanıcının sepetindeki
 ürünlerin toplam fiyatını hesaplayan ve indirim kodunu
 uygulayan bir Go programı yazın.

*/

package main

import (
	"fmt"
)

func main() {
	productPrices := make(map[string]float64)
	//ayakkabı = 123234 -> key value = map

	productQuantities := make(map[string]int)

	addProduct("Laptop", 5000, 2, productPrices, productQuantities)
	addProduct("Telefon", 3000, 1, productPrices, productQuantities)
	addProduct("Tablet", 2000, 3, productPrices, productQuantities)
	//ürün ismi, fiyatı, miktarı

	totalPrice := calculateTotal(productPrices, productQuantities)

	fmt.Println("Sepet Toplamı:", totalPrice)

	discountCode := "BEDAVA" //kullanıcıdan al

	applyDiscount(discountCode, totalPrice, productPrices, productQuantities)

}

func addProduct(productName string, price float64, quantity int, prices map[string]float64, quantities map[string]int) {
	prices[productName] = price
	quantities[productName] = quantity
}

func calculateTotal(prices map[string]float64, quantities map[string]int) float64 {
	total := 0.0
	for product, price := range prices {
		total += price * float64(quantities[product])
	}

	return total
}

func applyDiscount(code string, total float64, prices map[string]float64, quantities map[string]int) {
	switch code {
	case "INDIRIM10":
		total *= 0.9
		fmt.Println("INDIRIM10 kodu uygulandı. Indirim toplam :", total)
	case "BEDAVA":
		if total >= 3*prices["Tablet"] { // Eğer toplam fiyat 3 tabletten fazlaysa
			cheapestProduct := findCheapestProduct(prices, quantities)
			total -= prices[cheapestProduct] //en ucuz ürün bedava
			fmt.Println("Bedava kodu uygulandı. indiirm toplamı:", total)
		} else {
			fmt.Println("bedava kodu için minimum alışveriş tutarı sağlanmadı")
		}
	default:
		fmt.Println("Geçersiz veya tanımlanmamış indirim kodu")
	}
}

func findCheapestProduct(prices map[string]float64, quantities map[string]int) string {
	cheapest := ""
	minPrice := -1.0
	for product, price := range prices {
		if minPrice == -1 || price < minPrice {
			cheapest = product
			minPrice = price
		}
	}
	return cheapest
}
