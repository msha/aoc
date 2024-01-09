from itertools import combinations,permutations,product
import re
from functools import lru_cache

input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    with open(file, 'r') as file:
        for line in file:
            a = line.split(' ')[0].strip()
            b = tuple(map(int, line.split(' ')[1].strip().split(',')))
            array.append((a+('?'+a)*4,b*5)) 
    return array
values = open_file_to(input)

@lru_cache(maxsize=None)
def count_permutations(game):
    field, info = game
    if "?" not in field:
        return tuple(map(len, re.findall(r"#+", field))) == info

    question_mark_index = field.index("?")
    pattern = re.compile(r"#+\.")
    filled_segments = [(m.start(), m.end()) for m in pattern.finditer(field[:question_mark_index+1])]
    filled_lengths = [end-start-1 for start, end in filled_segments]

    if len(info) < len(filled_lengths) or any(info != length for info, length in zip(info, filled_lengths)):
        return 0

    next_filled_index = 0 if len(filled_segments) == 0 else filled_segments[-1][1]
    next_info = info[len(filled_lengths):]

    return sum(count_permutations((field[next_filled_index:question_mark_index] + fill + field[question_mark_index+1:], next_info)) for fill in ".#")

total = 0
for value in values:
    total += count_permutations(value)

print(total)