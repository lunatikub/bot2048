
MODULE = $(shell env GO111MODULE=on $(GO) list -m)
BIN = $(CURDIR)/bot2048
GO = go

# generated files
PATHGEN = bot/path.go
TRANSGEN = bot/transformation.go

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: $(PATHGEN) $(TRANSGEN)
	$(info $M build $(BIN))
	$Q $(GO) build -o $(BIN) main.go

$(PATHGEN): $(CURDIR)/path.py
	$(info $M generate $@)
	$Q $^ > $@

$(TRANSGEN): $(CURDIR)/transformation.py
	$(info $M generate $@)
	$Q $^ > $@

test:
	$(info $M test)
	$Q go test -v ./...

.PHONY: clean
clean:
	$(info $(M) cleaning…)
	@rm -rf $(BIN)
	@rm -rf $(PATHGEN) $(TRANSGEN)
