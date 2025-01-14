<!-- TOC -->
* [Parrot](#parrot)
  * [Features](#features)
  * [How to use](#how-to-use)
  * [Support](#support)
  * [Contributing](#contributing)
<!-- TOC -->

# Parrot
Parrot is a development tool to make the client and server independent during the _development cycles_ by acting as __mock server__. 
Both _front-end_ and _back-end_ developer could create a separate git repository and document their `request` and `response` easy and let the
__parrot__ do the rest until these two could be connected to each other.

## Features
- Easy configuration
- Handle JSON, HTML, and other responses as file
- Support config __hot-reload__ to prevent reruns

## How to use
First of all, you need to create a separate directory for your mock server files. Then download `parrot` executable file for your system by running the
command:
```bash
curl https://raw.githubusercontent.com/ARTM2000/parrot/refs/heads/master/scripts/download.bash -s | bash
```

After the parrot downloaded, you have to pass the mock server configuration to it. __parrot__ by default reads the configuration from `./config.yml` but
you can pass the config file path with `-c` flag. Here is the config file [sample](./examples/config.example.yml). 

After the config file prepared, you can validate it by running the following:
```bash
./parrot validate # for default path
# Or
./parrot validate -c /path/to/config.yml # for specific path
```

Run the parrot mock server by running the following:
```bash
./parrot run # for default path
# Or
./parrot run -c /path/to/config.yml # for specific path
```

## Support
Feel free to [open an issue](https://github.com/ARTM2000/parrot/issues/new) if you have questions, run into bugs, or have a feature request.

## Contributing
Contributions are welcome! Happy Coding :)
