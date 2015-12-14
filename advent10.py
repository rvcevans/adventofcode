import copy, os, requests
puzzle = requests.get('http://adventofcode.com/day/10/input',
                      cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


def look_and_say(puzzle):
    current = puzzle[:1]
    count = 1
    output = ''
    for i in xrange(1, len(puzzle)):
        if puzzle[i:i+1] == current:
            count += 1

        else:
            output += str(count) + current
            current = puzzle[i:i+1]
            count = 1

    output += str(count) + current
    return output


iterations = {40, 50}
answers = {}
j = 0
while not not iterations:
    j += 1
    puzzle = look_and_say(puzzle)
    if j in copy.copy(iterations):
        iterations.remove(j)
        answers[j] = len(puzzle)

print('\n'.join('Puzzle length after {} iterations: {}.'.format(key, val) for key, val in answers.items()))
