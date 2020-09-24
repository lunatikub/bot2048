#!/usr/bin/env python3
# coding: utf-8

from enum import Enum
import collections

SZ=4

class Direction(Enum):
    RIGHT = 1
    LEFT = 2
    UP = 3
    DOWN = 4


def right(y, x, dm, ds):
    """
    move to right direction
    """
    x += 1
    if x == SZ:
        y = (y + 1) if ds == Direction.DOWN else (y - 1)
        x = SZ - 1
        dm = Direction.LEFT
    return y, x, dm


def left(y, x, dm, ds):
    """
    move to left direction
    """
    x -= 1
    if x == -1:
        y = (y + 1) if ds == Direction.DOWN else (y - 1)
        x = 0
        dm = Direction.RIGHT
    return y, x, dm


def up(y, x, dm, ds):
    """
    move to up direction
    """
    y -= 1
    if y == -1:
        x = (x + 1) if ds == Direction.RIGHT else (x - 1)
        y = 0
        dm = Direction.DOWN
    return y, x, dm


def down(y, x, dm, ds):
    """
    move to down direction
    """
    y += 1
    if y == SZ:
        x = (x + 1) if ds == Direction.RIGHT else (x - 1)
        y = SZ - 1
        dm = Direction.UP
    return y, x, dm


def nextTile(tile, dm, ds):
    """
    Get the coordinate of the next tile in terms of the directions.
    dm is the main direction and ds the secondary direction.
    """
    x = tile[1]
    y = tile[0]
    if dm == Direction.RIGHT:
        y, x, dm = right(y, x, dm, ds)
    elif dm == Direction.LEFT:
        y, x, dm = left(y, x, dm, ds)
    elif dm == Direction.UP:
        y, x, dm = up(y, x, dm, ds)
    else:
        y, x, dm = down(y, x, dm, ds)
    return (y, x), dm, ds


def getMonotonicPath(tile, dm, ds):
    path = []
    for _ in range(SZ * SZ):
        path.append(tile)
        tile, dm, ds = nextTile(tile, dm, ds)
    return path


data=[
    #    [0][1][2][3]
    # [0] O--------+
    #              |
    # [1] +--------+
    #     |
    # [2] +--------+
    #              |
    # [3] <--------+
    [(0, 0), Direction.RIGHT, Direction.DOWN],

    #    [0][1][2][3]
    # [0] <--------+
    #              |
    # [1] +--------+
    #     |
    # [2] +--------+
    #              |
    # [3] O--------+
    [(3, 0), Direction.RIGHT, Direction.UP],

    #    [0][1][2][3]
    # [0] +--------O
    #     |
    # [1] +--------+
    #              |
    # [2] +--------+
    #     |
    # [3] +-------->
    [(0, 3), Direction.LEFT, Direction.DOWN],

    #    [0][1][2][3]
    # [0] +-------->
    #     |
    # [1] +--------+
    #              |
    # [2] +--------+
    #     |
    # [3] +--------O
    [(3, 3), Direction.LEFT, Direction.UP],


    #    [0][1][2][3]
    # [0] O  +--+  ^
    #     |  |  |  |
    # [1] |  |  |  |
    #     |  |  |  |
    # [2] |  |  |  |
    #     |  |  |  |
    # [3] +--+  +--+
    [(0, 0), Direction.DOWN, Direction.RIGHT],

    #    [0][1][2][3]
    # [0] ^  +--+  O
    #     |  |  |  |
    # [1] |  |  |  |
    #     |  |  |  |
    # [2] |  |  |  |
    #     |  |  |  |
    # [3] +--+  +--+
    [(0, 3), Direction.DOWN, Direction.LEFT],

    #    [0][1][2][3]
    # [0] +--+  +--+
    #     |  |  |  |
    # [1] |  |  |  |
    #     |  |  |  |
    # [2] |  |  |  |
    #     |  |  |  |
    # [3] O  +--+  v
    [(0, 3), Direction.UP, Direction.RIGHT],

    #    [0][1][2][3]
    # [0] +--+  +--+
    #     |  |  |  |
    # [1] |  |  |  |
    #     |  |  |  |
    # [2] |  |  |  |
    #     |  |  |  |
    # [3] v  +--+  O
    [(3, 3), Direction.UP, Direction.LEFT],
]

# Generate the file heuristic_path.go
print("package player\n")
print("// auto-generated file by `heuristic_path.py`.")
print("// Do not update it !\n")

print("var paths = [][]tile{")
for d in data:
    print("{", end='', sep='')
    path = getMonotonicPath(d[0], d[1], d[2])
    for p in path:
          print("{", p[0], ",", p[1], "},", end='', sep='')
    print("},")
print("}")


