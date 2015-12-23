import re, os, requests
from operator import add, mul, div
instructions = requests.get('http://adventofcode.com/day/23/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


class Command:
    def __init__(self, action, register=None, offset=None):
        self.action = action
        self.register = register
        self.offset = 0 if offset is None else int(offset)

    def apply(self, index, registers):
        index += {'hlf': 1, 'tpl': 1, 'inc': 1, 'jmp': self.offset}.get(self.action, 0)
        if self.register is not None:
            index += {'jie': 1}.get(self.action, 0)*{0: self.offset}.get(registers[self.register] % 2, 1)
            index += {'jio': 1}.get(self.action, 0)*{1: self.offset}.get(registers[self.register], 1)
            registers[self.register] = {'hlf': div, 'tpl': mul, 'inc': add}.get(self.action, add)\
                (registers[self.register], {'hlf': 2, 'tpl': 3, 'inc': 1}.get(self.action, 0))

        return index


commands = {}
for index, value in enumerate(instructions.splitlines()):
    rm = re.match('(\w{3}) (\w)?,? ?([+-]\d+)?', value)
    commands[index] = Command(rm.group(1), rm.group(2), rm.group(3))

indices = {1: 0, 2: 0}
registers = {1: {'a': 0, 'b': 0}, 2: {'a': 1, 'b': 0}}
for i in indices:
    while 0 <= indices[i] < len(commands):
        indices[i] = commands[indices[i]].apply(indices[i], registers[i])

print('\n'.join('Part {} register b value on exit: {}'.format(key, val['b']) for key, val in registers.items()))