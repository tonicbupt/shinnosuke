# Shinnosuke

野原しんのすけ, 野原新之助

![shinnosuke](https://raw.githubusercontent.com/tonicbupt/shinnosuke/master/images/shinnosuke.webp)

# Usage:

```
$ ./bin/shinnosuke -h
NAME:
   shinnosuke - 野原しんのすけ, helps to compress your images in JPEG / PNG

USAGE:
   shinnosuke [global options] command [command options] TARGET_SIZE (in kB/MB/GB)

VERSION:
   v1.1.7

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

$ cd /path/to/your/directory

$ ./bin/shinnosuke 1MB
```

就会把图片都压缩到最大 1MB, 不过 PNG 格式可能到不了, PNG 没有提供有损压缩算法...  
压缩效果好差=。=
