import pathlib
import sys


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def part1(data: list[str]) -> int:
    ans = 0
        
    for bank in data:
        first = 0
        second = 0
        for idx, battery in enumerate(bank):
            int_battery = int(battery)
            if int_battery > first and idx != len(bank)-1:
                first = int_battery
                second = int(bank[idx + 1])
            elif int_battery > second:
                second = int_battery
        ans += int(f"{first}{second}")
            
    return ans

def part2(data: list[str]) -> int:
    ans = 0
        
    for bank in data:
        bank_size = len(bank)
        start = 0
        max_joltage = ""
        for remaining in range(12, 0, -1):
            end = bank_size - remaining + 1
            idx, max_elm = max(enumerate(bank[start:end]), key=lambda x: x[1])
            start = start + idx + 1
            max_joltage += max_elm

        ans += int(max_joltage)
 
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
