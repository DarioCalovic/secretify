![Alt text](ui/static/outlook/icon-64.png?raw=true "Title")

A secure way to transfer or share secrets.

**Please note**: Still under heavy development. A working example of the current **beta release** can be found here https://beta.secretify.io

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* npm
* go

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/DarioCalovic/secretify.git
   ```
2. Install NPM packages
   ```sh
   npm install --prefix ./ui
   ```
3. Run API
   ```sh
   go run cmd/api.go
   ```
4. Run UI
   ```sh
   npm run dev --prefix ./ui
5. Open http://localhost:3000 in browser

## Roadpmap

- [x] First commit
- [ ] CI/CD pipeline
- [ ] Release stable version
- [ ] ..more to come

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. **Any contributions you make are greatly appreciated.**

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

## License

Distributed under the GNU License. See LICENSE.txt for more information.

## Acknowledgments

Our thanks are extended but not limited to the following people:

* [Reto Schelbert](https://github.com/hertus)
