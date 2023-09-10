#!/usr/bin/env python3
"""
https://adventofcode.com/2022/day/2
"""

import os
from enum import Enum


with open(os.path.join(os.path.dirname(__file__), "rounds.txt"), encoding="utf-8") as file:
    rounds = [tuple(moves.split()) for moves in file.readlines()]


class Move(Enum):
    """Move in a game of rock paper scissors"""

    ROCK = 1
    PAPER = 2
    SCISSORS = 3

    def beats(self):
        """Move the current move beats"""
        match self:
            case Move.ROCK:
                return Move.SCISSORS
            case Move.PAPER:
                return Move.ROCK
            case Move.SCISSORS:
                return Move.PAPER

    def loses_to(self):
        """Move the current move loses to"""
        match self:
            case Move.ROCK:
                return Move.PAPER
            case Move.PAPER:
                return Move.SCISSORS
            case Move.SCISSORS:
                return Move.ROCK

    def score(self):
        """Point value of the current move"""
        match self:
            case Move.ROCK:
                return 1
            case Move.PAPER:
                return 2
            case Move.SCISSORS:
                return 3

    @classmethod
    def from_str(cls, move: str) -> "Move":
        """Get move from a string identifier"""
        match move:
            case "A" | "X":
                return Move.ROCK
            case "B" | "Y":
                return Move.PAPER
            case "C" | "Z":
                return Move.SCISSORS
            case _:
                raise ValueError(f"{move} is not a valid move identifier")


# Part 1
def round_score_part1(opponent_move: Move, my_move: Move) -> int:
    """Get the score for a round given both moves"""
    if opponent_move.beats() == my_move:
        return 0 + my_move.score()
    if opponent_move == my_move:
        return 3 + my_move.score()
    if opponent_move.loses_to() == my_move:
        return 6 + my_move.score()
    raise ValueError(f"{opponent_move} and {my_move} are not both valid moves")


total_part1 = sum(round_score_part1(Move.from_str(rnd[0]), Move.from_str(rnd[1])) for rnd in rounds)
print(f"Total for all rounds in part 1: {total_part1}")


# Part 2
def round_score_part2(opponent_move: Move, outcome: str) -> int:
    """Get the score for a round given opponent's move and an outcome"""
    match outcome:
        case "X":
            return opponent_move.beats().score() + 0
        case "Y":
            return opponent_move.score() + 3
        case "Z":
            return opponent_move.loses_to().score() + 6
        case _:
            raise ValueError(f"{outcome} is not a valid outcome identifier")


total_part2 = sum(round_score_part2(Move.from_str(rnd[0]), rnd[1]) for rnd in rounds)
print(f"Total for all rounds in part 2: {total_part2}")
