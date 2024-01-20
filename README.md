# Hash-Map
Implementation of hash-map in Go/Python

When implementing a hash table, collisions can occur when two or more keys hash to the same index. There are several methods to handle collisions, each with its own advantages and disadvantages. Here are some common collision resolution methods:

1. Chaining:
    - Description: Chaining involves maintaining a linked list at each index in the hash table. When a collision occurs, the new key-value pair is added to the linked list at the corresponding index.
    - Advantages: Simple to implement, effective for handling a large number of collisions.
    - Disadvantages: Additional memory is required for the linked lists, and performance may degrade if the lists become too long.
2. Open Addressing:
    - Description: Open addressing involves placing the collided item in the next available slot within the hash table, rather than using a separate data structure like a linked list.
    - Types of Open Addressing:
      - Linear Probing: If a collision occurs at index i, the algorithm searches for the next available slot (i+1, i+2, and so on) until an empty slot is found.
      - Quadratic Probing: Similar to linear probing, but the interval between slots is increased quadratically (i+1^2, i-1^2, and so on).
      - Double Hashing: Uses a secondary hash function to determine the interval between slots.
    - Advantages: No additional data structures are needed, and it can be more memory-efficient than chaining.
    - Disadvantages: Tends to cause clustering, where consecutive slots are occupied, potentially leading to more collisions.
  

**NOTE: We are comparing only chaining and open addressing (LP) here**

## Benchmark score:

### Low Key Count

![Low Key Count](https://github.com/kumar-kunal/hash-map/blob/main/go/time_low_key_count.png)


### High Key Count
[![High Key Count](https://github.com/kumar-kunal/hash-map/blob/main/go/time_high_key_count.png)]

