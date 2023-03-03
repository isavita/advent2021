# 'day4_part2_revision_4/test.txt'
import numpy as np

def calculate_score(winning_board, last_number):
    unmarked_numbers = winning_board[winning_board == 0]
    score = last_number * unmarked_numbers.sum()
    return score

def check_rows(board):
    for i in range(board.shape[0]):
        if 0 not in board[i]:
            return True, i
    return False, None

def check_columns(board):
    for i in range(board.shape[1]):
        if 0 not in board[:, i]:
            return True, i
    return False, None

def check_diagonals(board):
    diagonal = board.diagonal()
    antidiagonal = np.fliplr(board).diagonal()
    if 0 not in diagonal:
        return True, 0
    elif 0 not in antidiagonal:
        return True, 1
    else:
        return False, None

def check_board(board):
    rows_win, row_index = check_rows(board)
    columns_win, col_index = check_columns(board)
    diagonals_win, diag_index = check_diagonals(board)
    
    if rows_win:
        return True, board[row_index, :]
    elif columns_win:
        return True, board[:, col_index]
    elif diagonals_win:
        if diag_index == 0:
            return True, board.diagonal()
        else:
            return True, np.fliplr(board).diagonal()
    else:
        return False, None

def main():
    # read input from file
    with open('day4_part2_revision_4/input.txt') as f:
        numbers = list(map(int, f.readline().strip().split(',')))
        boards = []
        for i in range(3):
            board = np.zeros((5, 5), dtype=int)
            for j in range(5):
                row = list(map(int, f.readline().strip().split()))
                for k in range(5):
                    board[j, k] = row[k]
            boards.append(board)
    
    # initialize variables
    winning_board = None
    last_number = None
    
    # play the game
    for number in numbers:
        for board in boards:
            board[board == number] = -1
            won, winning_line = check_board(board)
            if won:
                winning_board = board
                last_number = number
                break
        if winning_board is not None:
            break
    
    # calculate final score
    score = calculate_score(winning_board, last_number)
    
    # print output
    print(score)

if __name__ == '__main__':
    main()
