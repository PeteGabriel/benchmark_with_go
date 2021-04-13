package sequential

import (
	is2 "github.com/matryer/is"
	"testing"
)

func TestCounterWithExpectedValue(t *testing.T){
	is := is2.New(t)
	expNum := 8

	num, err := Counter(".././short.txt")
	is.NoErr(err)
	is.True(num == expNum)
}


func TestCounterWithNotFoundFile(t *testing.T){
	is := is2.New(t)

	_, err := Counter(".././not_found_file.txt")
	is.True(err != nil)
}
