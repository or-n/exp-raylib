from pyray import *
import random

class Map:
    max_x = 100
    max_y = 10
    dirt = load_texture("asset/dirt.png")
    size = Vector2(16, 16)

    def __init__(self):
        self.map = [[0 for i in range(Map.max_x)] for j in range(Map.max_y)]
        self.center = Vector2(0, 0)
        for y in range(Map.max_y):
            for x in range(Map.max_x):
                self.map[y][x] = random.choice([0, 1])

    def collide(self, rect, direction, is_grounded):
        jump_stop = False
        for y in range(Map.max_y):
            for x in range(Map.max_x):
                if self.map[y][x] == 1:
                    rec = Rectangle(self.center.x + x * Map.size.x, self.center.y + y * Map.size.y, Map.size.x, Map.size.y)
                    if check_collision_recs(rec, rect):
                        if rect.x + rect.width > rec.x + rec.width and direction.x < 0:
                            rect.x = rec.x + rec.width
                        if rect.y + rect.height > rec.y + rec.height and direction.y < 0:
                            rect.y = rec.y + rec.height
                            jump_stop = True
                        if rect.x < rec.x and direction.x > 0:
                            rect.x = rec.x - rect.width
                        if rect.y < rec.y and direction.y > 0:
                            rect.y = rec.y - rect.height
                            is_grounded = True
        return (rect, is_grounded, jump_stop)

    def draw(self):
        for y in range(Map.max_y):
            for x in range(Map.max_x):
                if self.map[y][x] == 1:
                    rec = Rectangle(0, 0, Map.size.x, Map.size.y)
                    position = Vector2(self.center.x + x * Map.size.x, self.center.y + y * Map.size.y)
                    draw_texture_rec(Map.dirt, rec, position, WHITE)
