import os, requests


class Present(object):
    def __init__(self, dimensions):
        self.x, self.y, self.z = sorted(dimensions)

    def wrapping(self):
        return 2 * (self.x * self.y + self.x * self.z + self.y * self.z) + self.x * self.y

    def ribbon(self):
        return 2 * self.x + 2 * self.y + self.x * self.y * self.z


input = requests.get('http://adventofcode.com/day/2/input',
                    cookies=dict(session=os.environ['ADVENT_SESSION'])).content

total_wrapping = 0
total_ribbon = 0
for line in input.strip().split('\n'):
    present = Present(map(int, line.split("x")))
    total_wrapping += present.wrapping()
    total_ribbon += present.ribbon()

print("Wrapping: %s, Ribbon: %s" % (total_wrapping, total_ribbon))
