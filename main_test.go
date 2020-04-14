package main

import (
  "testing"
)

func TestParsingNumberFunction(t *testing.T) {

	numbersAsString := "10.2,19,31.3"
	expected := []float32{10.2,19,31.3}
	actual := ParseNumbers(numbersAsString)
  
	for i, _ := range actual {
	  if actual[i] != expected[i] {
		  
		  t.Errorf("Expected: %v but got %v", expected, actual)
	  }
	}
  }

// TestParsingNumberFunctionWithString tests if function ignores text string in array
func TestParsingNumberFunctionWithString(t *testing.T) {

	numbersAsString := "batata,19,31.3"
	expected := []float32{19,31.3}
	actual := ParseNumbers(numbersAsString)
  
	for i, _ := range actual {
	  if actual[i] != expected[i] {
		  
		  t.Errorf("Expected: %v but got %v", expected, actual)
	  }
	}
  }
  
// TestParsingNumberFunctionWithNegative tests if function works with negative numbers
func TestParsingNumberFunctionWithNegative(t *testing.T) {

	numbersAsString := "-20,19,31.3"
	expected := []float32{-20,19,31.3}
	actual := ParseNumbers(numbersAsString)
  
	for i, _ := range actual {
	  if actual[i] != expected[i] {
		  
		  t.Errorf("Expected: %v but got %v", expected, actual)
	  }
	}
  }
  