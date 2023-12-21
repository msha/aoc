hands = []
input = "input.txt"
test = "test.txt"

card_values = {'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8,'7': 7, '6': 6,'5': 5, '4': 4,'3': 3, '2': 2,}

empty_hand = {key: 0 for key in card_values}

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
    hand_value = 1
    card_string = ''
    for card in hand["cards"]:
        values[card] += 1
        card_string += f"{card_values[card]:02d}"
    pairs = set()
    for value in values.values():
        if value > 0:
            if pairs.__contains__(2) and value == 2:
                pairs.add(6)
            pairs.add(value)
    if pairs.__contains__(5):
        hand_value = 7
    if pairs.__contains__(4):
        hand_value = 6
    if pairs.__contains__(3):
        hand_value = 4
    if pairs.__contains__(2):
        hand_value = 2
    if pairs.__contains__(3) and pairs.__contains__(2):
        hand_value = 5
    if pairs.__contains__(6):
        hand_value = 3
    check_sum = int(hand_value.__str__()+card_string)
    hand["check_sum"] = check_sum
    return check_sum


open_file_to(input,hands)

total_sum = 0
for i,play in enumerate(sort_hands(hands)):
    bet = play["bet"]
    total_sum += int(bet) * (i+1)
print(total_sum)
# print(hands)