# What ?

I do not know what kind of data structure to use.
Lets create an example set of data and test performances / easyness of several structures.

# data example set
- 30,000 images
- each image has 3 date tags
  - year => 10
  - month => 12
  - day => 31
- process tag:
  - 10% new
  - 3% to process
  - 2% processing
  - 10% to publish
  - 75% published
- publication target
  - each published image, 80% of "to publish" as 1-6 publication targets
- quality
  - randomly set on 60% of the images
- links
  - 90% of images have jpeg and raw
  - 20% of images has 2-5 versions
- path
  - all images are spread across a set of path
  - 1-10 level deep
  - each level have 1-20 values
- exif info for each image
  - each image get value from selection of
    - camera => 5
    - lens => 10
    - iso => 12
    - shutter speed => 20
    - aperture => 20
- random tags
  - 50% of image has 1-50 tags from as selection from 500 tags
