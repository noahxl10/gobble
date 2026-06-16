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
	// "reflect"
	"unicode/utf8"
	// "unsafe"
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
	} else {
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		numRecords := 100
		records = records[0:numRecords]
		
		sumSize := 0
		for _, row := range records {
			sumSize += calculateArraySize(row, false)
		}
		fmt.Println(sumSize/numRecords)
	}
}

func calculateArraySize(row []string, assumeASCII bool) int {
	// Calculates byte size of string array
	size := 0 

	for i, cell := range row {
		if i > 0 {
				size += 1 // add comma seperator byte, assume always ASCII
		}
		if assumeASCII {
			size += len(cell)
		} else {
			for _, c := range cell {
				bytes := utf8.RuneLen(c)
				// fmt.Printf("%c -> %d bytes\n", c, bytes)
				size += bytes
			}
		}
	}
	return size
}


func CSVtoTable(path string, schema schema) {
	
}



func main() {
	pth := "/Users/noahalex/develop/esii/ercot-zp15-612-zips/TEST_FILE.csv"
	readCSV(pth, false)
	// x := "10204049746606825 2703 BARON AVE  ABILENE TX 79606 TAYLOR 007923311 06 De-Energized Residential ERCOT TWILIGHT TWILIGHT Y  Residential Y AMSR N  "
	// fmt.Println(len(x))
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