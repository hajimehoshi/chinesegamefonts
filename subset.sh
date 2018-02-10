pyftsubset scregular.ttf --unicodes-file=unicodes.txt --text-file=cn.txt --layout-features='*' --output-file=/tmp/scregular.ttf
pyftsubset tcregular.ttf --unicodes-file=unicodes.txt --text-file=tw.txt --layout-features='*' --output-file=/tmp/tcregular.ttf
ttfautohint --dehint /tmp/scregular.ttf scregular.subset.ttf
ttfautohint --dehint /tmp/tcregular.ttf tcregular.subset.ttf
