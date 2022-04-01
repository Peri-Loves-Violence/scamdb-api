# scamdb-api
API-Library to interact with Scam-DB

# Supported Databases
* [Filesystem/Local](https://github.com/Peri-Loves-Violence/scamdb-api/wiki/Local): A local folder as a Scam-DB
* [Google Spreadsheets](https://github.com/Peri-Loves-Violence/scamdb-api/wiki/Google): Use a spreadsheet on Google Drive as a Scam-DB with the bonus of having an easy-to-use interface.
* [GitHub](https://github.com/Peri-Loves-Violence/scamdb-api/wiki/GitHub): The same as a local one except it's a repository on GitHub
* [MySQL/MariaDB](https://github.com/Peri-Loves-Violence/scamdb-api/wiki/MySQL): MySQL-like RDBMS as a Scam-DB
* [PostgresQL/CockroachDB](https://github.com/Peri-Loves-Violence/scamdb-api/wiki/PostgresQL): A Scam-DB powered by PostgresQL and alike

## Use cases
| Database    | Very Fast access*   | Reliability        | Easy-to-install    | Lightweight        |
| ----------- | ------------------- | ------------------ | ------------------ | ------------------ |
| GitHub      | :white_check_mark:  | :white_check_mark: | :white_check_mark: | :cloud:            |
| Local       | Depending on medium | :x: (except RAID)  | :white_check_mark: | :white_check_mark: |
| SQLite      | Depending on medium | :x: (except RAID)  | :white_check_mark: | :white_check_mark: |
| MySQL       | :x:                 | :x:                | :white_check_mark: | :white_check_mark: |
| CockroachDB | :x:                 | :white_check_mark: | N/A                | :x:/:cloud:        |

*: Database can handle a lot of traffic and is recommended for commercial/large use.
All of them are fast enough for private/small or moderate use with plug-and-play.

# Usage
[See the wiki for more](https://github.com/Peri-Loves-Violence/scamdb-api/wiki/Home)