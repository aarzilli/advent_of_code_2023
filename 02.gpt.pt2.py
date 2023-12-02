import re
import sys

def solve(file_name):
    with open(file_name, 'r') as f:
        lines = f.readlines()

    total_power = 0
    for line in lines:
        moves = line.split(':')[1].split(';')
        min_reds, min_blues, min_greens = 0, 0, 0
        for move in moves:
            cubes = re.findall(r'\d+ \w+', move)
            reds = sum(int(re.findall(r'\d+', c)[0]) for c in cubes if 'red' in c)
            blues = sum(int(re.findall(r'\d+', c)[0]) for c in cubes if 'blue' in c)
            greens = sum(int(re.findall(r'\d+', c)[0]) for c in cubes if 'green' in c)
            min_reds = max(min_reds, reds)
            min_blues = max(min_blues, blues)
            min_greens = max(min_greens, greens)
        total_power += min_reds * min_blues * min_greens
    return total_power


if __name__ == "__main__":
    print(solve(sys.argv[1]))

# total 5m15s -- success