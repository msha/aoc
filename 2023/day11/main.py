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
                        star = {"count":count, "x":x,"y":y}
                        xarray.append(star)
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
    for x,xc in enumerate(array):
        x_void = sum(999999 for number in empty_rows if number < x)
        for y,yc in enumerate(xc):
            y_void = sum(999999 for number in empty_columns if number < y)
            if yc != '.':
                yc['x'] = yc['x'] + y_void
                yc['y'] = yc['y'] + x_void
    
    return array
print('adjusting')

galaxy = adjust_void(np.array(lines))
print('finding')
import numpy as np

def find_routes(map, pairs):
    total = 0

    count_to_coords = {}
    for (x, y), star in np.ndenumerate(map):
        if type(star) is not str:
            count = star.get('count')
            count_to_coords[count] = (star.get('x'), star.get('y'))

    for pair in pairs:
        x1, y1 = count_to_coords[pair[0]]
        x2, y2 = count_to_coords[pair[1]]
        total += int(abs(x1-x2)+abs(y1-y2))

    return total


print(find_routes(galaxy,unique_pairs))

# print(galaxy)