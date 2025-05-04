package main

import (
	"math"
)

type f32 = float32
type f64 = float64

type i32 = int32
type i64 = int64

type u32 = uint32
type u64 = uint64

func Round(x f32) f32 {
	return f32(math.Round(f64(x)))
}

func BoolI32(x bool) i32 {
	var result i32
	if x {
		result = 1
	}
	return result
}
