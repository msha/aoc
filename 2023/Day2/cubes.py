strings = ["Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green","Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue","Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red","Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red","Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"]
lines = []
file_path = "input.txt"
with open(file_path, 'r') as file:
    # Read each line from the file and append to the list
    for line in file:
        lines.append(line.strip()) 
        
def read_lines(input):
    output = []
    for line in input:
        content = line.split(':')
        id = content[0].replace('Game ','')
        strings = content[1].split(';')
        maxred = 0
        maxblue = 0
        maxgreen = 0
        minred = 9999999999999
        minblue = 999999999999
        mingreen = 99999999999
        for string in strings:
            stuff = string.strip()
            values = stuff.split(',')
            coloramounts = []
            for value in values:
                value = value.strip()
                amount = int(value.split(' ')[0].strip())
                color = value.split(' ')[1].strip()
                if color == "green" and amount > maxgreen:
                    maxgreen = amount
                if color == "red" and amount > maxred:
                    maxred = amount
                if color == "blue" and amount > maxblue:
                    maxblue = amount
                if color == "green" and amount < mingreen:
                    mingreen = amount
                if color == "red" and amount < minred:
                    minred = amount
                if color == "blue" and amount < minblue:
                    minblue = amount
                coloramounts.append({"color":color,"amount":amount})
        game_stats = {"gameid":id,"maxred":maxred,"maxblue":maxblue,"maxgreen":maxgreen,"minred":minred,"minblue":minblue,"mingreen":mingreen}
        print(game_stats)
        output.append(game_stats)
    return output

def get_max_games(red,green,blue,input):
    possible_ids = []
    for game in read_lines(input):
        if game["maxblue"] <= blue and game["maxgreen"] <= green and game["maxred"] <= red:
            possible_ids.append(game["gameid"])
            
    return possible_ids

def get_id_sum(red,green,blue,input):
    sum = 0
    possible_games = get_max_games(red,green,blue,input)
    print(possible_games)
    for id in possible_games:
        sum += int(id)
    return sum

def get_power_of_games(input):
    sum = 0
    for game in read_lines(input):
        sum += game["maxblue"] * game["maxred"] * game["maxgreen"]
    return sum
    
# print(get_id_sum(12,13,14,lines))
print(get_power_of_games(lines))