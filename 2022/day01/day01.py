"""
https://adventofcode.com/2022/day/1
"""

# Part 1
with open("calories.txt", encoding="utf-8") as file:
    elves = [elf.split("\n") for elf in file.read().split("\n\n")]
summed_elves = [sum([int(calories) for calories in elf]) for elf in elves]
print(f"Elf with most calories has {max(summed_elves)}")

# Part 2
sorted_elves = sorted(summed_elves, reverse=True)
print(f"Total for top 3 elves is: {sum(sorted_elves[:3])}")
