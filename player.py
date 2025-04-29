from pyray import *
from util import *
from shot import *
from enemy import *
from input import *

class Player:
    position = Vector2(0, 0)
    radius = 16
    size = Vector2(220, 165)
    frames = 29
    frame = 0
    delay = 0.05
    next_frame = 0
    heart = load_texture("asset/heart.png")
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
        if get_time() > Player.next_frame:
            Player.next_frame = get_time() + Player.delay
            Player.frame = (Player.frame + 1) % Player.frames

    def constrain(window):
        Shots.constrain(window)
        radiuses = (Player.radius, Player.radius)
        available = vector2_subtract(window, radiuses)
        new = vector2_clamp(Player.position, radiuses, available)
        d = vector2_distance_sqr(Player.position, new)
        play_or_stop(Player.collide, d > 0.1)
        Player.position = new

    def draw():
        rec = Rectangle(Player.frame * Player.size.x, 0, Player.size.x, Player.size.y)
        position = vector2_subtract(Player.position, vector2_scale(Player.size, 0.5))
        draw_texture_rec(Player.heart, rec, position, GREEN)
