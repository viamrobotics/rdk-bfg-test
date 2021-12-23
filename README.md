# RDK (Robot Development Kit)

<p align="center">
  <a href="https://go.viam.com/pkg/go.viam.com/rdk/"><img src="https://pkg.go.dev/badge/go.viam.com/rdk" alt="PkgGoDev"></a>
  <a href="https://codecov.io/gh/viamrobotics/rdk"><img src="https://codecov.io/gh/viamrobotics/rdk/branch/master/graph/badge.svg?token=99YH0M8YOA" alt="CodeCov"></a>
</p>

* [Programs](#programs)
* [Dependencies](#dependencies)
* [Development](#development)

### API Documentation & more devices
To see more examples, check out the [Wiki](https://github.com/viamrobotics/rdk/wiki)

## Dependencies

* make
* Run `make setup`

### First time run

* Try out `go run web/cmd/server/main.go robots/configs/fake.json` and visit http://localhost:8080

### libvpx linux source build
If libvpx is not available on your distro, run the following:

1. `git clone git@github.com:webmproject/libvpx.git`
1. `cd libvpx`
1. `mkdir build; cd build`
1. `../configure --enable-runtime-cpu-detect --enable-vp8 --enable-postproc --enable-multi-res-encoding --enable-webm-io --enable-better-hw-compatibility --enable-onthefly-bitpacking --enable-pic`
1. `sudo make install`

## Development

### Conventions
1. Write tests!
2. Follow this [Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
3. Always run `make setup build lint` and test before pushing.
4. Usually merge and squash your PRs and more rarely do merge commits with each commit being a logical unit of work.
5. If you add a new package, please add it to this README.
6. If you add a new sample or command, please add it to this README.
7. Experiments should go in samples or any subdirectory with /samples/ in it. As "good" pieces get abstracted, put into a real package command directory.
8. Use imperative mood for commits (see [Git Documenation](https://git.kernel.org/pub/scm/git/git.git/tree/Documentation/SubmittingPatches?id=a5828ae6b52137b913b978e16cd2334482eb4c1f#n136)).
9. Try to avoid large merges unless you're really doing a big merge. Try to rebase (e.g. `git pull --rebase`).
10. Delete any non-release branches ASAP when done, or use a personal fork
11. Prefer metric SI prefixes where possible (e.g. millis) https://www.nist.gov/pml/weights-and-measures/metric-si-prefixes. The type of measurement (e.g. meters) is not necessary if it is implied (e.g. rulerLengthMillis).

### Resources

All resources implemented within the RDK follow the pattern of registering themselves within an `func init()` block. This requires the package they are implemented in be imported, but typically not explicitly used. The place where we currently put blank imports (`_ "pkgpath"`) is in [robot/impl/local_robot.go](./robot/impl/local_robot.go). They should be placed in the corresponding component's register package going forward.

### Protocol Buffers/gRPC

For API intercommunication, we use Protocol Buffers to serialize data and gRPC to communicate it. For more information on both technologies, see https://developers.google.com/protocol-buffers and https://grpc.io/.

Some guidelines on using these:
1. Follow the [Protobuf style guide](https://docs.buf.build/style-guide/).
1. After making changes to a `.proto` file, make sure to run `make buf` to generate new files. Make sure `protoc-gen-go-grpc` and `protoc-gen-go`, usually located in `~/go/bin`, are in your `$PATH`.

#### gRPC Language Samples

* [Go](./grpc) - See `client` and `server`.
* [Python](./grpc/python)
* [Java](./grpc/java)
* [C++](./grpc/cpp)

### Testing with big data

Let's assume big data is > 10KiB. This kind of data is annoying to slow to pull down with git and is typically not needed except for certain tests. In order to add large data test artifacts, you need to do the following:

```
# get ARTIFACT_GOOGLE_APPLICATION_CREDENTIALS by talking to Eliot or Eric
go install go.viam.com/utils/artifact/cmd/artifact
# place new artifacts in artifact_data
artifact push
git add .artifact
# commit the file at some point
```

### Testing from Github Actions

1. First make sure you have docker installed (https://docs.docker.com/get-docker/)
1. Install `act` with `brew install act`
1. Add `GIT_ACCESS_TOKEN` which is your GitHub Personal Access Token (repo scope) it to your .secrets file in the repo (see https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token and https://github.com/nektos/act#configuration)
1. Then just run `act`
