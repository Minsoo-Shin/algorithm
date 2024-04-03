def numJewelsInStones(jewels, stones) -> int:
    freqs = {}

    for jewel in jewels:
        freqs[jewel]=True

    count = 0 
    for stone in stones:
        if stone in freqs:
            count += 1
        
    print(count)
    

numJewelsInStones("aA", "aAAbbbb")    

import collections
def numJewelsInStonesDefaultDict(jewels, stones) -> int:
    freqs = collections.defaultdict(int)
    count = 0

    for stone in stones:
        freqs[stone] += 1
    
    for jewel in jewels:
        if jewel in freqs:
            count += freqs[jewel]
        
    print(count)
    
# 정답 7
numJewelsInStones("aAb", "aAAbbbb")    