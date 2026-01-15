package books;

import "testing";

func TestShowTitle(t *testing.T) {
	if got, want := showTitle("Go, Go, Golang"), "Title of this poem is Go, Go, Gophers"; got != want {
		t.Errorf("\n got =>>> %q\nwant =>>> %q", got, want);
	}
}