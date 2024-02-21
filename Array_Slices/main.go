package main

import "fmt"

func main() {
	//Arraylere boyut tanımlaya bilirizi.

	// var categories [11]string
	// //Arrayi boş yazdırdık.
	// //fmt.Println(categories)
	// categories[0] = "Kampanyalar"
	// categories[1] = "Temel Gıda"
	// categories[2] = "Sıvı Yağ & Margarin"
	// categories[3] = "Atıştırmalık"
	// categories[4] = "İçecek"
	// categories[5] = "Şarküteri & Kahvaltılık"
	// categories[6] = "Et Ürünleri ve Şarküteri"
	// categories[7] = "Unlu Mamüller"
	// categories[8] = "Bebek Ürünleri"
	// categories[9] = "Evcil Hayvan"
	// categories[10] = "Temizlik"

	// //fmt.Println(categories)

	// fmt.Println(categories[0:4])

	// var gender = [3]string{"erkek", "kadın", "çocuk"}

	// // fmt.Println(gender)

	// fmt.Println(gender[0:1])

	//slices : sınırlandırılmamış arraylardir.
	//append : slicese veri eklemek için kullanacağız.
	var users = []string{"ali", "ahmet", "mehmet", "mustafa", "rıza"}
	users = append(users, "oraz murat")
	users = append(users, "Muhammet murat")
	fmt.Println(users)

}
