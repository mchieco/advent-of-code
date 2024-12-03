import pathlib
import re
import sys

MUL_REGEX = r"mul\((\d{1,3}),(\d{1,3})\)"
DO_DONT_REGEX = r"do\(\)|don't\(\)"


def parse_data(puzzle_input: str):
    return puzzle_input


def part1(data: str):
    ans = 0
    for i in re.findall(MUL_REGEX, data):
        ans += int(i[0]) * int(i[1])
    return ans


def part2(data: str):
    ans = 0
    for i in re.finditer(MUL_REGEX, data):
        end_bound = i.start(0)
        do_dont_list = list(re.finditer(DO_DONT_REGEX, data[0:end_bound]))
        if do_dont_list:
            do_or_dont = do_dont_list[-1].group(0)
            if do_or_dont == "do()":
                ans += int(i.group(1)) * int(i.group(2))
        else:
            ans += int(i.group(1)) * int(i.group(2))

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
