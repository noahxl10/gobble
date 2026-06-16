// write

func createPage()

func findPage()

func checkWAL() {
	// Check size
	// Merge if needed
}


// Need write ahead log to never die
// So, write, keep in mem 
func touchWAL() {

}

func verifyWAL(walHash) {
	// compare data in written WAL
	// use checksum	
}


func createIndex() {

}

func createTable(name string, schema schema) boolean {

}

func insertIntoTable(table table, data ) {
	// touch WAL
	// create blocks
	// write blocks
	// checksum on blocks
}

func modifyTable() {
	// Touch WAL
	// Verify WAL
	// Finish transaction and rollback if fail
}


func joinTables() {

}
// read

func readTable(tableId, columns)


// EXECUTION PLAN??