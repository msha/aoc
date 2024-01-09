import copy
from collections import deque
import numpy as np

lines = []
input = "input.txt"
test = "testinput.txt"

def open_file_to(file,array):
    with open(file, 'r') as file:
        # Read each line from the file and append to the list
        for y,line in enumerate(file):
            xarray = []
            for x,char in enumerate(line):
                if char != '\n':
                    xarray.append((char,(x,y))) 
            array.append(xarray)
            
def is_viable(next,direction):
    if next[0] == '.':
        return False
    next_viables = viable_next_dir(next)
    if next_viables.__contains__(direction):
        return True
    return False

def viable_dir(current):
    match current[0]:
        case '|':
            return ['N','S']
        case '-':
            return ['E','W']
        case 'L':
            return ['N','E']
        case 'J':
            return ['N','W']
        case '7':
            return ['S','W']
        case 'F':
            return ['S','E']
        case 'S':
            return ['N','E','S','W']

def viable_next_dir(current):
    match current[0]:
        case '|':
            return ['N','S']
        case '-':
            return ['E','W']
        case 'L':
            return ['S','W']
        case 'J':
            return ['S','E']
        case '7':
            return ['N','E']
        case 'F':
            return ['N','W']
        case 'S':
            return ['N','E','S','W']
        
        
       
def look_for_route(start, map):
    darray = [[-1 for _ in range(len(map[0]))] for _ in range(len(map))]
    tovisit = deque([start])
    visited = set()  
    x,y = start[1]
    darray[y][x] = 0  
    visited.add((y, x))  
    while tovisit:
        current = tovisit.popleft()
        x, y = current[1]
        current_distance = darray[y][x]
        for dir in viable_dir(current):
            nx, ny = x, y  #
            if dir == 'N':
                ny -= 1
            elif dir == 'E':
                nx += 1
            elif dir == 'S':
                ny += 1
            elif dir == 'W':
                nx -= 1

            if (ny, nx) not in visited and is_viable(map[ny][nx], dir):
                tovisit.append(map[ny][nx])
                visited.add((ny, nx))
                darray[ny][nx] = current_distance + 1 

    return darray

open_file_to(input,lines)




start_loc = (0,0)
for line in lines:
    for item in line:
        if item[0] == 'S':
            start_loc = item

distances = look_for_route(start_loc,lines)
print(start_loc)
print(distances)
highest = 0
for y,line in enumerate(lines):
    string = ''
    for x,xline in enumerate(line):
        if distances[y][x] >= 0:
            dis = distances[y][x]
            if dis > highest:
                highest = dis
            string += str(dis)
        else :
            string +=  'a' #xline[0]
    print(string)
print(highest)