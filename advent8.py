import itertools, os, requests
edges = requests.get('http://adventofcode.com/day/8/input',
                     cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip().split('\n')


class Graph:
    def __init__(self):
        self.nodes = set()
        self.edges = {}

    def add_edge(self, start_node, end_node, weight):
        self.add_one_way_edge(start_node, end_node, weight)
        self.add_one_way_edge(end_node, start_node, weight)

    def add_one_way_edge(self, start_node, end_node, weight):
        self.nodes.add(start_node)
        self.nodes.add(end_node)
        if start_node not in self.edges.keys():
            self.edges[start_node] = {end_node: weight}

        else:
            self.edges[start_node][end_node] = weight

    def all_lengths(self):
        return [sum([G.edges[path[i]][path[i+1]] for i in xrange(len(path) - 1)])
                for path in itertools.permutations(G.nodes)]

G = Graph()
for edge in edges.split('\n'):
    values = edge.split(' ')
    G.add_edge(values[0], values[2], int(values[4]))

print("Minimum distance: %s, maximum distance: %s" % (min(G.all_lengths()), max(G.all_lengths())))