import pathlib
import sys
from collections import Counter


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def part1(data: list[str]) -> int:
    ans = 0
    left = []
    right = []
    for line in data:
        left_val, right_val = line.split()
        left.append(left_val)
        right.append(right_val)

    left.sort()
    right.sort()

    for left, right in zip(left, right):
        ans += abs(int(left) - int(right))

    return ans


def part2(data: list[str]):
    ans = 0
    left = []
    right = []
    for line in data:
        left_val, right_val = line.split()
        left.append(int(left_val))
        right.append(int(right_val))

    counter = Counter(right)
    for item in left:
        ans += item * counter.get(item, 0)

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
