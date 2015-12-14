import os, requests
descriptions = requests.get('http://adventofcode.com/day/14/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip().split('\n')


class Reindeer:
    def __init__(self, name, speed, fly, rest):
        self.name = name
        self.speed = speed
        self.fly = fly
        self.rest = rest
        self.time = 0
        self.distance = 0
        self.points = 0

    def race(self):
        current = self.time % (self.fly + self.rest)
        if current < self.fly:
            self.distance += self.speed

        self.time += 1

    def award(self, points):
        self.points += points

    def __lt__(self, other):
        return self.distance < other.distance

    def __eq__(self, other):
        return self.distance == other.distance


reindeers = []
for d in descriptions:
    values = d.split(' ')
    reindeers.append(Reindeer(values[0], int(values[3]), int(values[6]), int(values[13])))

race_time = 2503
time = 0
while time < race_time:
    for r in reindeers:
        r.race()

    time += 1
    max_reindeer = max(reindeers)
    for r in reindeers:
        if r == max_reindeer:
            r.award(1)


print('Max distance travelled after %s seconds: %s' % (race_time, max(reindeers).distance))
print('Max points after %s seconds: %s' % (race_time, max([r.points for r in reindeers])))