package models

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {

	alist := GetAllArticles()

	// Check that the lenght of the list of articles returned is the
	// same as the lenght of the global variable holding the list
	if len(alist) != len(Articlelist) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != Articlelist[i].Content ||
			v.Id != Articlelist[i].Id ||
			v.Title != Articlelist[i].Title {
			t.Fail()
			break
		}
	}

}
