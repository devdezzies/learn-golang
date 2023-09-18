package main;

import "fmt";

func main(){
	// single variable declaration
	var nama = "Abdullah";

	// multiple-variable declaration
	var (
		firstName = "Hello";
		lastName = "world";
		age int = 20;

		nilai8 int8 = 123;
	)
	// constant
	const usia int = 19;
	const usia1 = 19;

	// multiple-constant declaration

	const (
		initialValue = 1;
		initialBinding = 2;
	)

	var nilaiBaru = int64(nilai8);
	var name byte = lastName[0];
	var convertedName string = string(name);
	var intToString string = string(age);
	var byteToInt int = int(name);

	fmt.Println(firstName, lastName, age, nama);
	fmt.Println(initialBinding + initialValue);
	fmt.Println(nilaiBaru);
	fmt.Println(name,"->", convertedName);
	fmt.Println(age, intToString);
	fmt.Println(string(byteToInt + byteToInt), int(byteToInt + byteToInt), "123" + "123");

	// Alias
	type status bool;

	var bersekolah status = true;
	fmt.Println(bersekolah);

	// Augmented Assignments
	var (
		ten = 10;
		five = 5;
		iniBoolean = true;
	)
	fmt.Println(ten, five, iniBoolean);

	ten += 10; //20
	five++;
	fmt.Println(ten, -five, !iniBoolean);
}

// Questions
// 1. Ketika suatu variable dengan tipe data integer dideklarasikan lalu dikonversi ke string, maka nilai tersebut akan dibaca sebagai byte (karakter)
// sementara itu, tipe data byte tidak bisa dideklarasikan sebagai integer. Lalu bagaimana jika kita ingin mengkonversi data integer ke string (untuk keperluan concatenation)