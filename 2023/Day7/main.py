hands = []
input = "input.txt"
test = "test.txt"

card_values = {'A': 13, 'K': 12, 'Q': 11, 'T': 10, '9': 9, '8': 8,'7': 7, '6': 6,'5': 5, '4': 4,'3': 3, '2': 2, 'J': 1}

empty_hand = {key: 0 for key in card_values}
value_to_key = {value:key for key,value in card_values.items()}

hand_value = {'High card': 1, 'One pair': 2, 'Two pair': 3, 'Three of a kind': 4, 'Full house': 5, 'Four of a kind': 6, 'Five of a kind': 7}

def sort_hands(hands):
    hands.sort(key=calculate_value)
    return hands

def open_file_to(file,array):
    with open(file, 'r') as file:
        # Read each line from the file and append to the list
        for line in file:
            cards = line.split(' ')[0]
            bet = line.split(' ')[1].strip()
            array.append({"cards":cards, "bet": bet}) 

def calculate_value(hand):
    values = empty_hand.copy()
    values_without_j = empty_hand.copy()
    hand_value = 1
    card_string = ''
    j_amount = 0
    non_j = 0
    for card in hand["cards"]:
        if card == 'J':
            j_amount += 1
            for i in range(2,14):
                values[value_to_key[i]] += 1
        else:
            non_j += 1
            values_without_j[card] += 1
            values[card] += 1
        card_string += f"{card_values[card]:02d}"
    pairs = set()
    non_j_pairs = set()
    for value in values.values():
        if value > 1:
            if pairs.__contains__(2) and value == 2:
                pairs.add(22)
            pairs.add(value)
    for value in values_without_j.values():
        if value > 1:
            if non_j_pairs.__contains__(2) and value == 2:
                non_j_pairs.add(22)
            non_j_pairs.add(value)
    if pairs.__contains__(2):
        hand_value = 2
    if pairs.__contains__(22) and j_amount == 0:
        hand_value = 3
    if pairs.__contains__(3):
        hand_value = 4
    if pairs.__contains__(3) and pairs.__contains__(2) and not pairs.__contains__(22) :
        hand_value = 5
    if non_j == 4 and j_amount == 1 and non_j_pairs.__contains__(22):
        hand_value = 5
    if pairs.__contains__(4):
        hand_value = 6
    if pairs.__contains__(5):
        hand_value = 7
    
    check_sum = int(hand_value.__str__()+card_string)
    hand["check_sum"] = check_sum
    return check_sum


open_file_to(input,hands)

total_sum = 0
for i,play in enumerate(sort_hands(hands)):
    bet = play["bet"]
    print(play)
    total_sum += int(bet) * (i+1)
print(total_sum)