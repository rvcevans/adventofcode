import os, re, requests, itertools
from operator import eq, gt, lt
descriptions = requests.get('http://adventofcode.com/day/16/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


machine = {'children': 3, 'cats': 7, 'samoyeds': 2, 'pomeranians': 3, 'akitas': 0,
           'vizslas': 0, 'goldfish': 5, 'trees': 3, 'cars': 2, 'perfumes': 1}
operators = {'children': eq, 'cats': gt, 'samoyeds': eq, 'pomeranians': lt, 'akitas': eq,
             'vizslas': eq, 'goldfish': lt, 'trees': gt, 'cars': eq, 'perfumes': eq}

aunts = {}
for d in descriptions.splitlines():
    rm = re.match('Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)', d)
    aunts[int(rm.group(1))] = {rm.group(2 * j): int(rm.group(2 * j + 1)) for j in xrange(1, 4)}

for number, aunt in aunts.iteritems():
    for com, value in aunt.iteritems():
        if machine[com] != value:
            break

    else:
        print 'Fake Sue: %s' % number

    for com, value in aunt.iteritems():
        if not operators[com](value, machine[com]):
            break

    else:
        print 'Real Sue: %s' % number