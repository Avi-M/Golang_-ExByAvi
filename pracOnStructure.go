package main

import "fmt"

type Student struct {
	name   string
	rollno int
	class  string
	city   string
	pin    int
}

func main() {

	s1 := Student{name: "mahi reddy", rollno: 40, class: "B.Tech", city: "Azamgarh", pin: 474567}
	s2 := Student{"rahul ramapriya", 50, "BSC", "Ambedkar nagar", 224147}
	s3 := Student{"giloya singh", 55, "BA", "Ayodhya", 424568}
	s4 := Student{"shivam", 98, "BAMS", "varansi", 324589}
	mp := map[Student]int{s1: 1, s2: 2, s3: 3}
	mp[s4] = 4
	mp[Student{"Ram", 34, "B.Tech", "varansi", 224157}] = 5
	mp[Student{"Tina", 12, "B.Tech", "Mirjapur", 346578}]++
	fmt.Println(mp)
	for attr, val := range mp {
		fmt.Println(attr, val)
	}
	for attr := range mp {
		fmt.Println(attr.name, attr.class)
	}
	fmt.Println("The remaining map after deletion")
	delete(mp, s3)
	for attr, val := range mp {
		fmt.Println(attr, val)

	}
	_, check2 := mp[s2]
	fmt.Println("Is the key present:", check2)
}
