from pyray import *
import math

class Camera:
    zoom = 1
    camera = Camera2D(Vector2(0, 0), Vector2(0, 0), 0, zoom)

    def restart():
        Camera.camera = Camera2D(Vector2(0, 0), Vector2(0, 0), 0, Camera.zoom)

    def update(window):
        wheel = get_mouse_wheel_move()
        #mouseWorldPos = get_screen_to_world_2d(get_mouse_position(), Camera.camera)
        Camera.camera.offset = vector2_scale(window, 0.5)
        scale = 0.2 * wheel
        zoom = math.exp(math.log(Camera.camera.zoom) + scale)
        Camera.camera.zoom = clamp(zoom, 0.125, 64.0)