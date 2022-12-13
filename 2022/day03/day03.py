"""
https://adventofcode.com/2022/day/3
"""


def score(char: str) -> int:
    if char.islower():
        return ord(char) - 96
    return ord(char) - 38


# Part 1
with open("rucksacks.txt", encoding="utf-8") as file:
    rucksacks = file.read().splitlines()

part1_total = 0
for rucksack in rucksacks:
    for letter in set(rucksack[:int(len(rucksack) / 2)]).intersection(rucksack[int(len(rucksack) / 2):]):
        part1_total += score(letter)
print(f"Sum of priorities is {part1_total}")

# Part 2
part2_total = 0
for i in range(0, len(rucksacks), 3):
    group = rucksacks[i:i + 3]
    intersection = set(group[0]).intersection(group[1]).intersection(group[2])
    part2_total += score(set(group[0]).intersection(group[1]).intersection(group[2]).pop())
print(f"Sum of badge priorities is {part2_total}")
