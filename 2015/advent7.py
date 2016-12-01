import os, re, requests
from operator import and_, or_, rshift, lshift
puzzle_input = requests.get('http://adventofcode.com/day/7/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


def instructions(network):
    for d in network.splitlines():
        yield re.match('(?:(\w{1,2}|\d+) )?(AND|OR|LSHIFT|RSHIFT|NOT)? ?([\w]+) -> (\w+)', d).group(1, 2, 3, 4)


def get(node, nodes):
    return int(node) if node.isdigit() else nodes[node]


def get_nodes(network):
    nodes = {}
    for (ina, op, inb, out) in instructions(network):
        nodes[out] = None
        if not inb.isdigit():
            nodes[inb] = None

        if ina is not None:
            if not ina.isdigit():
                nodes[ina] = None

    return nodes


def evaluate(network, nodes, overrides=()):
    for o in overrides:
        nodes[o] = overrides[o]

    for (ina, op, inb, out) in instructions(network):
        if not nodes[out] is None:
            continue

        b = get(inb, nodes)
        if b is None:
            continue

        if ina is None:
            nodes[out] = ~ b % 65536 if op == 'NOT' else b

        else:
            a = get(ina, nodes)
            if a is None:
                continue

            nodes[out] = {'AND': and_, 'OR': or_, 'LSHIFT': lshift, 'RSHIFT': rshift}[op](a, b)

    return evaluate(network, nodes) if None in nodes.values() else nodes


networks = {1: evaluate(puzzle_input, get_nodes(puzzle_input))}
networks[2] = evaluate(puzzle_input, get_nodes(puzzle_input), {'b': networks[1]['a']})

print('\n'.join('Part {} value of wire a: {}'.format(key, val['a']) for key, val in networks.items()))
