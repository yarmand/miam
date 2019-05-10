# miam

[![Build Status](https://dev.azure.com/yann0602/miam/_apis/build/status/yarmand.miam?branchName=master)](https://dev.azure.com/yann0602/miam/_build/latest?definitionId=1&branchName=master)

image &amp; video manager local and cloud

This is a study project to teach myself golang in depth.

# Setup / data structures

## Env variables

**MIAM_DATA_DIR** => the root of the data. database and images aftert importation get there.

## Date folder structure
The root of the data dir is defined by the ENV variable `MIAM_DATA_DIR`

- /queues
- /queues/imports
- /queues/thumbnails.db
- /event_log.db
    - log of all operations that change the datas. Each row as a unique ID. This log can be use to sync several instances of Miam.
- /images.db
- /images
- /images/jpegs
- /images/raws

## multiple version of the same image

Each image can have:
- master raw
- master jpeg
- luminar raw
- luminar jpeg
- versions
- developed
- thumbnails

All the iteration of the same image are linked together

# targeted feature (not in particular order)

- import automatically media from a inserted SD card
  - select drive
  - import in date based folder `year/mont/day`.
    - file will remain forever in those date base folder
- link JPEG and RAW versions of the same image (can have several JPEG for the same RAW)
  - delete all the the JPEGs => delete the RAW
  - publish a developed/processed JPEG
  - published jpeg is marked to be able to find back the RAW later
- sync a designated folder to a cloud provider (OneDrive, DropBox...)
- web interface
  - dark table (quick select. delete, tag)
  - select files fast
  - move in folders (virtual folders)
  - create folders
  - manage attached drives
  - keep access to files via date forever (virtual folders)
- share folders
  - sharing remains if folder is moved
  - sharing type
    - permanent
    - temporary
    - destroy the folder after a time
- accessible from anywhere
- desktop client
  - easy open/save/sync
  - open files in external editor (photoshop, luminar...)
  - present the folder view as selected from the sync
  - local version of the web interface for offline management (import new files, edit, organize local files)
- image tagging
  - auto tagging from jpeg embedded infos

# Phase 1
- import files in date base folders
- web interface to present the folders contents
- tag images
- web interface to select/rate/delete files
- desktop client to open/save/upload

# selection of tools and librairies
- https://github.com/spf13/afero => A FileSystem Abstraction System for Go
- https://github.com/fsnotify/fsnotify => Cross-platform file system notifications for Go. https://fsnotify.org
- https://github.com/labstack/echo => High performance, minimalist Go web framework https://echo.labstack.com
- https://github.com/rwcarlsen/goexif => exif lib
- https://github.com/mattn/go-sqlite3 => sqlite3 driver to use with database/sql
- https://github.com/gobuffalo/pop => Pop makes it easy to do CRUD operations, run migrations, and build/execute queries.
- https://github.com/spf13/cobra => A Commander for modern Go CLI interactions
- https://github.com/onsi/ginkgo + gomega => BDD style testing (a bit heavy handed)
- https://github.com/google/go-cmp => Package for comparing Go values in tests
- https://github.com/stretchr/testify => A testing toolkit with common assertions and mocks that plays nicely with the standard library
- https://github.com/noelyahan/mergi => Fun and lightweight image programming tool + library (merge, crop, resize, watermark, animate, easing)
- https://github.com/esimov/caire => Content aware image resize library
