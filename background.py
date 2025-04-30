from pyray import *
import math
import numpy as np
import ctypes

def generate_sine_wave(frequency, duration, sample_rate):
    sample_count = int(sample_rate * duration)
    samples = np.zeros(sample_count, dtype=np.int16)
    for i in range(sample_count):
        time = i / sample_rate
        samples[i] = int(32767 * math.sin(2.0 * math.pi * frequency * time))
    audio_pointer = ffi.new("short[]", samples.tolist())
    return Wave(sample_count, sample_rate, 16, 1, audio_pointer)

def generate_background_music(sample_rate):
    notes = [261.63, 293.66, 329.63, 349.23, 392.00, 440.00, 493.88, 523.25]
    melody = []
    for freq in notes:
        wave = generate_sine_wave(freq, 0.5, sample_rate)
        sound = load_sound_from_wave(wave)
        melody.append(sound)
    rhythm = generate_sine_wave(100.0, 0.2, sample_rate)
    rhythm_sound = load_sound_from_wave(rhythm)
    return melody, rhythm_sound

class Bg:
    rhythm_interval = 0.5
    melody_interval = 0.5
    current_melody_index = 0
    time_since_rhythm = 0
    time_since_melody = 0
    melody, rhythm = generate_background_music(44100)

    def update():
        Bg.time_since_rhythm += get_frame_time()
        Bg.time_since_melody += get_frame_time()
        if Bg.time_since_rhythm >= Bg.rhythm_interval:
            play_sound(Bg.rhythm)
            Bg.time_since_rhythm = 0
        if Bg.time_since_melody >= Bg.melody_interval:
            play_sound(Bg.melody[Bg.current_melody_index])
            Bg.current_melody_index = (Bg.current_melody_index + 1) % len(Bg.melody)
            Bg.time_since_melody = 0
