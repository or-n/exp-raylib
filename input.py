from pyray import *
import util

class Input:
    change_action = None
    
    class DirBind:
        neg_y = KeyboardKey.KEY_W
        neg_x = KeyboardKey.KEY_A
        y = KeyboardKey.KEY_S
        x = KeyboardKey.KEY_D

    def update():
        if Input.change_action != None and util.LastPressed.key != None:
            setattr(Input.DirBind, Input.change_action, util.LastPressed.key)
            Input.change_action = None

    def draw():
        draw_text("Configure Keys (Press to Change)", 20, 20, 20, WHITE)
        y = 80
        for action in Input.actions:
            r = Rectangle(30, y - 5, 300, 25)
            key = getattr(Input.DirBind, action)
            name = "_" if Input.change_action == action else util.key_names[key]
            draw_rectangle(int(r.x), int(r.y), int(r.width), int(r.height), BLUE)
            draw_text(action, 40, y, 20, WHITE)
            draw_text(name, int(r.width - 40), y, 20, WHITE)
            if check_collision_point_rec(get_mouse_position(), r):
                if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
                    Input.change_action = action
            y += 40

Input.actions = [key for key in Input.DirBind.__dict__ if not key.startswith('__')]
