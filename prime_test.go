package prime

import (
  "testing"
  )

func TestPrimeOneShouldBeOne(t *testing.T) {
  factor := prime(1)
  if factor != "1" {
    t.Error("It should be 1 but got ", factor)
  }
}

func TestPrimeTwoShouldBeTwo(t *testing.T) {
  factor := prime(2)
  if factor != "2" {
    t.Error("It should be 2 but got ", factor)
  }
}

func TestPrimeThreeShouldBeThree(t *testing.T) {
  factor := prime(3)
  if factor != "3" {
    t.Error("It should be 3 but got ", factor)
  }
}

func TestPrimeFourShouldBeTwoTwo(t *testing.T) {
  factor := prime(4)
  if factor != "2,2" {
    t.Error("It should be 2,2 but got ", factor)
  }
}
