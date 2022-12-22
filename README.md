# ascii-art-reverse

### Description
Ascii-reverse consists on reversing the process, converting the graphic representation into a text. Program will create a text file containing a graphic representation of a random string given as an argument.

The argument will be a flag, `--reverse=<fileName>`, in which `--reverse` is the flag and `<fileName>` is the file name. The program will then print this string in normal text.

    
If there are other ascii-art optional projects implemented, the program should accept other correctly formatted [OPTION] and/or [BANNER].
Additionally, the program must still be able to run with a single [STRING] argument.
The flag must have exactly the same format as above, any other formats must return the following usage message:
```console
Usage: go run . [OPTION]

EX: go run . --reverse=<fileName>
```

### Instructions

There are other ascii-art optional projects implemented: 
- If you type go run . `--align=<type> [STRING] [BANNER]` change the alignment to the `left, right, center and justify`. 
##### Example: `$ go run . --align=center "hello" shadow` or `$ go run . --align=right "hello" thinkertoy`
- You can use `--output=<filename.txt>` flag. The file will be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name which will contain the output. 
  ##### Example: `$ go run . --output=sample.txt "Hello World" thinkertoy` or `$ go run . --output=sample.txt "Communicate"`
- You can also use `--color=<color>` flag. To specify which letters to color you can specify the by `--color=<color> <letters to be colored> <string>`. Colors can be chosen from here: white, black, red, green, blue, cyan, magenta,yellow, purple, and orange.
  ##### Example: `$ go run . --color=cyan "He" "Hello"` or `$ go run . --color=magenta "Think"`
- Additionally, the program is be able to run with a single [STRING] argument and `$ go run . [STRING] [BANNER]` format.
  ##### Example: `$ go run . "hello" shadow` or `$ go run . "hello"`

### Usage
```console
$ cat file.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run . --reverse=file.txt
hello
$
```

### Authors
made by ashagiro and TQ



