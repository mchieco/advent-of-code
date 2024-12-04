import pathlib
import sys


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def part1(data: list[str]):
    ans = 0
    lines = len(data)
    line_chars = len(data[0])
    for line in range(lines):
        for c in range(line_chars):
            if (
                c + 3 < line_chars
                and data[line][c] == "X"
                and data[line][c + 1] == "M"
                and data[line][c + 2] == "A"
                and data[line][c + 3] == "S"
            ):
                ans += 1
            if (
                line + 3 < lines
                and data[line][c] == "X"
                and data[line + 1][c] == "M"
                and data[line + 2][c] == "A"
                and data[line + 3][c] == "S"
            ):
                ans += 1
            if (
                line + 3 < lines
                and c + 3 < line_chars
                and data[line][c] == "X"
                and data[line + 1][c + 1] == "M"
                and data[line + 2][c + 2] == "A"
                and data[line + 3][c + 3] == "S"
            ):
                ans += 1
            if (
                c + 3 < line_chars
                and data[line][c] == "S"
                and data[line][c + 1] == "A"
                and data[line][c + 2] == "M"
                and data[line][c + 3] == "X"
            ):
                ans += 1
            if (
                line + 3 < lines
                and data[line][c] == "S"
                and data[line + 1][c] == "A"
                and data[line + 2][c] == "M"
                and data[line + 3][c] == "X"
            ):
                ans += 1
            if (
                line + 3 < lines
                and c + 3 < line_chars
                and data[line][c] == "S"
                and data[line + 1][c + 1] == "A"
                and data[line + 2][c + 2] == "M"
                and data[line + 3][c + 3] == "X"
            ):
                ans += 1
            if (
                line - 3 >= 0
                and c + 3 < line_chars
                and data[line][c] == "S"
                and data[line - 1][c + 1] == "A"
                and data[line - 2][c + 2] == "M"
                and data[line - 3][c + 3] == "X"
            ):
                ans += 1
            if (
                line - 3 >= 0
                and c + 3 < line_chars
                and data[line][c] == "X"
                and data[line - 1][c + 1] == "M"
                and data[line - 2][c + 2] == "A"
                and data[line - 3][c + 3] == "S"
            ):
                ans += 1

    return ans


def part2(data: list[str]):
    ans = 0
    lines = len(data)
    line_chars = len(data[0])
    for line in range(lines):
        for c in range(line_chars):
            if (
                line + 2 < line_chars
                and c + 2 < line_chars
                and data[line][c] == "M"
                and data[line + 1][c + 1] == "A"
                and data[line + 2][c] == "M"
                and data[line][c + 2] == "S"
                and data[line + 2][c + 2] == "S"
            ):
                ans += 1
            if (
                line + 2 < line_chars
                and c + 2 < line_chars
                and data[line][c] == "M"
                and data[line + 1][c + 1] == "A"
                and data[line + 2][c] == "S"
                and data[line][c + 2] == "M"
                and data[line + 2][c + 2] == "S"
            ):
                ans += 1
            if (
                line + 2 < line_chars
                and c + 2 < line_chars
                and data[line][c] == "S"
                and data[line + 1][c + 1] == "A"
                and data[line + 2][c] == "M"
                and data[line][c + 2] == "S"
                and data[line + 2][c + 2] == "M"
            ):
                ans += 1
            if (
                line + 2 < line_chars
                and c + 2 < line_chars
                and data[line][c] == "S"
                and data[line + 1][c + 1] == "A"
                and data[line + 2][c] == "S"
                and data[line][c + 2] == "M"
                and data[line + 2][c + 2] == "M"
            ):
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
