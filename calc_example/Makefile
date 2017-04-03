# Makefile

OBJS := bison.o flex.o main.o

PROGRAM := calc
OUTPUTS := *.o flex.c bison.c lex.yy.c calc.c tok.h calc.output y.tab.c

LEX := flex
LFLAGS :=
YACC := bison -y
YFLAGS := -d

CC := g++
CFLAGS := -g -Wall -ansi -pedantic

$(PROGRAM): $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@ -lfl

bison.o main.o: heading.h
calc.o main.o: tok.h

flex.o: flex.c

flex.c: calc.l tok.h
	$(LEX) $<
	mv lex.yy.c $@

y.tab.c y.tab.h: calc.y
	$(YACC) $(YFLAGS) calc.y

bison.c: y.tab.c
	mv $< $@

tok.h: y.tab.h
	mv $< $@

clean:
	$(RM) $(OUTPUTS) $(PROGRAM)

.PHONY: clean
