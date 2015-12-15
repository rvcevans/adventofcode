import os, re, requests, itertools
descriptions = requests.get('http://adventofcode.com/day/15/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


class Cookie:
    def __init__(self, ingredients):
        self.ingredients = ingredients
        self.quantities = {i.name: 0 for i in ingredients}

    def score(self):
        return max(sum([i.capacity * self.quantities[i.name] for i in self.ingredients]), 0) \
               * max(sum([i.durability * self.quantities[i.name] for i in self.ingredients]), 0) \
               * max(sum([i.flavor * self.quantities[i.name] for i in self.ingredients]), 0) \
               * max(sum([i.texture * self.quantities[i.name] for i in self.ingredients]), 0)

    def energy(self):
        return sum([i.calories * self.quantities[i.name] for i in self.ingredients])

    def alter(self, ingredient_name, quantity):
        self.quantities[ingredient_name] = quantity


class Ingredient:
    def __init__(self, name, capacity, durability, flavor, texture, calories):
        self.name = name
        self.capacity = capacity
        self.durability = durability
        self.flavor = flavor
        self.texture = texture
        self.calories = calories


def recipes(ingredients, teaspoons):
    if len(ingredients) == 1:
        return [{ingredients[0]: teaspoons}]

    recs = []
    for i in xrange(1, teaspoons - len(ingredients) + 2):
        for r in recipes(ingredients[1:], teaspoons - i):
            r[ingredients[0]] = i
            recs.append(r)

    return recs


ingredients = set()
for d in descriptions.splitlines():
    i = re.match('(\w+): capacity (\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)', d)
    ingredients.add(Ingredient(i.group(1), int(i.group(2)), int(i.group(3)), int(i.group(4)), int(i.group(5)), int(i.group(6))))

C = Cookie(ingredients)
max_score_1, max_score_2 = 0, 0
for recipe in recipes([i.name for i in C.ingredients], 100):
    for name, amount in recipe.iteritems():
        C.alter(name, amount)

    max_score_1 = max(max_score_1, C.score())
    if C.energy() == 500:
        max_score_2 = max(max_score_2, C.score())


print('Max score Part. 1: %s' % (max_score_1))
print('Max score Part. 2: %s' % (max_score_2))