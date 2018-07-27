package main

// This file only tests the functions that are performing real work, not looking at pretty print functions.
import (
	"testing"
)

func TestReadFiles(t *testing.T) {
	err := ReadFiles("test_names.txt", "test_list.txt")
    if err != nil {
    	t.Errorf("ReadFiles received an error, got: %v, wanted: %v", err, nil)
    }
}