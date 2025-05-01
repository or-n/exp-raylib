from pyray import *
import random
import util

import window
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
        Menu.draw()
        end_drawing()
    elif Menu.state == State.GAME:
        Player.update(dt)
        Player.constrain(map)
        Camera.camera.target = Player.position
        Enemies.update(Player.position, dt)
        Shots.update(dt)
        #Shots.constrain(window)
        for enemy in Enemies.xs:
            for shot in Shots.xs:
                d = vector2_distance_sqr(enemy.position, shot.position)
                if d < (Enemy.radius + Shot.radius) ** 2:
                    enemy.alive = shot.alive = False
                    stop_sound(enemy.step)
        if random.random() < 0.02:
            Enemies.new(Enemy.arbitrary())
        begin_drawing()
        clear_background(Color(127, 31, 255))
        draw_fps(30, 30)
        begin_mode_2d(Camera.camera)
        Player.draw()
        Enemies.draw()
        Shots.draw()
        map.draw()
        draw_border(Vector2(0, 0), window.size, RED)
        end_mode_2d()
        end_drawing()
    elif Menu.state == State.OPTIONS:
        begin_drawing()
        clear_background(Color(127, 31, 255))
        Menu.draw()
        end_drawing()
    elif Menu.state == State.EXIT:
        break
    LastPressed.update()
    Input.update()
    Camera.update()
    Bg.update()
close_window()
