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
- image tagging
  - auto tagging from jpeg embedded infos
