# binance-historical-data-crawler

A comprehensive tool for crawling and collecting historical price data from the Binance cryptocurrency exchange.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Contributing](#contributing)
5. [Next Steps](#next-steps)
6. [License](#license)

## Prerequisites

Before you begin, make sure you have the following installed:

- [Go](https://go.dev/doc/install)

## Installation

Clone the repository to your local machine (or download and unzip it):

```bash
git clone https://github.com/decentralizado/binance-historical-data-crawler.git
```

Navigate to the project directory:

```bash
cd binance-historical-data-crawler
```

Install the required Go modules:

```bash
go mod tidy
```

## Usage

1. **Create Binance API Keys:**

    You can skip this step if you already have one or follow the in-depth guide from Binance [here](https://www.binance.com/en-BH/support/faq/how-to-create-api-360002502072).

    _We recommend you to use an API Key with read-only permissions for this._

2. **Configure Your Crawler**

    **Configure Your API Key: (required)**

    Before you can use this tool, you need to configure your Binance API keys. Create a `.env` file in the project directory and add your API key and secret in the following format:

    ```dotenv
    BINANCE_API_KEY="YOUR_API_KEY"
    BINANCE_API_SECRET="YOUR_API_SECRET"
    ```

    **Specify Price Interval: (optional)**

    In the `.env` file, set the `INTERVAL` parameter to your preferred historical data interval (e.g., "1m" for 1-minute data, "1h" for 1-hour data).

    ```dotenv
    INTERVAL="1h"
    ```

    **Specify Symbol: (optional)**

    In the `.env` file, set the `SYMBOL` parameter to your preferred trading pair symbol (e.g., "BTCEUR" for BTC quote for EUR data, defaults to "BTCUSDT").

    ```dotenv
    SYMBOL="BTCUSDT"
    ```

    **Specify Date Range: (optional)**

    In the `.env` file, set the `START_TIME` and/or `END_TIME` parameters to your preferred date range to collect (e.g., "START_TIME" set to "2021-11-10T00:00:00Z" for collecting from Octer 11th 2021 to now, defaults to "from 30 days ago to now").

    ```dotenv
    START_TIME="2021-11-10T00:00:00Z"
    ```

    **Specify the output file name: (optional)**

    In the `.env` file, set the `FILE_NAME` parameter to your preferred file name (e.g., "history" for saving the results to `out/history.csv`, defaults to "historical_data").

    ```dotenv
    FILE_NAME="history"
    ```

3. **Run the Crawler:**

    Run the crawler program with the following command:

    ```bash
    go run .
    ```

    The tool will start downloading historical data for the specified symbol and date range and save it to the `out/historical_data.csv` directory in CSV format.

## Contributing

If you would like to contribute to the project, please follow these steps:

1. Fork the repository on GitHub.

2. Clone your fork locally:

   ```bash
   git clone https://github.com/your-username/binance-historical-data-crawler.git
   ```

3. Create a new branch for your feature or bug fix:

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. Make your changes, and then commit and push them to your fork.

5. Create a pull request to the original repository.

## Next Steps

We intend to support and implement the following features:

- [ ] Templating of output files (to dynamically support other formats)
- [ ] Configuration through YAML file
- [ ] Automated CI/CD pipelines with GitHub Actions (Lint, Test, Security Analysis, etc.)
- [ ] Automated releases of binaries (Windows, Linux, MacOS)

Also, feel free to submit a feature request through the GitHub Issues of this repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
