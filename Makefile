CC = clang
CFLAGS = -g -Wall -std=c18

TARGET = step2
OBJS =

.PHONY: all clean check run

all: $(TARGET)

clean:
	- rm -f a.out a.s *.o $(TARGET) test asrun

check: all test
	./test

$(TARGET): $(OBJS)

run:
	go run main.go
step2: asrun
	go run main.go |./asrun
