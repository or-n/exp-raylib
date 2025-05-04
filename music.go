package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
	"unsafe"
)

func generateSineWave(frequency f32, duration f32, sampleRate i32) Wave {
	sampleCount := i32(f32(sampleRate) * duration)
	samples := make([]int16, sampleCount)
	for i := i32(0); i < sampleCount; i++ {
		time := float64(f32(i) / f32(sampleRate))
		samples[i] = int16(32767 * math.Sin(2.0*math.Pi*float64(frequency)*time))
	}
	byteData := (*[1 << 30]byte)(unsafe.Pointer(&samples[0]))[:len(samples)*2]
	return NewWave(u32(sampleCount), u32(sampleRate), 16, 1, byteData)
}

func generateBackgroundMusic(sampleRate i32) ([]Sound, Sound) {
	notes := []f32{
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
	rhythmInterval     f32
	melodyInterval     f32
	currentMelodyIndex int
	timeSinceRhythm    f32
	timeSinceMelody    f32
	melody             []Sound
	rhythm             Sound
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
		SetSoundVolume(rhythm, MusicVolume)
		PlaySound(rhythm)
		timeSinceRhythm = 0
	}
	if timeSinceMelody >= melodyInterval {
		SetSoundVolume(melody[currentMelodyIndex], MusicVolume)
		PlaySound(melody[currentMelodyIndex])
		currentMelodyIndex = (currentMelodyIndex + 1) % len(melody)
		timeSinceMelody = 0
	}
}
