package controllers

import (
	"net/http"
	"testing"
)

func TestFindRoles(t *testing.T) {
	t.Log("Given")
	{
		t.Log("When")
		{
			resp, err := http.Get("localhost:8080/api/roles")
			if err != nil {
				t.Fatal("Should be")
			}
			defer resp.Body.Close()
		}
	}
}
