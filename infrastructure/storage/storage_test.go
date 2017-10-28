package storage

import (
	"fmt"
	"testing"

	"github.com/alicebob/miniredis"
)

func TestSave(t *testing.T) {
	s, _ := miniredis.Run()
	defer s.Close()

	Save("key", "value")
	myMap := Retrieve("key")

	fmt.Println(myMap)
}

func TestSave2(t *testing.T) {
	s, _ := miniredis.Run()
	defer s.Close()

	Save("keyx", "valuex")
	Save("keyxx", "valuexx")
	Save("keyxxx", "valuexxx")
	Save("keyxxxx", "valuexxxx")
	myMap := Retrieve("*")

	fmt.Println(myMap)
}
