class LinearProbingDictionary:

    def __init__(self):
        self.arr = [None for _ in range(2)]
        self.load_factor = 0
        self.used_keys = 0
        self.total_keys = 2
        self.resizing_count = 0
        self.resizing_sum = 0

    
    def __calculate_load_factor(self):
        self.load_factor = self.used_keys / self.total_keys
    
    def __should_resize(self):
        if self.load_factor >= 0.5:
            return True
        else:
            return False
    
    def __resize_array(self):
        self.total_keys = self.total_keys * 2
        self.__calculate_load_factor()
        new_arr = [None]*self.total_keys
        for i in range(len(self.arr)):
            if self.arr[i] is not None:
                new_arr[self.__get_hash_key(self.arr[i][0])] = self.arr[i]
                self.resizing_sum += 1
        self.arr = new_arr
        self.resizing_count += 1
    
    def __get_hash_key(self, key):
        # get hash key
        hash_key = hash(key) % self.total_keys
        return hash_key
    
    def __linear_probe(self, index):
        # Linear probing: move to the next slot
        return (index + 1) % self.total_keys

    def add(self, key, value):
        # print(self.arr, key, value, self.used_keys, self.total_keys)
        # print('')
        if self.__should_resize():
            self.__resize_array()
        
        hash_key = self.__get_hash_key(key)
        # print(hash_key)
        # check for collision, if yes, then linear probe
        while self.arr[hash_key] is not None:
            if self.arr[hash_key][0] == key:
                self.arr[hash_key][1] = value
                return
            hash_key = self.__linear_probe(hash_key)
        # add the key-value pair
        self.arr[hash_key] = [key, value]
        self.used_keys += 1
        self.__calculate_load_factor()
        # print(self.load_factor)
    
    def get(self, key):
        hash_key = self.__get_hash_key(key)
        # check for collision, if yes, then linear probe
        while self.arr[hash_key] is not None:
            # print(self.arr[hash_key], hash_key)
            if self.arr[hash_key][0] == key:
                return self.arr[hash_key][1]
            hash_key = (hash_key + 1) % self.total_keys
        return None

    def get_metadata(self):
        return {
            'total_keys': self.total_keys,
            'used_keys': self.used_keys,
            'load_factor': self.load_factor,
            'resizing_count': self.resizing_count,
            'resizing_sum': self.resizing_sum
        }