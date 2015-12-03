import copy
import requests
import os


class Location:
    def __init__(self, latitude, longitude):
        self.latitude = latitude
        self.longitude = longitude

    def __hash__(self):
        return hash((self.latitude, self.longitude))

    def __eq__(self, other):
        return (self.latitude, self.longitude) == (other.latitude, other.longitude)

    def move_north(self):
        self.latitude += 1

    def move_south(self):
        self.latitude -= 1

    def move_east(self):
        self.longitude += 1

    def move_west(self):
        self.longitude -= 1

    def move(self, direction):
        {'^': self.move_north, '>': self.move_east, 'v': self.move_south, '<': self.move_west}[direction]()


directions = requests.get('http://adventofcode.com/day/3/input',
                          cookies=dict(session=os.environ['ADVENT_SESSION'])).content

santa1, santa2, robo_santa = Location(0, 0), Location(0, 0), Location(0, 0)
visits1 = {copy.copy(santa1)}
visits2 = {copy.copy(santa1)}

for index, direction in enumerate(directions):
    santa1.move(direction)
    visits1.add(copy.copy(santa1))
    if index % 2 == 0:
        santa2.move(direction)
        visits2.add(copy.copy(santa2))

    else:
        robo_santa.move(direction)
        visits2.add(copy.copy(robo_santa))


print("Houses visited first year: %s, Houses visited second year: %s" % (len(visits1), len(visits2)))
