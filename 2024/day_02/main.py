import pathlib
import sys
from itertools import islice


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def within_bounds(num1: int, num2: int) -> bool:
    return 1 <= abs(num1 - num2) <= 3


def check_safety(items) -> bool:
    increasing = None
    for idx, item in enumerate(islice(items, len(items) - 1)):
        num = int(item)
        next_num = int(items[idx + 1])
        if num > next_num:
            if increasing is None:
                increasing = False
            if increasing is True or not within_bounds(num, next_num):
                return False
        elif num < next_num:
            if increasing is None:
                increasing = True
            if increasing is False or not within_bounds(num, next_num):
                return False
        else:
            return False

        if idx == len(items) - 2:
            return True

    return False


def part1(data: list[str]):
    ans = 0
    for line in data:
        items = line.split()
        safe = check_safety(items)
        if safe:
            ans += 1

    return ans


def part2(data: list[str]):
    ans = 0
    for line in data:
        items = line.split()
        safe = False
        for j in range(len(items)):
            items2 = items[:j] + items[j + 1 :]
            if check_safety(items2):
                safe = True
        if safe:
            ans += 1

    return ans


def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)


if __name__ == "__main__":
    path = sys.argv[1]
    print(f"\n{path}:")
    solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
    print("\n".join(str(solution) for solution in solutions))
