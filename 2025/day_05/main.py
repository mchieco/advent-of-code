import pathlib
import sys


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n\n")

def part1(data: list[str]) -> int:
    ans = 0

    ranges = data[0].split("\n")
    ingredients = [int(i) for i in data[1].split("\n")]

    for ingredient in ingredients:
        for r in ranges:
            start, end = r.split("-")
            if ingredient >= int(start) and ingredient <= int(end):
                ans += 1
                break

    return ans

def part2(data: list[str]) -> int:
    ans = 0
    ranges: list[tuple[int, int]] = []
    for r in data[0].split("\n"):
        start, end = r.split("-")
        start = int(start)
        end = int(end)
        ranges.append((int(start), int(end)))

    ranges = sorted(ranges)
    cur = 0
    for start, end in ranges:
        if cur >= start:
            start = cur + 1
        if start <= end:
            ans += end + 1 - start
        cur = max(cur, end)
 
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
