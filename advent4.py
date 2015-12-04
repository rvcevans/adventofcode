import copy
import hashlib
import os
import requests

key = requests.get('http://adventofcode.com/day/4/input',
                          cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

code = 0
zeros = {5, 6}
answers = {}

while not not zeros:
    code += 1
    md5hash = hashlib.md5(key + str(code)).hexdigest()
    for val in copy.copy(zeros):
        if md5hash[:val] == ('0'*val):
            answers[val] = code
            zeros.remove(val)

print('\n'.join('{} produces the first hash with {} leading zeros.'.format(val, key) for key, val in answers.items()))
