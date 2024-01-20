import random
import time

from chained_dictionary import ChainedDictionary
from linear_probing_dictionary import LinearProbingDictionary


def main():
    # generate random number in array
    arr = [random.randint(0, 1000) for i in range(1000000)]
    
    # record start time
    start_time = time.time()

    # store it in dictionary
    dict = {}
    for i in arr:
        dict[i] = 1
    
    # print the dictionary
    k = v = 0
    for key, value in dict.items():
        k = key
        v = value
        # print(key, ':', value)

    # record end time
    end_time = time.time()

    # print the time taken
    print('Time taken for dictionary:', end_time - start_time)

    # record start time
    start_time = time.time()

    # store it in dictionary class (separate chaining method)
    chaining_dict = ChainedDictionary()
    for i in arr:
        # check if i in chaining_dict
        chaining_dict.add(i, 1)
    
    # print the dictionary
    k = v = 0
    for i in arr:
        k = i
        v = chaining_dict.get(i)
        # print(i, ':', chaining_dict.get(i))

    # record end time
    end_time = time.time()

    # print the time taken
    print('Time taken for chaining dictionary:', end_time - start_time)
            
    # print the dictionary metadata
    # print(chaining_dict.get_metadata())
            
    # store it in dictionary class (linear probing method)

    # record start time
    start_time = time.time()

    linear_dict = LinearProbingDictionary()
    for i in arr:
        # check if i in linear_dict
        linear_dict.add(i, 1)
    
    # print data
    k = v = 0
    for i in arr:
        k = i
        v = linear_dict.get(i)
        # print(i, ':', chaining_dict.get(i))
    
    # record end time
    end_time = time.time()

    # print the time taken
    print('Time taken for linear probing dictionary:', end_time - start_time)
    print('Cache metadata:')
    print('Linear Probing Dictionary:')
    print(linear_dict.get_metadata())
    print('Chained Dictionary:')
    print(chaining_dict.get_metadata())
    

if __name__ == '__main__':
    main()