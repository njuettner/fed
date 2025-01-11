# Fed Data Collection and Visualization

https://juni.io/fed/

This repository contains a collection of tools for gathering FED data information such as money supply, consumer price index, ... .

<p align="center"><img src="screenshots/demo.png" width="150"></p>
<p align="center"><img src="screenshots/demo2.png" width="150"></p>
<p align="center"><img src="screenshots/demo3.png" width="150"></p>

## Features

- Collects different data from FED St.Louis Web Service.
- Updates every day automatically via GH action.
- Provides a webserver with all collected data, used as data source (JSON), for easy visualization in a Grafana dashboard.

## Dependencies

This repository relies on the following tools:

- [Grafana](https://grafana.com/) for data visualization.
- [Go](https://golang.org/) for running the data collection scripts.

## Contributing

If you would like to contribute to this repository, please submit a pull request or open an issue for discussion.

## License

This repository is licensed under the MIT License. See the `LICENSE` file for details.


