package main

import (
	"fmt"
	"math"
)

type Person struct {
	name    string
	age     int
	address string
}

type Employee struct {
	name     string
	address  string
	position string
}

// // func main() {
// // 	angka := 100
// // 	fmt.Print("Nomor Favorit saya adalah: ", rand.Intn(angka))
// // }

// // func main() {
// // 	var i, j int = 1, 2
// // 	var c, python bool
// // 	c = true
// // 	python = false
// // 	var java string = "nol"

// // 	fmt.Println(i, j, c, python, java)
// // }

// // function akar di parsing ke float, karena sqrt returnnya float
// // func main(){
// // 	x := float64(3)
// // 	y := float64(4)
// // 	z := math.Sqrt(x)
// // 	u := math.Sqrt(y)
// // 	fmt.Println(z, u)

// //  }

// // func main () {
// // i, j := 42, 2701
// // po := &i
// // fmt.Println(*po)
// // *po = 21
// // fmt.Println(po)
// // po = &j
// // j = (*po / 37)
// // fmt.Println(j)

// // }

// // func main() {
// // 	var score [3]int = [3]int{40, 50, 60}

// // 	for i := 0; 1 < len(score); i++ {
// // 		fmt.Println(score[i])

// // 	}

// // 	for _, value := range score {
// // 		fmt.Println(value)
// // 	}
// // }

// func main() {
// 	// uhuy := "aku"
// 	// fmt.Println(uhuy)

// 	// mahasiswa := []string{"Agus","Reza","Muhammad", "Akbar", "Rais"}
// 	// band := make ([]string, len(mahasiswa))

// 	// fmt.Println("Data Mahasiswa => ", mahasiswa)
// 	// fmt.Println("Data Mahasiswa => ", band)
// 	// fmt.Println("Data Mahasiswa => ", mahasiswa[0:1])
// 	// fmt.Println("Data Mahasiswa => ", len(mahasiswa))

// 	// copy(band, mahasiswa)
// 	// fmt.Println(band)

// 	// bandbaru := "tdwp"
// 	// semuaband := append(band, bandbaru)
// 	// fmt.Println("woy", semuaband)

// 	// //  ==============SLICE================

// 	// var name[3]string
// 	// name[0] = "Reza"
// 	// name[1] = "Mike"
// 	// name[2] = "hranice"

// 	// fmt.Println(name[0], name[1])

// 	// for i, v := range name {
// 	// 	fmt.Println("Index ke ", i, "=>", v)
// 	// }
// 	// 	// } for range

// 	// for index := 0; index < len(name); index++ {
// 	// 	fmt.Println("Index ke ", index, "=>", name[index])
// 	// }

// 	// name[len(name)-3] = ""

// 	// for index := 0; index < len(name); index++ {
// 	// 	fmt.Println("Index ke ", index, "=>", name[index])
// 	// }
// 	// // for

// 	// var nama = []string{"Reza", "Muhammad"}
// 	// fmt.Println("Nama Awal :", nama)
// 	// var potongnama = nama[0:1]
// 	// fmt.Println("Nama dipotong :", potongnama)

// 	// var tambahnama = append(potongnama, "Muhammad", "Akbar")
// 	// fmt.Println("Nama ditambah", tambahnama)

// 	// var tambahangka int = append(potongnama, 2)
// 	// fmt.Println(tambahangka)

// 	// 	var count []int
// 	// 	for i := 0; i < 10; i++ {
// 	// 		count = append(count, i)
// 	// 	}
// 	// 	fmt.Println(count)

// 	// 	//milsakan ada data buah dibawah ini
// 	// 	// Semangka, Pisang, Apel, Belimbing
// 	// 	// Simpan data array tersebut dalam array dinamis
// 	// 	// loop data array diatas dgn loop
// 	// 	//print loop array setelah di append di looping

// 	// := shorthand

// 	// var buah [4]string = [4]string{"Semangka","Pisang", "Apel", "Belimbing"}
// 	// var result []string // array dinamis adalah slices
// 	// for _, value := range buah {
// 	// 	result = append(result, value)
// 	// }

// 	// fmt.Println(result)
// 	// fmt.Println(len(result))

// 	// var ages map[string]int
// 	//ages = map[string]int{}
// 	// ages := make(map[string]int)

// 	// ages["Reza"] = 25
// 	// ages["muhammad"] = 27
// 	// //fmt.Println(ages)
// 	// for key, value := range ages {
// 	// 	fmt.Println(key, value)
// 	// }

// 	// var person []Person
// 	// person = append(person, Person{
// 		var arrPerson []Person
// 	arrPerson = append(arrPerson, Person{// array
// 		name:    "Reza",
// 		age:     25,
// 		address: "bandung",
// 	})
// 	fmt.Println(arrPerson)

// 	// var person Person
// 	// person.name ="Muhammad"
// 	// person.age = 25
// 	// person.address = "Bandung"

// 	//buat struct employee data
// 	//employee struct isi atribut name, position
// 	// buat array dari struct employee
// 	// isi array struct
// 	// print struct employee

// 	var arrEmployee []Employee
// 	arrEmployee = append(arrEmployee, Employee{
// 		name: "Reza",
// 		address: "Bandung",
// 		position: "staf",
// 	})
// 	fmt.Println(arrEmployee)
// 	fmt.Println(arrEmployee)

// 	for _, value := range arrEmployee {
// 		fmt.Println(value.name)
// // 	}
// 	func ShowName(name string) string {
// 		return name
// 	}

// 	func addition(numberOne int, numberTwo int)  {
// 		fmt.Println(numberOne + numberTwo) //returnnya
// 	}

// 	func main() {
// 	fmt.Println(ShowName("Reza"))
// 	addition(5, 6)
// 	}

// func tambah(angka1 int, angka2 int) int {
// 	a := angka1 + angka2
// 	return a
// }

// func kurang(angka1 int, angka2 int) int {
// 	a := angka1 - angka2
// 	return a
// }

// func kali(angka1 int, angka2 int) int {
// 	a := angka1 * angka2
// 	return a
// }

// func bagi(angka1 int, angka2 int) int {
// 	a := angka1 / angka2
// 	return a
// }

// func mod(angka1 int, angka2 int) int {
// 	a := angka1 % angka2
// 	return a
// }

// func main() {
// 	angka1 := 20
// 	angka2 := 5

// 	hasil := tambah(angka1, angka2)
// 	fmt.Println("Penjumlahan Antara", angka1, "+", angka2, "Adalah", hasil)

// 	hasilKurang := kurang(angka1, angka2)
// 	fmt.Println("Pengurangan Antara", angka1, "-", angka2, "Adalah", hasilKurang)

// 	hasilKali := kali(angka1, angka2)
// 	fmt.Println("Perkalian Antara", angka1, "*", angka2, "Adalah", hasilKali)

// 	hasilBagi := bagi(angka1, angka2)
// 	fmt.Println("Pembagian Antara", angka1, "/", angka2, "Adalah", hasilBagi)

// 	hasilMod := mod(angka1, angka2)
// 	fmt.Println("Modulasi Antara", angka1, "%", angka2, "Adalah", hasilMod)
// }

// func akar(angka1, angka2 int) (int, error) {
// 	return angka1 * angka2, nil
// }

// func main ()  {
// 	angka1 := 20
// 	angka2 := 20
// 	result, err := akar()

// 	if err != nil {
// 		return err, nil
// 	}

// 	hasil := akar(angka1, angka2)
// 	fmt.Println(hasil)
// } BELUM

// type Persegi struct {
// 	Sisi float64
// }

// //traditional function
// func Luas(sisi Persegi) float64 {
// 	return sisi.Sisi * sisi.Sisi
// }

// //method with receiver
// func (s Persegi) Luas() float64{
// 	return s.Sisi * s.Sisi
// }


// func main() {
// 	persegi1 := Persegi{6}
// 	fmt.Println(Luas(persegi1))
// 	fmt.Println(persegi1.Luas())

// }

// p x l

//method
// type PersegiPanjang struct {
// 	Panjang float64
// 	Lebar float64
// }

// //keliling
// func (pp PersegiPanjang) Keliling() float64{
// 	return 2 *  (pp.Panjang * pp.Lebar)
// }

// func (pp PersegiPanjang) Luas() float64  {
// 	return pp.Panjang * pp.Lebar
// }

// func main()  {
// 	pp := PersegiPanjang{Panjang: 4, Lebar: 5}
// 	fmt.Println("Kel : ", pp.Keliling())
// 	fmt.Println("Luas : ", pp.Luas())

	
// }

// type Lingkaran struct {
// 	Jari2 float64
// 	Diameter float64
	
// }

// func (l Lingkaran) LuasLingkaran() float64  {
// 	const PI float64 = 3.14
// 	return math.PI * l.Jari2
	
// }

// func main()  {
// 	l := Lingkaran{Panjang: 4, Lebar: 5} // belum

//interface //global method
type BangunDatar interface {
	Luas()
	Keliling()
}

type Persegi struct {
	Sisi float64
}

type PersegiPanjang struct {
	Panjang float64
	Lebar float64
}

type Lingkaran struct {
	Diameter float64
}

func (l Lingkaran) keliling() float64  {
	return math.Round(math.Pi * l.Diameter)
}

func (pp PersegiPanjang) Keliling() float64  {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Luas() float64  {
	return 2 * (pp.Panjang * pp.Lebar)
}

func (p Persegi) Luas() float64  {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

func hitungPersegi(p Persegi)  {
	fmt.Println(p.Keliling())
	fmt.Println(p.Luas())
}

func hitungPersegiPanjang(pp PersegiPanjang)  {
	fmt.Println(pp.Keliling())
	fmt.Println(pp.Luas())
}

func hitungLingkaran(l Lingkaran)  {
	fmt.Println(l.keliling())
}

func main()  {
	hitungPersegi(Persegi{Sisi: 6})
	hitungPersegiPanjang(PersegiPanjang{Panjang: 6, Lebar: 7})
	hitungLingkaran(Lingkaran{Diameter: 6})
	
}




// import "fmt"


// 	// fmt.Println("hello world")

// 	// var name string
// 	// name = "Reza"

// 	// fmt.Println("Reza" + name)

// 	// var age int 
// 	// age = 20

	// fmt.Println(age)
	// // var isPredicate bool
	// // isPredicate = false

	// fmt.Println("umur saya", age)

	// // // fmt.Println("Nama Saya: " + name + "umur saya" +age)

	// // // count := 20 //shorthand variabel

	// artist := "blackpink"

	// if artist == "Blackpink" {
	// 	fmt.Println("Korea")
	// } else {
	// 	fmt.Println("Not Korea")
	// }

	// switch artist {
	// case "Blackpink":
	// 	fmt.Println("Korea")
	// 	break
	// }

	// for i := 0; i < 10; i++{
	// 	fmt.Println(i)
	// }

	// // array := ["Apple", "Manggo"]
	// // for x, v := range array{
	// // 	fmt.Printf()

	// // } 

// 	// sum := 1
// 	// for sum < 20 {
// 	// 	sum += sum
// 	// }
// 	// fmt.Println(sum)

// 	number := 20

// 	pointerNumb := &number // alokasi data ke dalam virtual memori
// 	fmt.Println(*pointerNumb) // cara akses valuenya *+variabel
