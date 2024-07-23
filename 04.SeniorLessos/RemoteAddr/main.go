package main

import (
	"fmt"
	"net"
	"net/http"
)

//RemoteAddr : Kullanıcının ip adresini almak için kullanılmaktadır. Bu bilgi, güvenlik kontrolleri, loglama ve istemci trafiğini izleme gibi durumlarda faydalı olabilir.
//RemoteAddr, HTTP isteğinin http.Request nesnesinde yer alan bir alan olup,
// gelen bağlantının adresini temsil eder. Genellikle istemcinin IP adresini ve bağlantı noktasını içerir,
// ve bu alan net.Conn nesnesi tarafından sağlanır.
//IP adresi ve portu ayırmak için net.SplitHostPort fonksiyonunu kullanabilirsiniz:

func handler(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.RemoteAddr
	ip, port, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("IP: %s, Port: %s\n", ip, port)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
