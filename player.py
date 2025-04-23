from pyray import *
from util import *
from shot import *
from enemy import *
from input import *

class Player:
    position = Vector2(0, 0)
    radius = 16
    step = load_sound("asset/step.wav")
    collide = load_sound("asset/error_007.ogg")

    def update(dt):
        speed = 720 if is_key_down(KeyboardKey.KEY_LEFT_SHIFT) else 360
        direction = get_direction(Input.DirBind)
        if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
            d = vector2_subtract(get_mouse_position(), Player.position)
            mouse_dir = vector2_normalize(d)
            shot = Shot(Player.position, mouse_dir)
            Shots.new(shot)
        delta = vector2_scale(direction, speed * dt)
        d = vector2_length_sqr(delta)
        play_or_stop(Player.step, d > 0.1)
        Player.position = vector2_add(Player.position, delta)

    def constrain(window):
        Shots.constrain(window)
        radiuses = (Player.radius, Player.radius)
        available = vector2_subtract(window, radiuses)
        new = vector2_clamp(Player.position, radiuses, available)
        d = vector2_distance_sqr(Player.position, new)
        play_or_stop(Player.collide, d > 0.1)
        Player.position = new

    def draw():
        p = Player.position
        draw_circle(int(p.x), int(p.y), Player.radius, GREEN)
