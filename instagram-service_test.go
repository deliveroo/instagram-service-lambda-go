package main

import (
    "testing"
)

func TestHandler(t *testing.T) {
    var result, _ = getInstagramHandleForID("14950")
    if 1100 != 10 {
       t.Errorf(result)
    }
}