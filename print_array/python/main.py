import time


def timer(runner):
    start = time.perf_counter()
    runner()
    print()
    print(f"Function took: {(time.perf_counter() - start):0.8f} seconds")

def generate_numbers(end):
    numbers = []
    for i in range(0, end + 1):
        numbers.append(i)
    return numbers

def main():
    timer(lambda: print(generate_numbers(1_000_000)))

if __name__ == "__main__":
    main()