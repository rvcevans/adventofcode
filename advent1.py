import requests
import os

elevator = requests.get('http://adventofcode.com/day/1/input',
                        cookies=dict(session=os.environ['ADVENT_SESSION'])).content

floor = 0
for i in xrange(len(elevator)):
    if elevator[i] == '(':
        floor += 1

    else:
        floor -= 1

    if floor == -1:
        break

print ("Final floor: %s, First time in basement: %s" % (elevator.count('(') - elevator.count(')'), i + 1))
