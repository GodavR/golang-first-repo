/* menjalankan dengan cara masuk ke dir dengan comment dengan perintah go run test.go
perintah init pada awal project dengan go mod init <nama-project>Setiap file program harus memiliki package. Setiap project harus ada minimal
satu file dengan nama package main . File yang ber-package main , akan
dieksekusi pertama kali ketika program di jalankan.
Cara menentukan package dengan menggunakan keyword package , berikut
adalah contoh penulisannya.
package <nama-package>
package main */

package main

import "fmt"

/* Keyword import digunakan untuk meng-import atau memasukan package lain
ke dalam file program, agar isi dari package yang di-import bisa dimanfaatkan.
Package fmt merupakan salah satu package bawaan yang disediakan oleh Go,
isinya banyak fungsi untuk keperluan I/O yang berhubungan dengan text.
Berikut adalah skema penulisan keyword import :
import "<nama-package>"
import "fmt" */

func main() {
	/*komentar code
	menamppilkan pesan hello word */
	fmt.Println("Hello, World!")
	//fmt.Println("Hello, World!") baris ini tidak akan dieksekusi
}
