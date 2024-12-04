input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    with open(file, 'r') as file:
        for line in file:
            array.append(line.strip()) 
    return array
