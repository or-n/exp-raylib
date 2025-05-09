package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
)

var gradients = [][2]f64{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
}

var (
	p       [512]int
	texture Texture2D
)

func NoiseInit() {
	perm := rand.Perm(256)
	for i := 0; i < 256; i++ {
		p[i] = perm[i]
		p[i+256] = perm[i]
	}
}

func fade(t f64) f64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func mix(a, b, t f64) f64 {
	return a + t*(b-a)
}

func grad(hash int, x, y f64) f64 {
	g := gradients[hash%len(gradients)]
	return g[0]*x + g[1]*y
}

func noise(x, y f64) f64 {
	X := int(math.Floor(x)) & 255
	Y := int(math.Floor(y)) & 255
	x -= math.Floor(x)
	y -= math.Floor(y)
	u := fade(x)
	v := fade(y)
	aa := p[p[X]+Y]
	ab := p[p[X]+Y+1]
	ba := p[p[X+1]+Y]
	bb := p[p[X+1]+Y+1]
	n00 := grad(aa, x, y)
	n01 := grad(ab, x, y-1)
	n10 := grad(ba, x-1, y)
	n11 := grad(bb, x-1, y-1)
	x1 := mix(n00, n10, u)
	x2 := mix(n01, n11, u)
	return mix(x1, x2, v)
}

func OctaveNoise(x, y f64, n int, persistence f64) f64 {
	amplitude := 1.0
	frequency := 1.0
	maxAmplitude := 0.0
	total := 0.0
	for _ = range n {
		total += noise(x*frequency, y*frequency) * amplitude
		maxAmplitude += amplitude
		amplitude *= persistence
		frequency *= 2
	}
	return total / maxAmplitude
}
