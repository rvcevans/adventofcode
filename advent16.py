import operator, os, re, requests, itertools
descriptions = requests.get('http://adventofcode.com/day/16/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


machine = {'children': 3, 'cats': 7, 'samoyeds': 2, 'pomeranians': 3, 'akitas': 0,
           'vizslas': 0, 'goldfish': 5, 'trees': 3, 'cars': 2, 'perfumes': 1}
operators = {'children': operator.eq, 'cats': operator.gt, 'samoyeds': operator.eq, 'pomeranians': operator.lt,
             'akitas': operator.eq, 'vizslas': operator.eq, 'goldfish': operator.lt, 'trees': operator.gt,
             'cars': operator.eq, 'perfumes': operator.eq}

aunts = {}
for d in descriptions.splitlines():
    i = re.match('Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)', d)
    aunts[int(i.group(1))] = {i.group(2 * j + 2): int(i.group(2 * j + 3)) for j in xrange(3)}

for number, aunt in aunts.iteritems():
    for at, value in aunt.iteritems():
        if machine[at] != value:
            break

    else:
        print 'Fake Sue: %s' % number

    for at, value in aunt.iteritems():
        if not operators[at](value, machine[at]):
            break

    else:
        print 'Real Sue: %s' % number