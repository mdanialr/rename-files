# Rename Files
Rename a bunch of files in a directory using regex. I use this mainly to organize and rename episodes of anime, and that would be so much troublesome if I need to rename all those files one by one, therefor I create this little project.

# How to Use
* Download the binary file from Release and extract it.
* Run the binary file using necessary arguments.
* By default, all files would be printed out from `old` name and `new` name
* When you are satisfied with the result then you can __add__ `-r` argument in the command to do renaming files.

So basically you would be presented with the preview which shows `old` and `new` name of the files, then renamed it using `-r` argument.

## Arguments
`-h` *__boolean__*: show all available arguments.

`-d` *__string__*: the directory where the target files would be renamed.

`-p` *__string__*: the prefix string to add on the renamed files.

`-ext` *__string__*: filter by extension of the files. only rename files that contain defined extension.

`-no-end` *__boolean__*: include this argument to omit 'END' from the last file.

`-re` *__string__*: override the default regex pattern.

`-r` *__boolean__*: include this if you are ready to rename the target files.
