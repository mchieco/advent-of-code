import pathlib

import main
import pytest

PUZZLE_DIR = pathlib.Path(__file__).parent


@pytest.fixture
def example1_input():
    puzzle_input = (PUZZLE_DIR / "example1_input.txt").read_text().rstrip()
    return main.parse_data(puzzle_input)


@pytest.fixture
def example1_solution():
    puzzle_solution = (PUZZLE_DIR / "example1_solution.txt").read_text().rstrip()
    return int(puzzle_solution)


@pytest.fixture
def example2_input():
    puzzle_input = (PUZZLE_DIR / "example2_input.txt").read_text().rstrip()
    return main.parse_data(puzzle_input)


@pytest.fixture
def example2_solution():
    puzzle_solution = (PUZZLE_DIR / "example2_solution.txt").read_text().rstrip()
    return int(puzzle_solution)


def test_part1_example1(example1_input, example1_solution):
    assert main.part1(example1_input) == example1_solution


def test_part2_example2(example2_input, example2_solution):
    assert main.part2(example2_input) == example2_solution
