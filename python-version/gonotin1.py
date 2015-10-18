# coding: utf-8

#
# AUTHOR: GUSTAVO HENRIQUE
#


import sys

def add_line_to_array(file):
    with open(file, 'r') as f:
        array = [line for line in f]
    return set(array)

if __name__ == '__main__':
    data_a = add_line_to_array(sys.argv[1])
    data_b = add_line_to_array(sys.argv[2])
    for value in data_a:
        if not value in data_b:
            print value.rstrip()
