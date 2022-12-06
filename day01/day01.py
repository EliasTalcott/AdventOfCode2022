"""
https://adventofcode.com/2022/day/1

This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks, they need to know which Elf to ask:
they'd like to know how many Calories are being carried by the Elf carrying the most Calories.
In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
"""

# Part 1
with open("calories.txt", encoding="utf-8") as file:
    elves = [elf.split("\n") for elf in file.read().split("\n\n")]
summed_elves = [sum([int(calories) for calories in elf]) for elf in elves]
print(f"Elf with most calories has {max(summed_elves)}")

# Part 2
sorted_elves = sorted(summed_elves, reverse=True)
print(f"Total for top 3 elves is: {sum(sorted_elves[:3])}")
