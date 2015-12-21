import os, requests
instructions = requests.get('http://adventofcode.com/day/19/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

replacements = set()
medicine = ''
for line in instructions.splitlines():
    if '=>' in line:
        values = line.split(' ')
        replacements.add((values[0], values[2]))

    else:
        medicine = line

def fabricate(molecules):
    new_molecules = set()
    for m in molecules:
        for i in xrange(len(m)):
            for r in replacements:
                from_text = r[0]
                to_text = r[1]
                if len(m) - i < len(from_text):
                    continue

                if m[i: i + len(from_text)] == from_text:
                    new_molecules.add(m[:i] + to_text + m[i + len(from_text):])

    return new_molecules

def condense(molecules):
    new_molecules = set()
    for m in molecules:
        for i in xrange(len(m)):
            for r in replacements:
                from_text = r[1]
                to_text = r[0]
                if len(m) - i < len(from_text):
                    continue

                if m[i: i + len(from_text)] == from_text:
                    new_molecules.add(m[:i] + to_text + m[i + len(from_text):])

    return new_molecules


def condense(m, index, depth):
    if m == 'e':
        print depth

    for r in replacements:
        from_text = r[1]
        to_text = r[0]
        if len(m) - index < len(from_text):
            print(len(m), index, from_text)
            continue

        new_index = index - len(from_text) + len(to_text) + 1
        if new_index > len(from_text):
            new_index = 0

        print m[index: index + len(from_text)], from_text
        if m[index: index + len(from_text)] == from_text:
            condense(m[:index] + to_text + m[index + len(from_text):], new_index, depth + 1)


print(len(fabricate([medicine])))
condense(medicine, 0, 0)

