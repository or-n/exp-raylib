from pyray import *

def play_or_stop(sound, should_play):
    playing = is_sound_playing(sound)
    if should_play:
        if not playing:
            play_sound(sound)
    else:
        if playing:
            stop_sound(sound)

def axis_x(bind):
    return int(is_key_down(bind.x)) - int(is_key_down(bind.neg_x))

def axis_y(bind):
    return int(is_key_down(bind.y)) - int(is_key_down(bind.neg_y))

def draw_text_center(font, window, text, size, color):
    text_size = measure_text_ex(font, text, size, 1)
    available = vector2_subtract(window, text_size)
    position = vector2_scale(available, 0.5)
    draw_text_ex(font, text, (int(position.x), int(position.y)), size, 1, color)

def draw_split_x(window, color):
    draw_line(int(window.x / 2), 0, int(window.x / 2), int(window.y), color)

def draw_split_y(window, color):
    draw_line(0, int(window.y / 2), int(window.x), int(window.y / 2), color)

def draw_border(center, window, color):
    start = vector2_subtract(center, vector2_scale(window, 0.5))
    draw_line(int(start.x), int(start.x), int(start.x) + int(window.x), int(start.x), color)
    draw_line(int(start.x), int(start.x) + int(window.y), int(start.x) + int(window.x), int(start.x) + int(window.y), color)
    draw_line(int(start.x), int(start.x), int(start.x), int(start.x) + int(window.y), color)
    draw_line(int(start.x) + int(window.x), int(start.x), int(start.x) + int(window.x), int(start.x) + int(window.y), color)

key_names = {
    getattr(KeyboardKey, name): name.replace("KEY_", "")
    for name in dir(KeyboardKey) if name.startswith("KEY_")
}

class LastPressed:
    key = None

    def update():
        LastPressed.key = None
        for key in KeyboardKey:
            if is_key_pressed(key):
                LastPressed.key = key
        
