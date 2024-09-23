package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	var count int
	if fileContent, ok := cb[file]; ok {
		for _, x := range fileContent {
			if x {
				count++
			}
		}
	}
	return count

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int
	files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	for _, file := range files {
		if fileContent, ok := cb[file]; ok {
			for i, x := range fileContent {
				if i+1 == rank && x {
					count++
				}
			}
		}
	}
	return count

}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
	files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	for _, file := range files {
		for range cb[file] {
			count++
		}
	}
	return count

}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	for _, file := range files {
		if fileContent, ok := cb[file]; ok {
			for _, x := range fileContent {
				if x {
					count++
				}
			}
		}
	}
	return count
}
