#!/usr/bin/env python3
# coding: utf-8

# half-byte
nibbleMax=16

def trans(n, m):
    # merge
    if n == m and n != 0 and n < 15:
        return 0, n + 1, True
    # move
    if m == 0:
        return 0, n, False
    # nothing to do
    return n, m, True

def transRightDown(a, b, c, d):
    c, d, m1 = trans(c, d)
    b, c, m2 = trans(b, c)
    if not m1:
        c, d, m1 = trans(c, d)
    a, b, _ = trans(a, b)
    if not m2:
        b, c, _ = trans(b, c)
    if not m1:
        c, d, _ = trans(c, d)
    return a, b, c, d

def transLeftUp(a, b, c, d):
    b, a, m1 = trans(b, a)
    c, b, m2 = trans(c, b)
    if not m1:
        b, a, m1 = trans(b, a)
    d, c, _ = trans(d, c)
    if not m2:
        c, b, _ = trans(c, b)
    if not m1:
        b, a, _ = trans(b, a)
    return a, b, c, d

def encodeUint16(a, b, c, d):
    return a << 12 | b << 8 |  c << 4 | d

def addTrans(trans, t):
    trans.append(encodeUint16(t[0], t[1], t[2], t[3]))

TRB = [] # computed transitions right/bottom
TLT = [] # computed transitions left/top

for a in range(nibbleMax):
    for b in range(nibbleMax):
        for c in range(nibbleMax):
            for d in range(nibbleMax):
                addTrans(TRB, transRightDown(a, b, c, d))
                addTrans(TLT, transLeftUp(a, b, c, d))

print("package player")
print("// auto-generated file, do not update it !\n")

print("var transLeftUp = []uint16{")
for t in TLT:
    print(t, ", ", end='', sep='')
print("\n}")

print("var transRightDown = []uint16{")
for t in TRB:
    print(t, ", ", end='', sep='')
print("\n}")