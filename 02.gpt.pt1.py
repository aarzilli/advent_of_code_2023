import re
import sys

def solve(file_name):
    with open(file_name, 'r') as f:
        lines = f.readlines()

    valid_games = []
    for line in lines:
        game_id = int(re.findall(r'\d+', line.split(':')[0])[0])
        moves = line.split(':')[1].split(';')
        game_valid = True
        for move in moves:
            cubes = re.findall(r'\d+ \w+', move)
            reds = sum(int(re.findall(r'\d+', c)[0]) for c in cubes if 'red' in c)
            blues = sum(int(re.findall(r'\d+', c)[0]) for c in cubes if 'blue' in c)
            greens = sum(int(re.findall(r'\d+', c)[0]) for c in cubes if 'green' in c)
            if reds > 12 or blues > 14 or greens > 13:
                game_valid = False
                break
        if game_valid:
            valid_games.append(game_id)
    return sum(valid_games)


if __name__ == "__main__":
    print(solve(sys.argv[1]))
