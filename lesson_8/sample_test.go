package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup main")
	res := m.Run()
	fmt.Println("tear-down main")

	os.Exit(res)
}

func TestMultiple(t *testing.T) {
	// setup
	// insert test data within db
	t.Run("groupA", func(t *testing.T) {
		t.Log("A")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})

	t.Run("groupB", func(t *testing.T) {
		t.Log("B")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})

	t.Run("groupC", func(t *testing.T) {
		t.Log("C")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})

	t.Run("groupD", func(t *testing.T) {
		t.Log("D")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})
	// tearDown
	// delete test data within db
}

// func TestAdd(t *testing.T) {
// 	t.Run("simple", func(t *testing.T) {
// 		var x, y, result = 2, 2, 4

// 		realResult := Add(x, y)

// 		if realResult != result {
// 			t.Errorf("expected result: %d != %d real result", result, realResult)
// 		}

// 		t.Run("1", func(t *testing.T) {
// 			r := Add(1, 1)
// 			if r != 2 {
// 				t.Errorf("failed")
// 			}
// 		})
// 	})
// }
