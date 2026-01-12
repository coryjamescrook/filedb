<a id="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://github.com/github_username/repo_name">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">File DB</h3>

  <p align="center">
    A basic file database for hobby projects and simple needs.
    <!-- <br />
    <a href="https://github.com/coryjamescrook/filedb"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/coryjamescrook/filedb">View Demo</a>
    &middot;
    <a href="https://github.com/coryjamescrook/filedb/issues/new?labels=bug&template=bug-report.md">Report Bug</a>
    &middot;
    <a href="https://github.com/coryjamescrook/filedb/issues/new?labels=enhancement&template=feature-request.md">Request Feature</a> -->
  </p>
</div>

<!-- ABOUT THE PROJECT -->

## About The Project

Not everything needs a full database. Sometimes you just want something simple to persist some basic data - whether it be for a POC, a hobby project, or persisting basic settings data. <code>filedb</code> allows you to easily store your application model data to files.

**What this package is NOT for**: Any sort of highly concurrent uses of data reads & writes. Inevitably the basic implementation of this will result in data inconsistencies. This package was really create to ease POC creation and very simple hobby projects that don't justify using a full database, but require some amount of data persistence.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To start using `filedb`, install the go package in your go app:

```bash
go get github.com/coryjamescrook/filedb
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

The following simple example will demonstrate basic usage of this package. It defines a `Settings` data model which is stored to a `json` file.

```go
package main

import (
	"github.com/coryjamescrook/filedb"
	"github.com/coryjamescrook/filedb/translators"
)

type Settings struct {
	filedb.Model

	WelcomeMessage string `json:"welcome_message"`
}

func NewSettings(filePath string) *Settings {
	mySettings := &Settings{}
	err := mySettings.Configure(&filedb.Config{
		ModelObj:        mySettings,
		Path:            filePath,
		DefaultFileData: "{}\n",
		Translator:      translators.JsonTranslator{},
	})

	if err != nil {
		panic(err)
	}

	return mySettings
}

func main() {
	s := NewSettings("./data/settings.json")
	s.WelcomeMessage = "Hello, world!"
	s.Save()
}
```

What's happening here? Let's have a look by breaking the example down, piece by piece.

### 1. Package imports

```go
package main

import (
	"github.com/coryjamescrook/filedb"
	"github.com/coryjamescrook/filedb/translators"
)
```

In our example's `main` package, we import `github.com/coryjamescrook/filedb` and `github.com/coryjamescrook/filedb/translators`. This allows us to use the `filedb.Model` struct to embed in our data model, and imports the included translators for filedb. More about translators to come...

### 2. Define a data model

```go
type Settings struct {
	filedb.Model

	WelcomeMessage string `json:"welcome_message"`
}
```

This data model is defined as a `struct` and embeds the `filedb.Model` type. After that, the actual data for this model is defined, along with json marshal/unmarshal related struct tags (optional). Any public methods will be seriaized and deserialized in the `Save()` and `Load()` functionality. so **do not** define private fields on structs you wish to persist to disk.

### 3. Create a model constructor

```go
func NewSettings(filePath string) *Settings {
	mySettings := &Settings{}
	err := mySettings.Configure(&filedb.Config{
		ModelObj:        mySettings,
		Path:            filePath,
		DefaultFileData: "{}\n",
		Translator:      translators.JsonTranslator{},
	})

	if err != nil {
		panic(err)
	}

	return mySettings
}
```

We first instantiate a new instance of our `Settings` model and store a reference to it in the `mySettings` variable. Once we have an instance of our model, we **must** configure the model in order for it to synchronize with the file system at the defined path. Because we have embedded `filedb.Model`, it implements the `Configure` method, which expects a reference to a `filedb.Config` instance.

We provide a configuration that includes a reference to our model instance `mySettings`, the file path to the db file for this model, any default file data (optional), and a `filedb.Translator` instance - responsible for marshalling and unmarshalling the data to/from disk into our model. In this case, we are storing our data in the `json` file format, so we will use the built-in `JSON` translator from filedb.

If any errors occur, fail with a panic - for demo purposes.

Return a reference to the model, since this is a model constructor function.

### 4. The demo's `main()` function

```go
s := NewSettings("./data/settings.json")
```

This will configure our model and sync it to the provided file on the file system. Since this is the first time defining our model, it will be in sync with the file, with no need to call `Load()` to ensure synchronization with the file system. However, if we were to implement a nested http handler in this example file, we would want to call `Load()`, do any modifications we want (if necessary), and then call `Save()` to ensure that those changes are persisted.

```go
s.WelcomeMessage = "Hello, world!"
```

This reassigns the `WelcomeMessage` field to a new value - `"Hello, world!"`.

```go
s.Save()
```

This saves the state of the model to the file on the file system, persisting the data.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [x] Basic interface for file manipulation
- [x] Thread safety with mutex
- [x] JSON Translator
- [ ] yaml Translator

See the [open issues](https://github.com/coryjamescrook/filedb) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Top contributors:

<a href="https://github.com/coryjamescrook/filedb/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=coryjamescrook/filedb" alt="contrib.rocks image" />
</a>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Cory Crook - coryjamescrook@gmail.com

Project Link: [https://github.com/coryjamescrook/filedb](https://github.com/coryjamescrook/filedb)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

<!-- ## Acknowledgments

- []()
- []()
- []()

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->
