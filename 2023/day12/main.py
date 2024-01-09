from itertools import combinations,permutations,product

input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    with open(file, 'r') as file:
        for line in file:
            a = line.split(' ')[0].strip()
            b = line.split(' ')[1].strip()
            array.append({"field":a,"info":b}) 
    return array

values = open_file_to(input)

import itertools

def count_permutations(field, info):
    # Convert the info string into a list of integers
    info = [int(x) for x in info.split(',')]
    
    # Generate all possible fillings for '?'
    fill_options = [(char if char != '?' else ['0', '1']) for char in field]
    all_combinations = list(product(*fill_options))
    
    valid_permutations = 0
    for combination in all_combinations:
        # Convert the tuple to a string, replacing '0' with '.'
        filled_field = ''.join(combination).replace('0', '.')
        
        # Check if the filled field matches the info pattern
        if is_valid(filled_field, info):
            valid_permutations += 1
    
    return valid_permutations

def is_valid(filled_field, info):
    # Split the filled field by '.' and filter out empty strings
    groups = [group for group in filled_field.split('.') if group]
    
    # Check if the number of '#' matches the info pattern
    return [len(group) for group in groups] == info

total = 0
for value in values:
    total += count_permutations(value['field'],value['info'])
    
print(total)




# print(values)