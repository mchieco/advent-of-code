import pathlib
import sys
from collections import defaultdict
import math
from itertools import zip_longest

def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")

def part1(data: list[str]) -> int:
    ans = 0

    problems = defaultdict(list)

    last_row = len(data)-1 
    for idx, line in enumerate(data):
        for idj, column in enumerate(line.split()):
            if idx == last_row:
                if column == "*":
                    ans += math.prod(problems[idj])
                else:
                    ans += sum(problems[idj])
            else:
                problems[idj].append(int(column)) 

    return ans

def part2(data: list[str]) -> int:
    ans = 0
    
    problems = []
    for line in data:
        problems.append(list(line))
    
    current_problem = [] 
    for i in reversed(list(zip_longest(*problems, fillvalue=" "))):
        if all(x == " " for x in i):
            current_problem = []
        else:
            current_problem.append(int("".join(i[:-1]).strip())) 
        if i[-1] == "*":
            ans += math.prod(current_problem)
        elif i[-1] == "+":
            ans += sum(current_problem) 
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
