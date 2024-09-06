#  google-hello-app-logging-multiarch

This project is directly based on the code from the [GoogleCloudPlatform/kubernetes-engine-samples/quickstarts/hello-app](https://github.com/GoogleCloudPlatform/kubernetes-engine-samples/tree/main/quickstarts/hello-app)

The enhancement in this project is building this GoLang app with multiple architectural targets, so that it can also run on ARM64, specifically Apple Silicon as well as AMD64.

Also, it logs INFO, WARN, and ERROR level messages every 10 seconds for testing logging systems.  The message it sends can be modified based on passing the environment parameter 'whoAmI'.

## GoLang syntax check

If you have [GoLang installed locally](https://fabianlee.org/2022/10/29/golang-installing-the-go-programming-language-on-ubuntu-22-04/), you can run a local sanity check on the main.go syntax before having the remote pipeline do a full build.

```
make local-golang
```

## Github pipeline and published image

The github pipeline takes care of the multi-arch build, and publishes the image to the Github Container Registry.

```
docker pull ghcr.io/fabianlee/google-hello-app-logging-multiarch:latest
```

## Creating tag that invokes Github Action

```
newtag=v1.0.1
git commit -a -m "changes for new tag $newtag" && git push -o ci.skip
git tag $newtag && git push origin $newtag
```

## Deleting tag

```
# delete local tag, then remote
todel=v1.0.1
git tag -d $todel && git push -d origin $todel
```

