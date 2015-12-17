import os, re, requests, itertools
from operator import eq, gt, lt
descriptions = requests.get('http://adventofcode.com/day/17/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

containers = map(int, descriptions.splitlines())
ways = 0
for i in range(1, len(containers)):
    new_ways = sum(sum(combo) == 150 for combo in itertools.combinations(containers, i))
    if new_ways != 0 and ways == 0:
        print 'Combinations with the minimum number of containers: %s' % new_ways

    ways += new_ways

print 'Total combinations: %s' % ways