import os, requests
puzzle = requests.get('http://adventofcode.com/day/20/input',
                            cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip()

primes = []
max_prime = 10000
p = 1
while p < max_prime:
    p += 1
    for q in primes:
        if p % q == 0:
            break

    else:
        primes.append(p)


def prime_powers(n):
    powers = {}
    for i in primes:
        while n % i == 0:
            if i not in powers:
                powers[i] = 1

            else:
                powers[i] += 1

            n /= i

        if n == 1:
            break

    return powers


def divisors(powers):
    if not powers:
        return [1]

    divs = []
    prime, power = powers.popitem()
    for div in divisors(powers):
        for i in xrange(power + 1):
            divs.append(prime**i * div)

    return divs


def presents(house_number):
    total = 0
    for div in divisors(prime_powers(house_number)):
        total += 10 * div

    return total


def new_presents(house_number):
    total = 0
    min_elf = house_number // 50
    for div in divisors(prime_powers(house_number)):
        if div > min_elf:
            total += 11 * div

    return total


binary_house = 2
while presents(binary_house) < int(puzzle):
    binary_house *= 2

house, new_house = binary_house / 4, binary_house / 4
while presents(house) < int(puzzle):
    house += 1

while new_presents(new_house) < int(puzzle):
    new_house += 1

print 'Part. 1 minimum house: %s' % house
print 'Part. 2 minimum house: %s' % new_house
