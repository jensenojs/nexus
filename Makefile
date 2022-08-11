ROOT_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
.PHONY: yacc
yacc:
	@(cd $(ROOT_DIR)/pkg/spl/parser; ./gen)

.PHONY: clean
clean:
