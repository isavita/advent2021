import itertools
from collections import defaultdict

# Step 1: Parse input
def parse_input(filename):
    with open(filename) as file:
        raw_data = file.read().strip().split('\n\n')
    scanners = []
    for data in raw_data:
        beacons = [tuple(map(int, line.split(','))) for line in data.split('\n')[1:]]
        scanners.append(beacons)
    return scanners

# Step 2: Generate distance vectors
def generate_distance_vectors(beacons):
    return {(x2 - x1, y2 - y1, z2 - z1)
            for x1, y1, z1 in beacons
            for x2, y2, z2 in beacons}

# Step 3: Find overlapping scanners
def find_overlaps(scanners):
    overlap_info = {}
    for i, scanner_a in enumerate(scanners):
        for j, scanner_b in enumerate(scanners):
            if i >= j:
                continue
            vectors_a = generate_distance_vectors(scanner_a)
            vectors_b = generate_distance_vectors(scanner_b)
            common = vectors_a & vectors_b
            if len(common) >= 66:  # Choose 2 from 12 common beacons
                overlap_info[(i, j)] = common
    return overlap_info

# Step 4: Determine orientation and position
def determine_orientation_and_position(scanner_a, scanner_b, common_vectors):
    for perm in itertools.permutations((0, 1, 2)):
        for flips in itertools.product((-1, 1), repeat=3):
            transformed_b = {tuple(coord[perm[i]] * flips[i] for i in range(3)) for coord in scanner_b}
            vectors_b = generate_distance_vectors(transformed_b)
            if vectors_b & common_vectors == common_vectors:
                for beacon_a in scanner_a:
                    for beacon_b in transformed_b:
                        translation = tuple(beacon_a[i] - beacon_b[i] for i in range(3))
                        translated_b = {tuple(beacon[i] + translation[i] for i in range(3)) for beacon in transformed_b}
                        if len(translated_b & set(scanner_a)) >= 12:
                            return translated_b, translation
    return set(), (0, 0, 0)

# Step 5: Transform beacon coordinates
def transform_coordinates(scanners, overlaps):
    aligned = {0: scanners[0]}
    translations = {0: (0, 0, 0)}
    while len(aligned) < len(scanners):
        for (i, j), common_vectors in overlaps.items():
            if i in aligned and j not in aligned:
                transformed_b, translation = determine_orientation_and_position(aligned[i], scanners[j], common_vectors)
                if transformed_b:
                    aligned[j] = transformed_b
                    translations[j] = translation
    return aligned, translations

# Step 6: Combine all beacons
def combine_beacons(aligned):
    global_beacons = set()
    for beacons in aligned.values():
        global_beacons.update(beacons)
    return global_beacons

# Main execution
if __name__ == "__main__":
    scanners = parse_input("day19_part1/input.txt")
    overlaps = find_overlaps(scanners)
    aligned_scanners, scanner_positions = transform_coordinates(scanners, overlaps)
    global_beacons = combine_beacons(aligned_scanners)
    print(f"Number of beacons: {len(global_beacons)}")
