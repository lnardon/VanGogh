# Van Gogh

An image manipulation package written in Go.

## How to build

Make you sure you have golang installed on your machine and run the command below on the root folder to build it

```
go build
```

## How to run

Run the command below to run the application

```
./VanGogh <file_path or folder_path> <scaling_technique> <scaling_factor>
```

#### <file_path>

A string of the path to the image or folder the tool will modify
<br>
</br>

#### <scaling_factor>

Int from 1 to 32 and determines how much to resize the images (E.g. 10 will resize the image to one 10 times smaller)
<br>
</br>

#### <scaling_technique>

<p>String from the options (NearestNeighbor, ApproxBiLinear, BiLinear, CatmullRom)</p>
<p>The scaling techniques (from left to right =>) go from faster/worst_result to slower/best_result</p>
<p>Default is Nearest Neighbor</p>
