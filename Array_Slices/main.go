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

	// var jordans [6]string

	// jordans[0] = "All Jordan"
	// jordans[1] = "New Jordan"
	// jordans[2] = "Jorden Shoes"
	// jordans[3] = "Jorden 3"
	// jordans[4] = "Jorden 4"
	// jordans[5] = "Jorden 5"

	// fmt.Println(jordans)      //tamamı yazıldı
	// fmt.Println(jordans[0:4]) //kardeşim bana jordan 3 e kadar yazdır. bu veritabanından ayarla
	// fmt.Println(jordans[1:4]) //kardeşim bana jordan 3 e kadar yazdır. bu veritabanından ayarla

	var kids [8]string

	kids[0] = "Infant & Toddler Shoes"
	kids[1] = "Kids' Shoes"
	kids[2] = "Kids' Jordan Shoes"
	kids[3] = "Kids' Basketball Shoes"
	kids[4] = "Kids' Running Shoes"
	kids[5] = "Kids' Clothing"
	kids[6] = "Kids' Backpacks"
	kids[7] = "Kids' Backpacks"

	fmt.Println(kids) //tamamı yazıldı

	// fmt.Println(jordans)      //tamamı yazıldı
	// fmt.Println(jordans[0:4]) //kardeşim bana jordan 3 e kadar yazdır. bu veritabanından ayarla
	// fmt.Println(jordans[1:4]) //kardeşim bana jordan 3 e kadar yazdır. bu veritabanından ayarla

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
