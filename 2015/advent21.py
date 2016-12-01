import os, requests, itertools
attributes = requests.get('http://adventofcode.com/day/21/input',
                          cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

boss_values = {}
for at in attributes.splitlines():
    values = at.split(': ')
    boss_values[values[0]] = int(values[1])


class Player:
    def __init__(self, points, items):
        self.points = points
        self.items = items
        self.health = points

    def damage(self):
        return sum(a.damage for a in self.items)

    def armor(self):
        return sum(a.armor for a in self.items)

    def price(self):
        return sum(a.cost for a in self.items)


class Boss:
    def __init__(self, points, damage, armor):
        self.points = points
        self.damage = damage
        self.armor = armor
        self.health = points


class Item:
    def __init__(self, category, name, cost, damage, armor):
        self.category = category
        self.name = name
        self.cost = cost
        self.damage = damage
        self.armor = armor

weapons = {Item('weapon', "dagger", 8, 4, 0), Item('weapon', 'shortsword', 10, 5, 0),
        Item('weapon', 'warhammer', 25, 6, 0), Item('weapon', 'longsword', 40, 7, 0),
        Item('weapon', 'greataxe', 74, 8, 0)}
armors = {Item('armor', 'leather', 13, 0, 1), Item('armor', 'chainmail', 31, 0, 2),
          Item('armor', 'splntmail', 53, 0, 3), Item('armor', 'bandedmail', 75, 0, 4),
          Item('armor', 'platemail', 102, 0, 5)}
rings = {Item('ring', 'damage 1', 25, 1, 0), Item('ring', 'damage 2', 50, 2, 0),
        Item('ring', 'damage 3', 100, 3, 0), Item('ring', 'defense 1', 20, 0, 1),
        Item('ring', 'defense 2', 40, 0, 2), Item('ring', 'defense 3', 80, 0, 3)}


# returns bool if player wins
def simulate(player, boss):
    turn = 1
    while player.health > 0 and boss.health > 0:
        if turn % 2 == 1:
            boss.health -= max(1, player.damage() - boss.armor)

        else:
            player.health -= max(1, boss.damage - player.armor())

        turn += 1

    if player.health > 0:
        return True

    return False


setups = set()
for w in weapons:
    setups.add((w,))
    for a in armors:
        setups.add((a, w))
        for r in rings:
            setups.add((a, w, r))

        for rs in itertools.combinations(rings, 2):
            setups.add(tuple(rs) + (a, w))

    for rs in itertools.combinations(rings, 2):
            setups.add(tuple(rs) + (w,))

min_win = 1000
max_lose = 0
for outfit in setups:
    player = Player(100, outfit)
    boss = Boss(boss_values['Hit Points'], boss_values['Damage'], boss_values['Armor'])
    if simulate(player, boss):
        min_win = min(min_win, player.price())

    else:
        max_lose = max(max_lose, player.price())


print 'Cheapest way to win: %s' % min_win
print 'Most expensive way to lose: %s' % max_lose