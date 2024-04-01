package main

import "fmt"

func main() {
	//cihaz
	var temperature int16
	var windspeed int
	var cloudStatus string
	var rainStatus string
	fmt.Scan(&temperature)
	fmt.Scan(&windspeed)
	fmt.Scan(&cloudStatus)
	fmt.Scan(&rainStatus)

	//eşittir değer atamak için kullanılır. bir değişkenimiz var bunun değeri 10 a eşittir veya bu değişkenin değeri sarıdır.
	//color = sarı (color içerisinde sarı rengi depoladı)
	//cloudStatus = 10 // cloudStatus = müşteriden al yani fmtscan

	// eşit eşit nedir: istenilen değer o değişkenin içerisinde varmı, cloudstatus ün içerisinde açık değeri varmı

	//1. Sıcaklık 25 derecenin üzerinde, rüzgar hızı 10 km/saat'in altında ve bulut durumu "açık" ise:
	if temperature > 25 && windspeed < 10 && cloudStatus == "açık" {

		//a. Yağış durumu "dolu" ise: "Hava sıcak, güneşli ve dolu yağıyor, dikkatli ol."

		if rainStatus == "dolu" {
			fmt.Println("Hava sıcak, güneşli ve dolu yağıyor, dikkatli ol.")
			//b. Yağış durumu "yağmur" ise: "Hava sıcak, güneşli ve yağmur yağıyor, dikkatli ol."
		} else if rainStatus == "yağmur" {
			fmt.Println("Hava sıcak, güneşli ve yağmur yağıyor, dikkatli ol.")
			//c. Yağış durumu "kar" ise: "Hava sıcak, güneşli ve kar yağıyor, dikkatli ol."
		} else if rainStatus == "kar" {
			fmt.Println("Hava sıcak, güneşli ve kar yağıyor, dikkatli ol.")
			//d. Yağış durumu yoksa: "Hava sıcak, güneşli bir gün."
		} else {
			fmt.Println("Hava sıcak, güneşli bir gün.")
		}
		//2. Sıcaklık 25 derecenin üzerinde, rüzgar hızı 10 km/saat'in altında ve bulut durumu "kapalı" ise:
	} else if temperature > 25 && windspeed < 10 && cloudStatus == "kapalı" {
		//a. Yağış durumu "dolu" ise: "Hava sıcak, bulutlu ve dolu yağıyor, dikkatli ol."
		if rainStatus == "dolu" {
			fmt.Println("Hava sıcak, bulutlu ve dolu yağıyor, dikkatli ol.")
		}

	}
	/*
		   1. Sıcaklık 25 derecenin üzerinde, rüzgar hızı 10 km/saat'in altında ve bulut durumu "açık" ise:

					a. Yağış durumu "dolu" ise: "Hava sıcak, güneşli ve dolu yağıyor, dikkatli ol."
					b. Yağış durumu "yağmur" ise: "Hava sıcak, güneşli ve yağmur yağıyor, dikkatli ol."
					c. Yağış durumu "kar" ise: "Hava sıcak, güneşli ve kar yağıyor, dikkatli ol."
					d. Yağış durumu yoksa: "Hava sıcak, güneşli bir gün."


		   2. Sıcaklık 25 derecenin üzerinde, rüzgar hızı 10 km/saat'in altında ve bulut durumu "kapalı" ise:

					a. Yağış durumu "dolu" ise: "Hava sıcak, bulutlu ve dolu yağıyor, dikkatli ol."
					Yağış durumu "yağmur" ise: "Hava sıcak, bulutlu ve yağmur yağıyor, dikkatli ol."
					Yağış durumu "kar" ise: "Hava sıcak, bulutlu ve kar yağıyor, dikkatli ol."
					Yağış durumu yoksa: "Hava sıcak ama bulutlu bir gün."
		   Sıcaklık 25 derecenin üzerinde, rüzgar hızı 10 km/saat veya üzerinde ve bulut durumu "açık" ise:

		   Yağış durumu "dolu" ise: "Hava sıcak ve rüzgarlı, güneşli ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava sıcak ve rüzgarlı, güneşli ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava sıcak ve rüzgarlı, güneşli ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava sıcak ve rüzgarlı, güneşli bir gün."
		   Sıcaklık 25 derecenin üzerinde, rüzgar hızı 10 km/saat veya üzerinde ve bulut durumu "kapalı" ise:

		   Yağış durumu "dolu" ise: "Hava sıcak ve rüzgarlı, bulutlu ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava sıcak ve rüzgarlı, bulutlu ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava sıcak ve rüzgarlı, bulutlu ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava sıcak ve rüzgarlı, bulutlu bir gün."
		   Sıcaklık 15 ile 25 derece arasında, rüzgar hızı 10 km/saat'in altında ve bulut durumu "açık" ise:

		   Yağış durumu "dolu" ise: "Hava ılık, açık bir gün ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava ılık, açık bir gün ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava ılık, açık bir gün ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava ılık, açık bir gün."
		   Sıcaklık 15 ile 25 derece arasında, rüzgar hızı 10 km/saat'in altında ve bulut durumu "kapalı" ise:

		   Yağış durumu "dolu" ise: "Hava ılık ve bulutlu, dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava ılık ve bulutlu, yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava ılık ve bulutlu, kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava ılık ve bulutlu bir gün."
		   Sıcaklık 15 ile 25 derece arasında, rüzgar hızı 10 km/saat veya üzerinde ve bulut durumu "açık" ise:

		   Yağış durumu "dolu" ise: "Hava ılık ve rüzgarlı, açık bir gün ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava ılık ve rüzgarlı, açık bir gün ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava ılık ve rüzgarlı, açık bir gün ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava ılık ve rüzgarlı, açık bir gün."
		   Sıcaklık 15 ile 25 derece arasında, rüzgar hızı 10 km/saat veya üzerinde ve bulut durumu "kapalı" ise:

		   Yağış durumu "dolu" ise: "Hava ılık ve rüzgarlı, bulutlu ve dolu yağıyor, dikkatli ol."
		   - Yağış durumu "yağmur" ise: "Hava ılık ve rüzgarlı, bulutlu ve yağmur yağıyor, dikkatli ol."
		   - Yağış durumu "kar" ise: "Hava ılık ve rüzgarlı, bulutlu ve kar yağıyor, dikkatli ol."
		   - Yağış durumu yoksa: "Hava ılık ve rüzgarlı, bulutlu bir gün."

		   Sıcaklık 5 ile 15 derece arasında, rüzgar hızı 10 km/saat'in altında ve bulut durumu "açık" ise:

		   Yağış durumu "dolu" ise: "Hava serin, hafif rüzgarlı ve açık bir gün ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava serin, hafif rüzgarlı ve açık bir gün ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava serin, hafif rüzgarlı ve açık bir gün ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava serin, hafif rüzgarlı ve açık bir gün."
		   Sıcaklık 5 ile 15 derece arasında, rüzgar hızı 10 km/saat'in altında ve bulut durumu "kapalı" ise:

		   Yağış durumu "dolu" ise: "Hava serin ve bulutlu, hafif rüzgarlı ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava serin ve bulutlu, hafif rüzgarlı ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava serin ve bulutlu, hafif rüzgarlı ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava serin ve bulutlu, hafif rüzgarlı bir gün."
		   Sıcaklık 5 ile 15 derece arasında, rüzgar hızı 10 km/saat veya üzerinde ve bulut durumu "açık" ise:

		   Yağış durumu "dolu" ise: "Hava serin ve rüzgarlı, açık bir gün ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava serin ve rüzgarlı, açık bir gün ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava serin ve rüzgarlı, açık bir gün ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava serin ve rüzgarlı, açık bir gün."
		   Sıcaklık 5 ile 15 derece arasında, rüzgar hızı 10 km/saat veya üzerinde ve bulut durumu "kapalı" ise:

		   Yağış durumu "dolu" ise: "Hava serin ve rüzgarlı, bulutlu ve dolu yağıyor, dikkatli ol."
		   Yağış durumu "yağmur" ise: "Hava serin ve rüzgarlı, bulutlu ve yağmur yağıyor, dikkatli ol."
		   Yağış durumu "kar" ise: "Hava serin ve rüzgarlı, bulutlu ve kar yağıyor, dikkatli ol."
		   Yağış durumu yoksa: "Hava serin ve rüzgarlı, bulutlu bir gün."
		   Ayrıca, kullanıcı sıcaklık, rüzgar hızı, bulut durumu veya yağış durumu girişi olarak geçerli bir sayı veya geçerli bir kelime dışında bir değer girerse, "Geçersiz giriş! Lütfen sıcaklık, rüzgar hızı, bulut durumu ve yağış durumu için geçerli değerler girin." mesajını gösterin.
	*/
}
