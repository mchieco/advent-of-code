import itertools
import pathlib
import sys


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def part1(data: list[str]):
    ans = 0
    for line in data:
        test, n = line.split(":")
        nums = n.split()
        operands_combinations = itertools.product(["+", "*"], repeat=len(nums) - 1)
        for combinations in operands_combinations:
            result = 0
            for idx, operand in enumerate(combinations):
                if operand == "+":
                    if idx == 0:
                        result += int(nums[idx]) + int(nums[idx + 1])
                    else:
                        result += int(nums[idx + 1])
                else:
                    if idx == 0:
                        result += int(nums[idx]) * int(nums[idx + 1])
                    else:
                        result *= int(nums[idx + 1])
            if result == int(test):
                ans += result
                break

    return ans


def part2(data: list[str]):
    ans = 0
    for line in data:
        test, n = line.split(":")
        nums = n.split()
        operands_combinations = itertools.product(
            ["+", "*", "||"], repeat=len(nums) - 1
        )
        for combinations in operands_combinations:
            result = 0
            for idx, operand in enumerate(combinations):
                if operand == "+":
                    if idx == 0:
                        result += int(nums[idx]) + int(nums[idx + 1])
                    else:
                        result += int(nums[idx + 1])
                elif operand == "*":
                    if idx == 0:
                        result += int(nums[idx]) * int(nums[idx + 1])
                    else:
                        result *= int(nums[idx + 1])
                elif operand == "||":
                    if idx == 0:
                        result = int(str(nums[idx]) + nums[idx + 1])
                    else:
                        result = int(str(result) + nums[idx + 1])

            if result == int(test):
                ans += result
                break

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
