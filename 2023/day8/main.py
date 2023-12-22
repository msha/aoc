input = "input.txt"
test = "testinput.txt"
from math import gcd
from functools import reduce

def open_file(input_file):
    instruct = ''
    array = []
    with open(input_file, 'r') as file:
        for i,line in enumerate(file):
            if i == 0:
                instruct += line.strip()
            else:
                node = line.strip().split(' ')[0].strip()
                L = line.strip().split('(')[1][0:3]
                R = line.strip().split('(')[1][5:8]
                array.append({"node":node,"L":L,"R":R}) 
    return {"instructions": instruct, "map": array}

data = open_file(input)

instructions = data["instructions"].strip()
full_map = data["map"]
current_tiles = [d for d in full_map if d['node'][2:3] == 'A']
steps = 0


ghost_intervals = []

for ghost_path in current_tiles:
    steps = 0
    current_tile = ghost_path
    while not (current_tile['node'].strip()[2:3] == 'Z'): 
        current_instruction = instructions[steps%len(instructions)]
        new_node = [d for d in full_map if d['node'] == current_tile[current_instruction]]
        current_tile = new_node[0]
        steps += 1
    ghost_intervals.append(steps)

def lcm(a, b):
    return a * b // gcd(a, b)

def find_lcm(numbers):
    return reduce(lcm, numbers, 1)
    
print(find_lcm(ghost_intervals))