
# lsh

  

lsh is the [Latitude.sh](http://latitude.sh/) command-line tool that will help you make it easier to retrieve any data from your team or perform any action you need.

  

## [](https://dash.readme.com/project/control/v2023-06-01/docs/overview)Installation

  

#### MacOS / Linux / WSL

  

Installing the latest version

  

```bash

curl -sSL  https://raw.githubusercontent.com/latitudesh/lsh/feat/installation-script/install.sh | bash

```

  

#### Windows is not supported yet.

##

### From Github

  

Visit the Releases page and select any version you want to download.

  
  

## [](https://docs.latitude.sh/docs/getting-started)Getting Started

  

Log in into Latitude.sh. An API Key is required.

  

`lsh login <API_KEY>`

  

List your servers

  

`lsh servers list`

  

## [](https://docs.latitude.sh/docs/commands) Commands

  

The list of the available commands is available [here](https://docs.latitude.sh/docs/commands).

  
  

## [](https://docs.latitude.sh/docs/examples-1) Examples

  

List a server with a specific hostname

```bash

lsh servers list --hostname <HOSTNAME>

```

Create a server with Ubuntu 22 

```bash

lsh servers create --operating_system ubuntu_22_04_x64_lts --project <PROJECT_ID_OR_SLUG> --site <LOCATION> --hostname <HOSTNAME> --plan <PLAN>

```
  
List all GPU plans

```bash

lsh plans list --gpu true

```

You can see more examples [here](https://docs.latitude.sh/docs/examples-1). Reach out if you want to see other use cases on `lsh`.
  

## Docs

  

For more information, see the documentation.

- [lsh Docs](https://docs.latitude.sh/docs/getting-started)

- [Product Docs](https://docs.latitude.sh/docs)

- [API Docs](https://docs.latitude.sh/reference)

- [SDKs & Postman Collection](https://docs.latitude.sh/reference/client-libraries)

  

## Provide feedback and contribute

  

- [Open an issue](https://github.com/latitudesh/lsh/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc) for questions, feedback, bug reports or feature requests.

- We welcome pull requests for bug fixes, new features, and improvements to the examples.