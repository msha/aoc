import re

strings = ["two1nine","eightwothree","abcone2threexyz","xtwone3four","4nineeightseven2","zoneight234","7pqrstsixteen"]
digits= ["one","two","three","four","five","six","seven","eight","nine"]
lines = []
file_path = "input.txt"
with open(file_path, 'r') as file:
    # Read each line from the file and append to the list
    for line in file:
        lines.append(line.strip()) 



def find_matching_strings(main_string, string_list):
    # Create a regex pattern to match any of the strings in the list
    pattern = '|'.join(map(re.escape, string_list))
    
    # Split the main string using the pattern
    split_strings = re.split(pattern, main_string)

    # Filter out the matching strings
    matching_strings = [s for s in split_strings if s in string_list]

    return matching_strings

def get_calibrations(strings):
    total = 0
    total_sum = 0
    for string in strings:
        ostring = string
        first_int = None
        string = string.replace("one", "one1one")
        string = string.replace("two", "two2two")
        string = string.replace("three", "three3three")
        string = string.replace("four", "four4four")
        string = string.replace("five", "five5five")
        string = string.replace("six", "six6six")
        string = string.replace("seven", "seven7seven")
        string = string.replace("eight", "eight8eight")
        string = string.replace("nine", "nine9nine")
        for c in string:
            if c.isdigit():
                if first_int is None:
                    first_int = c
                last_digit = c
        total = int(f'{first_int}{last_digit}')
        print(f'{ostring} // {string} calibr is {total}')
        total_sum += total
    return total_sum
print(f'total sum {get_calibrations(lines)}')