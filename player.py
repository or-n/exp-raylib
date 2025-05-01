from pyray import *
from util import *
from shot import *
from enemy import *
from input import *

class Player:
    position = Vector2(0, 0)
    radius = 16
    size = Vector2(16, 16)
    bounds = Vector2(16, 16)
    frames = 1
    frame = 0
    delay = 0.05
    next_frame = 0
    heart = load_texture("asset/dirt.png")
    step = load_sound("asset/step.wav")
    collide = load_sound("asset/error_007.ogg")
    direction = Vector2(0, 0)
    jump_speed = 100
    gravity = 500
    is_grounded = False

    def update(dt, camera):
        speed = 720 if is_key_down(KeyboardKey.KEY_LEFT_SHIFT) else 360
        x = axis_x(Input.DirBind)
        x *= speed * dt
        y = 0
        if Player.is_grounded:
            y = axis_y(Input.DirBind)
            if y > 0:
                y = 0
            if y < 0:
                y *= Player.jump_speed
                Player.is_grounded = False
        else:
            y += Player.gravity * dt
        Player.direction = Vector2(x, y)
        if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
            mouseWorldPos = get_screen_to_world_2d(get_mouse_position(), camera)
            d = vector2_subtract(mouseWorldPos, Player.position)
            mouse_dir = vector2_normalize(d)
            shot = Shot(Player.position, mouse_dir)
            Shots.new(shot)
        delta = Player.direction
        d = vector2_length_sqr(delta)
        play_or_stop(Player.step, d > 0.1)
        Player.position = vector2_add(Player.position, delta)
        if get_time() > Player.next_frame:
            Player.next_frame = get_time() + Player.delay
            Player.frame = (Player.frame + 1) % Player.frames

    def constrain(window):
        Shots.constrain(window)
        half = vector2_scale(Player.bounds, 0.5)
        available = vector2_subtract(window, half)
        new = vector2_clamp(Player.position, half, available)
        d = vector2_distance_sqr(Player.position, new)
        play_or_stop(Player.collide, d > 0.1)
        Player.position = new

    def get_rect():
        half = vector2_scale(Player.size, 0.5)
        position = vector2_subtract(Player.position, half)
        return Rectangle(position.x, position.y, Player.size.x, Player.size.y)
    
    def set_rect(rect):
        half = vector2_scale(Player.size, 0.5)
        Player.position = vector2_add(Vector2(rect.x, rect.y), half)

    def draw():
        rec = Rectangle(Player.frame * Player.size.x, 0, Player.size.x, Player.size.y)
        half = vector2_scale(Player.size, 0.5)
        position = vector2_subtract(Player.position, half)
        draw_texture_rec(Player.heart, rec, position, WHITE)
