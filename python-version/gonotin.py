# coding: utf-8
import sys
import os


def add_line_to_array(file):
    with open(file, 'r') as f:
        array = [line for line in f]
    return set(array)


def mode1():
    data_a = add_line_to_array(sys.argv[1])
    data_b = add_line_to_array(sys.argv[2])
    print '> Data in A: %s | Data in B: %s' % (len(data_a), len(data_b))
    for value in data_a:
        if not value in data_b:
            print value.rstrip()


def add_line_to_set(file):
    with open(file, 'r') as f:
        return {line for line in f}


def mode2():
    data_b = add_line_to_set(sys.argv[2])
    i = 0
    with open(sys.argv[1], 'r') as f:
        for value in f:
            if not value in data_b:
                i += 1
                sys.stdout.write(value)
    sys.stdout.flush()
    print '> Data in A: %s | Data in B: %s' % (i, len(data_b))


if __name__ == '__main__':
    '''
    Usage: MODE=1 time python gonotin.py file1.txt file2.txt
    '''

    modes = {
        '1': mode1,
        '2': mode2
    }
    modes.get(os.getenv('MODE', '1'))()