import pathlib
import sys
from copy import deepcopy


def parse_data(puzzle_input: str):
    return puzzle_input.split("\n")


def part1(data: list[str]):
    lab_map = dict()
    guard_pos = (0, 0, "")
    for idl, line in enumerate(data):
        for idc, char in enumerate(line):
            lab_map[f"({idl,idc})"] = char
            if char in ["<", ">", "^", "v"]:
                guard_pos = (idl, idc, char)

    while True:
        if not lab_map.get(f"({guard_pos[0], guard_pos[1]})"):
            return sum(value == "X" for value in lab_map.values())

        match guard_pos[2]:
            case "^":
                next_char = lab_map.get(f"({guard_pos[0] - 1, guard_pos[1]})")
                if next_char == "#":
                    guard_pos = (guard_pos[0], guard_pos[1], ">")
                else:
                    lab_map[f"({guard_pos[0],guard_pos[1]})"] = "X"
                    guard_pos = (guard_pos[0] - 1, guard_pos[1], "^")
            case "v":
                next_char = lab_map.get(f"({guard_pos[0] + 1, guard_pos[1]})")
                if next_char == "#":
                    guard_pos = (guard_pos[0], guard_pos[1], "<")
                else:
                    lab_map[f"({guard_pos[0],guard_pos[1]})"] = "X"
                    guard_pos = (guard_pos[0] + 1, guard_pos[1], "v")
            case "<":
                next_char = lab_map.get(f"({guard_pos[0], guard_pos[1] - 1})")
                if next_char == "#":
                    guard_pos = (guard_pos[0], guard_pos[1], "^")
                else:
                    lab_map[f"({guard_pos[0],guard_pos[1]})"] = "X"
                    guard_pos = (guard_pos[0], guard_pos[1] - 1, "<")
            case ">":
                next_char = lab_map.get(f"({guard_pos[0], guard_pos[1] + 1})")
                if next_char == "#":
                    guard_pos = (guard_pos[0], guard_pos[1], "v")
                else:
                    lab_map[f"({guard_pos[0],guard_pos[1]})"] = "X"
                    guard_pos = (guard_pos[0], guard_pos[1] + 1, ">")


def part2(data: list[str]):
    ans = 0
    lab_map = dict()
    starting_guard_pos = (0, 0, "")
    for idl, line in enumerate(data):
        for idc, char in enumerate(line):
            lab_map[f"{idl},{idc}"] = char
            if char in ["<", ">", "^", "v"]:
                starting_guard_pos = (idl, idc, char)

    num_coords = len(lab_map) / 2
    for coords, char in lab_map.items():
        x, y = [int(coord) for coord in coords.split(",")]
        lab_map_copy = deepcopy(lab_map)
        lab_map_copy[f"{x},{y}"] = "#"
        guard_pos = starting_guard_pos
        seen = 0

        while True:
            if not lab_map_copy.get(f"{guard_pos[0]},{guard_pos[1]}"):
                break

            if seen > num_coords:
                seen = 0
                ans += 1
                break

            match guard_pos[2]:
                case "^":
                    next_char = lab_map_copy.get(f"{guard_pos[0] - 1},{guard_pos[1]}")
                    if next_char == "#":
                        guard_pos = (guard_pos[0], guard_pos[1], ">")
                    else:
                        lab_map_copy[f"{guard_pos[0]},{guard_pos[1]}"] = "X"
                        guard_pos = (guard_pos[0] - 1, guard_pos[1], "^")
                case "v":
                    next_char = lab_map_copy.get(f"{guard_pos[0] + 1},{guard_pos[1]}")
                    if next_char == "#":
                        guard_pos = (guard_pos[0], guard_pos[1], "<")
                    else:
                        lab_map_copy[f"{guard_pos[0]},{guard_pos[1]}"] = "X"
                        guard_pos = (guard_pos[0] + 1, guard_pos[1], "v")
                case "<":
                    next_char = lab_map_copy.get(f"{guard_pos[0]},{guard_pos[1] - 1}")
                    if next_char == "#":
                        guard_pos = (guard_pos[0], guard_pos[1], "^")
                    else:
                        lab_map_copy[f"{guard_pos[0]},{guard_pos[1]}"] = "X"
                        guard_pos = (guard_pos[0], guard_pos[1] - 1, "<")
                case ">":
                    next_char = lab_map_copy.get(f"{guard_pos[0]},{guard_pos[1] + 1}")
                    if next_char == "#":
                        guard_pos = (guard_pos[0], guard_pos[1], "v")
                    else:
                        lab_map_copy[f"{guard_pos[0]},{guard_pos[1]}"] = "X"
                        guard_pos = (guard_pos[0], guard_pos[1] + 1, ">")
            seen += 1
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
