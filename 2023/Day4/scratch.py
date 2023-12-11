file_path = "input.txt"
test_path = "testinput.txt"
def read_file(file_path):
    lines = []
    with open(file_path, 'r') as file:
        # Read each line from the file and append to the list
        for line in file:
            id = line.split(':')[0].replace('Card ','')
            numbers = [item for item in line.split(':')[1].split('|')[0].strip().split(' ') if item != '']
            winning_numbs = [item for item in line.split(':')[1].split('|')[1].strip().split(' ') if item != '']
            lines.append({"id":id,"numbers":numbers,"winning_nums":winning_numbs,"value":1}) 
    return lines

def check_lines(cards):
    sum = 0
    for index,card in enumerate(cards):
        new_cards = 0
        for num in card["numbers"]:
            if num in card["winning_nums"]:
                new_cards += 1
        sum += int(card["value"])
        for copies in range(1,new_cards+1):
    
            cards[index+copies]["value"] += cards[index]["value"]
  
    return sum

cards = read_file(file_path)
     
print(check_lines(cards))