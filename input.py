from pyray import *
import util

class Input:
    change_action = None
    
    class DirBind:
        neg_y = KeyboardKey.KEY_W
        neg_x = KeyboardKey.KEY_A
        y = KeyboardKey.KEY_S
        x = KeyboardKey.KEY_D
        jump = KeyboardKey.KEY_W
        sprint = KeyboardKey.KEY_LEFT_SHIFT

    def update():
        if Input.change_action and util.LastPressed.key:
            setattr(Input.DirBind, Input.change_action, util.LastPressed.key)
            Input.change_action = None

    def draw(window):
        text_y = 15
        text_spacing = 2
        font = get_font_default()
        font_size = 20
        spacing = 20
        button_size = Vector2(300, 50)
        n = len(Input.actions)
        size = vector2_multiply(button_size, Vector2(1, n))
        size.y += spacing * (n - 1)
        start = vector2_scale(vector2_subtract(window, size), 0.5)
        title = "Configure Keys (Press to Change)"
        title_size = measure_text_ex(font, title, font_size, text_spacing)
        start_title = (window.x - title_size.x) * 0.5
        draw_text(title, int(start_title), int(start.y + text_y), font_size, WHITE)
        i = 1
        pad = 40
        for action in Input.actions:
            y = i * (button_size.y + spacing)
            r = Rectangle(start.x, start.y + y, button_size.x, button_size.y)
            draw_rectangle(int(r.x), int(r.y), int(r.width), int(r.height), BLUE)
            draw_text(action, int(start.x + pad), int(start.y + y + text_y), font_size, WHITE)
            key = getattr(Input.DirBind, action)
            name = "_" if Input.change_action == action else util.key_names[key]
            name_size = measure_text_ex(font, name, font_size, text_spacing)
            position = Vector2(start.x + r.width - name_size.x - pad, start.y + y + text_y)
            draw_text_ex(font, name, position, font_size, 1, WHITE)
            if check_collision_point_rec(get_mouse_position(), r):
                if is_mouse_button_pressed(MouseButton.MOUSE_BUTTON_LEFT):
                    Input.change_action = action
            i += 1

Input.actions = [key for key in Input.DirBind.__dict__ if not key.startswith('__')]
