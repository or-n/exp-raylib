from pyray import *
from enum import Enum
from input import *

class State(Enum):
    MENU = 1
    GAME = 2
    OPTIONS = 3
    EXIT = 4

class Menu:
    start = "start"
    options = "options"
    exit = "exit"
    state = State.MENU
    button = Vector2(200, 100)

    def draw(window):
        if Menu.state == State.MENU:
            gui_set_style(GuiControl.DEFAULT, GuiDefaultProperty.TEXT_SIZE, 20)
            x = (window.x - Menu.button.x) * 0.5
            y = (window.y - Menu.button.y * 3) * 0.5
            rect = Rectangle(x, y, Menu.button.x, Menu.button.y)
            if gui_button(rect, Menu.start) == 1:
                Menu.state = State.GAME
            rect = Rectangle(x, y + Menu.button.y, Menu.button.x, Menu.button.y)
            if gui_button(rect, Menu.options) == 1:
                Menu.state = State.OPTIONS
            rect = Rectangle(x, y + Menu.button.y * 2, Menu.button.x, Menu.button.y)
            if gui_button(rect, Menu.exit) == 1:
                Menu.state = State.EXIT
        if Menu.state == State.OPTIONS:
            Input.draw()