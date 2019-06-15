# csvutil

Command line utility for transforming CSV files.


## Install

Make sure that [Go](https://golang.org/) 1.9+ is installed and your PATH includes GOBIN. Then run the following:

```
$ go get github.com/evantbyrne/csvutil
```

**Note:** The $ at the beginning of newlines in this document represents the bash shell prompt, and is not a part of the actual commands.


## Usage

```
$ csvutil source1.csv --operation1 --operation2 source2.csv --operation1 ...
```

There are two types of arguments that may be passed to csvutil:

1. Sources. These are CSV files to read from.
2. Operations denoted with the `--` prefix. These transform the current source.

Upon success, csvutil prints the resulting CSV data to STDOUT.


## Examples

Here are a few examples of working with csvutil.

**Reading files**

```
$ csvutil user.csv
username,id,group_id
foo,1,10
bar,2,
baz,3,20
four,4,20
five,5,
$ csvutil group.csv
name,id
Admin,10
Moderator,20
```

**Filtering**

```
$ csvutil user.csv --where username == baz --or group_id != 20
username,id,group_id
foo,1,10
bar,2,
baz,3,20
five,5,
```

**Join**

```
$ csvutil user.csv group.csv --join group_id == id
username,id,group_id,name,id
foo,1,10,Admin,10
baz,3,20,Moderator,20
four,4,20,Moderator,20
```

**Sort and select**

```
$ csvutil test/user.csv --sort username ALPHA DESC --select id,username
id,username
4,four
1,foo
5,five
3,baz
2,bar
```


## Operations

### Concat

Appends one source onto another.

No arguments.

```
$ csvutil foo.csv
id,name
1,Foo
2,Bar
$ csvutil bar.csv
id,name
3,Baz
$ csvutil foo.csv bar.csv --concat
id,name
1,Foo
2,Bar
3,Baz
```


### Count

Returns the total number of rows for the current source.

Arguments:
- Column name.

```
$ csvutil user.csv --count total
total
5
```


### Distinct

Removes duplicates from the current source.

Arguments:
- Comma-separated list of columns that must match to be considered duplicate. Use `"*"` to compare all columns.

```
$ csvutil foo.csv
id,name,group_id
1,Foo,10
2,Bar,10
3,Foo,20
4,Foo,10
5,Foobar,20
1,Baz,20
$ csvutil foo.csv --distinct name
id,name,group_id
1,Foo,10
2,Bar,10
5,Foobar,20
1,Baz,20
$ csvutil foo.csv --distinct name,group_id
id,name,group_id
1,Foo,10
2,Bar,10
3,Foo,20
5,Foobar,20
1,Baz,20
```


### Except

Remove rows from the previous source if they appear in the current source.

Arguments:
- Comma-separated list of columns that must match to be considered duplicate. Use `"*"` to compare all columns.

```
$ csvutil foo.csv
id,name
1,Foo
2,Bar
3,Baz
4,Foobar
$ csvutil bar.csv
id,name
2,Two
3,Three
5,Five
6,Six
$ csvutil foo.csv bar.csv --except id
id,name
1,Foo
4,Foobar
```


### Join

Perform an inner join on two sources.

Arguments:
- Column from previous source to compare.
- Operation. Choices: `==`, `!=`
- Column from the current source to compare.

```
$ csvutil user.csv
username,id,group_id
foo,1,10
bar,2,
baz,3,20
four,4,20
five,5,
$ csvutil group.csv
name,id
Admin,10
Moderator,20
$ csvutil user.csv group.csv --join group_id == id
username,id,group_id,name,id
foo,1,10,Admin,10
baz,3,20,Moderator,20
four,4,20,Moderator,20
```


### Select

Choose a subset of columns from the current source.

Arguments:
- Comma-separated list of columns to select.

```
$ csvutil foo.csv
id,name,group_id
1,Foo,10
2,Bar,10
3,Baz,20
$ csvutil foo.csv --select name,id
name,id
Foo,1
Bar,2
Baz,3
```


### Sort

Reorder the current source.

Arguments:
- Column to compare.
- Algorithm. Choices: `ALPHA`, `FLOAT`, `INT`
- Order. Choices: `ASC`, `DESC`

```
$ csvutil user.csv
username,id,group_id
foo,1,10
bar,2,
baz,3,20
four,4,20
five,5,
$ csvutil user.csv --sort username ALPHA ASC
username,id,group_id
bar,2,
baz,3,20
five,5,
foo,1,10
four,4,20
```

When using `FLOAT` and `INT`, values that cannot be converted will be considered lower than values which may be converted, and they will be compared to each other alphabetically. e.g.,

```
$ csvutil baz.csv
year,cost
10,500
2,200.10
2000,50.01
x,x
2019,1000.99
y,z
z,y
$ csvutil baz.csv --sort year INT DESC
year,cost
2019,1000.99
2000,50.01
10,500
2,200.10
z,y
y,z
x,x
```


### Values

Strip the header from the current source.

```
$ csvutil user.csv --count total
total
5
$ csvutil user.csv --count total --values
5
```


### Where

Filter the current source.

Arguments:
- Column name to compare.
- Operator. Choices: `==`, `!=`, `IN`, `NOT_IN`
- Value to look for.

```
$ csvutil user.csv
username,id,group_id
foo,1,10
bar,2,
baz,3,20
four,4,20
five,5,
$ csvutil user.csv --where username == baz --or group_id != 20
username,id,group_id
foo,1,10
bar,2,
baz,3,20
five,5,
```

The `IN` and `NOT_IN` operators take column-separated lists of values. e.g.,

```
$ csvutil user.csv --where id IN 1,3
username,id,group_id
foo,1,10
baz,3,20
```

Operations may be chained to perform "and" conditionals. e.g.,

```
$ csvutil user.csv --where username != baz --where id != 2
username,id,group_id
foo,1,10
four,4,20
five,5,
```
