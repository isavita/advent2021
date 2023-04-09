def read_instructions(file_path):
    with open(file_path, 'r') as f:
        return [line.strip().split() for line in f.readlines()]

def execute_instructions(instructions, inputs):
    variables = {'w': 0, 'x': 0, 'y': 0, 'z': 0}
    input_idx = 0

    for instr in instructions:
        op, a, *rest = instr
        b = rest[0] if rest else None

        if op == "inp":
            variables[a] = inputs[input_idx]
            input_idx += 1
        elif op == "add":
            variables[a] += int(b) if b.isdigit() else variables[b]
        elif op == "mul":
            variables[a] *= int(b) if b.isdigit() else variables[b]
        elif op == "div":
            divisor = int(b) if b.isdigit() else variables[b]
            if divisor != 0:
                variables[a] = variables[a] // divisor
        elif op == "mod":
            divisor = int(b) if b.isdigit() else variables[b]
            if 0 < divisor:
                variables[a] %= divisor
        elif op == "eql":
            variables[a] = 1 if variables[a] == (int(b) if b.isdigit() else variables[b]) else 0

    return variables

def find_largest_model_number(file_path):
    instructions = read_instructions(file_path)
    largest_number = 0

    def search(number, depth):
        nonlocal largest_number
        if depth == 14:
            variables = execute_instructions(instructions, [int(d) for d in str(number)])
            if variables['z'] == 0:
                largest_number = max(largest_number, number)
        else:
            for digit in range(1, 10):
                search(number * 10 + digit, depth + 1)

    search(0, 0)
    return largest_number

if __name__ == "__main__":
    file_path = "./test.txt"
    print(find_largest_model_number(file_path))
