package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
	"unsafe"
)

func generateSineWave(frequency float32, duration float32, sampleRate int32) Wave {
	sampleCount := int32(float32(sampleRate) * duration)
	samples := make([]int16, sampleCount)
	for i := int32(0); i < sampleCount; i++ {
		time := float64(float32(i) / float32(sampleRate))
		samples[i] = int16(32767 * math.Sin(2.0 * math.Pi * float64(frequency) * time))
	}
	byteData := (*[1 << 30]byte)(unsafe.Pointer(&samples[0]))[:len(samples)*2]
	return NewWave(uint32(sampleCount), uint32(sampleRate), 16, 1, byteData)
}

func generateBackgroundMusic(sampleRate int32) ([]Sound, Sound) {
	notes := []float32{
		261.63, 293.66, 329.63, 349.23, 392.00, 440.00, 493.88, 523.25,
	}
	melody := make([]Sound, 0, len(notes))
	for _, freq := range notes {
		wave := generateSineWave(freq, 0.5, sampleRate)
		sound := LoadSoundFromWave(wave)
		melody = append(melody, sound)
	}
	rhythmWave := generateSineWave(100.0, 0.2, sampleRate)
	rhythmSound := LoadSoundFromWave(rhythmWave)
	return melody, rhythmSound
}

var (
	rhythmInterval      float32
	melodyInterval      float32
	currentMelodyIndex  int
	timeSinceRhythm     float32
	timeSinceMelody     float32
	melody              []Sound
	rhythm              Sound
)

func MusicInit() {
	rhythmInterval = 0.5
	melodyInterval = 0.5
	melody, rhythm = generateBackgroundMusic(44100)
}

func MusicUpdate() {
	dt := GetFrameTime()
	timeSinceRhythm += dt
	timeSinceMelody += dt
	if timeSinceRhythm >= rhythmInterval {
		PlaySound(rhythm)
		timeSinceRhythm = 0
	}
	if timeSinceMelody >= melodyInterval {
		PlaySound(melody[currentMelodyIndex])
		currentMelodyIndex = (currentMelodyIndex + 1) % len(melody)
		timeSinceMelody = 0
	}
}
