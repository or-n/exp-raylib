from pyray import *
from util import *

class Shot:
    speed = 8
    radius = 8
    collide = load_sound("asset/error_007.ogg")

    def __init__(self, position, direction):
        self.position = position
        self.direction = direction
        self.alive = True

    def update(self):
        change = vector2_scale(self.direction, Shot.speed)
        self.position = vector2_add(self.position, change)
    
    def constrain(self, window):
        radiuses = (Shot.radius, Shot.radius)
        available = vector2_subtract(window, radiuses)
        new = vector2_clamp(self.position, radiuses, available)
        d = vector2_distance_sqr(self.position, new)
        if d > 0.1:
            play_sound(Shot.collide)
            self.alive = False

    def draw(self):
        draw_circle(int(self.position.x), int(self.position.y), Shot.radius, WHITE)

class Shots:
    xs = []

    def new(x):
        Shots.xs.append(x)

    def update():
        for x in Shots.xs:
            x.update()
        Shots.xs = [x for x in Shots.xs if x.alive]

    def constrain(window):
        for x in Shots.xs:
            x.constrain(window)

    def draw():
        for x in Shots.xs:
            x.draw()
        
