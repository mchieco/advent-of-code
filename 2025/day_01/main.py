import pathlib
import sys
import math


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def part1(data: list[str]) -> int:
    ans = 0
    current = 50
    for line in data:
        direction = line[0]
        assert direction in ["L", "R"], "Direction is not 'L' or 'R'"
        rotations = int(line[1:])
        if direction == "L":
            current -= rotations
        else:
            current += rotations
        current = current % 100 
        if current == 0:
            ans += 1

    return ans
    

def part2(data: list[str]) -> int:
    ans = 0
    current = 50
    for line in data:
        direction = line[0]
        assert direction in ["L", "R"], "Direction is not 'L' or 'R'"
        rotations = int(line[1:])
        started_at_0 = current == 0
        if direction == "L":
            current -= rotations 
        else:
            current += rotations

        times_hit_0 = abs(math.floor(current / 100))
        if started_at_0 and direction == "L":
            times_hit_0 = times_hit_0 - 1
        ans += times_hit_0

        current = current % 100
        if current == 0 and direction == "L":
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
