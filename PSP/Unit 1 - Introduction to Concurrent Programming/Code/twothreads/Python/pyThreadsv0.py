import time


def printNumber(name):
    for i in range(5):
        print(f"Function {name}: {i}")
        time.sleep(1)


if __name__ == "__main__":
    printNumber("A")
    printNumber("B")
