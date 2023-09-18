package main

import "fmt"

func main() {
	// var jumlahSiswa int = 10;
	var siswa [5]string;
	var newSiswa = [5]string{
		"halo",
		"halo",
		"bandung",		
	}

	siswa[1] = "John";
	siswa[2] = "English";

	fmt.Println(siswa[2]);
	fmt.Println(siswa[1]);

	fmt.Println(len(newSiswa[0]), newSiswa)
	newSiswa[0] = "haloha"
	fmt.Println(len(newSiswa[0]), len(newSiswa))
}