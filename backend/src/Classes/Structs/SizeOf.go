package structs

func SizeOfSuperBlock() int {
	return 68
}

func SizeOfJournal() int {
	return 212
}

func SizeOfInode() int {
	return 88
}

func SizeOfBlockPointers() int {
	return 64
}

func SizeOfBlockFolder() int {
	return 64
}

func SizeOfBlockFile() int {
	return 64
}
