import os, requests, itertools
from operator import mul
instructions = requests.get('http://adventofcode.com/day/24/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()
boxes = map(int, instructions.splitlines())


def entanglement(boxes):
    return reduce(mul, boxes, 1)


def arrangements(boxes, target):
    found = False
    for i in xrange(1, len(boxes) - 1):
        if found:
            return

        for c in itertools.combinations(boxes, i):
            if sum(c) == target:
                found = True
                yield c

print 'Part 1. %s' % min(entanglement(c) for c in arrangements(boxes, sum(boxes) / 3))
print 'Part 2. %s' % min(entanglement(c) for c in arrangements(boxes, sum(boxes) / 4))

