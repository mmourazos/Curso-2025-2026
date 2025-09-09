import time
import threading


def printNumbers(name):
    for i in range(5):
        print(f"Thread {name}: {i}")
        time.sleep(1)


if __name__ == "__main__":
    # Creamos los hilos usando la función printNumbers.
    t1 = threading.Thread(target=printNumbers, args=("A",))
    t2 = threading.Thread(target=printNumbers, args=("B",))

    # Iniciamos los hilos.
    t1.start()
    t2.start()

    # Esperamos a que los hilos terminen antes de continuar.
    t1.join()
    t2.join()

    print("The end!")
