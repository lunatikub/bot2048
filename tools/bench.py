#!/usr/bin/env python3
# coding: utf-8

from collections import defaultdict
from collections import OrderedDict

import subprocess

nrRun=20 # number of run by depth

def runBot2048(depth):
    process = subprocess.Popen(["./bot2048", "--stats", "--depth", str(depth)],
                        shell=False,
                        stdout=subprocess.PIPE)

    stdout = process.communicate()
    res=stdout[0].decode('utf-8').split()

    return int(res[0]), int(res[1]), int(res[2]), int(res[3]), int(res[4])


for depth in range(8, 10):
    time = score = nrEval = nrMove = 0
    tiles = defaultdict(int)
    for r in range(0, nrRun):
        t, m, s, n, nr = runBot2048(depth)
        time += t; score += s; nrEval += n; nrMove += nr
        tiles[m] += 1

    nrMove /= nrRun; nrEval /= nrRun; time /= nrRun; score /= nrRun
    print(depth, nrMove, time / 1e9, nrEval, score, sep=',', end='')
    od = {int(k) : v for k, v in tiles.items()}
    for k, v in sorted(od.items(), reverse=True):
        print(",", k, ":", (v*100)/nrRun, sep='', end='')
    print()

