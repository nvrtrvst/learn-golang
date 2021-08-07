package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

type Employee struct {
	name    string
	address string
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
// 	}


	func add(x int, y int) int {
	return x + y
	}

	func main() {
	fmt.Println(add(42, 13))
	}





