package main

import "fmt"

const (
	Malestr = "Male"
	Girlstr = "Girl"
)

type Person struct {
		Name      string
		Age       int
		Genre     string
		Paremetrs map[string]int
	}

func main() {
	

	Vanka := Person{
		Name:      "ivan",
		Age:       18,
		Genre:     Malestr,
		Paremetrs: map[string]int{"Weighr": 62, "height": 180},
	}

	fmt.Println("Age: ", Vanka.Age)

	Vanka.AgeInc()

	fmt.Println("Age: ", Vanka.Age)

	a:= 100
	Inc(&a)
	fmt.Println(a)
}
func Inc(p *int){
	*p++
}
// func Canculetet(a, b int, op string)(float64, error){
// 	// if op =="+"{
// 	// 	return float64(a + b), nil;
// 	// }
// 	// if op == "-"{
// 	// 	return float64(a - b), nil;
// 	// }
// 	// if op == "*"{
// 	// 	return float64(a * b), nil
// 	// }else{
// 	// 	return -1, fmt.Errorf("fignya")
// 	// }

// }


func (p Person) AgeInc(){
	p.Age++
	fmt.Println(p.Age)
}