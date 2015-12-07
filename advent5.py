import os, requests
strings = requests.get('http://adventofcode.com/day/5/input',
                       cookies=dict(session=os.environ['ADVENT_SESSION'])).content.strip().split('\n')

vowels = {'a', 'e', 'i', 'o', 'u'}
disallowed_strings = {'ab', 'cd', 'pq', 'xy'}


def vowel_count(word):
    count = 0
    for vowel in vowels:
        count += word.count(vowel)

    return count


def contains_repeated_letter(word, offset):
    for i in xrange(len(word) - 1 - offset):
        if word[i] == word[i + 1 + offset]:
            return True

    return False


def contains_disallowed(word):
    for s in disallowed_strings:
        if s in word:
            return True

    return False


def repeated_pair(word):
    for i in xrange(len(word) - 3):
        for j in xrange(i + 2, len(word) - 1):
            if word[i: i + 2] == word[j: j + 2]:
                return True

    return False


def nice1(word):
    return vowel_count(word) >= 3 and contains_repeated_letter(word, 0) and not contains_disallowed(word)


def nice2(word):
    return contains_repeated_letter(word, 1) and repeated_pair(word)

nice1_count, nice2_count = 0, 0
for word in strings:
    nice1_count += nice1(word)
    nice2_count += nice2(word)


print("Nice words for part 1: %s, Nice words for part 2: %s" % (nice1_count, nice2_count))
