package private
import (
	"encoding/binary"
	"errors"
	"fmt"
)

// s3 
//   /data
//     {table}/
//			segment.dat
//  /wal
// 
//  /indexes
// Reads are expensive in s3, but data storage is cheap... so page address is segment_id + page_number

const DEFAULT_SEGMENT_SIZE = 128,000,000 // BYTES, so 128MB default size

const DEFAULT_PAGE_SIZE = 4096 // BYTES - default size is 4kB
const DEFAULT_PAGE_HEADER_SIZE = 6 // bytes
const NUM_CELLS_HEADER_SPACE = 4 // bytes 1-4 are reserved for num cells
const FREE_SPACE_START_HEADER_SPACE = 2 // bytes 5-6 are reserved for free space location


type pageSegment struct {
	data [DEFAULT_SEGMENT_SIZE]byte
	// [ header_bytes ] [ free_space ] [ page1_bytes ] [ page2_bytes ]
	// max pages per segment = 128MB/4kB = ~31000 pages
	// so, page_count = 31k max (uint16)
	// free space start = uint32 (need to be able to reference up to 128 million bytes)
	// header = [bytes 1 - 2 for page count, bytes 3-6 for free space ref]
	//
}



// Stealing SQLite's "Page"-style blocks
type Page struct {
	data [DEFAULT_PAGE_SIZE]byte
	// Layout of data:
	// [ header_bytes ][ cell_pointer_array_bytes ][ free_space_bytes ][ record_data_bytes ]
	// header: 
	//	[b,b,...,b,b] -- first bytes say number of cells
	// second part of header says when free space starts (if there is free space in the block/page)
	// Records go at the back because moving pointers is cheap, but moving the entire block is much more expensive
	// So essentially space gets allocated to the block first, then gets filled up back to front

	// Cell pointer array is references to the data
}

type Record struct {
	id uint64,
	name string // do I need name?
}

// page build process: 
// Insert row
// Find page
// Does it fit?
//   Yes -> insert
//   No  -> split page

// what is a cell?
// - Key
// - Payload
// - Metadata

// methods belonging to the type/struct
// func (pg *Page) insertRecord(rc Record) error {

// }
func (pg *Page) fillPage() {
}

func (pg *Page) setHeader(rowCount uint32, freeSpaceStart uint16) {
	// Since row count could go beyond 64k (unlikely), support uint32, which is 4 bytes
	binary.BigEndian.PutUint16(pg.data[0:4], rowCount)
	// free space only needs to be a number 0-4096, essentially, since it is just a reference to which byte it starts at
	// uint16 is max 64k
	binary.BigEndian.PutUint32(pg.data[4:6], freeSpaceStart)
}

// 
func newPage() *Page {
	// declare and initialize a new page
	// * pointer, & is address of the object
	pg := &Page{}
	pg.setHeader()
	pg.
}


// // B tree storage

// type bTree struct {
// 	leftNode bTree,
// 	rightNode bTree,
// }

// func buildBTree {

// }

// // PAGES
// // Page Header
// // Cell Pointer Array
// // Record Data
// // Free Space

// type cell 

// type record struct {

// }

// type page struct {
// 	header,
// 	*cells []cell,

// }

// struct rawTable {

// }