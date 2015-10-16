# coding: utf-8
import sys
import os


def add_line_to_array(file):
    with open(file, 'r') as f:
        array = [line for line in f]
    return set(array)

# mode1 is from Gustavo
def mode1():
    data_a = add_line_to_array(sys.argv[1])
    data_b = add_line_to_array(sys.argv[2])
    for value in data_a:
        if not value in data_b:
            print value.rstrip()


def add_line_to_set(file):
    with open(file, 'r') as f:
        return {line for line in f}

# mode2 is from Justin - go-nuts group
def mode2():
    # it is generating wrong results
	print "it is generating wrong results"

	'''
    data_b = add_line_to_set(sys.argv[2])
    i = 0
    with open(sys.argv[1], 'r') as f:
        for value in f:
            if not value in data_b:
                i += 1
                sys.stdout.write(value)
    sys.stdout.flush()
    '''

# mode3 is from Manlio - go-nuts group
def mode3():
    data_a = set(open(sys.argv[1], 'r'))
    data_b = set(open(sys.argv[2], 'r'))

    for line in data_a - data_b:
        sys.stdout.write(line)

# mode4 is from Gustavo
def mode4():
    data_a = add_line_to_set(sys.argv[1])
    data_b = add_line_to_set(sys.argv[2])
    diff = set.difference(data_a, data_b)
    for line in diff:
        print line.rstrip()

if __name__ == '__main__':
    '''
    Usage: MODE=1 time python gonotin.py file1.txt file2.txt
    '''

    modes = {
        '1': mode1,
        '2': mode2,
        '3': mode3,
        '4': mode4
    }
    modes.get(os.getenv('MODE', '1'))()