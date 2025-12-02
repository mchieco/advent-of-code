import pathlib
import sys


def parse_data(puzzle_input: str):
    return puzzle_input.split(",")


def part1(data: list[str]) -> int:
    invalid_ids: set[int] = set()
    for id_range in data:
        start, end = id_range.split("-")
        for id in range(int(start), int(end) + 1):
            id_str = str(id)
            id_len = len(id_str)
            if id_len % 2 == 0:
                first = id_str[:id_len//2]
                second = id_str[id_len//2:id_len+1]
                if first == second:
                    invalid_ids.add(id)

    return sum(invalid_ids)

def part2(data: list[str]) -> int:
    invalid_ids: set[int] = set()
    for id_range in data:
        start, end = id_range.split("-")
        for id in range(int(start), int(end) + 1):
            id_str = str(id)
            id_len = len(id_str)
            for interval in range(1,id_len):
                combos = [id_str[i:i+interval] for i in range(0, len(id_str), interval)]
                if all(x==combos[0] for x in combos):
                    invalid_ids.add(id)
                    break

    return sum(invalid_ids)    

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)


if __name__ == "__main__":
    path = sys.argv[1]
    print(f"\n{path}:")
    solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
    print("\n".join(str(solution) for solution in solutions))
