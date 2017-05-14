// olb.go
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	var resp string
	resp = ""

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Error during getting hostname.\n")
	} else {
		resp += PadRight("Hostname:", " ", 30) + hostname + "\n"
	}

	resp += PadRight("-", "-", 80) + "\n"

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(w, PadRight("Error during getting network interfaces.", " ", 30)+"\n")
	} else {
		for _, intf := range interfaces {
			addrs, err := intf.Addrs()
			if err != nil {
				fmt.Fprintf(w, PadRight("Cannot get addresses for interface:", " ", 30)+intf.Name+".\n")
			} else {
				for _, addr := range addrs {
					resp += PadRight("Interface "+intf.Name+":", " ", 30) + addr.String() + "\n"
				}
			}
		}
	}

	resp += PadRight("-", "-", 80) + "\n"

	for k, v := range r.Header {
		resp += PadRight(k, " ", 30) + strings.Join(v, " ") + "\n"
	}

	fmt.Fprintf(w, resp)

}

func main() {

	// init webserver
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
