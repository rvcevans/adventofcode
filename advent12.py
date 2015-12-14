import json, os, requests, string
document = requests.get('http://adventofcode.com/day/12/input',
                        cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


def is_int(s):
    try:
        int(s)
        return True
    except ValueError:
        return False


def json_sum(a, exclude=None):
    if type(a) == type({}):
        for b in a:
            if b == exclude or a[b] == exclude:
                return 0

        return sum([json_sum(a[b], exclude) for b in a])

    if type(a) == type([]):
        return sum([json_sum(b, exclude) for b in a])

    if is_int(a):
        return int(a)

    return 0


print "Part. 1 sum: %s" % json_sum(json.loads(document))
print "Part. 2 sum: %s" % json_sum(json.loads(document), 'red')