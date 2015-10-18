# coding: utf-8

#
# AUTHOR: MANLIO (GONUTS)
#

import sys

if __name__ == '__main__':
    data_a = set(open(sys.argv[1], 'r'))
    data_b = set(open(sys.argv[2], 'r'))

    for line in data_a - data_b:
        sys.stdout.write(line)