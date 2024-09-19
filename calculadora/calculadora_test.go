package main

import (
	"testing"
)

func TestShouldSumCorrect(t *testing.T){

	result := sum(10.0, 10.0)

	expected := 20.0

	if result != expected{
		t.Error("expected value: ", expected, "value result: ", result)
	}

}

func TestShouldSumIncorrect(t *testing.T){

	result := sum(10.0, 10.0)

	expected := 21.0

	if result != expected{
		t.Error("expected sum value: ", expected, "value result: ", result)
	}

}

func TestShouldSubtractCorrect(t *testing.T){

	result := subtract(10.0, 2.0)

	expected := 8.0

	if result != expected{
		t.Error("expected subtracted value: ", expected, "value result: ", result)
	}

}

func TestShouldSubtractIncorrect(t *testing.T){

	result := subtract(10.0, 3.0)

	expected := 5.0

	if result != expected{
		t.Error("expected subtracted value: ", expected, "value result: ", result)
	}

}


func TestShouldDivisionCorrect(t *testing.T){

	result := division(10.0, 10.0)

	expected := 1.0

	if result != expected{
		t.Error("expected division value: ", expected, "value result: ", result)
	}

}

func TestShouldDivisionIncorrect(t *testing.T){

	result := division(10.0, 10.0)

	expected := 2.0

	if result != expected{
		t.Error("expected multiplication value: ", expected, "value result: ", result)
	}

}

func TestShouldMultiplyCorrect(t *testing.T){

	result := multiply(10.0, 10.0)

	expected := 100.0

	if result != expected{
		t.Error("expected multiplication value: ", expected, "value result: ", result)
	}

}

func TestShouldMultiplyIncorrect(t *testing.T){

	result := multiply(10.0, 10.0)

	expected := 150.0

	if result != expected{
		t.Error("expected multiplication value: ", expected, "value result: ", result)
	}

}