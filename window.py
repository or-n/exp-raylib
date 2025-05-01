from pyray import *

init_audio_device()
init_window(1920, 1080, "Hello")
toggle_fullscreen()
size = Vector2(get_monitor_width(0), get_monitor_height(0))
