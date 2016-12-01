import os, requests
instructions = requests.get('http://adventofcode.com/day/18/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

class Lights:
    def __init__(self, inital):
        self.grid = {}
        self.states = {}
        for (y, line) in enumerate(inital.splitlines()):
            for (x, val) in enumerate(line):
                self.grid[(x, y)] = {'#': True, '.': False}[val]

        self.row_count = max(p[0] for p in self.grid.keys()) + 1
        self.col_count = max(p[1] for p in self.grid.keys()) + 1
        self.evaluate()

    def toggle(self, p):
        self.grid[p] = not self.grid[p]

    def value(self, p):
        if p[0] < 0 or p[0] >= self.row_count or p[1] < 0 or p[1] >= self.col_count:
            return 0

        return self.grid[p]

    def neighbours(self, p):
        for i in xrange(-1, 2):
            for j in xrange(-1, 2):
                if (i, j) != (0, 0):
                    yield (p[0] - i, p[1] - j)

    def get_state(self, p):
        return sum(self.value(q) for q in self.neighbours(p))

    def evaluate(self):
        for p in self.grid.keys():
            self.states[p] = self.get_state(p)

    def next_state(self, p):
        state = self.states[p]
        val = self.grid[p]
        if val == True and not (state == 2 or state == 3):
            return False

        elif val == False and state == 3:
            return True

        return val

    def change(self):
        for p in self.grid.keys():
            self.grid[p] = self.next_state(p)

        self.evaluate()


    def brightness(self):
        return sum(self.grid.values())

    def print_lights(self):
        output = ''
        for y in xrange(0, self.col_count):
            for x in xrange(0, self.row_count):
                output += {True: '#', False: '.'}[self.grid[(x, y)]]

            output += '\n'

        return output

    def print_states(self):
        output = ''
        for y in xrange(0, self.col_count):
            for x in xrange(0, self.row_count):
                output += str(self.states[(x, y)])

            output += '\n'

        print output

class BrokenLights(object, Lights):
    def __init__(self, inital):
        Lights.__init__(self, inital)
        for x in {0, self.row_count - 1}:
            for y in {0, self.col_count - 1}:
                self.grid[(x, y)] = True

        self.evaluate()

    def next_state(self, p):
        if p in {(0, 0), (0, self.col_count - 1), (self.row_count - 1, 0), (self.row_count - 1, self.col_count - 1)}:
            return True

        return super(BrokenLights, self).next_state(p)

lights = Lights(instructions)
broken_lights = BrokenLights(instructions)
for i in xrange(100):
    lights.change()
    broken_lights.change()


print('Part. 1 brightness: %s' % lights.brightness())
print('part. 2 brightness: %s' % broken_lights.brightness())