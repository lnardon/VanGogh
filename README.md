# VanGogh

An image manipulation cli tool written in Go. Resize and convert images in the formats: jpeg, png and webp.

<p style="text-align:center">
    <img style="width: 400px" src="./logo.webp"></img>
</div>

## How to build

Make you sure you have golang installed on your machine and run the command below on the root folder to build it

```
go build
```

## How to run

There are two operations available, a file conversion or a file resize

#### File conversion

You can convert images on/to the formats jpeg, png and webp.

Run the command below to run the application

```
./VanGogh convert <file_path> <target_format>
```

#### File resize

You can also resize images in the formats like jpeg, png and webp.

```
./VanGogh resize <file_path> <scaling_technique> <scaling_factor>
```

| Parameter             | Description                                       | Type    | Range/Options                                         | Default         |
| --------------------- | ------------------------------------------------- | ------- | ----------------------------------------------------- | --------------- |
| `<file_path>`         | Path to the image or folder the tool will modify. | String  | Any valid file/folder path                            | -               |
| `<scaling_factor>`    | Determines how much to resize the images.         | Integer | 1 to 32                                               | -               |
| `<scaling_technique>` | Technique to use for image scaling.               | String  | NearestNeighbor, ApproxBiLinear, BiLinear, CatmullRom | NearestNeighbor |

Note: Scaling techniques progress from faster execution with potentially worse results (NearestNeighbor) to slower execution with potentially better results (CatmullRom).

### Examples

The command below converts the `./logo.png` file into a `logo.webp` file

```
./VanGogh convert ./logo.png .webp
```

The command below resizes all the files in the`/assets` folder into images 2 times smaller using the best quality

```
./VanGogh resize ./assets CatmullRom 2
```
