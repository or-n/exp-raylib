from pyray import *
import random
import util

init_audio_device()

from player import *
from enemy import *
from shot import *
from input import *

init_window(1920, 1080, "Hello")
toggle_fullscreen()
window = Vector2(get_monitor_width(0), get_monitor_height(0))
Player.position = vector2_scale(window, 0.5)

set_target_fps(180)
font = get_font_default()

while not window_should_close():
    dt = get_frame_time()
    Input.update()
    LastPressed.update()
    Player.update(dt)
    Enemies.update(Player.position, dt)
    Shots.update(dt)
    Player.constrain(window)
    Enemies.constrain(window)
    Shots.constrain(window)
    if random.random() < 0.002:
        value = random.random()
        side = random.choice([0, 1])
        x_or_y = random.choice([Vector2(side, value), Vector2(value, side)])
        position = vector2_multiply(window, x_or_y)
        Enemies.new(Enemy(position))

    begin_drawing()
    clear_background(Color(127, 31, 255))
    draw_fps(20, 0)
    Player.draw()
    Enemies.draw()
    Shots.draw()
    Input.draw()
    draw_split_x(window, RED)
    draw_split_y(window, RED)
    end_drawing()
close_window()
