# coding: utf-8

#
# AUTHOR: GUSTAVO HENRIQUE
#

import sys

def add_line_to_set(file):
    with open(file, 'r') as f:
        return {line for line in f}

if __name__ == '__main__':
    data_a = add_line_to_set(sys.argv[1])
    data_b = add_line_to_set(sys.argv[2])
    diff = set.difference(data_a, data_b)
    for line in diff:
        print line.rstrip()