
SHELL=/bin/bash
GOCMD=go 
SODACMD=soda

TEST_DB=miam_test
DEV_DB=miam_development

all: test

.PHONY: gcmd
gcmd:
	echo $(GOCMD)

.PHONY: test
test: test-createDB
	$(GOCMD) test ./...

.PHONY: getSoda
getSoda:
	which $(SODACMD) || $(GOCMD) get -u -v -tags sqlite github.com/gobuffalo/pop/... ; true
	which $(SODACMD) || $(GOCMD) install -tags sqlite github.com/gobuffalo/pop/soda

$(TEST_DB).sqlite: 
	cd db ; $(SODACMD) create  $(TEST_DB)

test-migrateDB: $(TEST_DB).sqlite
	cd db ; $(SODACMD) migrate -e test migrate

$(DEV_DB).sqlite: 
	cd db ; $(SODACMD) create  $(DEV_DB)

dev-migrateDB: $(DEV_DB).sqlite
	cd db ; $(SODACMD) migrate -e development 

