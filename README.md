<div style="display:flex; align-items:center; justify-content: center">
    <h1 style="font-size:3.5rem; font-weight: 700">Van Gogh</h1>
</div>

<div style="display:flex; align-items:center; justify-content: center">
<p style="font-size: 1.25rem; font-weight: 400">
    An image manipulation package written in Go.
</p>
</div>

<div style="display:flex; align-items:center; justify-content: center">
    <img style="width: 400px" src="./logo.webp"></img>
</div>

## How to build

Make you sure you have golang installed on your machine and run the command below on the root folder to build it

```
go build
```

## How to run

Run the command below to run the application

```
./VanGogh <file_path> <scaling_technique> <scaling_factor>
```

| Parameter             | Description                                       | Type    | Range/Options                                         | Default         |
| --------------------- | ------------------------------------------------- | ------- | ----------------------------------------------------- | --------------- |
| `<file_path>`         | Path to the image or folder the tool will modify. | String  | Any valid file/folder path                            | -               |
| `<scaling_factor>`    | Determines how much to resize the images.         | Integer | 1 to 32                                               | -               |
| `<scaling_technique>` | Technique to use for image scaling.               | String  | NearestNeighbor, ApproxBiLinear, BiLinear, CatmullRom | NearestNeighbor |

Note: Scaling techniques progress from faster execution with potentially worse results (NearestNeighbor) to slower execution with potentially better results (CatmullRom).
