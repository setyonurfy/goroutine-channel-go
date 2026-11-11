package goroutine_go

import (
	"reflect"
	"testing"
)

func TestChannelDasar(t *testing.T) {
	ch := make(chan string)

	go func() {
		ch <- "halo"
	}()

	got := <-ch
	if got != "halo" {
		t.Errorf("dapa %q, harusnya %q", got, "halo")
	}
}

func TestChannelBuffered(t *testing.T) {
	ch := make(chan int, 2) // kapasitas channel

	ch <- 1
	ch <- 2

	if len(ch) != 2 {
		t.Fatalf("isi channel %d, harusnya 2", len(ch))
	}
	if got := <-ch; got != 1 {
		t.Errorf("dapat %d, harusanya 1", got)
	}
	if got := <-ch; got != 2 {
		t.Errorf("dapat %d harusnya 2", got)
	}
}

func TestKuadrat(t *testing.T) {
	var got []int
	for hasil := range Kuadrat([]int{1, 2, 3, 4}) {
		got = append(got, hasil)
	}

	want := []int{1, 4, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dapat %v, harusnya %v", got, want)
	}
}

func TestGenap(t *testing.T) {
	tests := []struct {
		nama  string
		input []int
		want  []int
	}{
		{"campuran",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{2, 4, 6, 8}},
		{"tanpa genap",
			[]int{1, 3, 5},
			nil},
		{"negatif dan nol",
			[]int{-4, -3, 0, 7},
			[]int{-4, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.nama, func(t *testing.T) {
			var got []int
			for n := range Genap(tt.input) {
				got = append(got, n)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dapat %v, harusnya %v", got, tt.want)
			}
		})
	}
}
