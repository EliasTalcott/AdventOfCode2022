"""
https://adventofcode.com/2022/day/2
"""

# Part 1
with open("rounds.txt", encoding="utf-8") as file:
    rounds = [(combination[0], combination[2]) for combination in file.read().splitlines()]

PART1_OUTCOMES = {
    ("A", "X"): 1 + 3,  # Rock vs. Rock
    ("A", "Y"): 2 + 6,  # Rock vs. Paper
    ("A", "Z"): 3 + 0,  # Rock vs. Scissors
    ("B", "X"): 1 + 0,  # Paper vs. Rock
    ("B", "Y"): 2 + 3,  # Paper vs. Paper
    ("B", "Z"): 3 + 6,  # Paper vs. Scissors
    ("C", "X"): 1 + 6,  # Scissors vs. Rock
    ("C", "Y"): 2 + 0,  # Scissors vs. Paper
    ("C", "Z"): 3 + 3,  # Scissors vs. Scissors
}
total = 0
for rnd in rounds:
    total += PART1_OUTCOMES[rnd]
print(f"Total for part 1 rounds: {total}")

# Part 2
PART2_OUTCOMES = {
    ("A", "X"): 3 + 0,  # Rock + Lose (play scissors)
    ("A", "Y"): 1 + 3,  # Rock + Tie (play rock)
    ("A", "Z"): 2 + 6,  # Rock + Win (play paper)
    ("B", "X"): 1 + 0,  # Paper + Lose (play rock)
    ("B", "Y"): 2 + 3,  # Paper + Tie (play paper)
    ("B", "Z"): 3 + 6,  # Paper + Win (play scissors)
    ("C", "X"): 2 + 0,  # Scissors + Lose (play paper)
    ("C", "Y"): 3 + 3,  # Scissors + Tie (play scissors)
    ("C", "Z"): 1 + 6,  # Scissors + Win (play rock)
}
total = 0
for rnd in rounds:
    total += PART2_OUTCOMES[rnd]
print(f"Total for part 2 rounds: {total}")
