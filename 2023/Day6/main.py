lines = []
input = "input.txt"
test = "testinput.txt"

def open_file_to(file):
    array = []
    with open(file, 'r') as file:
        # Read each line from the file and append to the list
        line = file.read()
        times = int(line.split("\n")[0].split(':')[1].replace(" ",""))
        distance = int(line.split("\n")[1].split(':')[1].replace(" ",""))
        array.append(times) 
        array.append(distance)
    return array

def get_best_times(value):
    index = 0
    best_times = []
    while True:
        race_time = value[0]
        race_dist = value[1]
        new_best = set()
        for time in range(1,1000000000):
            new_dist = (race_time - time)*time
            if race_dist < new_dist:
                new_best.add(time)
                index += 1
            elif index > 0 and race_dist > new_dist:
                break
        best_times.append(new_best)
        break
    return best_times

num = 1
for times in get_best_times(open_file_to(input)):
    num = len(times)
print(num)
