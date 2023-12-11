strings = []
lines = []
file_path = "input.txt"
test_path = "testinput.txt"

def read_file_to(path,array):
    with open(path, 'r') as file:
        # Read each line from the file and append to the list
        id = -1
        for line in file:
            row = []
            num = ""
            id += 1
            for i,char in enumerate(line):
                value = {"value":None,"is_valid":False,"is_digit":False, "id": None}
                if not char.isdigit() and char == '*' and char != '\n' and char != '\r\n':
                    value = {"value":None,"is_valid":True,"is_digit":False,"id":None}
                if char.isdigit():
                    num += char
                    value = {"value":int(num),"is_valid":False,"is_digit":True,"id":id}
                    index = i
                    while index > 0:
                        index -= 1
                        if row[index]["is_digit"]:
                            row[index]["value"] = num
                        else:
                            break
                else:
                    num = ""
                    id += 1
                row.append(value)
            array.append(row)

def update_valid(array):
    seen_ids = set()
    final_array = []
    for yi,row in enumerate(array):
        for xi,value in enumerate(row):
            if value["is_valid"] and not value["is_digit"]:
                maybe_array = []
                for yindex in range(yi-1,yi+2):
                    if yindex >= 0 and yindex <= len(array):
                        for xindex in range(xi-1,xi+2):
                            if xindex >= 0 and xindex <= len(array[0]) and array[yindex][xindex]["is_digit"]:
                                array[yindex][xindex]["is_valid"] = True
                                if array[yindex][xindex]["id"] not in seen_ids:
                                    maybe_array.append(int(array[yindex][xindex]["value"]))
                                seen_ids.add(array[yindex][xindex]["id"])
                if len(maybe_array) == 2:
                    final_array.append(maybe_array[0]*maybe_array[1])
                                
    return final_array
                    
            
read_file_to(file_path,lines)
read_file_to(test_path,strings)

sum = 0
for num in update_valid(lines):
    sum += num
print(sum)