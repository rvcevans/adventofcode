import copy, os, requests, itertools
attributes = requests.get('http://adventofcode.com/day/22/input',
                          cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

boss_values = {}
for at in attributes.splitlines():
    values = at.split(': ')
    boss_values[values[0]] = int(values[1])


class Spell:
    def __init__(self, cost, damage, armor, healing, recharge, duration):
        self.cost = cost
        self.damage = damage
        self.armor = armor
        self.healing = healing
        self.recharge = recharge
        self.duration = duration

spell_names = {'magic missile', 'drain', 'shield', 'poison', 'recharge'}
weapons = {'magic missile': Spell(53, 4, 0, 0, 0, 0),
           'drain': Spell(73, 2, 0, 2, 0, 0),
           'shield': Spell(113, 0, 7, 0, 0, 6),
           'poison': Spell(173, 3, 0, 0, 0, 6),
           'recharge': Spell(229, 0, 0, 0, 101, 5)}


class Wizard:
    def __init__(self, points, mana, spells):
        self.points = points
        self.mana = mana
        self.health = points
        self.spells = spells
        self.purchases = []
        self.spend = 0

    def apply_spells(self, boss):
        for s in self.spells:
            self.apply(s, boss)

        self.reduce_duration()

    def apply(self, spell, boss):
        boss.health -= weapons[spell].damage
        self.health += weapons[spell].healing
        self.mana += weapons[spell].recharge

    def reduce_duration(self):
        for s in copy.copy(self.spells):
            self.spells[s] -= 1
            if self.spells[s] == 0:
                del self.spells[s]

    def purchase(self, spell, boss):
        self.mana -= weapons[spell].cost
        self.spend += weapons[spell].cost
        self.purchases.append(spell)
        if weapons[spell].duration == 0:
            self.apply(spell, boss)

        else:
            self.spells[spell] = weapons[spell].duration


class Boss:
    def __init__(self, points, damage):
        self.points = points
        self.damage = damage
        self.health = points

    def attack(self, wizard):
        wizard.health -= max(1, self.damage - sum(weapons[s].armor for s in wizard.spells))


def play_wizard(wizard, boss, hard_mode=False):
    if hard_mode:
        wizard.health -= 1
        if wizard.health <= 0:
            return 1000000

    wizard.apply_spells(boss)
    if boss.health <= 0:
        return wizard.spend

    min_cost = 1000000
    for s in spell_names:
        if s not in wizard.spells.keys() and weapons[s].cost < wizard.mana:
            new_wiz = copy.deepcopy(wizard)
            new_boss = copy.deepcopy(boss)
            new_wiz.purchase(s, new_boss)
            if new_wiz.spend > 1500:
                continue

            if new_boss.health <= 0:
                min_cost = min(min_cost, new_wiz.spend)

            else:
                min_cost = min(min_cost, play_boss(new_wiz, new_boss, hard_mode))

    return min_cost


def play_boss(wizard, boss, hard_mode=False):
    wizard.apply_spells(boss)
    if boss.health <= 0:
        return wizard.spend

    boss.attack(wizard)
    if wizard.health <= 0:
        return 1000000

    return play_wizard(wizard, boss, hard_mode)

print play_wizard(Wizard(50, 500, {}), Boss(boss_values['Hit Points'], boss_values['Damage']))
print play_wizard(Wizard(50, 500, {}), Boss(boss_values['Hit Points'], boss_values['Damage']), True)
