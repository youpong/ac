CC = clang
CFLAGS = -g -Wall -std=c18

TARGET = step0.1
OBJS =

.PHONY: all clean format check

all: $(TARGET)

clean:
	- rm -f a.out a.s *.o $(TARGET) test asrun

check: all test
	./test

$(TARGET): $(OBJS)

step0.1: asrun
	go run main.go |./asrun
step0: asrun
	cat 42.asm | ./asrun
