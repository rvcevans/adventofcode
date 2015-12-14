import os, requests, string
password = requests.get('http://adventofcode.com/day/11/input',
                        cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

characters = {i: x for i, x in enumerate([s for s in list(string.lowercase) if s not in ['i', 'o', 'l']])}
indices = {x: i for i, x in enumerate([s for s in list(string.lowercase) if s not in ['i', 'o', 'l']])}
base = len(characters) - 1


def test_straight(password):
    indexed = [indices[p] for p in password]
    straight, current = 1, indexed[0]
    for i in xrange(1, len(password)):
        if indexed[i] == current + 1:
            straight += 1
            if straight == 3:
                return True

        else:
            straight = 1

        current = indexed[i]

    return False


def test_double(password):
    current = password[0]
    double_count = 0
    last_letter_double = False
    for i in range(1, len(password)):
        if not last_letter_double and password[i] == current:
            double_count += 1
            last_letter_double = True

        else:
            last_letter_double = False

        current = password[i]

    if double_count >= 2:
        return True

    return False


def increment(password):
    indexed = [indices[p] for p in password]
    i = len(password) - 1
    for i in reversed(xrange(len(password))):
        if indexed[i] == base:
            indexed[i] = 0

        else:
            indexed[i] += 1
            break

    return ''.join([characters[i] for i in indexed])

passwords = []
while len(passwords) != 2:
    password = increment(password)
    if test_straight(password) and test_double(password):
        passwords.append(password)

for i, p in enumerate(passwords):
    print 'Answer to Part. %s: %s' % (i + 1, p)
