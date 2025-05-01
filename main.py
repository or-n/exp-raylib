from pyray import *
import random
import util
import math

init_audio_device()
init_window(1920, 1080, "Hello")
toggle_fullscreen()
window = Vector2(get_monitor_width(0), get_monitor_height(0))

from menu import *
from player import *
from enemy import *
from shot import *
from input import *
from background import *
from map import *
from restart import *

map = Map()

set_target_fps(600)
set_exit_key(0)

while not window_should_close():
    if is_key_down(KeyboardKey.KEY_ESCAPE):
        Menu.state = State.MENU
    dt = get_frame_time()
    if Menu.state == State.MENU:
        begin_drawing()
        clear_background(Color(127, 31, 255))
        draw_fps(30, 30)
        Menu.draw(window)
        end_drawing()
    elif Menu.state == State.GAME:
        Player.update(dt)
        rect = Player.get_rect()
        (rect, is_grounded, jump_stop) = map.collide(rect, Player.direction, Player.is_grounded)
        Player.set_rect(rect)
        Player.is_grounded = is_grounded
        if jump_stop:
            Player.jump_to = None
        Enemies.update(Player.position, dt)
        Shots.update(dt)
        #Player.constrain(window)
        #Enemies.constrain(window)
        #Shots.constrain(window)
        # if random.random() < 0.002:
        #     value = random.random()
        #     side = random.choice([0, 1])
        #     x_or_y = random.choice([Vector2(side, value), Vector2(value, side)])
        #     position = vector2_multiply(window, x_or_y)
        #     Enemies.new(Enemy(position))
        begin_drawing()
        clear_background(Color(127, 31, 255))
        draw_fps(30, 30)
        begin_mode_2d(Camera.camera)
        Player.draw()
        #Enemies.draw()
        Shots.draw()
        map.draw()
        draw_border(Vector2(0, 0), window, RED)
        end_mode_2d()
        end_drawing()
    elif Menu.state == State.OPTIONS:
        begin_drawing()
        clear_background(Color(127, 31, 255))
        Menu.draw(window)
        end_drawing()
    elif Menu.state == State.EXIT:
        break
    LastPressed.update()
    Input.update()
    Camera.update(window)
    Bg.update()
close_window()
