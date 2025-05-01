from pyray import *
from util import *
from shot import *
from enemy import *
from input import *
from camera import *

class Player:
    position = Vector2(0, -100)
    radius = 16
    size = Vector2(16, 32)
    bounds = Vector2(16, 32)
    frames = 1
    frame = 0
    delay = 0.05
    next_frame = 0
    texture = load_texture("asset/player.png")
    step = load_sound("asset/step.wav")
    collide = load_sound("asset/error_007.ogg")
    direction = Vector2(0, 0)
    jump_speed = 100
    gravity = 250
    is_grounded = False
    jump_to = None
    speed = Vector2(0, 0)

    def restart():
        stop_sound(Player.step)
        stop_sound(Player.collide)
        Player.position = Vector2(0, -100)
        Player.is_grounded = False
        Player.jump_to = None
        Player.speed = Vector2(0, 0)
        Player.direction = Vector2(0, 0)

    def update_direction():
        sprint = is_key_down(Input.DirBind.sprint)
        x = axis_x(Input.DirBind) * (400 if sprint else 200)
        if Player.jump_to:
            y = -Player.jump_speed
        elif Player.is_grounded:
            if is_key_down(Input.DirBind.jump):
                Player.is_grounded = False
                Player.jump_to = Player.position.y - 1.25 * 16
            y = 0
        else:
            y = Player.gravity
        Player.direction = Vector2(x, y)
        
    def update(dt):
        Player.update_direction()
        if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
            mouseWorldPos = get_screen_to_world_2d(get_mouse_position(), Camera.camera)
            d = vector2_subtract(mouseWorldPos, Player.position)
            mouse_dir = vector2_normalize(d)
            shot = Shot(Player.position, mouse_dir)
            Shots.new(shot)
        Player.speed = vector2_scale(Player.direction, dt)
        d = vector2_length_sqr(Player.speed)
        play_or_stop(Player.step, d > 0.1)
        Player.position = vector2_add(Player.position, Player.speed)
        if Player.jump_to and Player.position.y < Player.jump_to:
            Player.jump_to = None
        if get_time() > Player.next_frame:
            Player.next_frame = get_time() + Player.delay
            Player.frame = (Player.frame + 1) % Player.frames

    def constrain(map):
        # Shots.constrain(window)
        # half = vector2_scale(Player.bounds, 0.5)
        # available = vector2_subtract(window, half)
        # new = vector2_clamp(Player.position, half, available)
        # d = vector2_distance_sqr(Player.position, new)
        # play_or_stop(Player.collide, d > 0.1)
        # Player.position = new
        rect = Player.get_rect()
        (rect, is_grounded, jump_stop) = map.collide(rect, Player.direction, Player.is_grounded)
        Player.set_rect(rect)
        Player.is_grounded = is_grounded
        if jump_stop:
            Player.jump_to = None

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
        if Player.is_grounded:
            color = GREEN
        elif Player.jump_to:
            color = RED
        else:
            color = WHITE
        draw_texture_rec(Player.texture, rec, position, color)
        end = vector2_add(Player.position, vector2_scale(Player.direction, 1))
        draw_line(int(Player.position.x), int(Player.position.y), int(end.x), int(end.y), WHITE)
