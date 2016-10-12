# Chaincode Development Environment
The following is a list of dependencies and recommended tools that you should install in order to develop chaincode.

### Git
- [Git download page](https://git-scm.com/downloads)
- [Pro Git ebook](https://git-scm.com/book/en/v2)
- [Git Desktop (for those uncomfortable with git's CLI)](https://desktop.github.com/)

Git is a great version control tool to be familiarize yourself with, both for chaincode development and software development in general.  Also, git bash, which is installed with git on Windows, is an excellent alternative to the the Windows command prompt.

##### Instructions
After following the installation instructions above, you can verify that git is installed using the following command:

```
$ git version
git version 2.9.0.windows.1
```

Once you have git installed, go create an account for yourself on [GitHub](https://Github.com/).  The IBM Blockchain service on Bluemix currently requires that chaincode be in a GitHub repository in order to be deployed through the REST API.

### Go
- [Go 1.6 install](https://golang.org/dl/#go1.6.3)
- [Go installation instructions](https://golang.org/doc/install)
- [Go documentation and tutorials](https://golang.org/doc/)

Currently, Go is the only supported language for writing chaincode.  The Go installation installs a set of Go CLI tools which are very usefull when writing chaincode.  For example, the `go build` command allows you to check that your chaincode actually compiles before you attempt to deploy it to a network.  You should install Go 1.6, so that you have the same version of the language that the fabric is written against.

##### Instructions
Follow the installation instructions linked above.  You can verify that Go is installed properly by running the following commands.  Of course, the output of `go version` may change depending on your operating system.

```
$ go version
go version go1.6.3 windows/amd64

$ echo $GOPATH
C:\gopath
```

Your `GOPATH` does not need to match the one above.  It only matters that you have this variable set to a valid directory on your filesystem.  The installation instructions linked above should take you through the setup of this environment variable.  Why is this variable important?  When you run `go build` to test that your chaincode compiles, Go is going to look for the non-standard dependencies that you list in the `import` block of your chaincode in the `$GOPATH/src` directory.  

### Hyperledger fabric
- [v0.5-developer-preview Hyperledger fabric](https://github.com/hyperledger-archives/fabric/tree/v0.5-developer-preview)
- [v0.6-preview Hyperledger fabric](https://gerrit.hyperledger.org/r/gitweb?p=fabric.git;a=shortlog;h=refs/heads/v0.6)
- [master branch of the Hyperledger fabric](https://gerrit.hyperledger.org/r/gitweb?p=fabric.git;a=summary)

Any piece of chaincode that you write will need to import the chaincode shim from the Hyperledger fabric in order to be able to read and write data on the ledger.  To be able to compile chaincode locally, which you will be doing a lot, you will need to have the fabric code present in your `GOPATH`.

##### Instructions

Three different releases of the fabric are linked above.  The release you choose above should match the Hyperledger network you are deploying your chaincode on.  You will need to make sure that the fabric release you choose is stored under `$GOPATH/hyperledger/fabric`.

The instructions below should take you through the process of properly installing the v0.5 release on your `GOPATH`.
```

# Create the parent directories on your GOPATH
mkdir -p $GOPATH/github.com/hyperledger
cd $GOAPTH/github.com/hyperledger

# Clone the appropriate release codebase into $GOPATH/github.com/hyperledger/fabric
# Note that the v0.5 release is a branch of the repository
git clone -b v0.5-developer-preview https://github.com/hyperledger-archives/fabric.git
```

If the fabric is not installed properly on your `GOPATH`, you will see errors like the one below when building your chaincode:
```
$ go build .
chaincode_example02.go:27:2: cannot find package "github.com/hyperledger/fabric/core/chaincode/shim" in any of:
        C:\Go\src\github.com\hyperledger\fabric\core\chaincode\shim (from $GOROOT)
        C:\gopath\src\github.com\hyperledger\fabric\core\chaincode\shim (from $GOPATH)
```

A list of known specific releases is included below

- [Blockchain service on Bluemix](https://new-console.ng.bluemix.net/catalog/services/blockchain/) - use the v0.5-developer-preview release


### Postman
- [Home page](https://www.getpostman.com/)

Postman is a REST API testing tool.  Though it is deprecated, we still use the REST API in the fabric for this tutorial because it
allows you to deploy and test your chaincode without needing to use the fabric SDK.  You'll learn more about the fabric SDK in our other
examples.

##### Instructions
Go to the Postman website and download the tool.  Depending on your operating system, you may also need to install Chrome to use Postman.  Once you have the tool running, import the request collection included in this repository.  This collection contains requests for enrolling a user on a peer and deploying, invoking, and querying chaincode, which is all you should need to complete this tutorial.

### Node.js
- [Download links](https://nodejs.org/en/download/)

Node.js is NOT necessary to develop chaincode, but most of our demos are built on Node.js, so it might be handy to go ahead and install it now.  Also, you'll need it when you start using the fabric SDK.

##### Instructions
Download the appropriate installation package and make sure the following commands work on your machine:
```
$ node -v
v4.4.7

$ npm -v
3.10.5
```

## IDE Suggestions
### Visual Studio Code
- [Download links](https://code.visualstudio.com/#alt-downloads)

Visual Studio Code is a free IDE that supports both Node.js and Go through plugins.  All of our demos and examples use either one or both of these languages.  It also has tab support, git integration, and debugging support.

### Atom
- [Home page](https://atom.io/)

Like VS Code, Atom has plugins to support any of the languages needed to develop chaincode or modify our examples.