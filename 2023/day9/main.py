input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    with open(file, 'r') as file:
        # Read each line from the file and append to the list
        for line in file:
            values = []
            for value in line.split(' '):
                values.append(value.strip())
            values.reverse()
            array.append(values) 
    return array

def get_differences(array, accumulated_sum):
    differences = [int(array[i]) - int(array[i-1]) for i in range(1, len(array))]

    if len(differences) == 1:
        return differences[0] + accumulated_sum
    return get_differences(differences, differences[-1] + accumulated_sum)


data_set = open_file_to(input)


sum = 0

for array in data_set:
    sum += int(array[-1]) + get_differences(array, 0)

print(sum)