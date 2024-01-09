input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    with open(file, "r") as file:
        stofs = file.read().strip().split("\n\n")
    for line in stofs:
        mirror = []
        lines = line.split("\n")
        for row in lines:
            mirror.append(row.strip())
        array.append(mirror)
    return array

mirrors = open_file_to(input)

def result_for(mirror):
    result = 100 * get_rows(mirror)
    if result != 0:
        return result
    return get_cols(mirror)


def rotate(array):
    width = len(array[0])
    height = len(array)
    rotated = []
    for col in range(width):
        line = ""
        for row in range(height):
            line = array[row][col] + line
        rotated.append(line)
    return rotated

def get_cols(array):
    rotated = rotate(array)
    return get_rows(rotated)

def get_rows(array):
    height = len(array)
    row_num = 0
    for row in range(height - 1):
        if array[row] != array[row + 1] or test(array, row):
            continue
        row_num = row + 1
    return row_num

def test(mirror, start):
    a = start
    b = start + 1
    while True:
        a -= 1
        b += 1
        if a < 0 or b == len(mirror):
            return False
        if mirror[a] != mirror[b]:
            return True

total = 0
for mirror in mirrors:
    total += result_for(mirror)
print(total)