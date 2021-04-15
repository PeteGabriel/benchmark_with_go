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

func TestCountWordsWithExpectedValue_2(t *testing.T){
	is := is2.New(t)
	expNum := 1000

	num, err := CountWords(".././one_thousand_words.txt")
	is.NoErr(err)
	is.True(num == expNum)
}

func TestCountWordsWithNotFoundFile(t *testing.T){
	is := is2.New(t)

	_, err := CountWords(".././not_found_file.txt")
	is.True(err != nil)
}

func benchmarkCountWords(fn string, b *testing.B) int {
	var r int
	for n := 0; n < b.N; n++ {
		r, _ = CountWords(fn)
	}
	return r
}

var result int
func BenchmarkCountWords1000(b *testing.B){
	result = benchmarkCountWords(".././one_thousand_words.txt", b)
}

func BenchmarkCountWords2000(b *testing.B){
	result = benchmarkCountWords(".././two_thousand_words.txt", b)
}

func BenchmarkCountWords3000(b *testing.B){
	result = benchmarkCountWords(".././three_thousand_words.txt", b)
}

func BenchmarkCountWords4000(b *testing.B){
	result = benchmarkCountWords(".././four_thousand_words.txt", b)
}

func BenchmarkCountWords5000(b *testing.B){
	result = benchmarkCountWords(".././five_thousand_words.txt", b)
}

func BenchmarkCountWords6000(b *testing.B){
	result = benchmarkCountWords(".././six_thousand_words.txt", b)
}

func BenchmarkCountWords7000(b *testing.B){
	result = benchmarkCountWords(".././seven_thousand_words.txt", b)
}

func BenchmarkCountWords8000(b *testing.B){
	result = benchmarkCountWords(".././eight_thousand_words.txt", b)
}

func BenchmarkCountWords9000(b *testing.B){
	result = benchmarkCountWords(".././nine_thousand_words.txt", b)
}

func BenchmarkCountWords10000(b *testing.B){
	result = benchmarkCountWords(".././ten_thousand_words.txt", b)
}