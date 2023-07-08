import json

def parse_snail_fish_numbers(input_string):
    lines = input_string.split("\n")
    input_list = []
    for line in lines:
        line = line.replace(" ", "")
        line = line.replace("(", "[")
        line = line.replace(")", "]")
        try:
            snail_fish_number = json.loads(line)
            input_list.append(snail_fish_number)
        except json.JSONDecodeError as e:
            print(f"Invalid input: {line}")
            print(f"Error: {e}")
            return []
    return input_list

def explode(snail_fish_number, depth=0):
    if isinstance(snail_fish_number, int):
        return snail_fish_number, False
    left, exploded = explode(snail_fish_number[0], depth + 1)
    if exploded:
        return [left, snail_fish_number[1]], True
    right, exploded = explode(snail_fish_number[1], depth + 1)
    if exploded:
        return [snail_fish_number[0], right], True
    if depth >= 4:
        return [0, [snail_fish_number[0], snail_fish_number[1]], [0, [snail_fish_number[0], snail_fish_number[1]]]], True
    return snail_fish_number, False

def split(snail_fish_number):
    if type(snail_fish_number) is int:
        if snail_fish_number >= 10:
            return [snail_fish_number // 10, snail_fish_number % 10], True
        else:
            return snail_fish_number, False
    else:
        left, split_result_left = split(snail_fish_number[0])
        if split_result_left:
            return [left, snail_fish_number[1]], True
        right, split_result_right = split(snail_fish_number[1])
        if split_result_right:
            return [snail_fish_number[0], right], True
        return snail_fish_number, False


def reduce(snail_fish_number):
    while True:
        snail_fish_number, exploded = explode(snail_fish_number)
        if exploded:
            continue
        snail_fish_number, split_result = split(snail_fish_number)
        if split_result:
            continue
        return snail_fish_number


def add_snail_fish_numbers(snail_fish_numbers):
    sum = snail_fish_numbers[0]
    for snail_fish_number in snail_fish_numbers[1:]:
        sum = [sum, snail_fish_number]
    sum = reduce(sum)
    return sum

def magnitude(snail_fish_number):
    if isinstance(snail_fish_number, int):
        return snail_fish_number
    return magnitude(snail_fish_number[0]) + magnitude(snail_fish_number[1])

def main(input_string):
    snail_fish_numbers = parse_snail_fish_numbers(input_string)
    if not snail_fish_numbers:
        return "Invalid input: the snailfish numbers are not properly formatted."
    magnitudes = []
    for snail_fish_number in snail_fish_numbers:
        sum = add_snail_fish_numbers(snail_fish_number)
        final_magnitude = magnitude(sum)
        magnitudes.append(final_magnitude)
    return magnitudes

input_string = "[[[0, [5, 8]], [[1, 7], [9, 6]]], [[4, [1, 2]], [[1, 4], 2]]]\n[[[5, [2, 8]], 4], [5, [[9, 9], 0]]]"
print(main(input_string))