package goin

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNextString(t *testing.T) {
	sc = bufio.NewScanner(strings.NewReader("hello1\nhello2"))
	result := NextString()
	if result != "hello1" {
		t.Errorf("NextString() = %s; want 'hello1'", result)
		return
	}
	result = NextString()
	if result != "hello2" {
		t.Errorf("NextString() = %s; want 'hello2'", result)
	}
}

func TestNextStrings(t *testing.T) {
	sc = bufio.NewScanner(strings.NewReader("hello world\nfoo\nbar\n"))
	res := NextStrings(3)
	expected := []string{"hello world", "foo", "bar"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func ExampleNextString() {
	sc = bufio.NewScanner(strings.NewReader("hello world"))
	word := NextString()
	fmt.Println(word)
	// Output: hello world
}
