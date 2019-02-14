package mylib

import (
  "testing"
)

func TestAverage(t *testing.T)  {
  v := Average([]int{1, 2, 3, 4, 5})
  if v != 3{
    t.Error("３が来ると思っていたけど", v)
  }

}
