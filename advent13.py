import copy, itertools, os, requests
descriptions = requests.get('http://adventofcode.com/day/13/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip().split('\n')


class Feast:
    def __init__(self):
        self.guests = set()
        self.potentials = {}

    def add_potential(self, guest, neighbour, happiness):
        self.guests.add(guest)
        self.guests.add(neighbour)
        if guest not in self.potentials.keys():
            self.potentials[guest] = {neighbour: happiness}

        else:
            self.potentials[guest][neighbour] = happiness

    def all_arrangement_potentials(self):
        return [sum([self.potentials[path[i]][path[i+1]] for i in xrange(len(seating) - 1)])
                for seating in itertools.permutations(self.guests)]

    def happiness(self, arrangement):
        total = 0
        count = len(arrangement)
        for i in xrange(count):
            total += self.potentials[arrangement[i]][arrangement[(i+1) % count]]
            total += self.potentials[arrangement[i]][arrangement[(i-1) % count]]

        return total


F = Feast()
for d in descriptions:
    values = d.split(' ')
    if values[2] == 'gain':
        F.add_potential(values[0], values[10][:-1], int(values[3]))

    else:
        F.add_potential(values[0], values[10][:-1], - int(values[3]))

G = copy.deepcopy(F)
for g in F.guests:
    G.add_potential('me', g, 0)
    G.add_potential(g, 'me', 0)

print("Optimal arrangement change Part. 1: %s" % (max([F.happiness(a) for a in itertools.permutations(F.guests)])))
print("Optimal arrangement change Part. 2: %s" % (max([G.happiness(a) for a in itertools.permutations(G.guests)])))