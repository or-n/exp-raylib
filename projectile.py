from pyray import *
from util import *

class Projectile:
    speed = 8
    radius = 8
    collide = load_sound("asset/error_007.ogg")

    def __init__(self, position, direction):
        self.position = position
        self.direction = direction
        self.alive = True

    def update(self):
        change = vector2_scale(self.direction, Projectile.speed)
        self.position = vector2_add(self.position, change)
    
    def constrain(self, window):
        radiuses = (Projectile.radius, Projectile.radius)
        available = vector2_subtract(window, radiuses)
        new = vector2_clamp(self.position, radiuses, available)
        d = vector2_distance_sqr(self.position, new)
        if d > 0.1:
            play_sound(Projectile.collide)
            self.alive = False

    def draw(self):
        draw_circle(int(self.position.x), int(self.position.y), Projectile.radius, WHITE)

class Projectiles:
    xs = []

    def new(x):
        Projectiles.xs.append(x)

    def update():
        for x in Projectiles.xs:
            x.update()
        Projectiles.xs = [x for x in Projectiles.xs if x.alive]

    def constrain(window):
        for x in Projectiles.xs:
            x.constrain(window)

    def draw():
        for x in Projectiles.xs:
            x.draw()
        
