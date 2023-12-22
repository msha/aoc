input = "input.txt"
test = "testinput.txt"

def open_file(file):
    instruct = ''
    array = []
    with open(file, 'r') as file:
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
current_tile = [d for d in full_map if d['node'] == 'AAA'][0]
steps = 0

while current_tile["node"].strip() != 'ZZZ':

    current_instruction = instructions[steps%len(instructions)]
    new_node = [d for d in full_map if d['node'] == current_tile[current_instruction]]
    current_tile = new_node[0]
    
    steps += 1

print(steps)