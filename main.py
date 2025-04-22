from pyray import *
import random
import util

init_audio_device()

from player import *
from enemy import *
from projectile import *

window = Vector2(1920, 1080)
Player.position = vector2_scale(window, 0.5)

init_window(int(window.x), int(window.y), "Hello")
set_target_fps(60)
font = get_font_default()

while not window_should_close():
    LastPressed.update()
    Player.update()
    Enemies.update(Player.position)
    Projectiles.update()
    Player.constrain(window)
    Enemies.constrain(window)
    Projectiles.constrain(window)
    if random.random() < 0.002:
        value = random.random()
        side = random.choice([0, 1])
        x_or_y = random.choice([Vector2(side, value), Vector2(value, side)])
        position = vector2_multiply(window, x_or_y)
        Enemies.new(Enemy(position))

    begin_drawing()
    clear_background(Color(127, 31, 255))
    Player.draw()
    Enemies.draw()
    Projectiles.draw()
    draw_split_x(window, RED)
    draw_split_y(window, RED)
    end_drawing()
close_window()
