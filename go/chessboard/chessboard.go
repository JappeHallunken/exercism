package chessboard


// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int
	for _, v := range cb[file] {
		if v {
			counter++
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var counter int
	for _, v := range cb {
		for j, w := range v {
			if j+1 == rank && w {
				counter++
			}
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
  var counter int
  for i := range cb {
    for  range cb[i] {
      counter++
    }
  }
  return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
  var counter int
  for _, v := range cb {
    for _, w := range v {
      if w {
        counter ++
      }
    }
  }
  return counter
}
