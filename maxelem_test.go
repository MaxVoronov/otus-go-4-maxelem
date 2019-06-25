package maxelem

import "testing"

func TestFindMax(t *testing.T) {
	slice := []int{41, 72, 19, 36, 65}
	expected := 72

	result, err := FindMax(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if result != expected {
		t.Fatalf("Invalid max value in slice: expected %d, got %d", expected, result)
	}
}

func TestFindMaxInt(t *testing.T) {
	expected := 101
	result, err := FindMaxInt([]int{89, 22, 17, 35, 99, 0, 72, 101, 44})

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if result != expected {
		t.Fatalf("Invalid max value in int slice: expected %d, got %d", expected, result)
	}
}

func TestFindMaxFloat(t *testing.T) {
	var expected float32 = 1.208
	result, err := FindMaxFloat([]float32{0.137, 1.208, 0.993, 1.201, 0.135})

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if result != expected {
		t.Fatalf("Invalid max value in float slice: expected %f, got %f", expected, result)
	}
}

func TestFindMaxString(t *testing.T) {
	expected := "zz"
	result, err := FindMaxString([]string{"af", "se", "qx", "zz", "em", "gg"})

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if result != expected {
		t.Fatalf("Invalid max value in string slice: expected %s, got %s", expected, result)
	}
}
