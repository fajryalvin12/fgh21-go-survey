package project

import "fmt"

func Survey() {
	fmt.Println("Siapa Nama anda?")
	var name string 
	fmt.Scanln(&name)
	fmt.Println("Berapa umur anda?")
	var age string 
	fmt.Scanln(&age)
	fmt.Println("Apa jenis kelamin anda?")
	var gender string 
	fmt.Scanln(&gender)
	fmt.Println("Apakah anda perokok?")
	var isSmoker string
	fmt.Scanln(&isSmoker)
	variant := ("Jenis rokok yang pernah anda coba?")
	cigarVariant := []string{}
	cont := true
	for cont {
		if cont == false {
			fmt.Println("Apakah anda akan menginputkan data lagi?")
		} else {
			cigarVariant = append(cigarVariant, variant)
		}
	}
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("gender:", gender)
	fmt.Println("isSmoker:", isSmoker)
	fmt.Println("cigarVariant:", cigarVariant)
}