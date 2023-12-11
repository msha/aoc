
file_path = "input.txt"
test_path = "testinput.txt"
def read_file(file_path):
    lines = []
    with open(file_path, 'r') as file:
        # Read each line from the file and append to the list
        content = file.read()
        for line in content.split('\n\n'):
            key = line.split(':')[0].strip()
            values = line.split(':')[1].strip().split('\n')
            lines.append({"key":key,"value":values}) 
    return lines
values = read_file(file_path)
seeds = values[0]["value"][0].split(' ')
modifiers = [{"id":value["key"],"mods":[item.split(' ') for item in value["value"]]} for value in values if value["key"] != 'seeds']
best_in_mods = {}

for mod in modifiers:
    best_in_mods[mod["id"]] = {"seed_id":999999999999999999999,"value":999999999999999999999}

def eval_mods(mods):
    for mod in mods:
        for value in mod["mods"]:
            destiny_start = int(value[0])
            source_start = int(value[1])
            range = int(value[2])
    return mods

def get_mod_values(seed_id):
    seed_id = int(seed_id)
    for mod in modifiers:
        for value in mod["mods"]:
            destiny_start = int(value[0])
            source_start = int(value[1])
            range = int(value[2])
            if seed_id >= source_start and seed_id <= source_start + range:
                if destiny_start < source_start:
                    seed_id = seed_id + (destiny_start - source_start)
                    break
                elif destiny_start > source_start:
                    seed_id = seed_id - (source_start - destiny_start)
                    break
    return seed_id

def get_seed_values(location):
    location = int(location)
    for mod in reversed(modifiers):
        for value in mod["mods"]:
            destiny_start = int(value[0])
            source_start = int(value[1])
            range = int(value[2])
            if location >= destiny_start and location <= destiny_start + range:
                if destiny_start < source_start:
                    location = location - (destiny_start - source_start)
                    break
                elif destiny_start > source_start:
                    location = location + (source_start - destiny_start)
                    break
    return location

def make_seeds(seeds):
    seed_list = set()
    index = 0
    while index < len(seeds):
        start_seed = int(seeds[index])
        seed_range = int(seeds[index+1])
        for i in range(start_seed,start_seed+seed_range+1):
            seed_list.add(i)
        index += 2
    return seed_list

def get_min_max_seeds(seeds):
    seed_list = []
    index = 0
    while index < len(seeds):
        start_seed = int(seeds[index])
        range_seed = int(seeds[index+1])
        seed_list.append({"min":start_seed,"max":start_seed+range_seed})
        index += 2
    return seed_list


def get_low(seeds):
    lowest = 99999999999999999999999999
    for seed in seeds:
        seed_val = get_mod_values(seed)
        if seed_val < lowest:
            lowest = seed_val
    return lowest


# print(eval_mods(modifiers))

# print(best_in_mods)

#print(len(make_seeds(seeds)))
#print(get_min_max_seeds(seeds))
# print(make_seeds(seeds))

# print(get_low(['82']))
# print(make_seeds(seeds))
# print(get_min_max_seeds(seeds))
location = 10834441 #lowest found, but a lower value exists!
location = 10834440
while True:
    found = False
    seed = get_seed_values(location)
    for maxes in get_min_max_seeds(seeds):
        if seed > maxes['min'] and seed < maxes['max']:
            print(seed,location)
            found = True
            break
    if found:
        break
    if location % 100000 == 0:
        print('going over ', location)
    if location < 0:
        print('lowest was found!')
        break
    location -= 1