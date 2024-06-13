import random
import time

def generate_random_houses(num_houses):
    houses = []
    n=100                                    # ini adalah range number dari 1 hingga ke-n
    for i in range(num_houses):
        houses.append(random.randint(1, n))
    return houses

a=1000
houses = generate_random_houses(a)            # a artinya mengenerate rumah sebanyak 1000 dengan range number dari 1 hinnga n, a bisa disesuaikan dengan jumlah rumah yang diinginkan
print("Original houses:", houses)
