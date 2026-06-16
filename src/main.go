// NOTES
// func (t *type) myFunc() -- * is a pointer, so it accesses underlying structure
// Pass-by-Value: When you pass a struct (not a pointer) to a function, 
// Go creates a copy. The original struct remains unchanged, 
// providing a local form of immutability.
// Wither methods: func (u User) WithName(n string) User
//
//  ptr := &p     // & means "address of"
// *x means pointer to object x


package main

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
)

type schema struct {

}

func readCSV(path string, hasHeader bool) {
	
	file, err := os.Open(path)
	
	// check if file is nil/neg before calling a deferred close
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if hasHeader {
		header, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(header)
	}

	// records, err := reader.ReadAll()
	// records = records[0:10]
	// fmt.Println(records)
	// if err != nil {
	// 	log.Fatal(err)
	// }


	// print out
	// for _, row := range records {
	// 	fmt.Println(row)
	// }

}

// func CSVtoTable(path string, schema schema) {

// }



func main() {
	pth := "/Users/noahalex/develop/esii/ercot-zp15-612-zips/TEST_FILE.csv"
	readCSV(pth, false)
	// CSVtoTable()
}

// func main() {
	// x := test("bobby")
	
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(x)
	// }

	// names := []string{"Noah", "Hannah", "kiddo"}

	// for index, name := range names {
	// 	fmt.Println(index, name)
	// }

	// ages := map[string]int{
	// 	"Noah": 27,
	// 	"Hannah": 27,
	// }

	// for name, age := range ages {
	// 	fmt.Println(name, age)
	// }


// }