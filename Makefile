CC = clang
CFLAGS = -g -Wall -std=c18

TARGET = step1
OBJS =

.PHONY: all clean format check

all: $(TARGET)

clean:
	- rm -f a.out a.s *.o $(TARGET) test asrun

check: all test
	./test

$(TARGET): $(OBJS)

step1: asrun
	go run main.go |./asrun
