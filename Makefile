TARGET = ac

.PHONY: all clean check 

all: $(TARGET)

clean:
	- rm -f a.out a.s $(TARGET) test asrun

check: all test
	./test

$(TARGET): main.go
	go build -o $@ main.go
