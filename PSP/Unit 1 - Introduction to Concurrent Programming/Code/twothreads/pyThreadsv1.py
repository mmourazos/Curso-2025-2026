import time
import threading


def printNumbers(name):
    for i in range(5):
        print(f"Thread {name}: {i}")
        time.sleep(1)


if __name__ == "__main__":
    t1 = threading.Thread(target=printNumbers, args=("A",))
    t2 = threading.Thread(target=printNumbers, args=("B",))

    t1.start()
    t2.start()

    print("The end!")
