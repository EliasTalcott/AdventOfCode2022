#!/usr/bin/env python3
"""
https://adventofcode.com/2022/day/1
"""

import os


with open(os.path.join(os.path.dirname(__file__), "calories.txt"), encoding="utf-8") as file:
    elf_calories = [sum(map(int, elf.split("\n"))) for elf in file.read().split("\n\n")]

# Part 1
most_calories = max(elf_calories)
print(f"Elf with most calories has {most_calories}")

# Part 2
top_3_total = sum(sorted(elf_calories, reverse=True)[:3])
print(f"Total for top 3 elves is: {top_3_total}")
