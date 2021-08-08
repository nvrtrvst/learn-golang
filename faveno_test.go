package main

import "testing"

var (
	persegi      Persegi = Persegi{Sisi: 4}
	luasExpected float64 = 16
)

func TestLuasPersegi(t *testing.T)  {
	t.Logf("Luas Persegi : %v", persegi.Luas())

	if persegi.Luas() != luasExpected {
		t.Errorf("Salah, seharusnya : %v", luasExpected)
	}
}

var (
	persegipanjang PersegiPanjang = PersegiPanjang{panjang: 2}
	persegipanjang1 PersegiPanjang = PersegiPanjang{panjang: 4}
	hasilExpected float64 = 

)

func TestPenjumlahan(t *testing.T)  {
	t.Logf("Penjumlahannya adalah :  %v", tambah())

	if tambah() != luasExpected {
		t.Errorf("Salah, seharusnya : %v", jmlExpected)
	}

	
}


func (pp PersegiPanjang) Luas() float64  {
	return 2 * (pp.Panjang * pp.Lebar)
}
