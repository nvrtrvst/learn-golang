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