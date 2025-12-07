import pathlib
import sys


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")

def is_roll_of_paper(char: str) -> bool:
    if char == "@":
        return True
    return False

def rolls_of_paper_to_remove(data: list[str]) -> tuple[int, list[tuple[int, int]]]:
    print(data)
    rolls_to_remove = 0
    rolls_to_remove_coords: list[tuple[int, int]] = []
    for idx, line in enumerate(data):
        for idj, char in enumerate(line):
            rolls_of_paper = 0
            if char == "@":
                if idx != 0:
                    if is_roll_of_paper(data[idx - 1][idj]):
                        rolls_of_paper += 1
                if idx != len(data) -1: 
                    if is_roll_of_paper(data[idx + 1][idj]):
                        rolls_of_paper += 1
                if idj != 0:
                    if is_roll_of_paper(data[idx][idj - 1]):
                        rolls_of_paper += 1
                if idj != len(line) - 1:
                    if is_roll_of_paper(data[idx][idj + 1]):
                        rolls_of_paper += 1
                if idx != 0 and idj != 0:
                    if is_roll_of_paper(data[idx - 1][idj - 1]):
                        rolls_of_paper += 1
                if idx != 0 and idj != len(line) - 1: 
                    if is_roll_of_paper(data[idx - 1][idj + 1]):
                        rolls_of_paper += 1
                if idx != len(line) - 1 and idj != len(line) - 1:
                    if is_roll_of_paper(data[idx + 1][idj + 1]):
                        rolls_of_paper += 1
                if idx != len(line) - 1 and idj != 0:
                    if is_roll_of_paper(data[idx + 1][idj - 1]):
                        rolls_of_paper += 1

                if rolls_of_paper < 4:
                    rolls_to_remove += 1 
                    rolls_to_remove_coords.append((idx, idj))
    
    return rolls_to_remove, rolls_to_remove_coords

def part1(data: list[str]) -> int:
    return rolls_of_paper_to_remove(data)[0]

def part2(data: list[str]) -> int:
    ans = 0

    cur_state = data.copy()
    while True:
        rolls_to_remove, coords = rolls_of_paper_to_remove(cur_state)
        if rolls_to_remove:
            ans += rolls_to_remove
            for coord in coords:
                cur_state[coord[0]] = cur_state[coord[0]][:coord[1]] + "x" + cur_state[coord[0]][coord[1] + 1:]
        else:
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
