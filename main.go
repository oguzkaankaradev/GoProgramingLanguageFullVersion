package main

import (
	"fmt"
	// "github.com/gbrlsnchs/jwt"
)

// Her bir dosyamız mutlaka bir paket altında çalışması gerekmetedir.
// Pketler bizim doslarımızı gruplandırmak için go tarafından sunulan hizmettir.
// > go mod init name = bir paketi modüle çevirir.
func main() {
	fmt.Println("Üye girişiniz başarılı")
	fmt.Println("Profil Bilgileri")
	fmt.Println("Notification")
	fmt.Println("Ana sayfa")

	// jwt.HS256("IDCARTNUMBER")

}
