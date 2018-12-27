# miam
image &amp; video manager local and cloud

This is a study project to teach myself golang in depth.

# targeted feature (not in particular order)

- import automatically media from a inserted SD card
  - select drive
  - import in date based folder `year/mont/day`.
    - file will remain forever in those date base folder
- link JPEG and RAW versions of teh same image (can have several JPEG for the same RAW)
  - delete all the the JPEGs delete the RAW
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
- web interface to preset the folders contents
- web interface to select/rate/delete files
- desktop client to open/save/upload

# selection of tools and librairies
- https://github.com/spf13/afero => A FileSystem Abstraction System for Go
- https://github.com/fsnotify/fsnotify => Cross-platform file system notifications for Go. https://fsnotify.org
- 
