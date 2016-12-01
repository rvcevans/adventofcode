import os, requests
instructions = requests.get('http://adventofcode.com/day/6/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip().split('\n')


class Lights:
    def __init__(self):
        self.grid = {}
        for i in xrange(1000):
            for j in xrange(1000):
                self.grid[(i, j)] = 0

    def switch(self, x, y, on):
        for i in xrange(x[0], y[0] + 1):
            for j in xrange(x[1], y[1] + 1):
                self.grid[(i, j)] = on

    def toggle(self, x, y):
        for i in xrange(x[0], y[0] + 1):
            for j in xrange(x[1], y[1] + 1):
                self.grid[(i, j)] = not self.grid[(i, j)]

    def do(self, instruction):
        words = instruction.split(' ')
        if words[0] == 'toggle':
            self.toggle(map(int, words[1].split(',')), map(int, words[3].split(',')))

        else:
            self.switch(map(int, words[2].split(',')), map(int, words[4].split(',')), words[1] == 'on')

    def value(self):
        return sum(self.grid.values())


class BrightLights(Lights):
    def __init__(self):
        Lights.__init__(self)

    def switch(self, x, y, on):
        for i in xrange(x[0], y[0] + 1):
            for j in xrange(x[1], y[1] + 1):
                self.grid[(i, j)] = max(0, self.grid[(i, j)] + 2 * on - 1)

    def toggle(self, x, y):
        for i in xrange(x[0], y[0] + 1):
            for j in xrange(x[1], y[1] + 1):
                self.grid[(i, j)] += 2


lights = Lights()
brightLights = BrightLights()
for i in instructions:
    lights.do(i)
    brightLights.do(i)

print("Lights on in part 1: %s, brightness in part 2: %s" % (lights.value(), brightLights.value()))