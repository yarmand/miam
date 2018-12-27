# miam
image &amp; video manager local and cloud

This is a study project to teach myself golang in depth.

# targeted feature (not in particular order)

- import automatically media from a inserted SD card
- link JPEG and RAW versions of teh same image (can have several JPEG for the same RAW)
  - delete all the the JPEGs delete the RAW
  - publish a developed/processed JPEG
  - published jpeg is marked to be able to find back the RAW later
- sync a designated folder to a cloud provider (OneDrive, DropBox...)
- web interface
  - dark table (quick select. delete, tag)
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
- image tagging
- auto tagging from jpeg embedded infos
