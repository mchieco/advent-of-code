import pathlib
import sys
from collections import defaultdict


def parse_data(puzzle_input: str):
    parsed_input = puzzle_input.split("\n")
    parsed_input.remove("")
    return parsed_input


def part1(data: list[str]):
    ans = 0
    num_mapping = defaultdict(set)
    for line in data:
        first_items = line.split("|")
        if len(first_items) > 1:
            num_mapping[first_items[1]].add(first_items[0])
            continue
        second_items = line.split(",")
        for idx, num in enumerate(second_items):
            if not set(second_items[idx + 1 :]).isdisjoint(num_mapping.get(num, {})):
                break
            if idx == len(second_items) - 1:
                ans += int(second_items[(len(second_items) // 2)])
    return ans


def part2(data: list[str]):
    ans = 0
    num_mapping = defaultdict(set)
    for line in data:
        first_items = line.split("|")
        if len(first_items) > 1:
            num_mapping[first_items[0]].add(first_items[1])
            continue
        second_items = line.split(",")
        num_of_intersections = {}
        for idx, num in enumerate(second_items):
            num_of_intersections[num] = len(
                set((second_items[:idx] + second_items[idx + 1 :])).intersection(
                    num_mapping.get(num, {})
                )
            )

            if idx == len(second_items) - 1:
                sorted_second_items = list(
                    {
                        k: v
                        for k, v in sorted(
                            num_of_intersections.items(),
                            key=lambda item: item[1],
                            reverse=True,
                        )
                    }.keys()
                )
                if sorted_second_items != second_items:
                    ans += int(sorted_second_items[(len(sorted_second_items) // 2)])
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
