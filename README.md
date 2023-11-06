# msds301-wk7

### Setup (MAC i9, zshrc)
- **NOTE**: Executable files are included but I built them on an i9 Macbook
- Clone repo with `git clone --recursive git@github.com:jeremycruzz/msds301-wk7.git`
- Install GoLearn System Dependencies [instructions here](https://github.com/gonum/blas#installation)
    - Make sure you have `g++` and `MYSS2`[instructions here](https://code.visualstudio.com/docs/cpp/config-mingw)
    - Install OpenBLAS in any directory
        - `git clone https://github.com/xianyi/OpenBLAS`
        - `cd OpenBLAS`
        - `make`
        - alternatively `brew install openblas`
    - Link to library
        - `CGO_LDFLAGS="-L/path/to/OpenBLAS -lopenblas" go install github.com/gonum/blas/cgo`

### Building executable
- Run `go build ./cmd/isotree`


### Running Go executable
- Run `./isotree`

### Results

I decided to go with the GoLearn Package as I had experience with the package from the previous weeks assignment. I thought it was pretty straightforward when doing linear regression but it was not quite as straight forward since I did not load the data from a csv. Creating the dataframe from scratch proved to be a bit difficult. There was a weird way to add attributes and I was not able to add a class attribute as a categorical attribute but was forced to create a float attribute for the label. Unit tests were not done but I would regularly print out the data and rebuild the project as I was working on it. This is wear scripting languages have an advantage over go, as you can run a few lines and explore the data without needing to rebuild a project. After running the tests we can see that the results that Go produced were similar to those that python produced.

For the firm, I would reccomend that we try to explore a different go package that is easier to work with. The Go program was much faster than the R and python scripts and the speed would be an advantage in our data processing pipelines.


| Language | Average   | Median    | Min       | Max       |
|----------|-----------|-----------|-----------|-----------|
| R        | 0.265768  | 0.265265  | 0.187304  | 0.388845  |
| Python   | 0.443680  | 0.438875  | 0.371611  | 0.631608  |
| Go       | 0.445456  | 0.444328  | 0.304279  | 0.661401  |
