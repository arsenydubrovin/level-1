.PHONY: run

# Usage: make run TASK=<N>
run:
	go run "task_$(TASK)/main.go"
