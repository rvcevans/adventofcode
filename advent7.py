import os, re, requests
from operator import and_, or_, rshift, lshift
instructions = requests.get('http://adventofcode.com/day/7/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()


def get(node, nodes):
    return int(node) if node.isdigit() else nodes[node]


def get_nodes(instructions):
    nodes = {}
    for d in instructions.splitlines():
        rm = re.match('(?:(\w{1,2}|\d+) )?(AND|OR|LSHIFT|RSHIFT|NOT)? ?([\w]+) -> (\w+)', d)
        (ina, inb, out) = rm.group(1, 3, 4)
        nodes[out] = None
        if not inb.isdigit():
            nodes[inb] = None

        if ina is not None:
            if not ina.isdigit():
                nodes[ina] = None

    return nodes


def evaluate(instructions, nodes, overides=None):
    if overides is not None:
        for o in overides:
            nodes[o] = overides[o]

    for d in instructions.splitlines():
        rm = re.match('(?:(\w{1,2}|\d+) )?(AND|OR|LSHIFT|RSHIFT|NOT)? ?([\w]+) -> (\w+)', d)
        (ina, op, inb, out) = rm.group(1, 2, 3, 4)
        # if the out node has been evaluated then continue to the next instruction
        if not nodes[out] is None:
            continue

        bvalue = get(inb, nodes)
        if bvalue is None:
            continue

        if ina is None:
            if op == 'NOT':
                nodes[out] = ~ bvalue % 65536

            else:
                nodes[out] = bvalue

        else:
            avalue = get(ina, nodes)
            if avalue is None:
                continue

            nodes[out] = {'AND': and_, 'OR': or_, 'LSHIFT': lshift, 'RSHIFT': rshift}[op](avalue, bvalue)

    # print nodes
    if None in nodes.values():
        return evaluate(instructions, nodes)

    return nodes


networks = {1: evaluate(instructions, get_nodes(instructions))}
networks[2] = evaluate(instructions, get_nodes(instructions), {'b': networks[1]['a']})

print('\n'.join('Part {} value of wire a: {}'.format(key, val['a']) for key, val in networks.items()))
