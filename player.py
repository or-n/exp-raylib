from pyray import *
from util import *
from shot import *
from enemy import *

class Player:
    speed = 360
    position = Vector2(0, 0)
    radius = 16
    step = load_sound("asset/step.wav")
    collide = load_sound("asset/error_007.ogg")
    change_bind = None

    class DirBind:
        neg_y = KeyboardKey.KEY_W
        neg_x = KeyboardKey.KEY_A
        y = KeyboardKey.KEY_S
        x = KeyboardKey.KEY_D

    def update(dt):
        Player.speed = 720 if is_key_down(KeyboardKey.KEY_LEFT_SHIFT) else 360
        direction = get_direction(Player.DirBind)
        if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
            d = vector2_subtract(get_mouse_position(), Player.position)
            mouse_dir = vector2_normalize(d)
            shot = Shot(Player.position, mouse_dir)
            Shots.new(shot)
        delta = vector2_scale(direction, Player.speed * dt)
        d = vector2_length_sqr(delta)
        play_or_stop(Player.step, d > 0.1)
        Player.position = vector2_add(Player.position, delta)
        action = Player.change_bind
        if action != None and util.LastPressed.key != None:
            setattr(Player.DirBind, action, util.LastPressed.key)
            Player.change_bind = None

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
        draw_text("Configure Keys (Press to Change)", 20, 20, 20, WHITE)
        y = 80
        for action in Player.actions:
            r = Rectangle(30, y - 5, 300, 25)
            key = getattr(Player.DirBind, action)
            name = "_" if Player.change_bind == action else util.key_names[key]
            draw_rectangle(int(r.x), int(r.y), int(r.width), int(r.height), BLUE)
            draw_text(action, 40, y, 20, WHITE)
            draw_text(name, int(r.width - 40), y, 20, WHITE)
            if check_collision_point_rec(get_mouse_position(), r):
                if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
                    Player.change_bind = action
            y += 40

Player.actions = [key for key in Player.DirBind.__dict__ if not key.startswith('__')]
