PROGRAM:=golang

.PHONY : run
run : $(PROGRAM)
	cat samples/in_0.sql | ./$(PROGRAM)

$(PROGRAM) : main.go
	go build .

.PHONY : clean
clean :
	$(RM) $(PROGRAM)

.PHONY : check
check :
	@go test
