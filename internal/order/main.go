package main

import (
	"github.com/mfzzf/gorder-v2/common/config"
	"github.com/spf13/viper"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	str := viper.Get("order")
	value, _ := str.(map[string]interface{})
	log.Printf("%v", value["http-addr"])
	//log.Println("Listening on 8082")
	//mux := http.NewServeMux()
	//
	//mux.HandleFunc("/ping",
	//	func(w http.ResponseWriter, r *http.Request) {
	//		log.Printf("%v", r.RequestURI)
	//		_, _ = io.WriteString(w, "<h1>Welcome To Home Page</h1>")
	//	})
	//
	//if err := http.ListenAndServe(":8082", mux); err != nil {
	//	log.Fatal(err)
	//}

}
