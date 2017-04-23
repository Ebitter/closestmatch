package levenshtein

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/schollz/closestmatch/test"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(test.WordsToTest)
	}
}

func BenchmarkClosestOne(b *testing.B) {
	bText, _ := ioutil.ReadFile("../test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest)
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func ExampleMatching() {
	cm := New(test.WordsToTest)
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'emma by jane austen'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'hard times by charles dickens' matched 'hard times by charles dickens'
	// 'complete william shakespeare' matched 'the iliad by homer'
	// 'war by hg wells' matched 'beowulf'

}

func TestAccuray(t *testing.T) {
	cm := New(test.WordsToTest)
	accuracy := cm.Accuracy()
	if accuracy > 50 {
		t.Errorf("Accuracy should be higher than it usually is! %2.1f", accuracy)
	}
}
