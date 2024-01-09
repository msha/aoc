import numpy as np
from itertools import combinations
input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    count = 1
    with open(file, 'r') as file:
        # Read each line from the file and append to the list
        for y,line in enumerate(file):
            xarray = []
            for x,char in enumerate(line):
                if char != '\n':
                    if char != '#':
                        xarray.append(char)
                    else:
                        xarray.append(count)
                        count += 1
            array.append(xarray)
    return array,count

lines,count = open_file_to(input)

numbers = range(1, count) 
unique_pairs = list(combinations(numbers, 2))
visited = list()

def adjust_void(array):
    empty_columns = np.where((array == '.').all(axis=0))[0]
    empty_rows = np.where((array == '.').all(axis=1))[0]
    
    for column in reversed(empty_columns):
        new_void = np.full((1,array.shape[0]), '.')
        array = np.insert(array, column + 1, new_void, axis=1)
    
    for row in reversed(empty_rows):
        new_void = np.full((1, array.shape[1]),'.')
        array = np.insert(array, row + 1, new_void, axis=0)
    return array

galaxy = adjust_void(np.array(lines))

def find_routes(map,pairs):
    total = 0
    for pair in pairs:
        x = np.where(map == str(pair[0]))[0]
        y = np.where(map == str(pair[0]))[1]
        x2 = np.where(map == str(pair[1]))[0]
        y2 = np.where(map == str(pair[1]))[1]
        total += int(abs(x-x2)+abs(y-y2))
    return total

print(find_routes(galaxy,unique_pairs))

print(galaxy)