import os, re, requests
import operator
descriptions = requests.get('http://adventofcode.com/day/7/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

nodes = set()
node_values = {}

for d in descriptions.splitlines():
    print d
    rm = re.match('(\w+) (AND|OR|LSHIFT|RSHIFT|NOT) (\w+) -> (\w+)', d)
    print rm.group(1, 2, 3)