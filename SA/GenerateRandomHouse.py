import random
import time

def generate_random_houses(num_houses):
    houses = []
    for i in range(num_houses):
        houses.append(random.randint(1, 100))
    return houses

houses = generate_random_houses(1000)
print("Original houses:", houses)
