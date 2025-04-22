from pyray import *
from player import *
import util

class Enemy:
    speed = 1
    radius = 16
    collide = load_sound("asset/error_007.ogg")

    def __init__(self, position):
        self.position = position
        self.alive = True
        self.step = load_sound("step.wav")

    def update(self, target):
        if not self.alive:
            return
        delta = vector2_subtract(target, self.position)
        util.play_or_stop(self.step, vector2_length_sqr(delta) > Enemy.speed ** 2)
        direction = vector2_normalize(delta)
        change = vector2_scale(direction, Enemy.speed)
        self.position = vector2_add(self.position, change)

    def constrain(self, window):
        for x in Projectiles.xs:
            d = vector2_distance_sqr(self.position, x.position)
            if d < (Enemy.radius + Projectile.radius) ** 2:
                self.alive = x.alive = False
                stop_sound(self.step)

    def draw(self):
        draw_circle(int(self.position.x), int(self.position.y), Enemy.radius, RED)

class Enemies:
    xs = []

    def new(x):
        Enemies.xs.append(x)

    def update(target):
        for x in Enemies.xs:
            x.update(target)
        Enemies.xs = [x for x in Enemies.xs if x.alive]
    
    def constrain(window):
        for x in Enemies.xs:
            x.constrain(window)
            
    def draw():
        for x in Enemies.xs:
            x.draw()
