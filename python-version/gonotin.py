# coding: utf-8
import sys


def add_line_to_array(file):
    with open(file, 'r') as f:
        array = [line for line in f]
    return set(array)


data_a = add_line_to_array(sys.argv[1])
data_b = add_line_to_array(sys.argv[2])

print '> Data in A: %s | Data in B: %s' % (len(data_a), len(data_b))

for value in data_a:
    if not value in data_b:
        print value.rstrip()