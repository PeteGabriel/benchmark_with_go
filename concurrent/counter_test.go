package concurrent


import (
	is2 "github.com/matryer/is"
	"testing"
)

func TestCountWordsWithExpectedValue(t *testing.T){
	is := is2.New(t)
	expNum := 8

	num, err := CountWords(".././short.txt")
	is.NoErr(err)
	is.True(num == expNum)
}


func TestCountWordsWithNotFoundFile(t *testing.T){
	is := is2.New(t)

	_, err := CountWords(".././not_found_file.txt")
	is.True(err != nil)
}
