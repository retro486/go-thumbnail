Just messing around with golang and ImageMagick via https://github.com/gographics/imagick

API is all these C methods minus the leading "Magick" and the first \*wand argument: http://www.imagemagick.org/api/magick-image.php

Ideally would be a server listening on a port, simply taking in binary data then spitting out raw binary data for bulk image processing to replace a lot of the same thing in Ruby to save on RAM.

Of course, it'll also need all the insanity in PdfKit as well... and wkhtmltoimage...
