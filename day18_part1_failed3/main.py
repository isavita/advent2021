import ast
import copy

CIRCULAR_REFERENCE = "CIRCULAR_REFERENCE"

def is_regular_number(num):
    return not isinstance(num, list)

def parse_input(file_path):
    with open(file_path, "r") as file:
        data = file.readlines()
    snailfish_numbers = []
    for line in data:
        line = line.strip()
        line = line.replace("[...]", f'"{CIRCULAR_REFERENCE}"')
        snailfish_number = ast.literal_eval(line)
        snailfish_numbers.append(snailfish_number)
    return snailfish_numbers

def snailfish_add(a, b):
    result = [a, b]
    result = replace_circular_references(result, result)
    return result

def replace_circular_references(snailfish_number, replacement):
    if is_regular_number(snailfish_number):
        return snailfish_number
    if snailfish_number == CIRCULAR_REFERENCE:
        return replacement
    return [replace_circular_references(x, replacement) for x in snailfish_number]

def explode(snailfish_number):
    path = []
    exploded = True
    while exploded:
        path = []
        exploded = False
        def helper(num, depth=0):
            nonlocal exploded
            if is_regular_number(num):
                return
            if num != [0, 0]:
                depth += 1
            path.append((num, depth))
            if depth == 5:
                left, right = num[0], num[1]
                for p, _ in reversed(path[:-1]):
                    if is_regular_number(p[0]):
                        p[0] = snailfish_add(p[0], left)
                        break
                for p, _ in path[:-1]:
                    if is_regular_number(p[1]):
                        p[1] = snailfish_add(p[1], right)
                        break
                num[0], num[1] = 0, 0
                exploded = True
                return
            helper(num[0], depth)
            helper(num[1], depth)
            path.pop()
        helper(snailfish_number)
    return snailfish_number

def split(snailfish_number):
    if is_regular_number(snailfish_number):
        if snailfish_number >= 10:
            return [snailfish_number // 2, snailfish_number - snailfish_number // 2]
        else:
            return snailfish_number
    else:
        return [split(snailfish_number[0]), split(snailfish_number[1])]

def reduce(snailfish_number):
    while True:
        before = copy.deepcopy(snailfish_number)
        snailfish_number = explode(snailfish_number)
        snailfish_number = split(snailfish_number)
        if snailfish_number == before:
            break
    return snailfish_number

def magnitude(snailfish_number):
    if is_regular_number(snailfish_number):
        return snailfish_number
    else:
        return 3 * magnitude(snailfish_number[0]) + 2 * magnitude(snailfish_number[1])

def calculate_total_magnitude(snailfish_numbers):
    total = [0, 0]
    for number in snailfish_numbers:
        total = reduce(snailfish_add(total, number))
    return magnitude(total)

# Main function
def main():
    snailfish_numbers = parse_input("day18_part1/test.txt")
    total_magnitude = calculate_total_magnitude(snailfish_numbers)
    print(total_magnitude)

if __name__ == "__main__":
    main()