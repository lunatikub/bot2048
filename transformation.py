#!/usr/bin/env python3
# coding: utf-8

"""
Generate all transformations for all combinations
of each column/line.
"""

# half-byte
nibbleMax=16

def trans(n, m):
    """
    transformation between 2 tiles
    [n,m] -> [n',m']
    """
    # merge
    if n == m and n != 0 and n < 15:
        return 0, n + 1, True
    # move
    if m == 0:
        return 0, n, False
    # nothing to do
    return n, m, False


def transRightDown(a, b, c, d):
    """
    right/down transformation
    [a,b,c,d] -> [a',b',c',d']
    """
    # tile[2]
    c, d, m1 = trans(c, d)
    # tile[1]
    b, c, m2 = trans(b, c)
    if not m1 and not m2:
        c, d, m1 = trans(c, d)
    # tile[0]
    a, b, m3 = trans(a, b)
    if not m2 and not m3:
        b, c, _ = trans(b, c)
    if not m1 and not m2:
        c, d, _ = trans(c, d)
    return a, b, c, d


def transLeftUp(a, b, c, d):
    """
    left/up transformation
    [a,b,c,d] -> [a',b',c',d']
    """
    b, a, m1 = trans(b, a)
    c, b, m2 = trans(c, b)
    if not m1 and not m2:
        b, a, m1 = trans(b, a)
    d, c, m3 = trans(d, c)
    if not m2 and not m3:
        c, b, _ = trans(c, b)
    if not m1 and not m2:
        b, a, _ = trans(b, a)
    return a, b, c, d

def encodeUint16(a, b, c, d):
    return a << 12 | b << 8 |  c << 4 | d

def addTrans(trans, t):
    trans.append(encodeUint16(t[0], t[1], t[2], t[3]))

TRD = [] # computed transitions right/down
TLU = [] # computed transitions left/up

for a in range(nibbleMax):
    for b in range(nibbleMax):
        for c in range(nibbleMax):
            for d in range(nibbleMax):
                addTrans(TRD, transRightDown(a, b, c, d))
                addTrans(TLU, transLeftUp(a, b, c, d))

# Generate the file transition.go with the 2 slices.
print("package bot\n")
print("// auto-generated file by `transformation.py`.")
print("// Do not update it !\n")

print("var transLeftUp = []uint16{")
for t in TLU:
    print(t, ", ", end='', sep='')
print("\n}\n")

print("var transRightDown = []uint16{")
for t in TRD:
    print(t, ", ", end='', sep='')
print("\n}")
