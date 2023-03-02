import numpy as np

# read input file
with open('day4_part2_revision_3/test.txt') as f:
    data = f.readlines()

# split the data into board and numbers list
boards = []
numbers = []
for line in data:
    if ' ' in line:
        boards.append(line.strip())
    else:
        # strip any whitespace characters from both ends of the line before splitting it by commas
        nums = line.strip().split(',')
        if len(nums) > 1:
            numbers.extend([int(n) for n in nums])

# parse the board strings into numpy arrays
boards = [np.array([[int(cell) for cell in row.split()] for row in board.split('\n') if row]) for board in boards]

# set up some variables for keeping track of the game state
board_marks = [np.zeros_like(board, dtype=bool) for board in boards]
winning_board = None
winning_num = None
final_score = None

# loop over the drawn numbers
for num in numbers:
    # mark the number on all boards where it appears
    for board_mark in board_marks:
        board_mark[np.where(board_mark == False) and np.isin(boards, [board_mark])] = boards[np.isin(boards, [board_mark])] == num

    # check if any board has a winning row or column with this number
    for i, board in enumerate(boards):
        if np.any(np.all(board_marks[i][:,:5] == True, axis=0)):  # check columns
            winning_board = i
            winning_num = num
        elif np.any(np.all(board_marks[i][:,:5] == True, axis=1)):  # check rows
            winning_board = i
            winning_num = num

    # if there's a winning board, update the final score
    if winning_board is not None:
        final_score = np.sum(np.where(board_marks[winning_board], 0, boards[winning_board])) * winning_num

print(final_score)
