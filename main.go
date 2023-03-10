package main

import (
	"fmt"
	"net/http"

	"github.com/tuotoo/qrcode"
)

func main() {
	http.HandleFunc("/scan", scanQRCode)
	http.ListenAndServe(":8080", nil)
}

func scanQRCode(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get the image file
	file, _, err := r.FormFile("image")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Decode the image file into an image.Image object
	qr, err := qrcode.Decode(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Print the result to the console and return it as the HTTP response
	fmt.Println(qr.Content)
	w.Write([]byte(qr.Content))
}
