#!/usr/bin/env python3
"""
https://adventofcode.com/2022/day/3
"""

import os


with open(os.path.join(os.path.dirname(__file__), "rucksacks.txt"), encoding="utf-8") as file:
    rucksacks = file.read().splitlines()


def priority(char: str) -> int:
    """Calculate the priority of a character"""
    return ord(char) - 96 if char.islower() else ord(char) - 38


# Part 1
total_part1 = sum(
    sum(map(priority, shared))
    for shared in (set(rs[: len(rs) // 2]).intersection(rs[len(rs) // 2 :]) for rs in rucksacks)
)
print(f"Sum of priorities is {total_part1}")

# Part 2
total_part2 = sum(
    sum(map(priority, set.intersection(*list(map(set, group))).pop()))
    for group in (rucksacks[i : i + 3] for i in range(0, len(rucksacks), 3))
)
print(f"Sum of badge priorities is {total_part2}")
