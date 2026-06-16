// Core types


type createdModifiedMetadata struct {
	createdBy int,
	lastModifiedBy int
}

type timestampMetadata struct {
	createdAt int,
	lastModifiedAt int
}

type tablePermissions struct {
	read: int,
	write: int,
	modify: int,
	delete: int
}

type role struct {
	id int,
	name string,
	tablePermissions tablePermissions,
	timestampMetadata timestampMetadata
}

type user struct {
	id int,
	name string,
	roles []role,
	timestampMetadata timestampMetadata,
	createdModifiedMetadata createdModifiedMetadata
}



type table struct {
	id int, // random id
	name string, // table name: enforcements - no special chars, lowercase, no spaces
	mutable int, // bool 0/1 for if it is mutable
	size int, // mutable
	*roles []roles,
	*timestampMetadata timestampMetadata,
	*createdModifiedMetadata createdModifiedMetadata
}