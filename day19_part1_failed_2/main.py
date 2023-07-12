import numpy as np
from itertools import combinations

def parse_scanner_data(filename):
    with open(filename, 'r') as file:
        lines = file.read().splitlines()
    scanner_data = {}
    current_scanner = None
    for line in lines:
        if line.startswith('--- scanner'):
            current_scanner = int(line.split()[2])
            scanner_data[current_scanner] = []
        elif line != '':
            scanner_data[current_scanner].append(list(map(int, line.split(','))))
    return scanner_data

def find_overlapping_beacons(scanner_data, tolerance=1):
    overlapping_beacons = {}
    for scanner_id1, scanner_id2 in combinations(scanner_data.keys(), 2):
        overlap = [beacon1 for beacon1 in scanner_data[scanner_id1] if any(np.allclose(beacon1, beacon2, atol=tolerance) for beacon2 in scanner_data[scanner_id2])]
        if overlap:
            overlapping_beacons[(scanner_id1, scanner_id2)] = overlap
    return overlapping_beacons

def align_scanners(overlapping_beacons, scanner_data):
    alignments = {0: scanner_data[0]}
    for (scanner_id1, scanner_id2), overlap in overlapping_beacons.items():
        if scanner_id1 in alignments:
            reference_id = scanner_id1
            target_id = scanner_id2
        else:
            reference_id = scanner_id2
            target_id = scanner_id1
        P = np.asarray([beacon for beacon in alignments[reference_id] if beacon in overlap], dtype=float)
        Q = np.asarray([beacon for beacon in scanner_data[target_id] if beacon in overlap], dtype=float)
        centroid_P = np.mean(P, axis=0)
        centroid_Q = np.mean(Q, axis=0)
        P -= centroid_P
        Q -= centroid_Q
        C = np.dot(np.transpose(P), Q)
        V, _, Wt = np.linalg.svd(C)
        R = np.dot(V, Wt)
        if np.linalg.det(R) < 0:
            V[:, -1] *= -1
            R = np.dot(V, Wt)
        alignments[target_id] = [np.dot(beacon - centroid_Q, R) + centroid_P for beacon in scanner_data[target_id]]
    return alignments

def create_full_map(scanner_data, alignments):
    full_map = []
    for scanner_id, beacons in scanner_data.items():
        if scanner_id in alignments:
            full_map += alignments[scanner_id]
        else:
            full_map += beacons
    return full_map

def count_unique_beacons(full_map, tolerance=0.4):
    unique_beacons = []
    for beacon in full_map:
        if not any(np.allclose(beacon, unique_beacon, atol=tolerance) for unique_beacon in unique_beacons):
            unique_beacons.append(beacon)
    return len(unique_beacons)

scanner_data = parse_scanner_data('day19_part1/test1.txt')
overlapping_beacons = find_overlapping_beacons(scanner_data)
alignments = align_scanners(overlapping_beacons, scanner_data)
full_map = create_full_map(scanner_data, alignments)
num_unique_beacons = count_unique_beacons(full_map)
print(num_unique_beacons)
