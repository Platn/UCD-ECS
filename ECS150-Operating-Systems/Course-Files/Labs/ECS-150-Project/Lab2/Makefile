GXX := gcc -pedantic -Wall
GXXG := gcc -g -pedantic -Wall
GXXDEBUG := gcc -DDEBU -pedantic -Wall

SOURCE_FILES := src/*.c
MAIN := main.c
OUTPUT := procsim

all: $(MAIN) $(SOURCE_FILES)
	$(GXX) -o $(OUTPUT) $(MAIN)
	
allg: $(MAIN) $(SOURCE_FILES)
	$(GXXG) -o $(OUTPUT) $(MAIN)

%: %.c
	$(GXX) $< -o $@.out

%-debug: %.c
	$(GXXDEBUG) $< -o $@.out

clean: *.out
	rm $^
