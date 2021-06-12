
SHELL=/bin/bash
GOCMD=go 
SODACMD=soda

TEST_DB=miam_test
DEV_DB=miam_development

all: test build

build:
	go build -o miam .

.PHONY: gcmd
gcmd:
	echo $(GOCMD)

.PHONY: test
test: 
	$(GOCMD) test ./...

.PHONY: getSoda
get-soda:
	which $(SODACMD) || $(GOCMD) get -u -v -tags sqlite github.com/gobuffalo/pop/... ; true
	which $(SODACMD) || $(GOCMD) install -tags sqlite github.com/gobuffalo/pop/soda

$(TEST_DB).sqlite: 
	$(SODACMD) create  $(TEST_DB)

test-migrateDB: $(TEST_DB).sqlite
	$(SODACMD) migrate -e test

$(DEV_DB).sqlite: 
	$(SODACMD) create  $(DEV_DB)

dev-migrateDB: $(DEV_DB).sqlite
	$(SODACMD) migrate -e development 

